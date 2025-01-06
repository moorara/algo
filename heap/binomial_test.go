package heap

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getBinomialTests() []mergeableHeapTest[int, string] {
	tests := getMergeableHeapTests()

	tests[0].heap = "Binomial Min Heap"
	tests[0].merge = nil
	tests[0].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[1].heap = "Binomial Max Heap"
	tests[1].merge = nil
	tests[1].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];
}`

	tests[2].heap = "Binomial Min Heap"
	tests[2].merge = NewBinomial(cmpMin, eqVal)
	tests[2].merge.Insert(33, "Task#3.a")
	tests[2].merge.Insert(11, "Task#1.a")
	tests[2].merge.Insert(22, "Task#2.a")
	tests[2].expectedSize += 3
	tests[2].expectedPeek = KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[2].expectedContains = append(tests[2].expectedContains,
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
	)
	tests[2].expectedDelete = []KeyValue[int, string]{
		{Key: 10, Val: "Task#1"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 33, Val: "Task#3.a"},
	}
	tests[2].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="20 | Task#2"];
  2 [label="22 | Task#2.a"];
  3 [label="10 | Task#1"];
  4 [label="11 | Task#1.a"];
  5 [label="33 | Task#3.a"];
  6 [label="30 | Task#3"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 6 [color=red, style=dashed];
}`

	tests[3].heap = "Binomial Max Heap"
	tests[3].merge = NewBinomial(cmpMax, eqVal)
	tests[3].merge.Insert(11, "Task#1.a")
	tests[3].merge.Insert(33, "Task#3.a")
	tests[3].merge.Insert(22, "Task#2.a")
	tests[3].expectedSize += 3
	tests[3].expectedPeek = KeyValue[int, string]{Key: 33, Val: "Task#3.a"}
	tests[3].expectedContains = append(tests[3].expectedContains,
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[3].expectedDelete = []KeyValue[int, string]{
		{Key: 33, Val: "Task#3.a"},
		{Key: 30, Val: "Task#3"},
		{Key: 22, Val: "Task#2.a"},
		{Key: 20, Val: "Task#2"},
		{Key: 11, Val: "Task#1.a"},
		{Key: 10, Val: "Task#1"},
	}
	tests[3].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="22 | Task#2.a"];
  2 [label="20 | Task#2"];
  3 [label="33 | Task#3.a"];
  4 [label="30 | Task#3"];
  5 [label="10 | Task#1"];
  6 [label="11 | Task#1.a"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 6 [color=red, style=dashed];
}`

	tests[4].heap = "Binomial Min Heap"
	tests[4].merge = NewBinomial(cmpMin, eqVal)
	tests[4].merge.Insert(55, "Task#5.a")
	tests[4].merge.Insert(33, "Task#3.a")
	tests[4].merge.Insert(44, "Task#4.a")
	tests[4].merge.Insert(11, "Task#1.a")
	tests[4].merge.Insert(22, "Task#2.a")
	tests[4].expectedSize += 5
	tests[4].expectedPeek = KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[4].expectedContains = append(tests[4].expectedContains,
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
	)
	tests[4].expectedDelete = []KeyValue[int, string]{
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
	tests[4].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="20 | Task#2"];
  2 [label="22 | Task#2.a"];
  3 [label="10 | Task#1"];
  4 [label="11 | Task#1.a"];
  5 [label="33 | Task#3.a"];
  6 [label="55 | Task#5.a"];
  7 [label="44 | Task#4.a"];
  8 [label="30 | Task#3"];
  9 [label="50 | Task#5"];
  10 [label="40 | Task#4"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 8 [color=red, style=dashed];
  5 -> 6 [color=blue];
  5 -> 7 [color=red, style=dashed];
  8 -> 9 [color=blue];
  8 -> 10 [color=red, style=dashed];
}`

	tests[5].heap = "Binomial Max Heap"
	tests[5].merge = NewBinomial(cmpMax, eqVal)
	tests[5].merge.Insert(11, "Task#1.a")
	tests[5].merge.Insert(33, "Task#3.a")
	tests[5].merge.Insert(22, "Task#2.a")
	tests[5].merge.Insert(55, "Task#5.a")
	tests[5].merge.Insert(44, "Task#4.a")
	tests[5].expectedSize += 5
	tests[5].expectedPeek = KeyValue[int, string]{Key: 55, Val: "Task#5.a"}
	tests[5].expectedContains = append(tests[5].expectedContains,
		KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[5].expectedDelete = []KeyValue[int, string]{
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
	tests[5].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="44 | Task#4.a"];
  2 [label="40 | Task#4"];
  3 [label="55 | Task#5.a"];
  4 [label="50 | Task#5"];
  5 [label="30 | Task#3"];
  6 [label="10 | Task#1"];
  7 [label="20 | Task#2"];
  8 [label="33 | Task#3.a"];
  9 [label="11 | Task#1.a"];
  10 [label="22 | Task#2.a"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 8 [color=red, style=dashed];
  5 -> 6 [color=blue];
  5 -> 7 [color=red, style=dashed];
  8 -> 9 [color=blue];
  8 -> 10 [color=red, style=dashed];
}`

	tests[6].heap = "Binomial Min Heap"
	tests[6].merge = NewBinomial(cmpMin, eqVal)
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
	tests[6].expectedPeek = KeyValue[int, string]{Key: 10, Val: "Task#1"}
	tests[6].expectedContains = append(tests[6].expectedContains,
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		KeyValue[int, string]{Key: 66, Val: "Task#6.a"},
		KeyValue[int, string]{Key: 77, Val: "Task#7.a"},
		KeyValue[int, string]{Key: 88, Val: "Task#8.a"},
		KeyValue[int, string]{Key: 99, Val: "Task#9.a"},
	)
	tests[6].expectedDelete = []KeyValue[int, string]{
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
	tests[6].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="20 | Task#2"];
  2 [label="22 | Task#2.a"];
  3 [label="10 | Task#1"];
  4 [label="11 | Task#1.a"];
  5 [label="44 | Task#4.a"];
  6 [label="88 | Task#8.a"];
  7 [label="99 | Task#9.a"];
  8 [label="77 | Task#7.a"];
  9 [label="55 | Task#5.a"];
  10 [label="66 | Task#6.a"];
  11 [label="33 | Task#3.a"];
  12 [label="40 | Task#4"];
  13 [label="80 | Task#8"];
  14 [label="90 | Task#9"];
  15 [label="70 | Task#7"];
  16 [label="50 | Task#5"];
  17 [label="60 | Task#6"];
  18 [label="30 | Task#3"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 12 [color=red, style=dashed];
  5 -> 6 [color=blue];
  5 -> 9 [color=red, style=dashed];
  6 -> 7 [color=blue];
  6 -> 8 [color=red, style=dashed];
  9 -> 10 [color=blue];
  9 -> 11 [color=red, style=dashed];
  12 -> 13 [color=blue];
  12 -> 16 [color=red, style=dashed];
  13 -> 14 [color=blue];
  13 -> 15 [color=red, style=dashed];
  16 -> 17 [color=blue];
  16 -> 18 [color=red, style=dashed];
}`

	tests[7].heap = "Binomial Max Heap"
	tests[7].merge = NewBinomial(cmpMax, eqVal)
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
	tests[7].expectedPeek = KeyValue[int, string]{Key: 99, Val: "Task#9.a"}
	tests[7].expectedContains = append(tests[7].expectedContains,
		KeyValue[int, string]{Key: 99, Val: "Task#9.a"},
		KeyValue[int, string]{Key: 88, Val: "Task#8.a"},
		KeyValue[int, string]{Key: 77, Val: "Task#7.a"},
		KeyValue[int, string]{Key: 66, Val: "Task#6.a"},
		KeyValue[int, string]{Key: 55, Val: "Task#5.a"},
		KeyValue[int, string]{Key: 44, Val: "Task#4.a"},
		KeyValue[int, string]{Key: 33, Val: "Task#3.a"},
		KeyValue[int, string]{Key: 22, Val: "Task#2.a"},
		KeyValue[int, string]{Key: 11, Val: "Task#1.a"},
	)
	tests[7].expectedDelete = []KeyValue[int, string]{
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
	tests[7].expectedDOT = `strict digraph "Binomial Heap" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="88 | Task#8.a"];
  2 [label="80 | Task#8"];
  3 [label="99 | Task#9.a"];
  4 [label="90 | Task#9"];
  5 [label="50 | Task#5"];
  6 [label="30 | Task#3"];
  7 [label="10 | Task#1"];
  8 [label="20 | Task#2"];
  9 [label="60 | Task#6"];
  10 [label="40 | Task#4"];
  11 [label="70 | Task#7"];
  12 [label="55 | Task#5.a"];
  13 [label="33 | Task#3.a"];
  14 [label="11 | Task#1.a"];
  15 [label="22 | Task#2.a"];
  16 [label="66 | Task#6.a"];
  17 [label="44 | Task#4.a"];
  18 [label="77 | Task#7.a"];

  1 -> 2 [color=blue];
  1 -> 3 [color=red, style=dashed];
  3 -> 4 [color=blue];
  4 -> 5 [color=blue];
  4 -> 12 [color=red, style=dashed];
  5 -> 6 [color=blue];
  5 -> 9 [color=red, style=dashed];
  6 -> 7 [color=blue];
  6 -> 8 [color=red, style=dashed];
  9 -> 10 [color=blue];
  9 -> 11 [color=red, style=dashed];
  12 -> 13 [color=blue];
  12 -> 16 [color=red, style=dashed];
  13 -> 14 [color=blue];
  13 -> 15 [color=red, style=dashed];
  16 -> 17 [color=blue];
  16 -> 18 [color=red, style=dashed];
}`

	return tests
}

func TestBinomialHeap(t *testing.T) {
	tests := getBinomialTests()

	for _, tc := range tests {
		heap := NewBinomial(tc.cmpKey, tc.eqVal)
		runMergeableHeapTest(t, heap, tc)
	}
}
