package symboltable

import (
	"testing"

	"github.com/moorara/algo/generic"
)

var M = []int{4, 4, 4, 4}

func getBTreeTests() []orderedSymbolTableTest[string, int] {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "LLRB Tree"
	tests[0].expectedHeight = 2
	tests[0].equal = nil
	tests[0].expectedEqual = false
	tests[0].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDOT = ``

	tests[1].symbolTable = "LLRB Tree"
	tests[1].expectedHeight = 3
	tests[1].equal = NewBTree(M[1], cmpKey, eqVal)
	tests[1].expectedEqual = false
	tests[1].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDOT = ``

	tests[2].symbolTable = "LLRB Tree"
	tests[2].expectedHeight = 3
	tests[2].equal = NewBTree(M[2], cmpKey, eqVal)
	tests[2].expectedEqual = false
	tests[2].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 4}, {Key: "A", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 1}, {Key: "S", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 4}, {Key: "G", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 6}, {Key: "M", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "D", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 3}, {Key: "P", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "0", Val: 0}}
	tests[2].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "0", Val: 0}}
	tests[2].expectedDOT = ``

	tests[3].symbolTable = "LLRB Tree"
	tests[3].expectedHeight = 3
	tests[3].equal = NewBTree(M[3], cmpKey, eqVal)
	tests[3].expectedEqual = true
	tests[3].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDOT = ``

	return tests
}

func TestBTree(t *testing.T) {
	tests := getBTreeTests()

	for i, tc := range tests {
		rbt := NewBTree[string, int](M[i], tc.cmpKey, tc.eqVal)
		runOrderedSymbolTableTest(t, rbt, tc)
	}
}
