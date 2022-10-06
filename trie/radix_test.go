package trie

import (
	"testing"
)

func getRadixTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].symbolTable = "Radix"
	tests[0].expectedHeight = 0
	tests[0].expectedVLRTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedVRLTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedLVRTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedRVLTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedLRVTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedRLVTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedAscendingTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedDescendingTraverse = []KeyValue[int]{{"", 0}}
	tests[0].expectedDotCode = ``

	tests[1].symbolTable = "Radix"
	tests[1].expectedHeight = 0
	tests[1].expectedVLRTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedVRLTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedLVRTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedRVLTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedLRVTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedRLVTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedAscendingTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedDescendingTraverse = []KeyValue[int]{{"", 0}}
	tests[1].expectedDotCode = ``

	tests[2].symbolTable = "Radix"
	tests[2].expectedHeight = 0
	tests[2].expectedVLRTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedVRLTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedLVRTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedRVLTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedLRVTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedRLVTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedAscendingTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedDescendingTraverse = []KeyValue[int]{{"", 0}}
	tests[2].expectedDotCode = ``

	tests[3].symbolTable = "Radix"
	tests[3].expectedHeight = 0
	tests[3].expectedVLRTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedVRLTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedLVRTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedRVLTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedLRVTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedRLVTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedAscendingTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedDescendingTraverse = []KeyValue[int]{{"", 0}}
	tests[3].expectedDotCode = ``

	return tests
}

func TestRadix(t *testing.T) {

}
