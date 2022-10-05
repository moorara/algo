package trie

import "testing"

func getBinaryTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].symbolTable = "Binary Trie"
	tests[0].expectedHeight = 3
	tests[0].expectedVLRTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedVRLTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"B", 2}, {"C", 3}, {"", 0}}
	tests[0].expectedRVLTraverse = []KeyValue[int]{{"", 0}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedLRVTraverse = []KeyValue[int]{{"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[0].expectedRLVTraverse = []KeyValue[int]{{"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[0].expectedAscendingTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[int]{{"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[0].expectedDotCode = `strict digraph "Binary Trie" {
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

	tests[1].symbolTable = "Binary Trie"
	tests[1].expectedHeight = 5
	tests[1].expectedVLRTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedVRLTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}, {"", 0}}
	tests[1].expectedRVLTraverse = []KeyValue[int]{{"", 0}, {"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedLRVTraverse = []KeyValue[int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[1].expectedRLVTraverse = []KeyValue[int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[1].expectedAscendingTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}, {"", 0}}
	tests[1].expectedDotCode = `strict digraph "Binary Trie" {
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

	tests[2].symbolTable = "Binary Trie"
	tests[2].expectedHeight = 7
	tests[2].expectedVLRTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedVRLTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}, {"", 0}}
	tests[2].expectedRVLTraverse = []KeyValue[int]{{"", 0}, {"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedLRVTraverse = []KeyValue[int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}, {"", 0}}
	tests[2].expectedRLVTraverse = []KeyValue[int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}, {"", 0}}
	tests[2].expectedAscendingTraverse = []KeyValue[int]{{"", 0}, {"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}, {"", 0}}
	tests[2].expectedDotCode = `strict digraph "Binary Trie" {
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

	tests[3].symbolTable = "Binary Trie"
	tests[3].expectedHeight = 8
	tests[3].expectedVLRTraverse = []KeyValue[int]{{"", 0}, {"b", 0}, {"a", 0}, {"b", 0}, {"y", 5}, {"l", 0}, {"l", 0}, {"o", 0}, {"o", 0}, {"n", 17}, {"n", 0}, {"d", 11}, {"o", 0}, {"x", 2}, {"d", 0}, {"a", 0}, {"d", 3}, {"n", 0}, {"c", 0}, {"e", 13}, {"o", 0}, {"m", 0}, {"e", 7}}
	tests[3].expectedVRLTraverse = []KeyValue[int]{{"", 0}, {"b", 0}, {"d", 0}, {"a", 0}, {"o", 0}, {"m", 0}, {"e", 7}, {"d", 3}, {"n", 0}, {"c", 0}, {"e", 13}, {"a", 0}, {"o", 0}, {"x", 2}, {"b", 0}, {"l", 0}, {"n", 0}, {"d", 11}, {"l", 0}, {"o", 0}, {"o", 0}, {"n", 17}, {"y", 5}}
	tests[3].expectedLVRTraverse = []KeyValue[int]{{"y", 5}, {"b", 0}, {"n", 17}, {"o", 0}, {"o", 0}, {"l", 0}, {"l", 0}, {"d", 11}, {"n", 0}, {"a", 0}, {"x", 2}, {"o", 0}, {"b", 0}, {"d", 3}, {"e", 13}, {"c", 0}, {"n", 0}, {"a", 0}, {"e", 7}, {"m", 0}, {"o", 0}, {"d", 0}, {"", 0}}
	tests[3].expectedRVLTraverse = []KeyValue[int]{{"", 0}, {"d", 0}, {"o", 0}, {"m", 0}, {"e", 7}, {"a", 0}, {"n", 0}, {"c", 0}, {"e", 13}, {"d", 3}, {"b", 0}, {"o", 0}, {"x", 2}, {"a", 0}, {"n", 0}, {"d", 11}, {"l", 0}, {"l", 0}, {"o", 0}, {"o", 0}, {"n", 17}, {"b", 0}, {"y", 5}}
	tests[3].expectedLRVTraverse = []KeyValue[int]{{"y", 5}, {"n", 17}, {"o", 0}, {"o", 0}, {"l", 0}, {"d", 11}, {"n", 0}, {"l", 0}, {"b", 0}, {"x", 2}, {"o", 0}, {"a", 0}, {"e", 13}, {"c", 0}, {"n", 0}, {"d", 3}, {"e", 7}, {"m", 0}, {"o", 0}, {"a", 0}, {"d", 0}, {"b", 0}, {"", 0}}
	tests[3].expectedRLVTraverse = []KeyValue[int]{{"e", 7}, {"m", 0}, {"o", 0}, {"e", 13}, {"c", 0}, {"n", 0}, {"d", 3}, {"a", 0}, {"d", 0}, {"x", 2}, {"o", 0}, {"d", 11}, {"n", 0}, {"n", 17}, {"o", 0}, {"o", 0}, {"l", 0}, {"l", 0}, {"y", 5}, {"b", 0}, {"a", 0}, {"b", 0}, {"", 0}}
	tests[3].expectedAscendingTraverse = []KeyValue[int]{{"", 0}, {"b", 0}, {"a", 0}, {"b", 0}, {"y", 5}, {"l", 0}, {"l", 0}, {"o", 0}, {"o", 0}, {"n", 17}, {"n", 0}, {"d", 11}, {"o", 0}, {"x", 2}, {"d", 0}, {"a", 0}, {"d", 3}, {"n", 0}, {"c", 0}, {"e", 13}, {"o", 0}, {"m", 0}, {"e", 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[int]{{"e", 7}, {"m", 0}, {"o", 0}, {"e", 13}, {"c", 0}, {"n", 0}, {"d", 3}, {"a", 0}, {"d", 0}, {"x", 2}, {"o", 0}, {"d", 11}, {"n", 0}, {"n", 17}, {"o", 0}, {"o", 0}, {"l", 0}, {"l", 0}, {"y", 5}, {"b", 0}, {"a", 0}, {"b", 0}, {"", 0}}
	tests[3].expectedDotCode = `strict digraph "Binary Trie" {
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
		bin := NewBinary[int]()
		runTrieTest(t, bin, tc)
	}
}
