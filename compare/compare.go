// Package compare provides comparator functions for standard Go types.
package compare

import (
	"strings"
	"time"
)

// Func type is a function for comparing two values of the same type.
type Func func(interface{}, interface{}) int

// Uint compares two uint numbers.
func Uint(a, b interface{}) int {
	u1, _ := a.(uint)
	u2, _ := b.(uint)
	switch {
	case u1 < u2:
		return -1
	case u1 > u2:
		return 1
	default:
		return 0
	}
}

// Uint8 compares two uint8 numbers.
func Uint8(a, b interface{}) int {
	u1, _ := a.(uint8)
	u2, _ := b.(uint8)
	switch {
	case u1 < u2:
		return -1
	case u1 > u2:
		return 1
	default:
		return 0
	}
}

// Uint16 compares two uint16 numbers.
func Uint16(a, b interface{}) int {
	u1, _ := a.(uint16)
	u2, _ := b.(uint16)
	switch {
	case u1 < u2:
		return -1
	case u1 > u2:
		return 1
	default:
		return 0
	}
}

// Uint32 compares two uint32 numbers.
func Uint32(a, b interface{}) int {
	u1, _ := a.(uint32)
	u2, _ := b.(uint32)
	switch {
	case u1 < u2:
		return -1
	case u1 > u2:
		return 1
	default:
		return 0
	}
}

// Uint64 compares two uint64 numbers.
func Uint64(a, b interface{}) int {
	u1, _ := a.(uint64)
	u2, _ := b.(uint64)
	switch {
	case u1 < u2:
		return -1
	case u1 > u2:
		return 1
	default:
		return 0
	}
}

// Int compares two int numbers.
func Int(a, b interface{}) int {
	i1, _ := a.(int)
	i2, _ := b.(int)
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

// Int8 compares two int8 numbers.
func Int8(a, b interface{}) int {
	i1, _ := a.(int8)
	i2, _ := b.(int8)
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

// Int16 compares two int16 numbers.
func Int16(a, b interface{}) int {
	i1, _ := a.(int16)
	i2, _ := b.(int16)
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

// Int32 compares two int32 numbers.
func Int32(a, b interface{}) int {
	i1, _ := a.(int32)
	i2, _ := b.(int32)
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

// Int64 compares two int64 numbers.
func Int64(a, b interface{}) int {
	i1, _ := a.(int64)
	i2, _ := b.(int64)
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

// Float32 compares two float32 numbers.
func Float32(a, b interface{}) int {
	f1, _ := a.(float32)
	f2, _ := b.(float32)
	switch {
	case f1 < f2:
		return -1
	case f1 > f2:
		return 1
	default:
		return 0
	}
}

// Float64 compares two float32 numbers.
func Float64(a, b interface{}) int {
	f1, _ := a.(float64)
	f2, _ := b.(float64)
	switch {
	case f1 < f2:
		return -1
	case f1 > f2:
		return 1
	default:
		return 0
	}
}

// String compares two strings.
func String(a, b interface{}) int {
	s1, _ := a.(string)
	s2, _ := b.(string)
	return strings.Compare(s1, s2)
}

// Time compare two timestamps.
func Time(a, b interface{}) int {
	t1, _ := a.(time.Time)
	t2, _ := b.(time.Time)
	switch {
	case t1.Before(t2):
		return -1
	case t1.After(t2):
		return 1
	default:
		return 0
	}
}
