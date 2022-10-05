package heap

import "testing"

func getBinomialTests() []heapTest[int, string] {
	tests := getHeapTests()

	tests[0].heap = "Binomial Min Heap"
	tests[0].expectedGraphviz = ``

	tests[1].heap = "Binomial Max Heap"
	tests[1].expectedGraphviz = ``

	tests[2].heap = "Binomial Min Heap"
	tests[2].expectedGraphviz = ``

	tests[3].heap = "Binomial Max Heap"
	tests[3].expectedGraphviz = ``

	tests[4].heap = "Binomial Min Heap"
	tests[4].expectedGraphviz = ``

	tests[5].heap = "Binomial Max Heap"
	tests[5].expectedGraphviz = ``

	tests[6].heap = "Binomial Min Heap"
	tests[6].expectedGraphviz = ``

	tests[7].heap = "Binomial Max Heap"
	tests[7].expectedGraphviz = ``

	return tests
}

func TestBinomialHeap(t *testing.T) {
	tests := getBinomialTests()

	for _, tc := range tests {
		heap := NewBinomial[int, string](tc.cmpKey, tc.eqVal)
		runHeapTest(t, heap, tc)
	}
}
