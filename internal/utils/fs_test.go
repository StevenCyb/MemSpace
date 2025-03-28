package utils

import (
	"os"
	"testing"

	"github.com/StevenCyb/MemSpace/internal/models"
	"github.com/StevenCyb/MemSpace/internal/unit"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{input: "/path/to/file.txt", expected: "file.txt"},
		{input: "/path/to/directory/", expected: "directory"},
		{input: "relative/path/file.go", expected: "file.go"},
		{input: "file_without_path", expected: "file_without_path"},
		{input: "/", expected: "/"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got := GetName(tt.input)
			assert.Equal(t, tt.expected, got, "Base name does not match")
		})
	}
}

func TestFileSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setup       func() (string, func())
		expected    *unit.Size
		expectError bool
	}{
		{
			name: "Valid file",
			setup: func() (string, func()) {
				file, err := os.CreateTemp("", "testfile")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				content := []byte("Hello, World!")
				if _, err := file.Write(content); err != nil {
					t.Fatalf("Failed to write to temp file: %v", err)
				}
				file.Close()
				return file.Name(), func() { os.Remove(file.Name()) }
			},
			expected:    unit.NewFromBytes(13),
			expectError: false,
		},
		{
			name: "Non-existent file",
			setup: func() (string, func()) {
				return "/non/existent/file.txt", func() {}
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			path, cleanup := tt.setup()
			defer cleanup()

			got, err := FileSize(path)
			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error occurred")
				assert.Equal(t, tt.expected, got, "File sizes do not match")
			}
		})
	}
}

func TestWalkAndCollect(t *testing.T) {
	expected := &models.Item{
		Name:     "./test_data",
		Path:     ".",
		ItemType: models.ItemTypeDirectory,
		Size:     &unit.Size{Size: 10},
		Children: []*models.Item{
			{
				Name:     "a",
				Path:     "test_data/a",
				ItemType: models.ItemTypeFile,
				Size:     &unit.Size{Size: 2},
				Children: []*models.Item{},
			},
			{
				Name:     "b",
				Path:     "test_data/b",
				ItemType: models.ItemTypeFile,
				Size:     &unit.Size{Size: 3},
				Children: []*models.Item{},
			},
			{
				Name:     "c",
				Path:     "test_data/c",
				ItemType: models.ItemTypeDirectory,
				Size:     &unit.Size{Size: 5},
				Children: []*models.Item{
					{
						Name:     "c.txt",
						Path:     "test_data/c/c.txt",
						ItemType: models.ItemTypeFile,
						Size:     &unit.Size{Size: 2},
						Children: []*models.Item{},
					},
					{
						Name:     "d",
						Path:     "test_data/c/d",
						ItemType: models.ItemTypeDirectory,
						Size:     &unit.Size{Size: 3},
						Children: []*models.Item{
							{
								Name:     "d.dat",
								Path:     "test_data/c/d/d.dat",
								ItemType: models.ItemTypeFile,
								Size:     &unit.Size{Size: 3},
								Children: []*models.Item{},
							},
						},
					},
				},
			},
		},
	}
	parent := models.NewItem("./test_data", ".", models.ItemTypeDirectory)
	totalSize, err := WalkAndCollect(parent, "./test_data", 0)
	assert.NoError(t, err, "Unexpected error occurred")
	assert.Equal(t, &unit.Size{Size: 10}, totalSize)
	assert.Equal(t, expected, parent)
}
