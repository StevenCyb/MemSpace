package cli

import (
	"testing"

	"github.com/StevenCyb/MemSpace/internal/unit"

	"github.com/stretchr/testify/assert"
)

func TestHelpCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name: "Help flag long",
			args: []string{"--help"},
		},
		{
			name: "Help flag short",
			args: []string{"-h"},
		},
	}

	for _, tt_ := range tests {
		tt := tt_
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args)
			assert.Error(t, err, "Expected an error for help command but got none")
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		args      []string
		want      *Arguments
		expectErr bool
	}{
		{
			name: "Default values",
			args: []string{},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         nil,
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Custom base path",
			args: []string{"--path", "/tmp"},
			want: &Arguments{
				BasePath:      "/tmp",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         nil,
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Custom base path",
			args: []string{"-p", "/tmp"},
			want: &Arguments{
				BasePath:      "/tmp",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         nil,
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Directory only flag",
			args: []string{"--dir"},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: true,
				Recursive:     false,
				Depth:         nil,
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Recursive flag",
			args: []string{"-r"},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: false,
				Recursive:     true,
				Depth:         nil,
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Depth specified",
			args: []string{"--depth", "3"},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         intPtr(3),
				Threshold:     nil,
			},
			expectErr: false,
		},
		{
			name: "Threshold specified",
			args: []string{"--threshold", "10MB"},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         nil,
				Threshold:     &unit.Size{Size: 10 * 1024 * 1024},
			},
			expectErr: false,
		},
		{
			name: "Memory flag",
			args: []string{"--memory"},
			want: &Arguments{
				BasePath:      ".",
				DirectoryOnly: false,
				Recursive:     false,
				Depth:         nil,
				Threshold:     nil,
				Memory:        true,
			},
			expectErr: false,
		},
	}

	for _, tt_ := range tests {
		tt := tt_
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args)
			if tt.expectErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.Equal(t, tt.want, got, "Arguments do not match")
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}
