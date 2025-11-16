package symboltable

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

func getQuadHashTableTests() []symbolTableTest[string, int] {
	hashFunc := hash.HashFuncForString[string](nil)
	eqKey := generic.NewEqualFunc[string]()
	eqVal := generic.NewEqualFunc[int]()
	opts := HashOpts{}

	tests := getSymbolTableTests()

	tests[0].symbolTable = "Quadratic Probing Hash Table"
	tests[0].equal = NewQuadraticHashTable(hashFunc, eqKey, eqVal, opts)
	tests[0].equal.Put("Apple", 182)
	tests[0].equal.Put("Avocado", 200)
	tests[0].equal.Put("Banana", 120)
	tests[0].equal.Put("Coconut", 1500)
	tests[0].equal.Put("Dragon Fruit", 600)
	tests[0].equal.Put("Durian", 1500)
	tests[0].equal.Put("Guava", 180)
	tests[0].equal.Put("Kiwi", 75)
	tests[0].equal.Put("Lychee", 20)
	tests[0].equal.Put("Mango", 200)
	tests[0].equal.Put("Orange", 130)
	tests[0].equal.Put("Papaya", 1000)
	tests[0].equal.Put("Passion Fruit", 40)
	tests[0].equal.Put("Pineapple", 1200)
	tests[0].equal.Put("Watermelon", 9000)
	tests[0].expectedEqual = true

	tests[1].symbolTable = "Quadratic Probing Hash Table"
	tests[1].equal = NewQuadraticHashTable(hashFunc, eqKey, eqVal, opts)
	tests[1].equal.Put("Golden Pheasant", 15)
	tests[1].equal.Put("Harpy Eagle", 35)
	tests[1].equal.Put("Kingfisher", 15)
	tests[1].equal.Put("Mandarin Duck", 10)
	tests[1].equal.Put("Peacock", 20)
	tests[1].equal.Put("Quetzal", 25)
	tests[1].equal.Put("Scarlet Macaw", 50)
	tests[1].equal.Put("Snowy Owl", 10)
	tests[1].expectedEqual = true

	tests[2].symbolTable = "Quadratic Probing Hash Table"
	tests[2].equal = NewQuadraticHashTable(hashFunc, eqKey, eqVal, opts)
	tests[2].equal.Put("Accordion", 50)
	tests[2].equal.Put("Bassoon", 140)
	tests[2].equal.Put("Cello", 120)
	tests[2].equal.Put("Clarinet", 66)
	tests[2].equal.Put("Double Bass", 180)
	tests[2].equal.Put("Drum Set", 200)
	tests[2].equal.Put("Flute", 67)
	tests[2].equal.Put("Guitar", 100)
	tests[2].equal.Put("Harp", 170)
	tests[2].equal.Put("Organ", 300) // Extra
	tests[2].equal.Put("Piano", 150)
	tests[2].equal.Put("Saxophone", 80)
	tests[2].equal.Put("Trombone", 120)
	tests[2].equal.Put("Trumpet", 48)
	tests[2].equal.Put("Ukulele", 60)
	tests[2].equal.Put("Violin", 60)
	tests[2].expectedEqual = false

	tests[3].symbolTable = "Quadratic Probing Hash Table"
	tests[3].equal = NewQuadraticHashTable(hashFunc, eqKey, eqVal, opts)
	tests[3].equal.Put("Berlin", 10)
	// tests[3].equal.Put("London", 11)
	tests[3].equal.Put("Montreal", 6)
	tests[3].equal.Put("New York", 13)
	tests[3].equal.Put("Paris", 12)
	tests[3].equal.Put("Rome", 16)
	tests[3].equal.Put("Tehran", 17)
	tests[3].equal.Put("Tokyo", 16)
	tests[3].equal.Put("Toronto", 8)
	tests[3].equal.Put("Vancouver", 10)
	tests[3].expectedEqual = false

	tests[4].symbolTable = "Separate Chaining Hash Table"
	tests[4].equal = nil
	tests[4].expectedEqual = false

	return tests
}

func TestQuadraticHashTable(t *testing.T) {
	tests := getQuadHashTableTests()

	for _, tc := range tests {
		ht := NewQuadraticHashTable(tc.hashKey, tc.eqKey, tc.eqVal, tc.opts)
		runSymbolTableTest(t, ht, tc)
	}
}

func TestNewQuadraticHashTable_Panic(t *testing.T) {
	hashString := hash.HashFuncForString[string](nil)
	eqString := generic.NewEqualFunc[string]()
	eqInt := generic.NewEqualFunc[int]()

	assert.PanicsWithValue(t, "The hash table capacity must be a prime number", func() {
		NewQuadraticHashTable(hashString, eqString, eqInt, HashOpts{
			InitialCap: 69,
		})
	})
}
