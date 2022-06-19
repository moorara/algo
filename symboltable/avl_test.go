package symboltable

import "testing"

func getAVLTests() []orderedSymbolTableTest[string, int] {
	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "AVL"
	tests[0].expectedHeight = 1
	tests[0].expectedPreOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedInOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedPostOrderTraverse = []KeyValue[string, int]{{"", 0}}
	tests[0].expectedDotCode = `strict digraph AVL {
  node [shape=oval];

  0 [label=",0"];
}`

	tests[1].symbolTable = "AVL"
	tests[1].expectedHeight = 2
	tests[1].expectedPreOrderTraverse = []KeyValue[string, int]{{"B", 2}, {"A", 1}, {"C", 3}}
	tests[1].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[1].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"B", 2}}
	tests[1].expectedDotCode = `strict digraph AVL {
  node [shape=oval];

  1 [label="B,2"];
  0 [label="A,1"];
  2 [label="C,3"];

  1 -> 0 [];
  1 -> 2 [];
}`

	tests[2].symbolTable = "AVL"
	tests[2].expectedHeight = 3
	tests[2].expectedPreOrderTraverse = []KeyValue[string, int]{{"B", 2}, {"A", 1}, {"D", 4}, {"C", 3}, {"E", 5}}
	tests[2].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[2].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"C", 3}, {"E", 5}, {"D", 4}, {"B", 2}}
	tests[2].expectedDotCode = `strict digraph AVL {
  node [shape=oval];

  1 [label="B,2"];
  0 [label="A,1"];
  3 [label="D,4"];
  2 [label="C,3"];
  4 [label="E,5"];

  1 -> 0 [];
  1 -> 3 [];
  3 -> 2 [];
  3 -> 4 [];
}`

	tests[3].symbolTable = "AVL"
	tests[3].expectedHeight = 3
	tests[3].expectedPreOrderTraverse = []KeyValue[string, int]{{"J", 10}, {"D", 4}, {"A", 1}, {"G", 7}, {"P", 16}, {"M", 13}, {"S", 19}}
	tests[3].expectedInOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[3].expectedPostOrderTraverse = []KeyValue[string, int]{{"A", 1}, {"G", 7}, {"D", 4}, {"M", 13}, {"S", 19}, {"P", 16}, {"J", 10}}
	tests[3].expectedDotCode = `strict digraph AVL {
  node [shape=oval];

  3 [label="J,10"];
  1 [label="D,4"];
  0 [label="A,1"];
  2 [label="G,7"];
  5 [label="P,16"];
  4 [label="M,13"];
  6 [label="S,19"];

  3 -> 1 [];
  3 -> 5 [];
  1 -> 0 [];
  1 -> 2 [];
  5 -> 4 [];
  5 -> 6 [];
}`

	return tests
}

func TestAVL(t *testing.T) {
	tests := getAVLTests()

	for _, tc := range tests {
		avl := NewAVL[string, int](tc.cmpKey)
		runOrderedSymbolTableTest(t, avl, tc)
	}
}
