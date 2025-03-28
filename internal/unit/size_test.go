package unit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input     string
		want      *Size
		expectErr bool
	}{
		{input: "1B", want: &Size{Size: 1}, expectErr: false},
		{input: "10kb", want: &Size{Size: 10240}, expectErr: false},
		{input: "100MB", want: &Size{Size: 104857600}, expectErr: false},
		{input: "200gb", want: &Size{Size: 214748364800}, expectErr: false},
		{input: "3TB", want: &Size{Size: 3298534883328}, expectErr: false},
		{input: "4pb", want: &Size{Size: 4503599627370496}, expectErr: false},
		{input: "5", want: nil, expectErr: true},
		{input: "B", want: nil, expectErr: true},
		{input: "1 B", want: nil, expectErr: true},
	}

	for _, tt_ := range tests {
		tt := tt_
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := NewFromString(&tt.input)
			if tt.expectErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.Equal(t, tt.want, got, "Sizes do not match")
				assert.Equal(t, tt.want.Size, got.Size, "Raw size mismatch")
			}
		})
	}
}

func TestNewFromBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input int64
		want  *Size
	}{
		{input: -1, want: &Size{Size: 0}},
		{input: 0, want: &Size{Size: 0}},
		{input: 999, want: &Size{Size: 999}},
	}
	for _, tt_ := range tests {
		tt := tt_
		t.Run(fmt.Sprintf("%d bytes", tt.input), func(t *testing.T) {
			t.Parallel()

			got := NewFromBytes(tt.input)
			assert.Equal(t, tt.want, got, "Sizes do not match")
		})
	}
}

func TestSize_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		initial  *Size
		toAdd    *Size
		expected *Size
	}{
		{initial: &Size{Size: 1}, toAdd: &Size{Size: 1}, expected: &Size{Size: 2}},
		{initial: &Size{Size: 1024}, toAdd: &Size{Size: 1024}, expected: &Size{Size: 2048}},
		{initial: &Size{Size: 1048576}, toAdd: &Size{Size: 1048576}, expected: &Size{Size: 2097152}},
	}

	for _, tt_ := range tests {
		tt := tt_
		t.Run(fmt.Sprintf("%s+%s", tt.initial.RawSizeString(), tt.toAdd.RawSizeString()), func(t *testing.T) {
			t.Parallel()

			tt.initial.Add(tt.toAdd)
			assert.Equal(t, tt.expected, tt.initial, "Sizes do not match")
		})
	}
}

func TestSize_RawSizeString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *Size
		expected string
	}{
		{input: &Size{Size: 1}, expected: "1.00B"},
		{input: &Size{Size: 1024}, expected: "1.00KB"},
		{input: &Size{Size: 1048576}, expected: "1.00MB"},
		{input: &Size{Size: 1073741824}, expected: "1.00GB"},
		{input: &Size{Size: 1099511627776}, expected: "1.00TB"},
		{input: &Size{Size: 1125899906842624}, expected: "1.00PB"},
	}

	for _, tt_ := range tests {
		tt := tt_
		t.Run(fmt.Sprintf("%d to string", tt.input.Size), func(t *testing.T) {
			t.Parallel()

			got := tt.input.RawSizeString()
			assert.Equal(t, tt.expected, got, "String representation does not match")
		})
	}
}
