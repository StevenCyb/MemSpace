package main

import (
	"fmt"
	"memspace/internal/cli"
	"memspace/internal/models"
	"memspace/internal/print"
	"memspace/internal/utils"
	"os"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

func main() {
	arguments, err := cli.New(os.Args[1:])
	if err != nil {
		if flags.WroteHelp(err) {
			return
		}
		fmt.Fprintf(os.Stderr, color.RedString("failed to parse arguments: %s\n"), err)
		os.Exit(1)
	}

	if arguments.Memory {
		print.SystemMemory(arguments.BasePath)
	}

	root := models.NewItem(utils.GetName(arguments.BasePath), arguments.BasePath, models.ItemTypeDirectory)
	root.Root = true
	if _, err := utils.WalkAndCollect(root, arguments.BasePath, 0); err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("error walking the path: %s\n"), err)
	}

	print.Tree(root, arguments.Recursive, arguments.DirectoryOnly, arguments.Depth, arguments.Threshold, 0)
}
