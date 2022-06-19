package symboltable

import "testing"

func getRedBlackTests() []orderedSymbolTableTest[string, int] {
	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "LLRB Tree"
	tests[0].expectedHeight = 1
	tests[0].expectedPreOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedInOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedPostOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedDotCode = `strict digraph RedBlack {
  node [style=filled, shape=oval];

  0 [label=",0", color=black, fontcolor=white];
}`

	tests[1].symbolTable = "LLRB Tree"
	tests[1].expectedHeight = 2
	tests[1].expectedPreOrderTraverse = []KeyValue[string, int]{{"B", 2}, {"A", 1}, {"C", 3}}
	tests[1].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[1].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"B", 2}}
	tests[1].expectedDotCode = `strict digraph RedBlack {
  node [style=filled, shape=oval];

  1 [label="B,2", color=black, fontcolor=white];
  0 [label="A,1", color=black, fontcolor=white];
  2 [label="C,3", color=black, fontcolor=white];

  1 -> 0 [color=black];
  1 -> 2 [];
}`

	tests[2].symbolTable = "LLRB Tree"
	tests[2].expectedHeight = 3
	tests[2].expectedPreOrderTraverse = []KeyValue[string, int]{{"D", 4}, {"B", 2}, {"A", 1}, {"C", 3}, {"E", 5}}
	tests[2].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[2].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"B", 2}, {"E", 5}, {"D", 4}}
	tests[2].expectedDotCode = `strict digraph RedBlack {
  node [style=filled, shape=oval];

  3 [label="D,4", color=black, fontcolor=white];
  1 [label="B,2", color=red, fontcolor=white];
  0 [label="A,1", color=black, fontcolor=white];
  2 [label="C,3", color=black, fontcolor=white];
  4 [label="E,5", color=black, fontcolor=white];

  3 -> 1 [color=red];
  3 -> 4 [];
  1 -> 0 [color=black];
  1 -> 2 [];
}`

	tests[3].symbolTable = "LLRB Tree"
	tests[3].expectedHeight = 3
	tests[3].expectedPreOrderTraverse = []KeyValue[string, int]{{"J", 10}, {"D", 4}, {"A", 1}, {"G", 7}, {"P", 16}, {"M", 13}, {"S", 19}}
	tests[3].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[3].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"G", 7}, {"D", 4}, {"M", 13}, {"S", 19}, {"P", 16}, {"J", 10}}
	tests[3].expectedDotCode = `strict digraph RedBlack {
  node [style=filled, shape=oval];

  3 [label="J,10", color=black, fontcolor=white];
  1 [label="D,4", color=black, fontcolor=white];
  0 [label="A,1", color=black, fontcolor=white];
  2 [label="G,7", color=black, fontcolor=white];
  5 [label="P,16", color=black, fontcolor=white];
  4 [label="M,13", color=black, fontcolor=white];
  6 [label="S,19", color=black, fontcolor=white];

  3 -> 1 [color=black];
  3 -> 5 [];
  1 -> 0 [color=black];
  1 -> 2 [];
  5 -> 4 [color=black];
  5 -> 6 [];
}`

	return tests
}

func TestRedBlack(t *testing.T) {
	tests := getRedBlackTests()

	for _, tc := range tests {
		rbt := NewRedBlack[string, int](tc.cmpKey)
		runOrderedSymbolTableTest(t, rbt, tc)
	}
}
