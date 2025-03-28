package models

import "github.com/stevencyb/memspace/internal/unit"

// ItemType represents a custom type used to define different categories or types of items.
// It is implemented as a byte to minimize memory usage and improve performance.
type ItemType byte

// ItemTypeDirectory represents an item type that corresponds to a directory.
const (
	ItemTypeDirectory ItemType = iota
	ItemTypeFile
)

// Item represents a hierarchical structure that can be used to model
// files, directories, or other similar entities. Each Item can have
// child Items, forming a tree-like structure.
//
// Fields:
//   - Root: Indicates whether the Item is the root of the hierarchy.
//   - Name: The name of the Item.
//   - Path: The full path to the Item.
//   - ItemType: The type of the Item (e.g., file, directory).
//   - Size: The size of the Item, represented as a pointer to a unit.Size.
//   - Children: A slice of child Items, representing the hierarchical
//     relationship.
type Item struct {
	Root     bool
	Name     string
	Path     string
	ItemType ItemType
	Size     *unit.Size
	Children []*Item
}

// NewItem creates and returns a new Item instance with the specified name, path, and item type.
// It initializes the Children slice as an empty slice of Item pointers.
//
// Parameters:
//   - name: The name of the item.
//   - path: The path associated with the item.
//   - itemType: The type of the item, represented as an ItemType.
//
// Returns:
//
//	A pointer to the newly created Item.
func NewItem(name, path string, itemType ItemType) *Item {
	return &Item{
		Name:     name,
		Path:     path,
		ItemType: itemType,
		Children: make([]*Item, 0),
	}
}

// NewItemWithSize creates a new Item instance with the specified name, path, item type, and size.
// It initializes the Children field as an empty slice of Item pointers.
//
// Parameters:
//   - name: The name of the item.
//   - path: The file system path of the item.
//   - itemType: The type of the item (e.g., file, directory).
//   - size: A pointer to the size of the item.
//
// Returns:
//
//	A pointer to the newly created Item instance.
func NewItemWithSize(name, path string, itemType ItemType, size *unit.Size) *Item {
	return &Item{
		Name:     name,
		Path:     path,
		ItemType: itemType,
		Size:     size,
		Children: make([]*Item, 0),
	}
}
