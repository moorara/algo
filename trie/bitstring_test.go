package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitString(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{
			name: "OK",
			s:    "Alice",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := newBitString(tc.s)
			assert.NotNil(t, b)
		})
	}
}

func TestBitString_Len(t *testing.T) {
	tests := []struct {
		name        string
		b           *bitString
		expectedLen int
	}{
		{
			name:        "OK",
			b:           newBitString("Alice"),
			expectedLen: 40,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedLen, tc.b.Len())
		})
	}
}

func TestBitString_String(t *testing.T) {
	tests := []struct {
		name           string
		b              *bitString
		expectedString string
	}{
		{
			name:           "OK",
			b:              newBitString("Alice"),
			expectedString: "Alice",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.b.String())
		})
	}
}

func TestBitString_BitString(t *testing.T) {
	tests := []struct {
		name              string
		b                 *bitString
		expectedBitString string
	}{
		{
			name:              "OK",
			b:                 newBitString("Alice"),
			expectedBitString: "01000001 01101100 01101001 01100011 01100101",
		},
		{
			name: "Partial",
			b: &bitString{
				bits: []byte{0b01000001, 0b01100000},
				len:  12,
			},
			expectedBitString: "01000001 0110",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedBitString, tc.b.BitString())
		})
	}
}

func TestBitString_Bit(t *testing.T) {
	tests := []struct {
		name         string
		b            *bitString
		expectedBits []int
	}{
		{
			name: "OK",
			b:    newBitString("Alice"),
			expectedBits: []int{
				-1,
				0, 1, 0, 0, 0, 0, 0, 1,
				0, 1, 1, 0, 1, 1, 0, 0,
				0, 1, 1, 0, 1, 0, 0, 1,
				0, 1, 1, 0, 0, 0, 1, 1,
				0, 1, 1, 0, 0, 1, 0, 1,
				0, 0, 0, 0, 0, 0, 0, 0, // Trailing zeros
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for pos, expectedBit := range tc.expectedBits {
				assert.Equal(t, expectedBit, tc.b.Bit(pos))
			}
		})
	}
}

func TestBitString_DiffPos(t *testing.T) {
	tests := []struct {
		name            string
		b, c            *bitString
		expectedDiffPos int
	}{
		{
			name:            "Alice_Alice",
			b:               newBitString("Alice"),
			c:               newBitString("Alice"),
			expectedDiffPos: 0,
		},
		{
			name:            "Alice_Alex",
			b:               newBitString("Alice"),
			c:               newBitString("Alex"),
			expectedDiffPos: 21,
		},
		{
			name:            "Alice_Bob",
			b:               newBitString("Alice"),
			c:               newBitString("Bob"),
			expectedDiffPos: 7,
		},
		{
			name:            "Alice_Charlie",
			b:               newBitString("Alice"),
			c:               newBitString("Charlie"),
			expectedDiffPos: 7,
		},
		{
			name:            "Alice_David",
			b:               newBitString("Alice"),
			c:               newBitString("David"),
			expectedDiffPos: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDiffPos, tc.b.DiffPos(tc.c))
		})
	}
}

func TestBitString_Equals(t *testing.T) {
	tests := []struct {
		name           string
		b, c           *bitString
		expectedEquals bool
	}{
		{
			name:           "Alice_Alice",
			b:              newBitString("Alice"),
			c:              newBitString("Alice"),
			expectedEquals: true,
		},
		{
			name:           "Alice_Alex",
			b:              newBitString("Alice"),
			c:              newBitString("Alex"),
			expectedEquals: false,
		},
		{
			name:           "Alice_Bob",
			b:              newBitString("Alice"),
			c:              newBitString("Bob"),
			expectedEquals: false,
		},
		{
			name:           "Alice_Charlie",
			b:              newBitString("Alice"),
			c:              newBitString("Charlie"),
			expectedEquals: false,
		},
		{
			name:           "Alice_David",
			b:              newBitString("Alice"),
			c:              newBitString("David"),
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.b.Equals(tc.c))
		})
	}
}

func TestBitString_Sub(t *testing.T) {
	tests := []struct {
		name        string
		b           *bitString
		start       int
		end         int
		expectedSub *bitString
	}{
		{
			name:        "First",
			b:           newBitString("Alice"),
			start:       0,
			end:         10,
			expectedSub: nil,
		},
		{
			name:        "Second",
			b:           newBitString("Alice"),
			start:       20,
			end:         10,
			expectedSub: nil,
		},
		{
			name:        "Third",
			b:           newBitString("Alice"),
			start:       1,
			end:         100,
			expectedSub: nil,
		},
		{
			name:  "Fourth",
			b:     newBitString("Alice"),
			start: 1,
			end:   16,
			expectedSub: &bitString{
				bits: []byte{'A', 'l'},
				len:  16,
			},
		},
		{
			name:  "Fifth",
			b:     newBitString("Alice"),
			start: 10,
			end:   30,
			expectedSub: &bitString{
				bits: []byte{0b11011000, 0b11010010, 0b11000000},
				len:  21,
			},
		},
		{
			name:  "Sixth",
			b:     newBitString("Alice"),
			start: 1,
			end:   40,
			expectedSub: &bitString{
				bits: []byte{'A', 'l', 'i', 'c', 'e'},
				len:  40,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSub, tc.b.Sub(tc.start, tc.end))
		})
	}
}

func TestBitString_Concat(t *testing.T) {
	tests := []struct {
		name           string
		b, c           *bitString
		expectedConcat *bitString
	}{
		{
			name: "Alice_Green",
			b:    newBitString("Alice"),
			c:    newBitString("Green"),
			expectedConcat: &bitString{
				bits: []byte{'A', 'l', 'i', 'c', 'e', 'G', 'r', 'e', 'e', 'n'},
				len:  80,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedConcat, tc.b.Concat(tc.c))
		})
	}
}

func TestBitString_HasPrefix(t *testing.T) {
	tests := []struct {
		name           string
		b, c           *bitString
		expectedResult bool
	}{
		{
			name:           "Alice_B",
			b:              newBitString("Alice"),
			c:              newBitString("B"),
			expectedResult: false,
		},
		{
			name:           "Alice_Al",
			b:              newBitString("Alice"),
			c:              newBitString("Al"),
			expectedResult: true,
		},
		{
			name:           "Alice_Alice",
			b:              newBitString("Alice"),
			c:              newBitString("Alice"),
			expectedResult: true,
		},
		{
			name: "PartialPrefix",
			b:    newBitString("Alice"),
			c: &bitString{
				bits: []byte{0b01000001, 0b01100000},
				len:  12,
			},
			expectedResult: true,
		},
		{
			name: "PartialKeyAndPrefix",
			b: &bitString{
				bits: []byte{0b01000001, 0b01101000},
				len:  13,
			},
			c: &bitString{
				bits: []byte{0b01000001, 0b01100000},
				len:  12,
			},
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.b.HasPrefix(tc.c))
		})
	}
}
