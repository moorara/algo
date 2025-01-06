package trie

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getBinaryTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].trie = "Binary Trie"
	tests[0].expectedHeight = 3
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "", Val: 0}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[0].expectedDOT = `strict digraph "Binary Trie" {
  concentrate=false;
  node [shape=circle];

  1 [label="•"];
  2 [label="A,1", color=black, style=filled, fontcolor=white];
  3 [label="B,2", color=black, style=filled, fontcolor=white];
  4 [label="C,3", color=black, style=filled, fontcolor=white];

  1 -> 2 [color=blue];
  2 -> 3 [color=red];
  3 -> 4 [color=red];
}`

	tests[1].trie = "Binary Trie"
	tests[1].expectedHeight = 5
	tests[1].equals = NewBinary[int](nil)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}, {Key: "", Val: 0}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[1].expectedDOT = `strict digraph "Binary Trie" {
  concentrate=false;
  node [shape=circle];

  1 [label="•"];
  2 [label="A,1", color=black, style=filled, fontcolor=white];
  3 [label="B,2", color=black, style=filled, fontcolor=white];
  4 [label="C,3", color=black, style=filled, fontcolor=white];
  5 [label="D,4", color=black, style=filled, fontcolor=white];
  6 [label="E,5", color=black, style=filled, fontcolor=white];

  1 -> 2 [color=blue];
  2 -> 3 [color=red];
  3 -> 4 [color=red];
  4 -> 5 [color=red];
  5 -> 6 [color=red];
}`

	tests[2].trie = "Binary Trie"
	tests[2].expectedHeight = 7
	tests[2].equals = NewBinary[int](nil)
	tests[2].equals.Put("A", 1)
	tests[2].equals.Put("D", 4)
	tests[2].equals.Put("G", 7)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}, {Key: "", Val: 0}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "", Val: 0}}
	tests[2].expectedDOT = `strict digraph "Binary Trie" {
  concentrate=false;
  node [shape=circle];

  1 [label="•"];
  2 [label="A,1", color=black, style=filled, fontcolor=white];
  3 [label="D,4", color=black, style=filled, fontcolor=white];
  4 [label="G,7", color=black, style=filled, fontcolor=white];
  5 [label="J,10", color=black, style=filled, fontcolor=white];
  6 [label="M,13", color=black, style=filled, fontcolor=white];
  7 [label="P,16", color=black, style=filled, fontcolor=white];
  8 [label="S,19", color=black, style=filled, fontcolor=white];

  1 -> 2 [color=blue];
  2 -> 3 [color=red];
  3 -> 4 [color=red];
  4 -> 5 [color=red];
  5 -> 6 [color=red];
  6 -> 7 [color=red];
  7 -> 8 [color=red];
}`

	tests[3].trie = "Binary Trie"
	tests[3].expectedHeight = 8
	tests[3].equals = NewBinary[int](nil)
	tests[3].equals.Put("box", 2)
	tests[3].equals.Put("dad", 3)
	tests[3].equals.Put("baby", 5)
	tests[3].equals.Put("dome", 7)
	tests[3].equals.Put("band", 11)
	tests[3].equals.Put("dance", 13)
	tests[3].equals.Put("balloon", 17)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "b", Val: 0}, {Key: "a", Val: 0}, {Key: "b", Val: 0}, {Key: "y", Val: 5}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "n", Val: 17}, {Key: "n", Val: 0}, {Key: "d", Val: 11}, {Key: "o", Val: 0}, {Key: "x", Val: 2}, {Key: "d", Val: 0}, {Key: "a", Val: 0}, {Key: "d", Val: 3}, {Key: "n", Val: 0}, {Key: "c", Val: 0}, {Key: "e", Val: 13}, {Key: "o", Val: 0}, {Key: "m", Val: 0}, {Key: "e", Val: 7}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "b", Val: 0}, {Key: "d", Val: 0}, {Key: "a", Val: 0}, {Key: "o", Val: 0}, {Key: "m", Val: 0}, {Key: "e", Val: 7}, {Key: "d", Val: 3}, {Key: "n", Val: 0}, {Key: "c", Val: 0}, {Key: "e", Val: 13}, {Key: "a", Val: 0}, {Key: "o", Val: 0}, {Key: "x", Val: 2}, {Key: "b", Val: 0}, {Key: "l", Val: 0}, {Key: "n", Val: 0}, {Key: "d", Val: 11}, {Key: "l", Val: 0}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "n", Val: 17}, {Key: "y", Val: 5}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{Key: "y", Val: 5}, {Key: "b", Val: 0}, {Key: "n", Val: 17}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "d", Val: 11}, {Key: "n", Val: 0}, {Key: "a", Val: 0}, {Key: "x", Val: 2}, {Key: "o", Val: 0}, {Key: "b", Val: 0}, {Key: "d", Val: 3}, {Key: "e", Val: 13}, {Key: "c", Val: 0}, {Key: "n", Val: 0}, {Key: "a", Val: 0}, {Key: "e", Val: 7}, {Key: "m", Val: 0}, {Key: "o", Val: 0}, {Key: "d", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "d", Val: 0}, {Key: "o", Val: 0}, {Key: "m", Val: 0}, {Key: "e", Val: 7}, {Key: "a", Val: 0}, {Key: "n", Val: 0}, {Key: "c", Val: 0}, {Key: "e", Val: 13}, {Key: "d", Val: 3}, {Key: "b", Val: 0}, {Key: "o", Val: 0}, {Key: "x", Val: 2}, {Key: "a", Val: 0}, {Key: "n", Val: 0}, {Key: "d", Val: 11}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "n", Val: 17}, {Key: "b", Val: 0}, {Key: "y", Val: 5}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{Key: "y", Val: 5}, {Key: "n", Val: 17}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "l", Val: 0}, {Key: "d", Val: 11}, {Key: "n", Val: 0}, {Key: "l", Val: 0}, {Key: "b", Val: 0}, {Key: "x", Val: 2}, {Key: "o", Val: 0}, {Key: "a", Val: 0}, {Key: "e", Val: 13}, {Key: "c", Val: 0}, {Key: "n", Val: 0}, {Key: "d", Val: 3}, {Key: "e", Val: 7}, {Key: "m", Val: 0}, {Key: "o", Val: 0}, {Key: "a", Val: 0}, {Key: "d", Val: 0}, {Key: "b", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{Key: "e", Val: 7}, {Key: "m", Val: 0}, {Key: "o", Val: 0}, {Key: "e", Val: 13}, {Key: "c", Val: 0}, {Key: "n", Val: 0}, {Key: "d", Val: 3}, {Key: "a", Val: 0}, {Key: "d", Val: 0}, {Key: "x", Val: 2}, {Key: "o", Val: 0}, {Key: "d", Val: 11}, {Key: "n", Val: 0}, {Key: "n", Val: 17}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "y", Val: 5}, {Key: "b", Val: 0}, {Key: "a", Val: 0}, {Key: "b", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "b", Val: 0}, {Key: "a", Val: 0}, {Key: "b", Val: 0}, {Key: "y", Val: 5}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "n", Val: 17}, {Key: "n", Val: 0}, {Key: "d", Val: 11}, {Key: "o", Val: 0}, {Key: "x", Val: 2}, {Key: "d", Val: 0}, {Key: "a", Val: 0}, {Key: "d", Val: 3}, {Key: "n", Val: 0}, {Key: "c", Val: 0}, {Key: "e", Val: 13}, {Key: "o", Val: 0}, {Key: "m", Val: 0}, {Key: "e", Val: 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "e", Val: 7}, {Key: "m", Val: 0}, {Key: "o", Val: 0}, {Key: "e", Val: 13}, {Key: "c", Val: 0}, {Key: "n", Val: 0}, {Key: "d", Val: 3}, {Key: "a", Val: 0}, {Key: "d", Val: 0}, {Key: "x", Val: 2}, {Key: "o", Val: 0}, {Key: "d", Val: 11}, {Key: "n", Val: 0}, {Key: "n", Val: 17}, {Key: "o", Val: 0}, {Key: "o", Val: 0}, {Key: "l", Val: 0}, {Key: "l", Val: 0}, {Key: "y", Val: 5}, {Key: "b", Val: 0}, {Key: "a", Val: 0}, {Key: "b", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDOT = `strict digraph "Binary Trie" {
  concentrate=false;
  node [shape=circle];

  1 [label="•"];
  2 [label="b"];
  3 [label="a"];
  4 [label="b"];
  5 [label="y,5", color=black, style=filled, fontcolor=white];
  6 [label="l"];
  7 [label="l"];
  8 [label="o"];
  9 [label="o"];
  10 [label="n,17", color=black, style=filled, fontcolor=white];
  11 [label="n"];
  12 [label="d,11", color=black, style=filled, fontcolor=white];
  13 [label="o"];
  14 [label="x,2", color=black, style=filled, fontcolor=white];
  15 [label="d"];
  16 [label="a"];
  17 [label="d,3", color=black, style=filled, fontcolor=white];
  18 [label="n"];
  19 [label="c"];
  20 [label="e,13", color=black, style=filled, fontcolor=white];
  21 [label="o"];
  22 [label="m"];
  23 [label="e,7", color=black, style=filled, fontcolor=white];

  1 -> 2 [color=blue];
  2 -> 3 [color=blue];
  2 -> 15 [color=red];
  3 -> 4 [color=blue];
  3 -> 13 [color=red];
  4 -> 5 [color=blue];
  4 -> 6 [color=red];
  6 -> 7 [color=blue];
  6 -> 11 [color=red];
  7 -> 8 [color=blue];
  8 -> 9 [color=blue];
  9 -> 10 [color=blue];
  11 -> 12 [color=blue];
  13 -> 14 [color=blue];
  15 -> 16 [color=blue];
  16 -> 17 [color=blue];
  16 -> 21 [color=red];
  17 -> 18 [color=red];
  18 -> 19 [color=blue];
  19 -> 20 [color=blue];
  21 -> 22 [color=blue];
  22 -> 23 [color=blue];
}`

	return tests
}

func TestBinaryTrie(t *testing.T) {
	tests := getBinaryTests()

	for _, tc := range tests {
		bin := NewBinary[int](tc.eqVal)
		runTrieTest(t, bin, tc)
	}
}
