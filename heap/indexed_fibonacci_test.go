package heap

import "testing"

func getIndexedFibonacciTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Fibonacci Min Heap"
	tests[0].expectedGraphviz = ``

	tests[1].heap = "Indexed Fibonacci Max Heap"
	tests[1].expectedGraphviz = ``

	return tests
}

func TestIndexedFibonacciHeap(t *testing.T) {
	tests := getIndexedFibonacciTests()

	for _, tc := range tests {
		heap := NewIndexedFibonacci[int, string](tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
