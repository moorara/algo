package trie

import (
	"fmt"
	"strings"
)

var (
	empty = &bitString{}

	// FIXME: seems like a linter bug!
	// nolint: unused
	zero = &bitString{
		bits: []byte{0x00},
		len:  1,
	}

	// FIXME: seems like a linter bug!
	// nolint: unused
	one = &bitString{
		bits: []byte{0x80},
		len:  1,
	}
)

// bitString is a sequence of bits padded with trailing zeros.
type bitString struct {
	bits []byte
	len  int
}

func newBitString(s string) *bitString {
	return &bitString{
		bits: []byte(s),
		len:  len(s) * 8,
	}
}

func (b *bitString) Len() int {
	return b.len
}

func (b *bitString) String() string {
	vals := make([]string, len(b.bits))
	for i, v := range b.bits {
		if j := (i + 1) * 8; j <= b.len {
			vals[i] = string(v)
		} else { // last partial byte
			vals[i] = fmt.Sprintf("0x%x", string(v>>(j-b.len)))
		}
	}

	return strings.Join(vals, "")
}

func (b *bitString) BitString() string {
	bits := make([]string, len(b.bits))
	for i, v := range b.bits {
		if j := (i + 1) * 8; j <= b.len {
			bits[i] = fmt.Sprintf("%08b", v)
		} else { // last partial byte
			bits[i] = fmt.Sprintf(
				fmt.Sprintf("%%0%db", b.len-j+8),
				v>>(j-b.len),
			)
		}
	}

	return strings.Join(bits, " ")
}

// Bit returns a given bit by its position (position starts from one).
func (b *bitString) Bit(pos int) bool {
	if pos > b.len {
		return false
	}

	i := pos - 1
	var mask byte = 0x80 >> (i % 8)
	return b.bits[i/8]&mask != 0
}

// DiffPos compares two bitstrings and returns the position of the leftmost bit at which they differ.
func (b *bitString) DiffPos(c *bitString) int {
	var i, j int
	var x, y byte

	for x == y {
		// The bitstrings are the same and there is no difference
		if i >= len(b.bits) && i >= len(c.bits) {
			return 0
		}

		if i < len(b.bits) {
			x = b.bits[i]
		} else {
			x = 0
		}

		if i < len(c.bits) {
			y = c.bits[i]
		} else {
			y = 0
		}

		i++
	}

	for xor := x ^ y; xor != 0; xor >>= 1 {
		j++
	}

	return i*8 - j + 1
}

// Equals determines whether or not two bitstrings are equal.
func (b *bitString) Equals(c *bitString) bool {
	if b.len != c.len {
		return false
	}

	for i := range b.bits {
		if b.bits[i] != c.bits[i] {
			return false
		}
	}

	return true
}

// Sub returns a sub-bitstring by start and end positions.
func (b *bitString) Sub(start, end int) *bitString {
	if start < 1 || end > b.len || start > end {
		return empty
	}

	var mask, bb byte
	n := 0
	sub := &bitString{
		len: end - start + 1,
	}

	for i := start - 1; i < end; i++ {
		if mask = 0x80 >> (i % 8); b.bits[i/8]&mask == 0 {
			bb = bb << 1
		} else {
			bb = (bb << 1) + 1
		}

		if n++; n == 8 {
			sub.bits = append(sub.bits, bb)
			bb, n = 0, 0
		}
	}

	if n > 0 {
		for n < 8 { // Pad with zero
			bb = bb << 1
			n++
		}
		sub.bits = append(sub.bits, bb)
	}

	return sub
}

// Concat concatenates two bitstrings and returns a new bitstring.
func (b *bitString) Concat(c *bitString) *bitString {
	var mask, bb byte
	n := 0
	new := &bitString{
		len: b.Len() + c.Len(),
	}

	for i := 0; i < b.Len(); i++ {
		if mask = 0x80 >> (i % 8); b.bits[i/8]&mask == 0 {
			bb = bb << 1
		} else {
			bb = (bb << 1) + 1
		}

		if n++; n == 8 {
			new.bits = append(new.bits, bb)
			bb, n = 0, 0
		}
	}

	for i := 0; i < c.Len(); i++ {
		if mask = 0x80 >> (i % 8); c.bits[i/8]&mask == 0 {
			bb = bb << 1
		} else {
			bb = (bb << 1) + 1
		}

		if n++; n == 8 {
			new.bits = append(new.bits, bb)
			bb, n = 0, 0
		}
	}

	if n > 0 {
		for ; n < 8; n++ { // Pad with zero
			bb = bb << 1
		}
		new.bits = append(new.bits, bb)
	}

	return new
}

// HasPrefix determines whether or not a bitstring is a prefix of another bitstring.
func (b *bitString) HasPrefix(c *bitString) bool {
	for i, y := range c.bits {
		var x byte = 0
		if i < len(b.bits) {
			x = b.bits[i]
		}

		var mask byte = 0xFF
		if j := (i + 1) * 8; j > c.len { // last byte
			mask = mask << (j - c.len)
		}

		if x&mask != y&mask {
			return false
		}
	}

	return true
}
