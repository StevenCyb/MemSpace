package unit

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidSize = errors.New("invalid size")

type UnitType string

const (
	UnitTypeByte     UnitType = "B"
	UnitTypeKilobyte UnitType = "KB"
	UnitTypeMegabyte UnitType = "MB"
	UnitTypeGigabyte UnitType = "GB"
	UnitTypeTerabyte UnitType = "TB"
	UnitTypePetabyte UnitType = "PB"
)

type Size struct {
	Size int64
}

func NewFromString(s *string) (*Size, error) {
	if s == nil || *s == "" {
		return nil, nil
	}

	re := regexp.MustCompile(`(?i)^(\d+)(b|kb|mb|gb|tb|pb)$`)
	matches := re.FindStringSubmatch(*s)
	if len(matches) != 3 {
		return nil, ErrInvalidSize
	}

	size, _ := strconv.Atoi(matches[1])
	unit := UnitType(strings.ToUpper(matches[2]))
	var rawSize int64

	switch unit {
	case UnitTypeByte:
		rawSize = int64(size)
	case UnitTypeKilobyte:
		rawSize = int64(size * (1 << 10))
	case UnitTypeMegabyte:
		rawSize = int64(size * (1 << 20))
	case UnitTypeGigabyte:
		rawSize = int64(size * (1 << 30))
	case UnitTypeTerabyte:
		rawSize = int64(size * (1 << 40))
	case UnitTypePetabyte:
		rawSize = int64(size * (1 << 50))
	default:
		rawSize = 0
	}

	sizeObj := &Size{
		Size: rawSize,
	}

	return sizeObj, nil
}

func NewFromBytes(bytes int64) *Size {
	if bytes < 0 {
		bytes = 0
	}

	return &Size{
		Size: bytes,
	}
}

func (s *Size) RawSizeString() string {
	var unit UnitType
	var size float64

	switch {
	case s.Size >= 1<<50:
		unit = UnitTypePetabyte
		size = float64(s.Size) / float64(1<<50)
	case s.Size >= 1<<40:
		unit = UnitTypeTerabyte
		size = float64(s.Size) / float64(1<<40)
	case s.Size >= 1<<30:
		unit = UnitTypeGigabyte
		size = float64(s.Size) / float64(1<<30)
	case s.Size >= 1<<20:
		unit = UnitTypeMegabyte
		size = float64(s.Size) / float64(1<<20)
	case s.Size >= 1<<10:
		unit = UnitTypeKilobyte
		size = float64(s.Size) / float64(1<<10)
	default:
		unit = UnitTypeByte
		size = float64(s.Size)
	}

	return fmt.Sprintf("%.2f%s", size, unit)
}

func (s *Size) Add(size *Size) {
	if size == nil {
		return
	}

	s.Size += size.Size
}
