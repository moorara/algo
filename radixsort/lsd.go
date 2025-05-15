package radixsort

import "math/bits"

// LSDString is the LSD (least significant digit) sorting algorithm for string keys with fixed length w.
func LSDString(a []string, w int) {
	const R = 256

	n := len(a)
	aux := make([]string, n)

	// sort by key-indexed counting on dth char (stable)
	for d := w - 1; d >= 0; d-- {
		count := make([]int, R+1)

		// compute frequency counts
		for _, s := range a {
			count[s[d]+1]++
		}

		// compute cumulative counts
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// distribute keys to aux
		for _, s := range a {
			aux[count[s[d]]] = s
			count[s[d]]++
		}

		// copy back aux to a
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}

// LSDInt is the LSD (least significant digit) sorting algorithm for integer numbers (signed).
func LSDInt(a []int) {
	const (
		ByteSize = 8
		IntSize  = bits.UintSize
		W        = IntSize / ByteSize
		R        = 1 << ByteSize
		Mask     = R - 1
	)

	n := len(a)
	aux := make([]int, n)

	// sort by key-indexed counting on dth char (stable)
	for d := 0; d < W; d++ {
		count := make([]int, R+1)
		shift := ByteSize * d

		// compute frequency counts
		for _, v := range a {
			c := (v >> shift) & Mask
			count[c+1]++
		}

		// compute cumulative counts
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// for most significant byte, 0x80-0xFF comes before 0x00-0x7F
		if d == W-1 {
			shift1 := count[R] - count[R/2]
			shift2 := count[R/2]

			for r := 0; r < R/2; r++ {
				count[r] += shift1
			}

			for r := R / 2; r < R; r++ {
				count[r] -= shift2
			}
		}

		// distribute keys to aux
		for _, v := range a {
			c := (v >> shift) & Mask
			aux[count[c]] = v
			count[c]++
		}

		// copy back aux to a
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}

// LSDUint is the LSD (least significant digit) sorting algorithm for integer numbers (unsigned).
func LSDUint(a []uint) {
	const (
		ByteSize = 8
		UintSize = bits.UintSize
		W        = UintSize / ByteSize
		R        = 1 << ByteSize
		Mask     = R - 1
	)

	n := len(a)
	aux := make([]uint, n)

	// sort by key-indexed counting on dth char (stable)
	for d := 0; d < W; d++ {
		count := make([]int, R+1)
		shift := ByteSize * d

		// compute frequency counts
		for _, v := range a {
			c := (v >> shift) & Mask
			count[c+1]++
		}

		// compute cumulative counts
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// distribute keys to aux
		for _, v := range a {
			c := (v >> shift) & Mask
			aux[count[c]] = v
			count[c]++
		}

		// copy back aux to a
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}
