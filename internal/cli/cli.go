package cli

import (
	"fmt"
	"os"

	"github.com/StevenCyb/MemSpace/internal/unit"

	"github.com/jessevdk/go-flags"
)

// Arguments represents the configuration options for a CLI command.
// It includes the following fields:
//
// - BasePath: The base directory path where the operation will start.
// - DirectoryOnly: A flag indicating whether to process only directories.
// - Recursive: A flag indicating whether to process directories recursively.
// - Depth: An optional pointer to an integer specifying the maximum depth for recursion.
// - Threshold: An optional pointer to a unit.Size value specifying a size threshold for filtering.
// - Memory: A flag indicating whether to show driver memory.
type Arguments struct {
	BasePath      string
	DirectoryOnly bool
	Recursive     bool
	Depth         *int
	Threshold     *unit.Size
	Memory        bool
}

// New creates a new instance of Arguments by parsing the provided command-line arguments.
// It uses the flags package to define and parse the options available to the CLI.
//
// Parameters:
//   - args: A slice of strings representing the command-line arguments.
//
// Returns:
//   - A pointer to an Arguments struct populated with the parsed values.
//   - An error if the arguments cannot be parsed, the threshold value is invalid, or
//     the resulting Arguments fail verification.
//
// The supported options are:
//   - -p, --path: The base path to start scanning from (default: ".").
//   - -d, --dir: If set, only calculates the size of directories.
//   - -r, --recursive: If set, recursively calculates the size of directories.
//   - -e, --depth: Specifies the depth of recursion (default: -1 for unlimited depth).
//   - -t, --threshold: Specifies a threshold value to alert on.
//   - -m, --memory: If set, shows driver memory.
//
// Example usage:
//
//	args := []string{"--path", "/home/user", "--recursive", "--depth", "2"}
//	arguments, err := New(args)
//	if err != nil {
//	    log.Fatalf("Error: %v", err)
//	}
func New(args []string) (*Arguments, error) {
	var opts struct {
		Path      string `short:"p" long:"path" default:"." description:"The base path to start scanning from"`
		Dir       bool   `short:"d" long:"dir" description:"Only show directories"`
		Recursive bool   `short:"r" long:"recursive" description:"Show files (and directories) Recursively"`
		Depth     int    `short:"e" long:"depth" default:"-1" description:"The depth of recursion"`
		Threshold string `short:"t" long:"threshold" default:"" description:"Show only files or directories larger than the threshold"`
		Memory    bool   `short:"m" long:"memory" description:"Show driver memory"`
	}

	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.ParseArgs(args); flags.WroteHelp(err) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to parse arguments: %w", err)
	}

	arguments := &Arguments{
		BasePath:      opts.Path,
		DirectoryOnly: opts.Dir,
		Recursive:     opts.Recursive,
		Memory:        opts.Memory,
	}

	if opts.Depth >= 0 {
		arguments.Depth = &opts.Depth
	}

	var err error
	arguments.Threshold, err = unit.NewFromString(&opts.Threshold)
	if err != nil {
		return nil, fmt.Errorf("invalid threshold: %s", err)
	}

	if err := arguments.Verify(); err != nil {
		return nil, err
	}

	return arguments, nil
}

// Verify checks the validity of the Arguments struct by ensuring that the BasePath field
// is not empty and that the specified path exists in the filesystem.
// It returns an error if the BasePath is empty or if the path does not exist.
func (a Arguments) Verify() error {
	if a.BasePath == "" {
		return fmt.Errorf("base path cannot be empty")
	}

	if _, err := os.Stat(a.BasePath); os.IsNotExist(err) {
		return fmt.Errorf("base path does not exist: %s", a.BasePath)
	}

	return nil
}
