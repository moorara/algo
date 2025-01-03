package heap

import "testing"

func getBinaryTests() []heapTest[int, string] {
	tests := getHeapTests()

	tests[0].heap = "Binary Min Heap"
	tests[0].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Binary Max Heap"
	tests[1].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Binary Min Heap"
	tests[2].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1"];
  2 [label="30 | Task#3"];
  3 [label="20 | Task#2"];

  1 -> 2 [];
  1 -> 3 [];
}`

	tests[3].heap = "Binary Max Heap"
	tests[3].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="30 | Task#3"];
  2 [label="10 | Task#1"];
  3 [label="20 | Task#2"];

  1 -> 2 [];
  1 -> 3 [];
}`

	tests[4].heap = "Binary Min Heap"
	tests[4].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1"];
  2 [label="20 | Task#2"];
  3 [label="40 | Task#4"];
  4 [label="50 | Task#5"];
  5 [label="30 | Task#3"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
}`

	tests[5].heap = "Binary Max Heap"
	tests[5].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="50 | Task#5"];
  2 [label="40 | Task#4"];
  3 [label="20 | Task#2"];
  4 [label="10 | Task#1"];
  5 [label="30 | Task#3"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
}`

	tests[6].heap = "Binary Min Heap"
	tests[6].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1"];
  2 [label="20 | Task#2"];
  3 [label="40 | Task#4"];
  4 [label="30 | Task#3"];
  5 [label="70 | Task#7"];
  6 [label="80 | Task#8"];
  7 [label="60 | Task#6"];
  8 [label="90 | Task#9"];
  9 [label="50 | Task#5"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
  3 -> 6 [];
  3 -> 7 [];
  4 -> 8 [];
  4 -> 9 [];
}`

	tests[7].heap = "Binary Max Heap"
	tests[7].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="90 | Task#9"];
  2 [label="80 | Task#8"];
  3 [label="60 | Task#6"];
  4 [label="70 | Task#7"];
  5 [label="30 | Task#3"];
  6 [label="20 | Task#2"];
  7 [label="50 | Task#5"];
  8 [label="10 | Task#1"];
  9 [label="40 | Task#4"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
  3 -> 6 [];
  3 -> 7 [];
  4 -> 8 [];
  4 -> 9 [];
}`

	return tests
}

func TestBinaryHeap(t *testing.T) {
	tests := getBinaryTests()

	for _, tc := range tests {
		heap := NewBinary(tc.size, tc.cmpKey, tc.eqVal)
		runHeapTest(t, heap, tc)
	}
}
