package radixsort

import "math/bits"

// MSDString is the MSD (most significant digit) sorting algorithm for string keys with variable length.
func MSDString(a []string) {
	n := len(a)
	aux := make([]string, n)
	msdString(a, aux, 0, n-1, 0)
}

func msdString(a, aux []string, lo, hi, d int) {
	const (
		CUTOFF = 15
		R      = 256
	)

	// cutoff to insertion sort for small subarrays
	if hi <= lo+CUTOFF {
		insertion(a, lo, hi)
		return
	}

	count := make([]int, R+2)

	// compute frequency counts
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		count[c+2]++
	}

	// transform counts to indicies
	for r := 0; r < R+1; r++ {
		count[r+1] += count[r]
	}

	// distribute keys to aux
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back aux to a
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < R; r++ {
		msdString(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

// MSDInt is the MSD (most significant digit) sorting algorithm for integer numbers (signed).
func MSDInt(a []int) {
	n := len(a)
	aux := make([]int, n)
	msdInt(a, aux, 0, n-1, 0)
}

func msdInt(a, aux []int, lo, hi, d int) {
	const (
		CutOff   = 15
		ByteSize = 8
		IntSize  = bits.UintSize
		W        = IntSize / ByteSize
		R        = 1 << ByteSize
		Mask     = R - 1
	)

	// cutoff to insertion sort for small subarrays
	if hi <= lo+CutOff {
		insertion[int](a, lo, hi)
		return
	}

	count := make([]int, R+1)
	shift := IntSize - ByteSize - ByteSize*d

	// compute frequency counts
	for i := lo; i <= hi; i++ {
		c := (a[i] >> shift) & Mask
		count[c+1]++
	}

	// transform counts to indicies
	for r := 0; r < R; r++ {
		count[r+1] += count[r]
	}

	// for most significant byte, 0x80-0xFF comes before 0x00-0x7F
	if d == 0 {
		shift1 := count[R] - count[R/2]
		shift2 := count[R/2]

		// to simplify recursive calls later
		count[R] = shift1 + count[1]

		for r := 0; r < R/2; r++ {
			count[r] += shift1
		}

		for r := R / 2; r < R; r++ {
			count[r] -= shift2
		}
	}

	// distribute keys to aux
	for i := lo; i <= hi; i++ {
		c := (a[i] >> shift) & Mask
		aux[count[c]] = a[i]
		count[c]++
	}

	// copy back aux to a
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// no more bits
	if d == W-1 {
		return
	}

	// special case for most significant byte
	if d == 0 && count[R/2] > 0 {
		msdInt(a, aux, lo, lo+count[R/2]-1, d+1)
	}

	// special case for other bytes
	if d != 0 && count[0] > 0 {
		msdInt(a, aux, lo, lo+count[0]-1, d+1)
	}

	// recursively sort for each digit (could skip r = R/2 for d = 0 and skip r = R for d > 0)
	for r := 0; r < R; r++ {
		if count[r+1] > count[r] {
			msdInt(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
		}
	}
}

// MSDUint is the MSD (most significant digit) sorting algorithm for integer numbers (unsigned).
func MSDUint(a []uint) {
	n := len(a)
	aux := make([]uint, n)
	msdUint(a, aux, 0, n-1, 0)
}

func msdUint(a, aux []uint, lo, hi, d int) {
	const (
		CutOff   = 15
		ByteSize = 8
		IntSize  = bits.UintSize
		W        = IntSize / ByteSize
		R        = 1 << ByteSize
		Mask     = R - 1
	)

	// cutoff to insertion sort for small subarrays
	if hi <= lo+CutOff {
		insertion(a, lo, hi)
		return
	}

	count := make([]int, R+1)
	shift := IntSize - ByteSize - ByteSize*d

	// compute frequency counts
	for i := lo; i <= hi; i++ {
		c := (a[i] >> shift) & Mask
		count[c+1]++
	}

	// transform counts to indicies
	for r := 0; r < R; r++ {
		count[r+1] += count[r]
	}

	// distribute keys to aux
	for i := lo; i <= hi; i++ {
		c := (a[i] >> shift) & Mask
		aux[count[c]] = a[i]
		count[c]++
	}

	// copy back aux to a
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// no more bits
	if d == W-1 {
		return
	}

	// special case for most significant byte
	if d == 0 && count[R/2] > 0 {
		msdUint(a, aux, lo, lo+count[R/2]-1, d+1)
	}

	// special case for other bytes
	if d != 0 && count[0] > 0 {
		msdUint(a, aux, lo, lo+count[0]-1, d+1)
	}

	// recursively sort for each digit (could skip r = R/2 for d = 0 and skip r = R for d > 0)
	for r := 0; r < R; r++ {
		if count[r+1] > count[r] {
			msdUint(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
		}
	}
}
