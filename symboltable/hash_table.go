package symboltable

// HashOpts represents configuration options for a hash table.
type HashOpts struct {
	// The initial capacity of the hash table (must be a power of 2 for efficient hashing).
	InitialCap int
	// The minimum load factor before resizing (shrinking) the hash table.
	MinLoadFactor float32
	// The maximum load factor before resizing (expanding) the hash table.
	MaxLoadFactor float32
}

// hashTableEntry represents an entry in a non-linear probing hash table that requires soft deletion.
type hashTableEntry[K, V any] struct {
	key     K
	val     V
	deleted bool
}

// gcd computes the greatest common divisor of two numbers.
// It implements the Euclidean algorithm.
func gcd(a, b uint64) uint64 {
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

// isPowerOf2 determines whether or not a given integer n is a power of 2.
func isPowerOf2(n int) bool {
	return n&(n-1) == 0
}

// isPrime determines whether or not a given integer n is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Check for prime numbers less 100 directly
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

// largestPrimeSmallerThan finds the largest prime number smaller than a given arbitray number.
// If n is less than 2, the function returns -1 (2 is the first prime number).
// If n is prime, the function returns n.
func largestPrimeSmallerThan(n int) int {
	if n < 2 {
		return -1 // No prime number smaller than 2
	}

	for p := n; p >= 2; p-- {
		if isPrime(p) {
			return p
		}
	}

	return -1
}

// smallestPrimeLargerThan finds the smallest prime number larger than a given arbitray number.
// If n is prime, the function returns n.
func smallestPrimeLargerThan(n int) int {
	for p := n; ; p++ {
		if isPrime(p) {
			return p
		}
	}
}
