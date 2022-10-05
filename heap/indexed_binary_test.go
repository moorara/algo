package heap

import "testing"

func getIndexedBinaryTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Binary Min Heap"
	tests[0].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Indexed Binary Max Heap"
	tests[1].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Indexed Binary Min Heap"
	tests[2].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 10 | ten } }"];
  0 [label="{ 0 | { 30 | thirty } }"];
  2 [label="{ 2 | { 20 | twenty } }"];

  1 -> 0 [];
  1 -> 2 [];
}`

	tests[3].heap = "Indexed Binary Max Heap"
	tests[3].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 30 | thirty } }"];
  2 [label="{ 2 | { 20 | twenty } }"];
  0 [label="{ 0 | { 10 | ten } }"];

  1 -> 2 [];
  1 -> 0 [];
}`

	tests[4].heap = "Indexed Binary Min Heap"
	tests[4].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  3 [label="{ 3 | { 10 | ten } }"];
  4 [label="{ 4 | { 20 | twenty } }"];
  1 [label="{ 1 | { 30 | thirty } }"];
  0 [label="{ 0 | { 50 | fifty } }"];
  2 [label="{ 2 | { 40 | forty } }"];

  3 -> 4 [];
  3 -> 1 [];
  4 -> 0 [];
  4 -> 2 [];
}`

	tests[5].heap = "Indexed Binary Max Heap"
	tests[5].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  3 [label="{ 3 | { 50 | fifty } }"];
  4 [label="{ 4 | { 40 | forty } }"];
  2 [label="{ 2 | { 20 | twenty } }"];
  0 [label="{ 0 | { 10 | ten } }"];
  1 [label="{ 1 | { 30 | thirty } }"];

  3 -> 4 [];
  3 -> 2 [];
  4 -> 0 [];
  4 -> 1 [];
}`

	tests[6].heap = "Indexed Binary Min Heap"
	tests[6].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  7 [label="{ 7 | { 10 | ten } }"];
  8 [label="{ 8 | { 20 | twenty } }"];
  4 [label="{ 4 | { 50 | fifty } }"];
  6 [label="{ 6 | { 30 | thirty } }"];
  2 [label="{ 2 | { 70 | seventy } }"];
  1 [label="{ 1 | { 80 | eighty } }"];
  5 [label="{ 5 | { 60 | sixty } }"];
  0 [label="{ 0 | { 90 | ninety } }"];
  3 [label="{ 3 | { 40 | forty } }"];

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
	tests[7].expectedGraphviz = `strict digraph "Indexed Binary Heap" {
  concentrate=false;
  node [shape=Mrecord];

  7 [label="{ 7 | { 90 | ninety } }"];
  8 [label="{ 8 | { 80 | eighty } }"];
  5 [label="{ 5 | { 60 | sixty } }"];
  6 [label="{ 6 | { 70 | seventy } }"];
  4 [label="{ 4 | { 40 | forty } }"];
  3 [label="{ 3 | { 50 | fifty } }"];
  2 [label="{ 2 | { 20 | twenty } }"];
  0 [label="{ 0 | { 10 | ten } }"];
  1 [label="{ 1 | { 30 | thirty } }"];

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
		heap := NewIndexedBinary[int, string](tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
