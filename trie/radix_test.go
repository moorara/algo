package trie

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func getRadixTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].trie = "Radix Trie"
	tests[0].expectedHeight = 3
	tests[0].equal = nil
	tests[0].expectedEqual = false
	tests[0].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDOT = ``

	tests[1].trie = "Radix Trie"
	tests[1].expectedHeight = 5
	tests[1].equal = NewRadix[int](nil)
	tests[1].expectedEqual = false
	tests[1].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDOT = ``

	tests[2].trie = "Radix Trie"
	tests[2].expectedHeight = 7
	tests[2].equal = NewRadix[int](nil)
	tests[2].expectedEqual = false
	tests[2].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedDOT = ``

	tests[3].trie = "Radix Trie"
	tests[3].expectedHeight = 8
	tests[3].equal = NewRadix[int](nil)
	tests[3].expectedEqual = true
	tests[3].expectedVLRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedVRLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 7}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLVRTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRVLTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLRVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRLVTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedAscendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDescendingTraverse = []generic.KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDOT = ``

	return tests
}

func TestRadixTrie(t *testing.T) {
	tests := getRadixTests()

	for _, tc := range tests {
		bin := NewRadix(tc.eqVal)
		runTrieTest(t, bin, tc)
	}
}
