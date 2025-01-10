package heap

import "testing"

func getIndexedFibonacciTests() []indexedHeapTest[int, string] {
	tests := getIndexedHeapTests()

	tests[0].heap = "Indexed Fibonacci Min Heap"
	tests[0].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Indexed Fibonacci Max Heap"
	tests[1].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Indexed Fibonacci Min Heap"
	tests[2].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 10 | Task#1 } }", color=limegreen, style=filled];
  2 [label="{ 0 | { 30 | Task#3 } }"];
  3 [label="{ 2 | { 20 | Task#2 } }"];

  1 -> 2 [color=red, style=dashed];
  1 -> 3 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 1 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
}`

	tests[3].heap = "Indexed Fibonacci Max Heap"
	tests[3].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 1 | { 30 | Task#3 } }", color=limegreen, style=filled];
  2 [label="{ 0 | { 10 | Task#1 } }"];
  3 [label="{ 2 | { 20 | Task#2 } }"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  1 -> 3 [color=orange, style=dashed];
  2 -> 2 [color=red, style=dashed];
  2 -> 2 [color=orange, style=dashed];
  3 -> 1 [color=red, style=dashed];
  3 -> 1 [color=orange, style=dashed];
}`

	tests[4].heap = "Indexed Fibonacci Min Heap"
	tests[4].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 3 | { 10 | Task#1 } }", color=limegreen, style=filled];
  2 [label="{ 1 | { 30 | Task#3 } }"];
  3 [label="{ 0 | { 50 | Task#5 } }"];
  4 [label="{ 4 | { 20 | Task#2 } }"];
  5 [label="{ 2 | { 40 | Task#4 } }"];

  1 -> 2 [color=blue];
  1 -> 5 [color=red, style=dashed];
  1 -> 5 [color=orange, style=dashed];
  2 -> 3 [color=blue];
  2 -> 4 [color=red, style=dashed];
  2 -> 4 [color=orange, style=dashed];
  3 -> 3 [color=red, style=dashed];
  3 -> 3 [color=orange, style=dashed];
  4 -> 2 [color=red, style=dashed];
  4 -> 2 [color=orange, style=dashed];
  5 -> 1 [color=red, style=dashed];
  5 -> 1 [color=orange, style=dashed];
}`

	tests[5].heap = "Indexed Fibonacci Max Heap"
	tests[5].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 3 | { 50 | Task#5 } }", color=limegreen, style=filled];
  2 [label="{ 2 | { 20 | Task#2 } }"];
  3 [label="{ 0 | { 10 | Task#1 } }"];
  4 [label="{ 1 | { 30 | Task#3 } }"];
  5 [label="{ 4 | { 40 | Task#4 } }"];

  1 -> 2 [color=blue];
  1 -> 5 [color=red, style=dashed];
  1 -> 5 [color=orange, style=dashed];
  2 -> 3 [color=blue];
  2 -> 4 [color=red, style=dashed];
  2 -> 4 [color=orange, style=dashed];
  3 -> 3 [color=red, style=dashed];
  3 -> 3 [color=orange, style=dashed];
  4 -> 2 [color=red, style=dashed];
  4 -> 2 [color=orange, style=dashed];
  5 -> 1 [color=red, style=dashed];
  5 -> 1 [color=orange, style=dashed];
}`

	tests[6].heap = "Indexed Fibonacci Min Heap"
	tests[6].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 7 | { 10 | Task#1 } }", color=limegreen, style=filled];
  2 [label="{ 8 | { 20 | Task#2 } }"];
  3 [label="{ 6 | { 30 | Task#3 } }"];
  4 [label="{ 3 | { 40 | Task#4 } }"];
  5 [label="{ 1 | { 80 | Task#8 } }"];
  6 [label="{ 0 | { 90 | Task#9 } }"];
  7 [label="{ 2 | { 70 | Task#7 } }"];
  8 [label="{ 4 | { 50 | Task#5 } }"];
  9 [label="{ 5 | { 60 | Task#6 } }"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  1 -> 9 [color=orange, style=dashed];
  2 -> 2 [color=red, style=dashed];
  2 -> 2 [color=orange, style=dashed];
  3 -> 4 [color=blue];
  3 -> 9 [color=red, style=dashed];
  3 -> 1 [color=orange, style=dashed];
  4 -> 5 [color=blue];
  4 -> 8 [color=red, style=dashed];
  4 -> 8 [color=orange, style=dashed];
  5 -> 6 [color=blue];
  5 -> 7 [color=red, style=dashed];
  5 -> 7 [color=orange, style=dashed];
  6 -> 6 [color=red, style=dashed];
  6 -> 6 [color=orange, style=dashed];
  7 -> 5 [color=red, style=dashed];
  7 -> 5 [color=orange, style=dashed];
  8 -> 4 [color=red, style=dashed];
  8 -> 4 [color=orange, style=dashed];
  9 -> 1 [color=red, style=dashed];
  9 -> 3 [color=orange, style=dashed];
}`

	tests[7].heap = "Indexed Fibonacci Max Heap"
	tests[7].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ 7 | { 90 | Task#9 } }", color=limegreen, style=filled];
  2 [label="{ 6 | { 70 | Task#7 } }"];
  3 [label="{ 1 | { 30 | Task#3 } }"];
  4 [label="{ 0 | { 10 | Task#1 } }"];
  5 [label="{ 3 | { 50 | Task#5 } }"];
  6 [label="{ 4 | { 40 | Task#4 } }"];
  7 [label="{ 2 | { 20 | Task#2 } }"];
  8 [label="{ 5 | { 60 | Task#6 } }"];
  9 [label="{ 8 | { 80 | Task#8 } }"];

  1 -> 2 [color=blue];
  1 -> 9 [color=red, style=dashed];
  1 -> 9 [color=orange, style=dashed];
  2 -> 3 [color=blue];
  2 -> 6 [color=red, style=dashed];
  2 -> 8 [color=orange, style=dashed];
  3 -> 4 [color=blue];
  3 -> 5 [color=red, style=dashed];
  3 -> 5 [color=orange, style=dashed];
  4 -> 4 [color=red, style=dashed];
  4 -> 4 [color=orange, style=dashed];
  5 -> 3 [color=red, style=dashed];
  5 -> 3 [color=orange, style=dashed];
  6 -> 7 [color=blue];
  6 -> 8 [color=red, style=dashed];
  6 -> 2 [color=orange, style=dashed];
  7 -> 7 [color=red, style=dashed];
  7 -> 7 [color=orange, style=dashed];
  8 -> 2 [color=red, style=dashed];
  8 -> 6 [color=orange, style=dashed];
  9 -> 1 [color=red, style=dashed];
  9 -> 1 [color=orange, style=dashed];
}`

	return tests
}

func TestIndexedFibonacciHeap(t *testing.T) {
	tests := getIndexedFibonacciTests()

	for _, tc := range tests {
		heap := NewIndexedFibonacci(tc.cap, tc.cmpKey, tc.eqVal)
		runIndexedHeapTest(t, heap, tc)
	}
}
