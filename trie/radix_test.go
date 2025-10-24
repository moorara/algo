package trie

import (
	"testing"

	. "github.com/moorara/algo/generic"
)

func getRadixTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].trie = "Radix Trie"
	tests[0].expectedHeight = 3
	tests[0].equals = nil
	tests[0].expectedEquals = false
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[0].expectedDOT = ``

	tests[1].trie = "Radix Trie"
	tests[1].expectedHeight = 5
	tests[1].equals = NewRadix[int](nil)
	tests[1].expectedEquals = false
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[1].expectedDOT = ``

	tests[2].trie = "Radix Trie"
	tests[2].expectedHeight = 7
	tests[2].equals = NewRadix[int](nil)
	tests[2].expectedEquals = false
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[2].expectedDOT = ``

	tests[3].trie = "Radix Trie"
	tests[3].expectedHeight = 8
	tests[3].equals = NewRadix[int](nil)
	tests[3].expectedEquals = true
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 7}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLVRTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRVLTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}, {Key: "", Val: 0}}
	tests[3].expectedDOT = ``

	return tests
}

func TestRadixTrie(t *testing.T) {
	tests := getRadixTests()

	for _, tc := range tests {
		bin := NewRadix[int](tc.eqVal)
		runTrieTest(t, bin, tc)
	}
}
