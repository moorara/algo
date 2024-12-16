package heap

import "testing"

func getFibonacciTests() []mergeableHeapTest[int, string] {
	tests := getMergeableHeapTest()

	tests[0].heap = "Fibonacci Min Heap"
	tests[0].expectedGraphviz = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Fibonacci Max Heap"
	tests[1].expectedGraphviz = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	return tests
}

func TestFibonacciHeap(t *testing.T) {
	tests := getFibonacciTests()

	for _, tc := range tests {
		heap := NewFibonacci[int, string](tc.cmpKey, tc.eqVal)
		runMergeableHeapTest(t, heap, tc)
	}
}
