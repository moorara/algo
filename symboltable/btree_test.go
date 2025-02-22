package symboltable

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

var M = []int{4, 4, 4, 4}

func getBTreeTests() []orderedSymbolTableTest[string, int] {
	cmpKey := NewCompareFunc[string]()
	eqVal := NewEqualFunc[int]()

	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "LLRB Tree"
	tests[0].expectedHeight = 2
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDOT = ``

	tests[1].symbolTable = "LLRB Tree"
	tests[1].expectedHeight = 3
	tests[1].equals = NewBTree[string, int](M[1], cmpKey, eqVal)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDOT = ``

	tests[2].symbolTable = "LLRB Tree"
	tests[2].expectedHeight = 3
	tests[2].equals = NewBTree[string, int](M[2], cmpKey, eqVal)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 4}, {Key: "A", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 1}, {Key: "S", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 4}, {Key: "G", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 6}, {Key: "M", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "D", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 3}, {Key: "P", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "0", Val: 0}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "0", Val: 0}}
	tests[2].expectedDOT = ``

	tests[3].symbolTable = "LLRB Tree"
	tests[3].expectedHeight = 3
	tests[3].equals = NewBTree[string, int](M[3], cmpKey, eqVal)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
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
