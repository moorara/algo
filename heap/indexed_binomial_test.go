package heap

import "testing"

func getIndexedBinomialTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Binomial Min Heap"
	tests[0].expectedGraphviz = ``

	tests[1].heap = "Indexed Binomial Max Heap"
	tests[1].expectedGraphviz = ``

	return tests
}

func TestIndexedBinomialHeap(t *testing.T) {
	tests := getIndexedBinomialTests()

	for _, tc := range tests {
		heap := NewIndexedBinomial[int, string](tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
