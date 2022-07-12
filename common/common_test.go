package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "x", Min[string]("x", "y"))
		assert.Equal(t, "x", Min[string]("y", "x"))
	})

	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float32(1.1), Min[float32](1.1, 1.2))
		assert.Equal(t, float32(1.1), Min[float32](1.2, 1.1))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(1.1), Min[float64](1.1, 1.2))
		assert.Equal(t, float64(1.1), Min[float64](1.2, 1.1))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, int(10), Min[int](10, 20))
		assert.Equal(t, int(10), Min[int](20, 10))
	})

	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, int8(10), Min[int8](10, 20))
		assert.Equal(t, int8(10), Min[int8](20, 10))
	})

	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, int16(10), Min[int16](10, 20))
		assert.Equal(t, int16(10), Min[int16](20, 10))
	})

	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, int32(10), Min[int32](10, 20))
		assert.Equal(t, int32(10), Min[int32](20, 10))
	})

	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, int64(10), Min[int64](10, 20))
		assert.Equal(t, int64(10), Min[int64](20, 10))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(10), Min[uint](10, 20))
		assert.Equal(t, uint(10), Min[uint](20, 10))
	})

	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, uint8(10), Min[uint8](10, 20))
		assert.Equal(t, uint8(10), Min[uint8](20, 10))
	})

	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, uint16(10), Min[uint16](10, 20))
		assert.Equal(t, uint16(10), Min[uint16](20, 10))
	})

	t.Run("uint32", func(t *testing.T) {
		assert.Equal(t, uint32(10), Min[uint32](10, 20))
		assert.Equal(t, uint32(10), Min[uint32](20, 10))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, uint64(10), Min[uint64](10, 20))
		assert.Equal(t, uint64(10), Min[uint64](20, 10))
	})

	t.Run("uintptr", func(t *testing.T) {
		assert.Equal(t, uintptr(10), Min[uintptr](10, 20))
		assert.Equal(t, uintptr(10), Min[uintptr](20, 10))
	})
}

func TestMax(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "y", Max[string]("x", "y"))
		assert.Equal(t, "y", Max[string]("y", "x"))
	})

	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float32(1.2), Max[float32](1.1, 1.2))
		assert.Equal(t, float32(1.2), Max[float32](1.2, 1.1))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(1.2), Max[float64](1.1, 1.2))
		assert.Equal(t, float64(1.2), Max[float64](1.2, 1.1))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, int(20), Max[int](10, 20))
		assert.Equal(t, int(20), Max[int](20, 10))
	})

	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, int8(20), Max[int8](10, 20))
		assert.Equal(t, int8(20), Max[int8](20, 10))
	})

	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, int16(20), Max[int16](10, 20))
		assert.Equal(t, int16(20), Max[int16](20, 10))
	})

	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, int32(20), Max[int32](10, 20))
		assert.Equal(t, int32(20), Max[int32](20, 10))
	})

	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, int64(20), Max[int64](10, 20))
		assert.Equal(t, int64(20), Max[int64](20, 10))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(20), Max[uint](10, 20))
		assert.Equal(t, uint(20), Max[uint](20, 10))
	})

	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, uint8(20), Max[uint8](10, 20))
		assert.Equal(t, uint8(20), Max[uint8](20, 10))
	})

	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, uint16(20), Max[uint16](10, 20))
		assert.Equal(t, uint16(20), Max[uint16](20, 10))
	})

	t.Run("uint32", func(t *testing.T) {
		assert.Equal(t, uint32(20), Max[uint32](10, 20))
		assert.Equal(t, uint32(20), Max[uint32](20, 10))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, uint64(20), Max[uint64](10, 20))
		assert.Equal(t, uint64(20), Max[uint64](20, 10))
	})

	t.Run("uintptr", func(t *testing.T) {
		assert.Equal(t, uintptr(20), Max[uintptr](10, 20))
		assert.Equal(t, uintptr(20), Max[uintptr](20, 10))
	})
}

func TestNewEqualFunc(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		eq := NewEqualFunc[string]()

		assert.False(t, eq("aa", "ab"))
		assert.True(t, eq("aa", "aa"))
	})

	t.Run("float32", func(t *testing.T) {
		eq := NewEqualFunc[float32]()

		assert.False(t, eq(1.1, 1.2))
		assert.True(t, eq(1.1, 1.1))
	})

	t.Run("float64", func(t *testing.T) {
		eq := NewEqualFunc[float64]()

		assert.False(t, eq(1.1, 1.2))
		assert.True(t, eq(1.1, 1.1))
	})

	t.Run("int", func(t *testing.T) {
		eq := NewEqualFunc[int]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("int8", func(t *testing.T) {
		eq := NewEqualFunc[int8]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("int16", func(t *testing.T) {
		eq := NewEqualFunc[int16]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("int32", func(t *testing.T) {
		eq := NewEqualFunc[int32]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("int64", func(t *testing.T) {
		eq := NewEqualFunc[int64]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uint", func(t *testing.T) {
		eq := NewEqualFunc[uint]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uint8", func(t *testing.T) {
		eq := NewEqualFunc[uint8]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uint16", func(t *testing.T) {
		eq := NewEqualFunc[uint16]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uint32", func(t *testing.T) {
		eq := NewEqualFunc[uint32]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uint64", func(t *testing.T) {
		eq := NewEqualFunc[uint64]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})

	t.Run("uintptr", func(t *testing.T) {
		eq := NewEqualFunc[uintptr]()

		assert.False(t, eq(10, 20))
		assert.True(t, eq(10, 10))
	})
}

func TestNewCompareFunc(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		cmp := NewCompareFunc[string]()

		assert.Equal(t, -1, cmp("aa", "ab"))
		assert.Equal(t, 0, cmp("aa", "aa"))
		assert.Equal(t, 1, cmp("ab", "aa"))
	})

	t.Run("float32", func(t *testing.T) {
		cmp := NewCompareFunc[float32]()

		assert.Equal(t, -1, cmp(1.1, 1.2))
		assert.Equal(t, 0, cmp(1.1, 1.1))
		assert.Equal(t, 1, cmp(2.1, 1.1))
	})

	t.Run("float64", func(t *testing.T) {
		cmp := NewCompareFunc[float64]()

		assert.Equal(t, -1, cmp(1.1, 1.2))
		assert.Equal(t, 0, cmp(1.1, 1.1))
		assert.Equal(t, 1, cmp(2.1, 1.1))
	})

	t.Run("int", func(t *testing.T) {
		cmp := NewCompareFunc[int]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("int8", func(t *testing.T) {
		cmp := NewCompareFunc[int8]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("int16", func(t *testing.T) {
		cmp := NewCompareFunc[int16]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("int32", func(t *testing.T) {
		cmp := NewCompareFunc[int32]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("int64", func(t *testing.T) {
		cmp := NewCompareFunc[int64]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uint", func(t *testing.T) {
		cmp := NewCompareFunc[uint]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uint8", func(t *testing.T) {
		cmp := NewCompareFunc[uint8]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uint16", func(t *testing.T) {
		cmp := NewCompareFunc[uint16]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uint32", func(t *testing.T) {
		cmp := NewCompareFunc[uint32]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uint64", func(t *testing.T) {
		cmp := NewCompareFunc[uint64]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})

	t.Run("uintptr", func(t *testing.T) {
		cmp := NewCompareFunc[uintptr]()

		assert.Equal(t, -1, cmp(10, 20))
		assert.Equal(t, 0, cmp(10, 10))
		assert.Equal(t, 1, cmp(20, 10))
	})
}
