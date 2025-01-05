package heap

import "testing"

func getIndexedBinomialTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Binomial Min Heap"
	tests[0].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Indexed Binomial Max Heap"
	tests[1].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Indexed Binomial Min Heap"
	tests[2].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  2 [label="{ 2 | { 20 | Task#2 } }"];
  1 [label="{ 1 | { 10 | Task#1 } }"];
  0 [label="{ 0 | { 30 | Task#3 } }"];

  2 -> 1 [color=red, style=dashed];
  1 -> 0 [color=blue];
  0 -> 1 [color=turquoise, style=dashed];
}`

	tests[3].heap = "Indexed Binomial Max Heap"
	tests[3].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  2 [label="{ 2 | { 20 | Task#2 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];

  2 -> 1 [color=red, style=dashed];
  1 -> 0 [color=blue];
  0 -> 1 [color=turquoise, style=dashed];
}`

	tests[4].heap = "Indexed Binomial Min Heap"
	tests[4].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  4 [label="{ 4 | { 20 | Task#2 } }"];
  3 [label="{ 3 | { 10 | Task#1 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];
  0 [label="{ 0 | { 50 | Task#5 } }"];
  2 [label="{ 2 | { 40 | Task#4 } }"];

  4 -> 3 [color=red, style=dashed];
  3 -> 1 [color=blue];
  1 -> 3 [color=turquoise, style=dashed];
  1 -> 0 [color=blue];
  1 -> 2 [color=red, style=dashed];
  0 -> 1 [color=turquoise, style=dashed];
  2 -> 3 [color=turquoise, style=dashed];
}`

	tests[5].heap = "Indexed Binomial Max Heap"
	tests[5].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  4 [label="{ 4 | { 40 | Task#4 } }"];
  3 [label="{ 3 | { 50 | Task#5 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];

  4 -> 3 [color=red, style=dashed];
  3 -> 1 [color=blue];
  1 -> 3 [color=turquoise, style=dashed];
  1 -> 0 [color=blue];
  1 -> 2 [color=red, style=dashed];
  0 -> 1 [color=turquoise, style=dashed];
  2 -> 3 [color=turquoise, style=dashed];
}`

	tests[6].heap = "Indexed Binomial Min Heap"
	tests[6].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  8 [label="{ 8 | { 20 | Task#2 } }"];
  7 [label="{ 7 | { 10 | Task#1 } }"];
  3 [label="{ 3 | { 40 | Task#4 } }"];
  1 [label="{ 1 | { 80 | Task#8 } }"];
  0 [label="{ 0 | { 90 | Task#9 } }"];
  2 [label="{ 2 | { 70 | Task#7 } }"];
  6 [label="{ 6 | { 30 | Task#3 } }"];
  5 [label="{ 5 | { 60 | Task#6 } }"];
  4 [label="{ 4 | { 50 | Task#5 } }"];

  8 -> 7 [color=red, style=dashed];
  7 -> 3 [color=blue];
  3 -> 7 [color=turquoise, style=dashed];
  3 -> 1 [color=blue];
  3 -> 6 [color=red, style=dashed];
  1 -> 3 [color=turquoise, style=dashed];
  1 -> 0 [color=blue];
  1 -> 2 [color=red, style=dashed];
  0 -> 1 [color=turquoise, style=dashed];
  2 -> 3 [color=turquoise, style=dashed];
  6 -> 7 [color=turquoise, style=dashed];
  6 -> 5 [color=blue];
  6 -> 4 [color=red, style=dashed];
  5 -> 6 [color=turquoise, style=dashed];
  4 -> 7 [color=turquoise, style=dashed];
}`

	tests[7].heap = "Indexed Binomial Max Heap"
	tests[7].expectedGraphviz = `strict digraph "Indexed Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  8 [label="{ 8 | { 80 | Task#8 } }"];
  7 [label="{ 7 | { 90 | Task#9 } }"];
  3 [label="{ 3 | { 50 | Task#5 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];
  5 [label="{ 5 | { 60 | Task#6 } }"];
  4 [label="{ 4 | { 40 | Task#4 } }"];
  6 [label="{ 6 | { 70 | Task#7 } }"];

  8 -> 7 [color=red, style=dashed];
  7 -> 3 [color=blue];
  3 -> 7 [color=turquoise, style=dashed];
  3 -> 1 [color=blue];
  3 -> 5 [color=red, style=dashed];
  1 -> 3 [color=turquoise, style=dashed];
  1 -> 0 [color=blue];
  1 -> 2 [color=red, style=dashed];
  0 -> 1 [color=turquoise, style=dashed];
  2 -> 3 [color=turquoise, style=dashed];
  5 -> 7 [color=turquoise, style=dashed];
  5 -> 4 [color=blue];
  5 -> 6 [color=red, style=dashed];
  4 -> 5 [color=turquoise, style=dashed];
  6 -> 7 [color=turquoise, style=dashed];
}`

	return tests
}

func TestIndexedBinomialHeap(t *testing.T) {
	tests := getIndexedBinomialTests()

	for _, tc := range tests {
		heap := NewIndexedBinomial(tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
