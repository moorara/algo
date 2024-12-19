package trie

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getPatriciaTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].trie = "Patricia"
	tests[0].expectedHeight = 2
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "C", Val: 3}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "C", Val: 3}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "C", Val: 3}, {Key: "B", Val: 2}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "A", Val: 1}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "A", Val: 1}, {Key: "B", Val: 2}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "A", Val: 1}, {Key: "B", Val: 2}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[0].expectedGraphviz = `strict digraph "Patricia Trie" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ B,2 | { <l>• | 0 | 01000010 | <r>• } }"];
  2 [label="{ A,1 | { <l>• | 7 | 01000001 | <r>• } }"];
  3 [label="{ C,3 | { <l>• | 8 | 01000011 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 2 [color=red, style=dashed];
  2:r -> 3 [color=blue];
  3:l -> 1 [color=red, style=dashed];
  3:r -> 3 [color=red, style=dashed];
}`

	tests[1].trie = "Patricia"
	tests[1].expectedHeight = 3
	tests[1].equals = NewPatricia[int](nil)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "C", Val: 3}, {Key: "E", Val: 5}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "D", Val: 4}, {Key: "E", Val: 5}, {Key: "A", Val: 1}, {Key: "C", Val: 3}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}, {Key: "B", Val: 2}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "A", Val: 1}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "A", Val: 1}, {Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "B", Val: 2}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "C", Val: 3}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "B", Val: 2}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[1].expectedGraphviz = `strict digraph "Patricia Trie" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ B,2 | { <l>• | 0 | 01000010 | <r>• } }"];
  2 [label="{ D,4 | { <l>• | 6 | 01000100 | <r>• } }"];
  3 [label="{ A,1 | { <l>• | 7 | 01000001 | <r>• } }"];
  4 [label="{ C,3 | { <l>• | 8 | 01000011 | <r>• } }"];
  5 [label="{ E,5 | { <l>• | 8 | 01000101 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 5 [color=blue];
  3:l -> 3 [color=red, style=dashed];
  3:r -> 4 [color=blue];
  4:l -> 1 [color=red, style=dashed];
  4:r -> 4 [color=red, style=dashed];
  5:l -> 2 [color=red, style=dashed];
  5:r -> 5 [color=red, style=dashed];
}`

	tests[2].trie = "Patricia"
	tests[2].expectedHeight = 4
	tests[2].equals = NewPatricia[int](nil)
	tests[2].equals.Put("A", 1)
	tests[2].equals.Put("D", 4)
	tests[2].equals.Put("G", 7)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{Key: "J", Val: 10}, {Key: "P", Val: 16}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "G", Val: 7}, {Key: "M", Val: 13}, {Key: "S", Val: 19}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{Key: "J", Val: 10}, {Key: "P", Val: 16}, {Key: "S", Val: 19}, {Key: "D", Val: 4}, {Key: "M", Val: 13}, {Key: "A", Val: 1}, {Key: "G", Val: 7}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}, {Key: "J", Val: 10}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{Key: "J", Val: 10}, {Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "A", Val: 1}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{Key: "G", Val: 7}, {Key: "A", Val: 1}, {Key: "M", Val: 13}, {Key: "D", Val: 4}, {Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "J", Val: 10}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "M", Val: 13}, {Key: "G", Val: 7}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "P", Val: 16}, {Key: "J", Val: 10}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}}
	tests[2].expectedGraphviz = `strict digraph "Patricia Trie" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ J,10 | { <l>• | 0 | 01001010 | <r>• } }"];
  2 [label="{ P,16 | { <l>• | 4 | 01010000 | <r>• } }"];
  3 [label="{ D,4 | { <l>• | 5 | 01000100 | <r>• } }"];
  4 [label="{ A,1 | { <l>• | 6 | 01000001 | <r>• } }"];
  5 [label="{ G,7 | { <l>• | 7 | 01000111 | <r>• } }"];
  6 [label="{ M,13 | { <l>• | 6 | 01001101 | <r>• } }"];
  7 [label="{ S,19 | { <l>• | 7 | 01010011 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 7 [color=blue];
  3:l -> 4 [color=blue];
  3:r -> 6 [color=blue];
  4:l -> 4 [color=red, style=dashed];
  4:r -> 5 [color=blue];
  5:l -> 3 [color=red, style=dashed];
  5:r -> 5 [color=red, style=dashed];
  6:l -> 1 [color=red, style=dashed];
  6:r -> 6 [color=red, style=dashed];
  7:l -> 2 [color=red, style=dashed];
  7:r -> 7 [color=red, style=dashed];
}`

	tests[3].trie = "Patricia"
	tests[3].expectedHeight = 4
	tests[3].equals = NewPatricia[int](nil)
	tests[3].equals.Put("box", 2)
	tests[3].equals.Put("dad", 3)
	tests[3].equals.Put("baby", 5)
	tests[3].equals.Put("dome", 7)
	tests[3].equals.Put("band", 11)
	tests[3].equals.Put("dance", 13)
	tests[3].equals.Put("balloon", 17)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{Key: "box", Val: 2}, {Key: "dad", Val: 3}, {Key: "band", Val: 11}, {Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "dome", Val: 7}, {Key: "dance", Val: 13}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{Key: "box", Val: 2}, {Key: "dad", Val: 3}, {Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "band", Val: 11}, {Key: "baby", Val: 5}, {Key: "balloon", Val: 17}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "band", Val: 11}, {Key: "dad", Val: 3}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}, {Key: "box", Val: 2}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{Key: "box", Val: 2}, {Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "dad", Val: 3}, {Key: "band", Val: 11}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{Key: "balloon", Val: 17}, {Key: "baby", Val: 5}, {Key: "band", Val: 11}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}, {Key: "dad", Val: 3}, {Key: "box", Val: 2}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{Key: "dance", Val: 13}, {Key: "dome", Val: 7}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}, {Key: "band", Val: 11}, {Key: "dad", Val: 3}, {Key: "box", Val: 2}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "band", Val: 11}, {Key: "box", Val: 2}, {Key: "dad", Val: 3}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "dad", Val: 3}, {Key: "box", Val: 2}, {Key: "band", Val: 11}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}}
	tests[3].expectedGraphviz = `strict digraph "Patricia Trie" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ box,2 | { <l>• | 0 | 01100010 01101111 01111000 | <r>• } }"];
  2 [label="{ dad,3 | { <l>• | 6 | 01100100 01100001 01100100 | <r>• } }"];
  3 [label="{ band,11 | { <l>• | 13 | 01100010 01100001 01101110 01100100 | <r>• } }"];
  4 [label="{ baby,5 | { <l>• | 21 | 01100010 01100001 01100010 01111001 | <r>• } }"];
  5 [label="{ balloon,17 | { <l>• | 23 | 01100010 01100001 01101100 01101100 01101111 01101111 01101110 | <r>• } }"];
  6 [label="{ dome,7 | { <l>• | 13 | 01100100 01101111 01101101 01100101 | <r>• } }"];
  7 [label="{ dance,13 | { <l>• | 21 | 01100100 01100001 01101110 01100011 01100101 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 6 [color=blue];
  3:l -> 4 [color=blue];
  3:r -> 1 [color=red, style=dashed];
  4:l -> 4 [color=red, style=dashed];
  4:r -> 5 [color=blue];
  5:l -> 5 [color=red, style=dashed];
  5:r -> 3 [color=red, style=dashed];
  6:l -> 7 [color=blue];
  6:r -> 6 [color=red, style=dashed];
  7:l -> 2 [color=red, style=dashed];
  7:r -> 7 [color=red, style=dashed];
}`

	return tests
}

func TestPatricia(t *testing.T) {
	tests := getPatriciaTests()

	for _, tc := range tests {
		pat := NewPatricia[int](tc.eqVal)
		runTrieTest(t, pat, tc)
	}
}
