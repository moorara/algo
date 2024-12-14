package symboltable

import (
	"testing"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

func getLinearHashTableTests() []symbolTableTest[string, int] {
	hashFunc := HashFuncForString[string](nil)
	eqKey := NewEqualFunc[string]()
	eqVal := NewEqualFunc[int]()
	opts := HashOpts{}

	tests := getSymbolTableTests()

	tests[0].symbolTable = "Linear Probing Hash Table"
	tests[0].equals = NewLinearHashTable[string, int](hashFunc, eqKey, eqVal, opts)
	tests[0].equals.Put("Apple", 182)
	tests[0].equals.Put("Avocado", 200)
	tests[0].equals.Put("Banana", 120)
	tests[0].equals.Put("Coconut", 1500)
	tests[0].equals.Put("Dragon Fruit", 600)
	tests[0].equals.Put("Durian", 1500)
	tests[0].equals.Put("Guava", 180)
	tests[0].equals.Put("Kiwi", 75)
	tests[0].equals.Put("Lychee", 20)
	tests[0].equals.Put("Mango", 200)
	tests[0].equals.Put("Orange", 130)
	tests[0].equals.Put("Papaya", 1000)
	tests[0].equals.Put("Passion Fruit", 40)
	tests[0].equals.Put("Pineapple", 1200)
	tests[0].equals.Put("Watermelon", 9000)
	tests[0].expectedEquals = true

	tests[1].symbolTable = "Linear Probing Hash Table"
	tests[1].equals = NewLinearHashTable[string, int](hashFunc, eqKey, eqVal, opts)
	tests[1].equals.Put("Golden Pheasant", 15)
	tests[1].equals.Put("Harpy Eagle", 35)
	tests[1].equals.Put("Kingfisher", 15)
	tests[1].equals.Put("Mandarin Duck", 10)
	tests[1].equals.Put("Peacock", 20)
	tests[1].equals.Put("Quetzal", 25)
	tests[1].equals.Put("Scarlet Macaw", 50)
	tests[1].equals.Put("Snowy Owl", 10)
	tests[1].expectedEquals = true

	tests[2].symbolTable = "Linear Probing Hash Table"
	tests[2].equals = NewLinearHashTable[string, int](hashFunc, eqKey, eqVal, opts)
	tests[2].equals.Put("Accordion", 50)
	tests[2].equals.Put("Bassoon", 140)
	tests[2].equals.Put("Cello", 120)
	tests[2].equals.Put("Clarinet", 66)
	tests[2].equals.Put("Double Bass", 180)
	tests[2].equals.Put("Drum Set", 200)
	tests[2].equals.Put("Flute", 67)
	tests[2].equals.Put("Guitar", 100)
	tests[2].equals.Put("Harp", 170)
	tests[2].equals.Put("Organ", 300) // Extra
	tests[2].equals.Put("Piano", 150)
	tests[2].equals.Put("Saxophone", 80)
	tests[2].equals.Put("Trombone", 120)
	tests[2].equals.Put("Trumpet", 48)
	tests[2].equals.Put("Ukulele", 60)
	tests[2].equals.Put("Violin", 60)
	tests[2].expectedEquals = false

	tests[3].symbolTable = "Linear Probing Hash Table"
	tests[3].equals = NewLinearHashTable[string, int](hashFunc, eqKey, eqVal, opts)
	tests[3].equals.Put("Berlin", 10)
	// tests[3].equals.Put("London", 11)
	tests[3].equals.Put("Montreal", 6)
	tests[3].equals.Put("New York", 13)
	tests[3].equals.Put("Paris", 12)
	tests[3].equals.Put("Rome", 16)
	tests[3].equals.Put("Tehran", 17)
	tests[3].equals.Put("Tokyo", 16)
	tests[3].equals.Put("Toronto", 8)
	tests[3].equals.Put("Vancouver", 10)
	tests[3].expectedEquals = false

	tests[4].symbolTable = "Separate Chaining Hash Table"
	tests[4].equals = nil
	tests[4].expectedEquals = false

	return tests
}

func TestLinearHashTable(t *testing.T) {
	tests := getLinearHashTableTests()

	for _, tc := range tests {
		ht := NewLinearHashTable[string, int](tc.hashKey, tc.eqKey, tc.eqVal, tc.opts)
		runSymbolTableTest(t, ht, tc)
	}
}
