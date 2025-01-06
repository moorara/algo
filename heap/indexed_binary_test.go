package heap

import "testing"

func getIndexedBinaryTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Binary Min Heap"
	tests[0].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Indexed Binary Max Heap"
	tests[1].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Indexed Binary Min Heap"
	tests[2].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 10 | Task#1 } }"];
  0 [label="{ 0 | { 30 | Task#3 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];

  1 -> 0 [];
  1 -> 2 [];
}`

	tests[3].heap = "Indexed Binary Max Heap"
	tests[3].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 30 | Task#3 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];

  1 -> 2 [];
  1 -> 0 [];
}`

	tests[4].heap = "Indexed Binary Min Heap"
	tests[4].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  3 [label="{ 3 | { 10 | Task#1 } }"];
  4 [label="{ 4 | { 20 | Task#2 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];
  0 [label="{ 0 | { 50 | Task#5 } }"];
  2 [label="{ 2 | { 40 | Task#4 } }"];

  3 -> 4 [];
  3 -> 1 [];
  4 -> 0 [];
  4 -> 2 [];
}`

	tests[5].heap = "Indexed Binary Max Heap"
	tests[5].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  3 [label="{ 3 | { 50 | Task#5 } }"];
  4 [label="{ 4 | { 40 | Task#4 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];

  3 -> 4 [];
  3 -> 2 [];
  4 -> 0 [];
  4 -> 1 [];
}`

	tests[6].heap = "Indexed Binary Min Heap"
	tests[6].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  7 [label="{ 7 | { 10 | Task#1 } }"];
  8 [label="{ 8 | { 20 | Task#2 } }"];
  4 [label="{ 4 | { 50 | Task#5 } }"];
  6 [label="{ 6 | { 30 | Task#3 } }"];
  2 [label="{ 2 | { 70 | Task#7 } }"];
  1 [label="{ 1 | { 80 | Task#8 } }"];
  5 [label="{ 5 | { 60 | Task#6 } }"];
  0 [label="{ 0 | { 90 | Task#9 } }"];
  3 [label="{ 3 | { 40 | Task#4 } }"];

  7 -> 8 [];
  7 -> 4 [];
  8 -> 6 [];
  8 -> 2 [];
  4 -> 1 [];
  4 -> 5 [];
  6 -> 0 [];
  6 -> 3 [];
}`

	tests[7].heap = "Indexed Binary Max Heap"
	tests[7].expectedDOT = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  7 [label="{ 7 | { 90 | Task#9 } }"];
  8 [label="{ 8 | { 80 | Task#8 } }"];
  5 [label="{ 5 | { 60 | Task#6 } }"];
  6 [label="{ 6 | { 70 | Task#7 } }"];
  4 [label="{ 4 | { 40 | Task#4 } }"];
  3 [label="{ 3 | { 50 | Task#5 } }"];
  2 [label="{ 2 | { 20 | Task#2 } }"];
  0 [label="{ 0 | { 10 | Task#1 } }"];
  1 [label="{ 1 | { 30 | Task#3 } }"];

  7 -> 8 [];
  7 -> 5 [];
  8 -> 6 [];
  8 -> 4 [];
  5 -> 3 [];
  5 -> 2 [];
  6 -> 0 [];
  6 -> 1 [];
}`

	return tests
}

func TestIndexedBinaryHeap(t *testing.T) {
	tests := getIndexedBinaryTests()

	for _, tc := range tests {
		heap := NewIndexedBinary(tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
