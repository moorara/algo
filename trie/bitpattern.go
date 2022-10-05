package trie

// bitPattern is an extension of bitstring for pattern matching.
type bitPattern struct {
	*bitString
}

func newBitPattern(pattern string) *bitPattern {
	return &bitPattern{
		bitString: newBitString(pattern),
	}
}

// Bit returns a given bit by its position (position starts from one).
func (b *bitPattern) Bit(pos int) byte {
	if pos > b.len {
		return '0'
	}

	i := pos - 1
	var mask byte = 0x80 >> (i % 8)

	if b.bits[i/8] == '*' {
		return '*'
	}

	if b.bits[i/8]&mask == 0 {
		return '0'
	} else {
		return '1'
	}
}
