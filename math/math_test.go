package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name        string
		a, b        uint64
		expectedGCD uint64
	}{
		{
			name:        "GCDOne",
			a:           64,
			b:           61,
			expectedGCD: 1,
		},
		{
			name:        "GCDGreaterThanOne",
			a:           48,
			b:           64,
			expectedGCD: 16,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedGCD, GCD(tc.a, tc.b))
		})
	}
}

func TestPower2(t *testing.T) {
	tests := []struct {
		name           string
		n              int
		expectedResult int
	}{
		{
			name:           "2⁶",
			n:              6,
			expectedResult: 64,
		},
		{
			name:           "2¹⁰",
			n:              10,
			expectedResult: 1024,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, Power2(tc.n))
		})
	}
}

func TestIsPowerOf2(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		expectedOK bool
	}{
		{
			name:       "PowerOf2",
			n:          64,
			expectedOK: true,
		},
		{
			name:       "NotPowerOf2",
			n:          69,
			expectedOK: false,
		},
		{
			name:       "Negative",
			n:          -16,
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedOK, IsPowerOf2(tc.n))
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name            string
		n               int
		expectedIsPrime bool
	}{
		{
			name:            "Negative",
			n:               -1,
			expectedIsPrime: false,
		},
		{
			name:            "Zero",
			n:               0,
			expectedIsPrime: false,
		},
		{
			name:            "One",
			n:               1,
			expectedIsPrime: false,
		},
		{
			name:            "PrimeLessThan10",
			n:               7,
			expectedIsPrime: true,
		},
		{
			name:            "NotPrimeLessThan10",
			n:               8,
			expectedIsPrime: false,
		},
		{
			name:            "PrimeLessThan100",
			n:               97,
			expectedIsPrime: true,
		},
		{
			name:            "NotPrimeLessThan100",
			n:               64,
			expectedIsPrime: false,
		},
		{
			name:            "PrimeLessThan1000",
			n:               997,
			expectedIsPrime: true,
		},
		{
			name:            "NotPrimeLessThan1000",
			n:               666,
			expectedIsPrime: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsPrime, IsPrime(tc.n))
		})
	}
}

func TestLargestPrimeSmallerThan(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		expectedPrime int
	}{
		{
			name:          "Negative",
			n:             -1,
			expectedPrime: -1,
		},
		{
			name:          "Zero",
			n:             0,
			expectedPrime: -1,
		},
		{
			name:          "One",
			n:             1,
			expectedPrime: -1,
		},
		{
			name:          "Two",
			n:             2,
			expectedPrime: 2,
		},
		{
			name:          "TwentySeven",
			n:             27,
			expectedPrime: 23,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedPrime, LargestPrimeSmallerThan(tc.n))
		})
	}
}

func TestSmallestPrimeLargerThan(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		expectedPrime int
	}{
		{
			name:          "Negative",
			n:             -1,
			expectedPrime: 2,
		},
		{
			name:          "Zero",
			n:             0,
			expectedPrime: 2,
		},
		{
			name:          "One",
			n:             1,
			expectedPrime: 2,
		},
		{
			name:          "Two",
			n:             2,
			expectedPrime: 2,
		},
		{
			name:          "TwentySeven",
			n:             27,
			expectedPrime: 29,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedPrime, SmallestPrimeLargerThan(tc.n))
		})
	}
}
