package symboltable

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getAVLTests() []orderedSymbolTableTest[string, int] {
	cmpKey := NewCompareFunc[string]()
	eqVal := NewEqualFunc[int]()

	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "AVL"
	tests[0].expectedHeight = 2
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "C", Val: 3}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "A", Val: 1}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "C", Val: 3}, {Key: "B", Val: 2}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "A", Val: 1}, {Key: "B", Val: 2}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[0].expectedGraphviz = `strict digraph "AVL" {
  concentrate=false;
  node [shape=oval];

  1 [label="B,2"];
  2 [label="A,1"];
  3 [label="C,3"];

  1 -> 2 [];
  1 -> 3 [];
}`

	tests[1].symbolTable = "AVL"
	tests[1].expectedHeight = 3
	tests[1].equals = NewAVL[string, int](cmpKey, eqVal)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "E", Val: 5}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{Key: "B", Val: 2}, {Key: "D", Val: 4}, {Key: "E", Val: 5}, {Key: "C", Val: 3}, {Key: "A", Val: 1}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "C", Val: 3}, {Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "B", Val: 2}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "B", Val: 2}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "B", Val: 2}, {Key: "C", Val: 3}, {Key: "D", Val: 4}, {Key: "E", Val: 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "E", Val: 5}, {Key: "D", Val: 4}, {Key: "C", Val: 3}, {Key: "B", Val: 2}, {Key: "A", Val: 1}}
	tests[1].expectedGraphviz = `strict digraph "AVL" {
  concentrate=false;
  node [shape=oval];

  1 [label="B,2"];
  2 [label="A,1"];
  3 [label="D,4"];
  4 [label="C,3"];
  5 [label="E,5"];

  1 -> 2 [];
  1 -> 3 [];
  3 -> 4 [];
  3 -> 5 [];
}`

	tests[2].symbolTable = "AVL"
	tests[2].expectedHeight = 3
	tests[2].equals = NewAVL[string, int](cmpKey, eqVal)
	tests[2].equals.Put("D", 4)
	tests[2].equals.Put("J", 10)
	tests[2].equals.Put("P", 16)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{Key: "J", Val: 10}, {Key: "D", Val: 4}, {Key: "A", Val: 1}, {Key: "G", Val: 7}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "S", Val: 19}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{Key: "J", Val: 10}, {Key: "P", Val: 16}, {Key: "S", Val: 19}, {Key: "M", Val: 13}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "A", Val: 1}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "M", Val: 13}, {Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "J", Val: 10}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "G", Val: 7}, {Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "J", Val: 10}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "A", Val: 1}, {Key: "D", Val: 4}, {Key: "G", Val: 7}, {Key: "J", Val: 10}, {Key: "M", Val: 13}, {Key: "P", Val: 16}, {Key: "S", Val: 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "S", Val: 19}, {Key: "P", Val: 16}, {Key: "M", Val: 13}, {Key: "J", Val: 10}, {Key: "G", Val: 7}, {Key: "D", Val: 4}, {Key: "A", Val: 1}}
	tests[2].expectedGraphviz = `strict digraph "AVL" {
  concentrate=false;
  node [shape=oval];

  1 [label="J,10"];
  2 [label="D,4"];
  3 [label="A,1"];
  4 [label="G,7"];
  5 [label="P,16"];
  6 [label="M,13"];
  7 [label="S,19"];

  1 -> 2 [];
  1 -> 5 [];
  2 -> 3 [];
  2 -> 4 [];
  5 -> 6 [];
  5 -> 7 [];
}`

	tests[3].symbolTable = "AVL"
	tests[3].expectedHeight = 3
	tests[3].equals = NewAVL[string, int](cmpKey, eqVal)
	tests[3].equals.Put("box", 2)
	tests[3].equals.Put("dad", 3)
	tests[3].equals.Put("baby", 5)
	tests[3].equals.Put("dome", 7)
	tests[3].equals.Put("band", 11)
	tests[3].equals.Put("dance", 13)
	tests[3].equals.Put("balloon", 17)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{Key: "box", Val: 2}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}, {Key: "band", Val: 11}, {Key: "dance", Val: 13}, {Key: "dad", Val: 3}, {Key: "dome", Val: 7}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{Key: "box", Val: 2}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}, {Key: "dad", Val: 3}, {Key: "balloon", Val: 17}, {Key: "band", Val: 11}, {Key: "baby", Val: 5}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "band", Val: 11}, {Key: "box", Val: 2}, {Key: "dad", Val: 3}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "dad", Val: 3}, {Key: "box", Val: 2}, {Key: "band", Val: 11}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{Key: "baby", Val: 5}, {Key: "band", Val: 11}, {Key: "balloon", Val: 17}, {Key: "dad", Val: 3}, {Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "box", Val: 2}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{Key: "dome", Val: 7}, {Key: "dad", Val: 3}, {Key: "dance", Val: 13}, {Key: "band", Val: 11}, {Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "box", Val: 2}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "baby", Val: 5}, {Key: "balloon", Val: 17}, {Key: "band", Val: 11}, {Key: "box", Val: 2}, {Key: "dad", Val: 3}, {Key: "dance", Val: 13}, {Key: "dome", Val: 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "dome", Val: 7}, {Key: "dance", Val: 13}, {Key: "dad", Val: 3}, {Key: "box", Val: 2}, {Key: "band", Val: 11}, {Key: "balloon", Val: 17}, {Key: "baby", Val: 5}}
	tests[3].expectedGraphviz = `strict digraph "AVL" {
  concentrate=false;
  node [shape=oval];

  1 [label="box,2"];
  2 [label="balloon,17"];
  3 [label="baby,5"];
  4 [label="band,11"];
  5 [label="dance,13"];
  6 [label="dad,3"];
  7 [label="dome,7"];

  1 -> 2 [];
  1 -> 5 [];
  2 -> 3 [];
  2 -> 4 [];
  5 -> 6 [];
  5 -> 7 [];
}`

	return tests
}

func TestAVL(t *testing.T) {
	tests := getAVLTests()

	for _, tc := range tests {
		avl := NewAVL[string, int](tc.cmpKey, tc.eqVal)
		runOrderedSymbolTableTest(t, avl, tc)
	}
}
