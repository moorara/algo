// Package radixsort implements common radix sorting algorithms.
// Radix sorts are key-indexed counting for sorting keys with integer digits between 0 and R-1 (R is a small number).
package radixsort

const r = 256 // uint8 (byte) size
