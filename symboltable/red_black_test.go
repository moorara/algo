package symboltable

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getRedBlackTests() []orderedSymbolTableTest[string, int] {
	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "LLRB Tree"
	tests[0].expectedHeight = 2
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{"B", 2}, {"A", 1}, {"C", 3}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{"B", 2}, {"C", 3}, {"A", 1}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"B", 2}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{"C", 3}, {"A", 1}, {"B", 2}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedGraphviz = `strict digraph "Red-Black" {
  concentrate=false;
  node [style=filled, shape=oval];

  1 [label="B,2", color=black, fontcolor=white];
  2 [label="A,1", color=black, fontcolor=white];
  3 [label="C,3", color=black, fontcolor=white];

  1 -> 2 [color=black];
  1 -> 3 [];
}`

	tests[1].symbolTable = "LLRB Tree"
	tests[1].expectedHeight = 3
	tests[1].equals = NewRedBlack[string, int](NewCompareFunc[string](), nil)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{"D", 4}, {"B", 2}, {"A", 1}, {"C", 3}, {"E", 5}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{"D", 4}, {"E", 5}, {"B", 2}, {"C", 3}, {"A", 1}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"B", 2}, {"E", 5}, {"D", 4}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{"E", 5}, {"C", 3}, {"A", 1}, {"B", 2}, {"D", 4}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedGraphviz = `strict digraph "Red-Black" {
  concentrate=false;
  node [style=filled, shape=oval];

  1 [label="D,4", color=black, fontcolor=white];
  2 [label="B,2", color=red, fontcolor=white];
  3 [label="A,1", color=black, fontcolor=white];
  4 [label="C,3", color=black, fontcolor=white];
  5 [label="E,5", color=black, fontcolor=white];

  1 -> 2 [color=red];
  1 -> 5 [];
  2 -> 3 [color=black];
  2 -> 4 [];
}`

	tests[2].symbolTable = "LLRB Tree"
	tests[2].expectedHeight = 3
	tests[2].equals = NewRedBlack[string, int](NewCompareFunc[string](), nil)
	tests[2].equals.Put("D", 4)
	tests[2].equals.Put("J", 10)
	tests[2].equals.Put("P", 16)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{"J", 10}, {"D", 4}, {"A", 1}, {"G", 7}, {"P", 16}, {"M", 13}, {"S", 19}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{"J", 10}, {"P", 16}, {"S", 19}, {"M", 13}, {"D", 4}, {"G", 7}, {"A", 1}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"G", 7}, {"D", 4}, {"M", 13}, {"S", 19}, {"P", 16}, {"J", 10}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{"S", 19}, {"M", 13}, {"P", 16}, {"G", 7}, {"A", 1}, {"D", 4}, {"J", 10}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedGraphviz = `strict digraph "Red-Black" {
  concentrate=false;
  node [style=filled, shape=oval];

  1 [label="J,10", color=black, fontcolor=white];
  2 [label="D,4", color=black, fontcolor=white];
  3 [label="A,1", color=black, fontcolor=white];
  4 [label="G,7", color=black, fontcolor=white];
  5 [label="P,16", color=black, fontcolor=white];
  6 [label="M,13", color=black, fontcolor=white];
  7 [label="S,19", color=black, fontcolor=white];

  1 -> 2 [color=black];
  1 -> 5 [];
  2 -> 3 [color=black];
  2 -> 4 [];
  5 -> 6 [color=black];
  5 -> 7 [];
}`

	tests[3].symbolTable = "LLRB Tree"
	tests[3].expectedHeight = 3
	tests[3].equals = NewRedBlack[string, int](NewCompareFunc[string](), nil)
	tests[3].equals.Put("box", 2)
	tests[3].equals.Put("dad", 3)
	tests[3].equals.Put("baby", 5)
	tests[3].equals.Put("dome", 7)
	tests[3].equals.Put("band", 11)
	tests[3].equals.Put("dance", 13)
	tests[3].equals.Put("balloon", 17)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{"box", 2}, {"balloon", 17}, {"baby", 5}, {"band", 11}, {"dance", 13}, {"dad", 3}, {"dome", 7}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{"box", 2}, {"dance", 13}, {"dome", 7}, {"dad", 3}, {"balloon", 17}, {"band", 11}, {"baby", 5}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{"baby", 5}, {"band", 11}, {"balloon", 17}, {"dad", 3}, {"dome", 7}, {"dance", 13}, {"box", 2}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{"dome", 7}, {"dad", 3}, {"dance", 13}, {"band", 11}, {"baby", 5}, {"balloon", 17}, {"box", 2}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedGraphviz = `strict digraph "Red-Black" {
  concentrate=false;
  node [style=filled, shape=oval];

  1 [label="box,2", color=black, fontcolor=white];
  2 [label="balloon,17", color=black, fontcolor=white];
  3 [label="baby,5", color=black, fontcolor=white];
  4 [label="band,11", color=black, fontcolor=white];
  5 [label="dance,13", color=black, fontcolor=white];
  6 [label="dad,3", color=black, fontcolor=white];
  7 [label="dome,7", color=black, fontcolor=white];

  1 -> 2 [color=black];
  1 -> 5 [];
  2 -> 3 [color=black];
  2 -> 4 [];
  5 -> 6 [color=black];
  5 -> 7 [];
}`

	return tests
}

func TestRedBlack(t *testing.T) {
	tests := getRedBlackTests()

	for _, tc := range tests {
		rbt := NewRedBlack[string, int](tc.cmpKey, tc.eqVal)
		runOrderedSymbolTableTest(t, rbt, tc)
	}
}
