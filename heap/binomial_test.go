package heap

import "testing"

func getBinomialTests() []mergeableHeapTest[int, string] {
	tests := getMergeableHeapTest()

	tests[0].heap = "Binomial Min Heap"
	tests[0].expectedGraphviz = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Binomial Max Heap"
	tests[1].expectedGraphviz = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	return tests
}

func TestBinomialHeap(t *testing.T) {
	tests := getBinomialTests()

	for _, tc := range tests {
		heap := NewBinomial[int, string](tc.cmpKey, tc.eqVal)
		runMergeableHeapTest(t, heap, tc)
	}
}
