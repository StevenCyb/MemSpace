package utils

import (
	"fmt"
	"memspace/internal/models"
	"memspace/internal/unit"
	"os"
	"path/filepath"
)

var ErrRecursionEnd = fmt.Errorf("recursion end")

// GetName returns the base name of the given file path.
// It extracts the last element of the path, which is typically
// the file or directory name.
//
// Parameters:
//   - path: The file path from which to extract the base name.
//
// Returns:
//
//	The base name of the provided file path.
//	- error: An error if the file cannot be opened or its metadata cannot be retrieved.
func GetName(path string) string {
	return filepath.Base(path)
}

// FileSize returns the size of the file at the specified path as a *unit.Size.
// It opens the file, retrieves its metadata, and calculates its size in bytes.
// If an error occurs while opening the file or retrieving its metadata, it
// returns the error.
//
// Parameters:
//   - path: The file path as a string.
//
// Returns:
//   - *unit.Size: The size of the file in bytes wrapped in a unit.Size object.
//   - error: An error if the file cannot be opened or its metadata cannot be retrieved.
func FileSize(path string) (*unit.Size, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return unit.NewFromBytes(stat.Size()), nil
}

// WalkAndCollect traverses the directory tree starting from the specified path,
// collects information about files and directories, and calculates their sizes.
// It populates the provided parent *models.Item with its children and their sizes.
//
// Parameters:
//   - parent: A pointer to a models.Item representing the parent directory.
//   - path: The file system path to start traversing from.
//   - currentDepth: The current depth of traversal (used for recursive calls).
//
// Returns:
//   - *unit.Size: The total size of all files and directories under the given path.
//   - error: An error if any issues occur during directory traversal or file size calculation.
//
// The function reads the contents of the directory at the given path. For each entry:
//   - If the entry is a directory, it recursively calls WalkAndCollect to process its contents.
//   - If the entry is a file, it calculates its size and adds it to the parent's children.
//
// The parent *models.Item is updated with its children and their respective sizes.
// The total size of all files and directories is returned.
func WalkAndCollect(parent *models.Item, path string, currentDepth int) (*unit.Size, error) {
	totalSize := unit.NewFromBytes(0)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			relativePath := filepath.Join(path, entry.Name())
			newParent := models.NewItem(entry.Name(), relativePath, models.ItemTypeDirectory)

			size, err := WalkAndCollect(newParent, filepath.Join(path, entry.Name()), currentDepth+1)
			if err != nil {
				return nil, err
			}
			newParent.Size = size
			parent.Children = append(parent.Children, newParent)

			totalSize.Add(size)
		} else {
			size, err := FileSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return nil, err
			}

			totalSize.Add(size)
			relativePath := filepath.Join(path, entry.Name())
			parent.Children = append(parent.Children, models.NewItemWithSize(entry.Name(), relativePath, models.ItemTypeFile, size))
		}
	}

	parent.Size = totalSize

	return totalSize, nil
}
