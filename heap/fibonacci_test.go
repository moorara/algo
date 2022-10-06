package heap

import "testing"

func getFibonacciTests() []heapTest[int, string] {
	tests := getHeapTests()

	tests[0].heap = "Fibonacci Min Heap"
	tests[0].expectedGraphviz = ``

	tests[1].heap = "Fibonacci Max Heap"
	tests[1].expectedGraphviz = ``

	tests[2].heap = "Fibonacci Min Heap"
	tests[2].expectedGraphviz = ``

	tests[3].heap = "Fibonacci Max Heap"
	tests[3].expectedGraphviz = ``

	tests[4].heap = "Fibonacci Min Heap"
	tests[4].expectedGraphviz = ``

	tests[5].heap = "Fibonacci Max Heap"
	tests[5].expectedGraphviz = ``

	tests[6].heap = "Fibonacci Min Heap"
	tests[6].expectedGraphviz = ``

	tests[7].heap = "Fibonacci Max Heap"
	tests[7].expectedGraphviz = ``

	return tests
}

func TestFibonacciHeap(t *testing.T) {
	tests := getFibonacciTests()

	for _, tc := range tests {
		heap := NewFibonacci[int, string](tc.cmpKey, tc.eqVal)
		runHeapTest(t, heap, tc)
	}
}
