package heap

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func getFibonacciTests() []mergeableHeapTest[int, string] {
	tests := getMergeableHeapTests()

	tests[0].heap = "Fibonacci Min Heap"
	tests[0].merge = nil
	tests[0].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Fibonacci Max Heap"
	tests[1].merge = nil
	tests[1].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Fibonacci Min Heap"
	tests[2].merge = NewFibonacci(cmpMin, eqVal)
	tests[2].merge.Insert(33, "Task#3.a")
	tests[2].merge.Insert(11, "Task#1.a")
	tests[2].merge.Insert(22, "Task#2.a")
	tests[2].expectedSize += 3
	tests[2].expectedPeek = generic.KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[2].expectedContains = append(tests[2].expectedContains,
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
	)
	tests[2].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 10, Val: "Task#1"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 33, Val: "Task#3.a"},
	}
	tests[2].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1", color=limegreen, style=filled];
  2 [label="30 | Task#3"];
  3 [label="20 | Task#2"];
  4 [label="33 | Task#3.a"];
  5 [label="22 | Task#2.a"];
  6 [label="11 | Task#1.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 6 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 1 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
}`

	tests[3].heap = "Fibonacci Max Heap"
	tests[3].merge = NewFibonacci(cmpMax, eqVal)
	tests[3].merge.Insert(11, "Task#1.a")
	tests[3].merge.Insert(33, "Task#3.a")
	tests[3].merge.Insert(22, "Task#2.a")
	tests[3].expectedSize += 3
	tests[3].expectedPeek = generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"}
	tests[3].expectedContains = append(tests[3].expectedContains,
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[3].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 33, Val: "Task#3.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 10, Val: "Task#1"},
	}
	tests[3].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="33 | Task#3.a", color=limegreen, style=filled];
  2 [label="30 | Task#3"];
  3 [label="10 | Task#1"];
  4 [label="20 | Task#2"];
  5 [label="11 | Task#1.a"];
  6 [label="22 | Task#2.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 6 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 1 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
}`

	tests[4].heap = "Fibonacci Min Heap"
	tests[4].merge = NewFibonacci(cmpMin, eqVal)
	tests[4].merge.Insert(55, "Task#5.a")
	tests[4].merge.Insert(33, "Task#3.a")
	tests[4].merge.Insert(44, "Task#4.a")
	tests[4].merge.Insert(11, "Task#1.a")
	tests[4].merge.Insert(22, "Task#2.a")
	tests[4].expectedSize += 5
	tests[4].expectedPeek = generic.KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[4].expectedContains = append(tests[4].expectedContains,
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		generic.KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		generic.KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
	)
	tests[4].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 10, Val: "Task#1"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 33, Val: "Task#3.a"},
		{Key: 40, Val: "Task#4"},
		{Key: 44, Val: "Task#4.a"},
		{Key: 50, Val: "Task#5"},
		{Key: 55, Val: "Task#5.a"},
	}
	tests[4].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1", color=limegreen, style=filled];
  2 [label="30 | Task#3"];
  3 [label="50 | Task#5"];
  4 [label="40 | Task#4"];
  5 [label="20 | Task#2"];
  6 [label="33 | Task#3.a"];
  7 [label="55 | Task#5.a"];
  8 [label="44 | Task#4.a"];
  9 [label="22 | Task#2.a"];
  10 [label="11 | Task#1.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 10 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 7 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
  7 -> 8 [color=red, style=dashed];
  7 -> 6 [color=orange, style=dashed];
  8 -> 9 [color=red, style=dashed];
  8 -> 7 [color=orange, style=dashed];
  9 -> 10 [color=red, style=dashed];
  9 -> 8 [color=orange, style=dashed];
  10 -> 1 [color=red, style=dashed];
  10 -> 9 [color=orange, style=dashed];
}`

	tests[5].heap = "Fibonacci Max Heap"
	tests[5].merge = NewFibonacci(cmpMax, eqVal)
	tests[5].merge.Insert(11, "Task#1.a")
	tests[5].merge.Insert(33, "Task#3.a")
	tests[5].merge.Insert(22, "Task#2.a")
	tests[5].merge.Insert(55, "Task#5.a")
	tests[5].merge.Insert(44, "Task#4.a")
	tests[5].expectedSize += 5
	tests[5].expectedPeek = generic.KeyValue[int, string]{Key: 55, Val: "Task#5.a"}
	tests[5].expectedContains = append(tests[5].expectedContains,
		generic.KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		generic.KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[5].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 55, Val: "Task#5.a"},
		{Key: 50, Val: "Task#5"},
		{Key: 44, Val: "Task#4.a"},
		{Key: 40, Val: "Task#4"},
		{Key: 33, Val: "Task#3.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 10, Val: "Task#1"},
	}
	tests[5].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="55 | Task#5.a", color=limegreen, style=filled];
  2 [label="50 | Task#5"];
  3 [label="30 | Task#3"];
  4 [label="10 | Task#1"];
  5 [label="20 | Task#2"];
  6 [label="40 | Task#4"];
  7 [label="33 | Task#3.a"];
  8 [label="11 | Task#1.a"];
  9 [label="22 | Task#2.a"];
  10 [label="44 | Task#4.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 10 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 7 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
  7 -> 8 [color=red, style=dashed];
  7 -> 6 [color=orange, style=dashed];
  8 -> 9 [color=red, style=dashed];
  8 -> 7 [color=orange, style=dashed];
  9 -> 10 [color=red, style=dashed];
  9 -> 8 [color=orange, style=dashed];
  10 -> 1 [color=red, style=dashed];
  10 -> 9 [color=orange, style=dashed];
}`

	tests[6].heap = "Fibonacci Min Heap"
	tests[6].merge = NewFibonacci(cmpMin, eqVal)
	tests[6].merge.Insert(99, "Task#9.a")
	tests[6].merge.Insert(88, "Task#8.a")
	tests[6].merge.Insert(77, "Task#7.a")
	tests[6].merge.Insert(44, "Task#4.a")
	tests[6].merge.Insert(55, "Task#5.a")
	tests[6].merge.Insert(66, "Task#6.a")
	tests[6].merge.Insert(33, "Task#3.a")
	tests[6].merge.Insert(11, "Task#1.a")
	tests[6].merge.Insert(22, "Task#2.a")
	tests[6].expectedSize += 9
	tests[6].expectedPeek = generic.KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[6].expectedContains = append(tests[6].expectedContains,
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		generic.KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		generic.KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		generic.KeyValue[int, string]{Key: 66, Val: "Task#6.a"},
		generic.KeyValue[int, string]{Key: 77, Val: "Task#7.a"},
		generic.KeyValue[int, string]{Key: 88, Val: "Task#8.a"},
		generic.KeyValue[int, string]{Key: 99, Val: "Task#9.a"},
	)
	tests[6].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 10, Val: "Task#1"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 33, Val: "Task#3.a"},
		{Key: 40, Val: "Task#4"},
		{Key: 44, Val: "Task#4.a"},
		{Key: 50, Val: "Task#5"},
		{Key: 55, Val: "Task#5.a"},
		{Key: 60, Val: "Task#6"},
		{Key: 66, Val: "Task#6.a"},
		{Key: 70, Val: "Task#7"},
		{Key: 77, Val: "Task#7.a"},
		{Key: 80, Val: "Task#8"},
		{Key: 88, Val: "Task#8.a"},
		{Key: 90, Val: "Task#9"},
		{Key: 99, Val: "Task#9.a"},
	}
	tests[6].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="10 | Task#1", color=limegreen, style=filled];
  2 [label="30 | Task#3"];
  3 [label="40 | Task#4"];
  4 [label="70 | Task#7"];
  5 [label="80 | Task#8"];
  6 [label="90 | Task#9"];
  7 [label="50 | Task#5"];
  8 [label="60 | Task#6"];
  9 [label="20 | Task#2"];
  10 [label="33 | Task#3.a"];
  11 [label="44 | Task#4.a"];
  12 [label="77 | Task#7.a"];
  13 [label="88 | Task#8.a"];
  14 [label="99 | Task#9.a"];
  15 [label="55 | Task#5.a"];
  16 [label="66 | Task#6.a"];
  17 [label="22 | Task#2.a"];
  18 [label="11 | Task#1.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 18 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 7 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
  7 -> 8 [color=red, style=dashed];
  7 -> 6 [color=orange, style=dashed];
  8 -> 9 [color=red, style=dashed];
  8 -> 7 [color=orange, style=dashed];
  9 -> 10 [color=red, style=dashed];
  9 -> 8 [color=orange, style=dashed];
  10 -> 11 [color=red, style=dashed];
  10 -> 9 [color=orange, style=dashed];
  11 -> 12 [color=red, style=dashed];
  11 -> 10 [color=orange, style=dashed];
  12 -> 13 [color=red, style=dashed];
  12 -> 11 [color=orange, style=dashed];
  13 -> 14 [color=red, style=dashed];
  13 -> 12 [color=orange, style=dashed];
  14 -> 15 [color=red, style=dashed];
  14 -> 13 [color=orange, style=dashed];
  15 -> 16 [color=red, style=dashed];
  15 -> 14 [color=orange, style=dashed];
  16 -> 17 [color=red, style=dashed];
  16 -> 15 [color=orange, style=dashed];
  17 -> 18 [color=red, style=dashed];
  17 -> 16 [color=orange, style=dashed];
  18 -> 1 [color=red, style=dashed];
  18 -> 17 [color=orange, style=dashed];
}`

	tests[7].heap = "Fibonacci Max Heap"
	tests[7].merge = NewFibonacci(cmpMax, eqVal)
	tests[7].merge.Insert(11, "Task#1.a")
	tests[7].merge.Insert(33, "Task#3.a")
	tests[7].merge.Insert(22, "Task#2.a")
	tests[7].merge.Insert(55, "Task#5.a")
	tests[7].merge.Insert(44, "Task#4.a")
	tests[7].merge.Insert(66, "Task#6.a")
	tests[7].merge.Insert(77, "Task#7.a")
	tests[7].merge.Insert(99, "Task#9.a")
	tests[7].merge.Insert(88, "Task#8.a")
	tests[7].expectedSize += 9
	tests[7].expectedPeek = generic.KeyValue[int, string]{Key: 99, Val: "Task#9.a"}
	tests[7].expectedContains = append(tests[7].expectedContains,
		generic.KeyValue[int, string]{Key: 99, Val: "Task#9.a"},
		generic.KeyValue[int, string]{Key: 88, Val: "Task#8.a"},
		generic.KeyValue[int, string]{Key: 77, Val: "Task#7.a"},
		generic.KeyValue[int, string]{Key: 66, Val: "Task#6.a"},
		generic.KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		generic.KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		generic.KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		generic.KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		generic.KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[7].expectedDelete = []generic.KeyValue[int, string]{
		{Key: 99, Val: "Task#9.a"},
		{Key: 90, Val: "Task#9"},
		{Key: 88, Val: "Task#8.a"},
		{Key: 80, Val: "Task#8"},
		{Key: 77, Val: "Task#7.a"},
		{Key: 70, Val: "Task#7"},
		{Key: 66, Val: "Task#6.a"},
		{Key: 60, Val: "Task#6"},
		{Key: 55, Val: "Task#5.a"},
		{Key: 50, Val: "Task#5"},
		{Key: 44, Val: "Task#4.a"},
		{Key: 40, Val: "Task#4"},
		{Key: 33, Val: "Task#3.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 10, Val: "Task#1"},
	}
	tests[7].expectedDOT = `strict digraph "Fibonacci Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="99 | Task#9.a", color=limegreen, style=filled];
  2 [label="90 | Task#9"];
  3 [label="70 | Task#7"];
  4 [label="60 | Task#6"];
  5 [label="50 | Task#5"];
  6 [label="30 | Task#3"];
  7 [label="10 | Task#1"];
  8 [label="20 | Task#2"];
  9 [label="40 | Task#4"];
  10 [label="80 | Task#8"];
  11 [label="77 | Task#7.a"];
  12 [label="66 | Task#6.a"];
  13 [label="55 | Task#5.a"];
  14 [label="33 | Task#3.a"];
  15 [label="11 | Task#1.a"];
  16 [label="22 | Task#2.a"];
  17 [label="44 | Task#4.a"];
  18 [label="88 | Task#8.a"];

  1 -> 2 [color=red, style=dashed];
  1 -> 18 [color=orange, style=dashed];
  2 -> 3 [color=red, style=dashed];
  2 -> 1 [color=orange, style=dashed];
  3 -> 4 [color=red, style=dashed];
  3 -> 2 [color=orange, style=dashed];
  4 -> 5 [color=red, style=dashed];
  4 -> 3 [color=orange, style=dashed];
  5 -> 6 [color=red, style=dashed];
  5 -> 4 [color=orange, style=dashed];
  6 -> 7 [color=red, style=dashed];
  6 -> 5 [color=orange, style=dashed];
  7 -> 8 [color=red, style=dashed];
  7 -> 6 [color=orange, style=dashed];
  8 -> 9 [color=red, style=dashed];
  8 -> 7 [color=orange, style=dashed];
  9 -> 10 [color=red, style=dashed];
  9 -> 8 [color=orange, style=dashed];
  10 -> 11 [color=red, style=dashed];
  10 -> 9 [color=orange, style=dashed];
  11 -> 12 [color=red, style=dashed];
  11 -> 10 [color=orange, style=dashed];
  12 -> 13 [color=red, style=dashed];
  12 -> 11 [color=orange, style=dashed];
  13 -> 14 [color=red, style=dashed];
  13 -> 12 [color=orange, style=dashed];
  14 -> 15 [color=red, style=dashed];
  14 -> 13 [color=orange, style=dashed];
  15 -> 16 [color=red, style=dashed];
  15 -> 14 [color=orange, style=dashed];
  16 -> 17 [color=red, style=dashed];
  16 -> 15 [color=orange, style=dashed];
  17 -> 18 [color=red, style=dashed];
  17 -> 16 [color=orange, style=dashed];
  18 -> 1 [color=red, style=dashed];
  18 -> 17 [color=orange, style=dashed];
}`

	return tests
}

func TestFibonacciHeap(t *testing.T) {
	tests := getFibonacciTests()

	for _, tc := range tests {
		heap := NewFibonacci(tc.cmpKey, tc.eqVal)
		runMergeableHeapTest(t, heap, tc)
	}
}
