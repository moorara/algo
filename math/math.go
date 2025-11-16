// Package math provides utility mathematical functions.
package math

// GCD computes the greatest common divisor of two numbers.
// It implements the Euclidean algorithm.
func GCD(a, b uint64) uint64 {
	// Ensure a ≥ b
	a, b = max(a, b), min(a, b)

	/*
	 * Let the quotient be q and the remainder be r, so that a = b × q + r
	 * Replace a with b and b with r
	 * Repeat this until the remainder r becomes 0
	 * The GCD is the last non-zero remainder
	 */

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Power2 computes 2 raised to the power of n.
func Power2(n int) int {
	return 1 << n
}

// IsPowerOf2 determines whether or not a given integer n is a power of 2.
func IsPowerOf2(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// IsPrime determines whether or not a given integer n is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Check for prime numbers less than 100 directly
	if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 || n == 23 || n == 29 || n == 31 || n == 37 ||
		n == 41 || n == 43 || n == 47 || n == 53 || n == 59 || n == 61 || n == 67 || n == 71 || n == 73 || n == 79 || n == 83 || n == 89 || n == 97 {
		return true
	} else if n <= 100 {
		return false
	}

	// Check if n is prime using trial division
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// LargestPrimeSmallerThan finds the largest prime number equal to or smaller than a given arbitrary number.
// If n is less than 2, the function returns -1 (2 is the first prime number).
// If n is prime, the function returns n.
func LargestPrimeSmallerThan(n int) int {
	if n < 2 {
		return -1 // No prime number smaller than 2
	}

	for p := n; p >= 2; p-- {
		if IsPrime(p) {
			return p
		}
	}

	return -1
}

// SmallestPrimeLargerThan finds the smallest prime number equal to or larger than a given arbitrary number.
// If n is prime, the function returns n.
func SmallestPrimeLargerThan(n int) int {
	for p := n; ; p++ {
		if IsPrime(p) {
			return p
		}
	}
}
