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

  1 [label="10 | ten"];
  2 [label="30 | thirty"];
  3 [label="20 | twenty"];

  1 -> 2 [];
  1 -> 3 [];
}`

	tests[3].heap = "Binary Max Heap"
	tests[3].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="30 | thirty"];
  2 [label="10 | ten"];
  3 [label="20 | twenty"];

  1 -> 2 [];
  1 -> 3 [];
}`

	tests[4].heap = "Binary Min Heap"
	tests[4].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | ten"];
  2 [label="20 | twenty"];
  3 [label="40 | forty"];
  4 [label="50 | fifty"];
  5 [label="30 | thirty"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
}`

	tests[5].heap = "Binary Max Heap"
	tests[5].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="50 | fifty"];
  2 [label="40 | forty"];
  3 [label="20 | twenty"];
  4 [label="10 | ten"];
  5 [label="30 | thirty"];

  1 -> 2 [];
  1 -> 3 [];
  2 -> 4 [];
  2 -> 5 [];
}`

	tests[6].heap = "Binary Min Heap"
	tests[6].expectedGraphviz = `strict digraph "Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | ten"];
  2 [label="20 | twenty"];
  3 [label="40 | forty"];
  4 [label="30 | thirty"];
  5 [label="70 | seventy"];
  6 [label="80 | eighty"];
  7 [label="60 | sixty"];
  8 [label="90 | ninety"];
  9 [label="50 | fifty"];

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

  1 [label="90 | ninety"];
  2 [label="80 | eighty"];
  3 [label="60 | sixty"];
  4 [label="70 | seventy"];
  5 [label="30 | thirty"];
  6 [label="20 | twenty"];
  7 [label="50 | fifty"];
  8 [label="10 | ten"];
  9 [label="40 | forty"];

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
		heap := NewBinary[int, string](tc.size, tc.cmpKey, tc.eqVal)
		runHeapTest(t, heap, tc)
	}
}
