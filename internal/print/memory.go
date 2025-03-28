package print

import (
	"fmt"
	"syscall"

	"github.com/StevenCyb/MemSpace/internal/unit"

	"github.com/fatih/color"
)

// SystemMemory retrieves and prints the system memory statistics for the given file system path.
// It calculates and displays the total size, free space, available space, and used space in bytes,
// along with their respective percentages. The output is formatted with color for better readability.
//
// Parameters:
//   - path: A string representing the file system path to retrieve memory statistics for.
//
// The function uses the syscall.Statfs system call to gather file system statistics and
// formats the output using the `unit` and `color` packages. If an error occurs during the
// syscall, the function will panic.
func SystemMemory(path string) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		panic(err)
	}

	size := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	avail := stat.Bavail * uint64(stat.Bsize)
	used := size - free

	freePercentage := (float64(free) / float64(size)) * 100
	usedPercentage := (float64(used) / float64(size)) * 100
	availablePercentage := (float64(avail) / float64(size)) * 100

	fmt.Println("Size:     ", color.YellowString(unit.NewFromBytes(int64(size)).RawSizeString()))
	fmt.Printf("Free:      %s - %.2f%%\n", color.GreenString(unit.NewFromBytes(int64(free)).RawSizeString()), freePercentage)
	fmt.Printf("Available: %s - %.2f%%\n", color.GreenString(unit.NewFromBytes(int64(avail)).RawSizeString()), availablePercentage)
	fmt.Printf("Used:      %s - %.2f%%\n", color.RedString(unit.NewFromBytes(int64(used)).RawSizeString()), usedPercentage)
}
