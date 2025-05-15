// Package hash provides hash functions for standard Go types.
// It leverages the built-in "hash" package, with FNV-1 (Fowler–Noll–Vo) as the default hashing algorithm.
//
// The functions in this package are designed for high performance,
// following a similar approach to the built-in "encoding/binary" package,
// but avoiding type assertions or reflection to minimize overhead.
package hash

import (
	"hash"
	"hash/fnv"
	"unsafe"
)

// Hasher is a generic interface that defines a method for computing a hash value for any type T.
type Hasher interface {
	Hash() uint64
}

// HashFunc defines a generic function type for hashing a key of type K.
type HashFunc[T any] func(T) uint64

func ensureHasher(h hash.Hash64) hash.Hash64 {
	if h == nil {
		h = fnv.New64()
	}
	return h
}

// HashFuncForBool creates a HashFunc for boolean-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForBool[T ~bool](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 1)

	return func(v T) uint64 {
		h.Reset()

		if v {
			b[0] = 1
		} else {
			b[0] = 0
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForBoolSlice creates a HashFunc for slice types with boolean elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForBoolSlice[T ~[]bool](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, len(v))
		for i, x := range v {
			if x {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt8 creates a HashFunc for int8-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt8[T ~int8](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 1)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt8Slice creates a HashFunc for slice types with int8 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt8Slice[T ~[]int8](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, len(v))
		for i, x := range v {
			// Little-endian
			b[i] = byte(x)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt16 creates a HashFunc for int16-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt16[T ~int16](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 2)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt16Slice creates a HashFunc for slice types with int16 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt16Slice[T ~[]int16](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 2*len(v))
		for i, x := range v {
			// Little-endian
			b[2*i] = byte(x)
			b[2*i+1] = byte(x >> 8)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt32 creates a HashFunc for int32-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt32[T ~int32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 4)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt32Slice creates a HashFunc for slice types with int32 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt32Slice[T ~[]int32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 4*len(v))
		for i, x := range v {
			// Little-endian
			b[4*i] = byte(x)
			b[4*i+1] = byte(x >> 8)
			b[4*i+2] = byte(x >> 16)
			b[4*i+3] = byte(x >> 24)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt64 creates a HashFunc for int64-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt64[T ~int64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 8)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)
		b[4] = byte(v >> 32)
		b[5] = byte(v >> 40)
		b[6] = byte(v >> 48)
		b[7] = byte(v >> 56)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt64Slice creates a HashFunc for slice types with int64 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt64Slice[T ~[]int64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 8*len(v))
		for i, x := range v {
			// Little-endian
			b[8*i] = byte(x)
			b[8*i+1] = byte(x >> 8)
			b[8*i+2] = byte(x >> 16)
			b[8*i+3] = byte(x >> 24)
			b[8*i+4] = byte(x >> 32)
			b[8*i+5] = byte(x >> 40)
			b[8*i+6] = byte(x >> 48)
			b[8*i+7] = byte(x >> 56)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForInt creates a HashFunc for int-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForInt[T ~int](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))
	b := make([]byte, size)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		for i := range size {
			b[i] = byte(v >> (i * 8))
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForIntSlice creates a HashFunc for slice types with int elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForIntSlice[T ~[]int](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, size*len(v))
		for i, x := range v {
			// Little-endian
			for j := range size {
				b[size*i+j] = byte(x >> (j * 8))
			}
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint8 creates a HashFunc for uint8-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint8[T ~uint8](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 1)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint8Slice creates a HashFunc for slice types with uint8 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint8Slice[T ~[]uint8](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, len(v))
		for i, x := range v {
			// Little-endian
			b[i] = byte(x)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint16 creates a HashFunc for uint16-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint16[T ~uint16](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 2)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint16Slice creates a HashFunc for slice types with uint16 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint16Slice[T ~[]uint16](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 2*len(v))
		for i, x := range v {
			// Little-endian
			b[2*i] = byte(x)
			b[2*i+1] = byte(x >> 8)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint32 creates a HashFunc for uint32-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint32[T ~uint32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 4)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint32Slice creates a HashFunc for slice types with uint32 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint32Slice[T ~[]uint32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 4*len(v))
		for i, x := range v {
			// Little-endian
			b[4*i] = byte(x)
			b[4*i+1] = byte(x >> 8)
			b[4*i+2] = byte(x >> 16)
			b[4*i+3] = byte(x >> 24)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint64 creates a HashFunc for uint64-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint64[T ~uint64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 8)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)
		b[4] = byte(v >> 32)
		b[5] = byte(v >> 40)
		b[6] = byte(v >> 48)
		b[7] = byte(v >> 56)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint64Slice creates a HashFunc for slice types with uint64 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint64Slice[T ~[]uint64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, 8*len(v))
		for i, x := range v {
			// Little-endian
			b[8*i] = byte(x)
			b[8*i+1] = byte(x >> 8)
			b[8*i+2] = byte(x >> 16)
			b[8*i+3] = byte(x >> 24)
			b[8*i+4] = byte(x >> 32)
			b[8*i+5] = byte(x >> 40)
			b[8*i+6] = byte(x >> 48)
			b[8*i+7] = byte(x >> 56)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUintptr creates a HashFunc for uintptr-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUintptr[T ~uintptr](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))
	b := make([]byte, size)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		for i := range size {
			b[i] = byte(v >> (i * 8))
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUintptrSlice creates a HashFunc for slice types with uintptr elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUintptrSlice[T ~[]uintptr](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, size*len(v))
		for i, x := range v {
			// Little-endian
			for j := range size {
				b[size*i+j] = byte(x >> (j * 8))
			}
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUint creates a HashFunc for uint-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUint[T ~uint](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))
	b := make([]byte, size)

	return func(v T) uint64 {
		h.Reset()

		// Little-endian
		for i := range size {
			b[i] = byte(v >> (i * 8))
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForUintSlice creates a HashFunc for slice types with uint elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForUintSlice[T ~[]uint](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	var v T
	size := int(unsafe.Sizeof(v))

	return func(v T) uint64 {
		h.Reset()

		b := make([]byte, size*len(v))
		for i, x := range v {
			// Little-endian
			for j := range size {
				b[size*i+j] = byte(x >> (j * 8))
			}
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForFloat32 creates a HashFunc for float32-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForFloat32[T ~float32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 4)

	return func(f T) uint64 {
		h.Reset()

		// The IEEE 754 binary representation of f,
		//   where the sign bit of f is preserved in the same bit position in the resulting uint32 value.
		v := *(*uint32)(unsafe.Pointer(&f))

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForFloat32Slice creates a HashFunc for slice types with float32 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForFloat32Slice[T ~[]float32](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(f T) uint64 {
		h.Reset()

		b := make([]byte, 4*len(f))
		for i, x := range f {
			// The IEEE 754 binary representation of f,
			//   where the sign bit of f is preserved in the same bit position in the resulting uint32 value.
			v := *(*uint32)(unsafe.Pointer(&x))

			// Little-endian
			b[4*i] = byte(v)
			b[4*i+1] = byte(v >> 8)
			b[4*i+2] = byte(v >> 16)
			b[4*i+3] = byte(v >> 24)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForFloat64 creates a HashFunc for float64-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForFloat64[T ~float64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 8)

	return func(f T) uint64 {
		h.Reset()

		// The IEEE 754 binary representation of f,
		//   where the sign bit of f is preserved in the same bit position in the resulting uint64 value.
		v := *(*uint64)(unsafe.Pointer(&f))

		// Little-endian
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)
		b[4] = byte(v >> 32)
		b[5] = byte(v >> 40)
		b[6] = byte(v >> 48)
		b[7] = byte(v >> 56)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForFloat64Slice creates a HashFunc for slice types with float64 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForFloat64Slice[T ~[]float64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(f T) uint64 {
		h.Reset()

		b := make([]byte, 8*len(f))
		for i, x := range f {
			// The IEEE 754 binary representation of f,
			//   where the sign bit of f is preserved in the same bit position in the resulting uint64 value.
			v := *(*uint64)(unsafe.Pointer(&x))

			// Little-endian
			b[8*i] = byte(v)
			b[8*i+1] = byte(v >> 8)
			b[8*i+2] = byte(v >> 16)
			b[8*i+3] = byte(v >> 24)
			b[8*i+4] = byte(v >> 32)
			b[8*i+5] = byte(v >> 40)
			b[8*i+6] = byte(v >> 48)
			b[8*i+7] = byte(v >> 56)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForComplex64 creates a HashFunc for complex64-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForComplex64[T ~complex64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 8)

	return func(c T) uint64 {
		h.Reset()

		c64 := complex64(c)
		re, im := real(c64), imag(c64)

		// The IEEE 754 binary representation of f,
		//   where the sign bit of f is preserved in the same bit position in the resulting uint32 value.
		u, v := *(*uint32)(unsafe.Pointer(&re)), *(*uint32)(unsafe.Pointer(&im))

		// Little-endian
		b[0] = byte(u)
		b[1] = byte(u >> 8)
		b[2] = byte(u >> 16)
		b[3] = byte(u >> 24)
		b[4] = byte(v)
		b[5] = byte(v >> 8)
		b[6] = byte(v >> 16)
		b[7] = byte(v >> 24)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForComplex64Slice creates a HashFunc for slice types with complex64 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForComplex64Slice[T ~[]complex64](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(c T) uint64 {
		h.Reset()

		b := make([]byte, 8*len(c))
		for i, x := range c {
			c64 := complex64(x)
			re, im := real(c64), imag(c64)

			// The IEEE 754 binary representation of f,
			//   where the sign bit of f is preserved in the same bit position in the resulting uint32 value.
			u, v := *(*uint32)(unsafe.Pointer(&re)), *(*uint32)(unsafe.Pointer(&im))

			// Little-endian
			b[8*i] = byte(u)
			b[8*i+1] = byte(u >> 8)
			b[8*i+2] = byte(u >> 16)
			b[8*i+3] = byte(u >> 24)
			b[8*i+4] = byte(v)
			b[8*i+5] = byte(v >> 8)
			b[8*i+6] = byte(v >> 16)
			b[8*i+7] = byte(v >> 24)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForComplex128 creates a HashFunc for complex128-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForComplex128[T ~complex128](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)
	b := make([]byte, 16)

	return func(c T) uint64 {
		h.Reset()

		c128 := complex128(c)
		re, im := real(c128), imag(c128)

		// The IEEE 754 binary representation of f,
		//   where the sign bit of f is preserved in the same bit position in the resulting uint64 value.
		u, v := *(*uint64)(unsafe.Pointer(&re)), *(*uint64)(unsafe.Pointer(&im))

		// Little-endian
		b[0] = byte(u)
		b[1] = byte(u >> 8)
		b[2] = byte(u >> 16)
		b[3] = byte(u >> 24)
		b[4] = byte(u >> 32)
		b[5] = byte(u >> 40)
		b[6] = byte(u >> 48)
		b[7] = byte(u >> 56)
		b[8] = byte(v)
		b[9] = byte(v >> 8)
		b[10] = byte(v >> 16)
		b[11] = byte(v >> 24)
		b[12] = byte(v >> 32)
		b[13] = byte(v >> 40)
		b[14] = byte(v >> 48)
		b[15] = byte(v >> 56)

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForComplex128Slice creates a HashFunc for slice types with complex128 elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForComplex128Slice[T ~[]complex128](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(c T) uint64 {
		h.Reset()

		b := make([]byte, 16*len(c))
		for i, x := range c {
			c128 := complex128(x)
			re, im := real(c128), imag(c128)

			// The IEEE 754 binary representation of f,
			//   where the sign bit of f is preserved in the same bit position in the resulting uint64 value.
			u, v := *(*uint64)(unsafe.Pointer(&re)), *(*uint64)(unsafe.Pointer(&im))

			// Little-endian
			b[16*i] = byte(u)
			b[16*i+1] = byte(u >> 8)
			b[16*i+2] = byte(u >> 16)
			b[16*i+3] = byte(u >> 24)
			b[16*i+4] = byte(u >> 32)
			b[16*i+5] = byte(u >> 40)
			b[16*i+6] = byte(u >> 48)
			b[16*i+7] = byte(u >> 56)
			b[16*i+8] = byte(v)
			b[16*i+9] = byte(v >> 8)
			b[16*i+10] = byte(v >> 16)
			b[16*i+11] = byte(v >> 24)
			b[16*i+12] = byte(v >> 32)
			b[16*i+13] = byte(v >> 40)
			b[16*i+14] = byte(v >> 48)
			b[16*i+15] = byte(v >> 56)
		}

		// Hash.Write never returns an error
		_, _ = h.Write(b)
		return h.Sum64()
	}
}

// HashFuncForString creates a HashFunc for string-like types.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForString[T ~string](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(s T) uint64 {
		h.Reset()

		// Hash.Write never returns an error
		_, _ = h.Write([]byte(s))
		return h.Sum64()
	}
}

// HashFuncForStringSlice creates a HashFunc for slice types with string elements.
// If h is nil, a default hash.Hash64 implementation will be used.
func HashFuncForStringSlice[T ~[]string](h hash.Hash64) HashFunc[T] {
	h = ensureHasher(h)

	return func(s T) uint64 {
		h.Reset()

		for _, x := range s {
			// Hash.Write never returns an error
			_, _ = h.Write([]byte(x))
		}

		return h.Sum64()
	}
}
