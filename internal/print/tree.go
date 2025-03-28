package print

import (
	"fmt"
	"strings"

	"github.com/stevencyb/memspace/internal/models"
	"github.com/stevencyb/memspace/internal/unit"

	"github.com/fatih/color"
)

// Tree prints a visual representation of a directory tree structure starting from the given item.
// It supports recursive traversal, filtering by directory-only items, and limiting depth or size thresholds.
//
// Parameters:
//   - item: The root item of the tree to be printed. It must be of type *models.Item.
//   - recursive: A boolean indicating whether to traverse the tree recursively.
//   - dirOnly: A boolean indicating whether to include only directories in the output.
//   - depth: A pointer to an integer specifying the maximum depth to traverse. If nil, no depth limit is applied.
//   - threshold: A pointer to a unit.Size specifying the minimum size of items to include. If nil, no size threshold is applied.
//   - currentDepth: An integer representing the current depth of traversal (used internally for recursion).
//
// Behavior:
//   - If the item is marked as the root, it prints the root directory with its size.
//   - Traverses the children of the item and prints them with appropriate prefixes to indicate tree structure.
//   - Applies the depth and size thresholds to filter items.
//   - If dirOnly is true, only directories are included in the output.
//   - Uses visual indicators (e.g., "📁" for directories and "📄" for files) and colors for better readability.
//
// Example:
//
//	Tree(rootItem, true, false, nil, nil, 0)
func Tree(item *models.Item, recursive bool, dirOnly bool, depth *int, threshold *unit.Size, currentDepth int) {
	if item.Root {
		fmt.Printf("📁%s [%s]\n", color.GreenString(item.Name), color.YellowString(item.Size.RawSizeString()))
	}

	if (!recursive && currentDepth > 0) || (depth != nil && currentDepth > *depth) {
		return
	}

	for i, child := range item.Children {
		isLast := i == len(item.Children)-1
		prefix := "│-"
		if isLast {
			prefix = "└-"
		}

		if child.ItemType == models.ItemTypeDirectory {
			if (threshold == nil || threshold.Size <= child.Size.Size) && (depth == nil || currentDepth <= *depth) {
				fmt.Printf("%s%s📁%s [%s]\n", strings.Repeat("│ ", currentDepth), prefix, color.GreenString(child.Name), color.YellowString(child.Size.RawSizeString()))
			}
			Tree(child, recursive, dirOnly, depth, threshold, currentDepth+1)
		} else if !dirOnly && (threshold == nil || threshold.Size <= child.Size.Size) && (depth == nil || currentDepth <= *depth) {
			fmt.Printf("%s%s📄%s [%s]\n", strings.Repeat("│ ", currentDepth), prefix, color.BlueString(child.Name), color.YellowString(child.Size.RawSizeString()))
		}
	}
}
