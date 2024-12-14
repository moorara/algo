package symboltable

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

type symbolTableTest[K, V any] struct {
	name              string
	symbolTable       string
	hashKey           HashFunc[K]
	eqKey             EqualFunc[K]
	eqVal             EqualFunc[V]
	opts              HashOpts
	keyVals           []KeyValue[K, V]
	expectedSize      int
	expectedIsEmpty   bool
	expectedStrings   []string
	equals            SymbolTable[K, V]
	expectedEquals    bool
	expectedAll       []KeyValue[K, V]
	anyMatchPredicate Predicate2[K, V]
	expectedAnyMatch  bool
	allMatchPredicate Predicate2[K, V]
	expectedAllMatch  bool
}

type orderedSymbolTableTest[K, V any] struct {
	name            string
	symbolTable     string
	cmpKey          CompareFunc[K]
	eqVal           EqualFunc[V]
	keyVals         []KeyValue[K, V]
	expectedSize    int
	expectedHeight  int
	expectedIsEmpty bool

	expectedMinKey             K
	expectedMinVal             V
	expectedMinOK              bool
	expectedMaxKey             K
	expectedMaxVal             V
	expectedMaxOK              bool
	floorKey                   string
	expectedFloorKey           K
	expectedFloorVal           V
	expectedFloorOK            bool
	ceilingKey                 string
	expectedCeilingKey         K
	expectedCeilingVal         V
	expectedCeilingOK          bool
	selectRank                 int
	expectedSelectKey          K
	expectedSelectVal          V
	expectedSelectOK           bool
	rankKey                    string
	expectedRank               int
	rangeKeyLo                 string
	rangeKeyHi                 string
	expectedRange              []KeyValue[K, V]
	expectedRangeSize          int
	expectedString             string
	equals                     SymbolTable[K, V]
	expectedEquals             bool
	expectedAll                []KeyValue[K, V]
	anyMatchPredicate          Predicate2[K, V]
	expectedAnyMatch           bool
	allMatchPredicate          Predicate2[K, V]
	expectedAllMatch           bool
	expectedVLRTraverse        []KeyValue[K, V]
	expectedVRLTraverse        []KeyValue[K, V]
	expectedLVRTraverse        []KeyValue[K, V]
	expectedRVLTraverse        []KeyValue[K, V]
	expectedLRVTraverse        []KeyValue[K, V]
	expectedRLVTraverse        []KeyValue[K, V]
	expectedAscendingTraverse  []KeyValue[K, V]
	expectedDescendingTraverse []KeyValue[K, V]
	expectedGraphviz           string
}

func getSymbolTableTests() []symbolTableTest[string, int] {
	hashKey := HashFuncForString[string](nil)
	eqKey := NewEqualFunc[string]()
	eqVal := NewEqualFunc[int]()

	return []symbolTableTest[string, int]{
		{
			name:    "FruitWeight",
			hashKey: hashKey,
			eqKey:   eqKey,
			eqVal:   eqVal,
			opts:    HashOpts{},
			keyVals: []KeyValue[string, int]{
				{Key: "Apple", Val: 182},
				{Key: "Banana", Val: 120},
				{Key: "Mango", Val: 200},
				{Key: "Pineapple", Val: 1200},
				{Key: "Papaya", Val: 1000},
				{Key: "Kiwi", Val: 75},
				{Key: "Orange", Val: 130},
				{Key: "Guava", Val: 180},
				{Key: "Dragon Fruit", Val: 600},
				{Key: "Coconut", Val: 1500},
				{Key: "Lychee", Val: 20},
				{Key: "Durian", Val: 1500},
				{Key: "Passion Fruit", Val: 40},
				{Key: "Watermelon", Val: 9000},
				{Key: "Avocado", Val: 200},
			},
			expectedSize:    15,
			expectedIsEmpty: false,
			expectedStrings: []string{
				"<Apple:182>",
				"<Avocado:200>",
				"<Banana:120>",
				"<Coconut:1500>",
				"<Dragon Fruit:600>",
				"<Durian:1500>",
				"<Guava:180>",
				"<Kiwi:75>",
				"<Lychee:20>",
				"<Mango:200>",
				"<Orange:130>",
				"<Papaya:1000>",
				"<Passion Fruit:40>",
				"<Pineapple:1200>",
				"<Watermelon:9000>",
			},
			expectedAll: []KeyValue[string, int]{
				{Key: "Apple", Val: 182},
				{Key: "Avocado", Val: 200},
				{Key: "Banana", Val: 120},
				{Key: "Coconut", Val: 1500},
				{Key: "Dragon Fruit", Val: 600},
				{Key: "Durian", Val: 1500},
				{Key: "Guava", Val: 180},
				{Key: "Kiwi", Val: 75},
				{Key: "Lychee", Val: 20},
				{Key: "Mango", Val: 200},
				{Key: "Orange", Val: 130},
				{Key: "Papaya", Val: 1000},
				{Key: "Passion Fruit", Val: 40},
				{Key: "Pineapple", Val: 1200},
				{Key: "Watermelon", Val: 9000},
			},
			anyMatchPredicate: func(k string, v int) bool { return k == "Sour Cherry" },
			expectedAnyMatch:  false,
			allMatchPredicate: func(k string, v int) bool { return v > 100 },
			expectedAllMatch:  false,
		},
		{
			name:    "BirdLifespan",
			hashKey: hashKey,
			eqKey:   eqKey,
			eqVal:   eqVal,
			opts:    HashOpts{},
			keyVals: []KeyValue[string, int]{
				{Key: "Peacock", Val: 20},
				{Key: "Scarlet Macaw", Val: 50},
				{Key: "Golden Pheasant", Val: 15},
				{Key: "Mandarin Duck", Val: 10},
				{Key: "Harpy Eagle", Val: 35},
				{Key: "Kingfisher", Val: 15},
				{Key: "Snowy Owl", Val: 10},
				{Key: "Quetzal", Val: 25},
			},
			expectedSize:    8,
			expectedIsEmpty: false,
			expectedStrings: []string{
				"<Golden Pheasant:15>",
				"<Harpy Eagle:35>",
				"<Kingfisher:15>",
				"<Mandarin Duck:10>",
				"<Peacock:20>",
				"<Quetzal:25>",
				"<Scarlet Macaw:50>",
				"<Snowy Owl:10>",
			},
			expectedAll: []KeyValue[string, int]{
				{Key: "Golden Pheasant", Val: 15},
				{Key: "Harpy Eagle", Val: 35},
				{Key: "Kingfisher", Val: 15},
				{Key: "Mandarin Duck", Val: 10},
				{Key: "Peacock", Val: 20},
				{Key: "Quetzal", Val: 25},
				{Key: "Scarlet Macaw", Val: 50},
				{Key: "Snowy Owl", Val: 10},
			},
			anyMatchPredicate: func(k string, v int) bool { return k == "Cardinal" },
			expectedAnyMatch:  false,
			allMatchPredicate: func(k string, v int) bool { return v >= 10 },
			expectedAllMatch:  true,
		},
		{
			name:    "InstrumentLength",
			hashKey: hashKey,
			eqKey:   eqKey,
			eqVal:   eqVal,
			opts:    HashOpts{},
			keyVals: []KeyValue[string, int]{
				{Key: "Violin", Val: 60},
				{Key: "Guitar", Val: 100},
				{Key: "Piano", Val: 150},
				{Key: "Flute", Val: 67},
				{Key: "Trumpet", Val: 48},
				{Key: "Drum Set", Val: 200},
				{Key: "Saxophone", Val: 80},
				{Key: "Clarinet", Val: 66},
				{Key: "Cello", Val: 120},
				{Key: "Double Bass", Val: 180},
				{Key: "Harp", Val: 170},
				{Key: "Trombone", Val: 120},
				{Key: "Bassoon", Val: 140},
				{Key: "Ukulele", Val: 60},
				{Key: "Accordion", Val: 50},
			},
			expectedSize:    15,
			expectedIsEmpty: false,
			expectedStrings: []string{
				"<Accordion:50>",
				"<Bassoon:140>",
				"<Cello:120>",
				"<Clarinet:66>",
				"<Double Bass:180>",
				"<Drum Set:200>",
				"<Flute:67>",
				"<Guitar:100>",
				"<Harp:170>",
				"<Piano:150>",
				"<Saxophone:80>",
				"<Trombone:120>",
				"<Trumpet:48>",
				"<Ukulele:60>",
				"<Violin:60>",
			},
			expectedAll: []KeyValue[string, int]{
				{Key: "Accordion", Val: 50},
				{Key: "Bassoon", Val: 140},
				{Key: "Cello", Val: 120},
				{Key: "Clarinet", Val: 66},
				{Key: "Double Bass", Val: 180},
				{Key: "Drum Set", Val: 200},
				{Key: "Flute", Val: 67},
				{Key: "Guitar", Val: 100},
				{Key: "Harp", Val: 170},
				{Key: "Piano", Val: 150},
				{Key: "Saxophone", Val: 80},
				{Key: "Trombone", Val: 120},
				{Key: "Trumpet", Val: 48},
				{Key: "Ukulele", Val: 60},
				{Key: "Violin", Val: 60},
			},
			anyMatchPredicate: func(k string, v int) bool { return k == "Saxophone" },
			expectedAnyMatch:  true,
			allMatchPredicate: func(k string, v int) bool { return v < 100 },
			expectedAllMatch:  false,
		},
		{
			name:    "CityTemperature",
			hashKey: hashKey,
			eqKey:   eqKey,
			eqVal:   eqVal,
			opts:    HashOpts{},
			keyVals: []KeyValue[string, int]{
				{Key: "Toronto", Val: 8},
				{Key: "Montreal", Val: 6},
				{Key: "Vancouver", Val: 10},
				{Key: "New York", Val: 13},
				{Key: "London", Val: 11},
				{Key: "Paris", Val: 12},
				{Key: "Rome", Val: 16},
				{Key: "Berlin", Val: 10},
				{Key: "Tokyo", Val: 16},
				{Key: "Tehran", Val: 17},
			},
			expectedSize:    10,
			expectedIsEmpty: false,
			expectedStrings: []string{
				"<Berlin:10>",
				"<London:11>",
				"<Montreal:6>",
				"<New York:13>",
				"<Paris:12>",
				"<Rome:16>",
				"<Tehran:17>",
				"<Tokyo:16>",
				"<Toronto:8>",
				"<Vancouver:10>",
			},
			expectedAll: []KeyValue[string, int]{
				{Key: "Berlin", Val: 10},
				{Key: "London", Val: 11},
				{Key: "Montreal", Val: 6},
				{Key: "New York", Val: 13},
				{Key: "Paris", Val: 12},
				{Key: "Rome", Val: 16},
				{Key: "Tehran", Val: 17},
				{Key: "Tokyo", Val: 16},
				{Key: "Toronto", Val: 8},
				{Key: "Vancouver", Val: 10},
			},
			anyMatchPredicate: func(k string, v int) bool { return k == "Toronto" },
			expectedAnyMatch:  true,
			allMatchPredicate: func(k string, v int) bool { return v > 4 && v < 24 },
			expectedAllMatch:  true,
		},
		{
			name:    "AnimalLifespan",
			hashKey: hashKey,
			eqKey:   eqKey,
			eqVal:   eqVal,
			opts:    HashOpts{},
			keyVals: []KeyValue[string, int]{
				{Key: "Elephant", Val: 70},
				{Key: "Blue Whale", Val: 80},
				{Key: "Galapagos Tortoise", Val: 100},
				{Key: "Macaw", Val: 60},
				{Key: "Bald Eagle", Val: 20},
				{Key: "Horse", Val: 25},
				{Key: "Dog", Val: 13},
				{Key: "Cat", Val: 15},
				{Key: "Chimpanzee", Val: 40},
				{Key: "Rabbit", Val: 9},
				{Key: "Goldfish", Val: 10},
				{Key: "Parrot", Val: 50},
				{Key: "Kangaroo", Val: 23},
				{Key: "Lion", Val: 14},
				{Key: "Tiger", Val: 16},
				{Key: "Giraffe", Val: 26},
				{Key: "Penguin", Val: 20},
				{Key: "Wolf", Val: 14},
				{Key: "Zebra", Val: 25},
				{Key: "Cheetah", Val: 12},
				{Key: "Dolphin", Val: 40},
				{Key: "Polar Bear", Val: 25},
				{Key: "Brown Bear", Val: 30},
				{Key: "Crocodile", Val: 70},
				{Key: "Shark", Val: 30},
				{Key: "Frog", Val: 10},
				{Key: "Salamander", Val: 20},
				{Key: "Tarantula", Val: 20},
				{Key: "Owl", Val: 25},
				{Key: "Swan", Val: 30},
				{Key: "Peacock", Val: 20},
				{Key: "Raven", Val: 15},
				{Key: "Snake", Val: 20},
				{Key: "Lizard", Val: 10},
				{Key: "Hamster", Val: 3},
				{Key: "Guinea Pig", Val: 6},
				{Key: "Ferret", Val: 7},
				{Key: "Hedgehog", Val: 5},
				{Key: "Bat", Val: 30},
				{Key: "Koala", Val: 20},
				{Key: "Platypus", Val: 17},
				{Key: "Octopus", Val: 5},
				{Key: "Crab", Val: 10},
				{Key: "Lobster", Val: 50},
				{Key: "Starfish", Val: 35},
				{Key: "Sea Turtle", Val: 100},
				{Key: "Jellyfish", Val: 1},
				{Key: "Ant", Val: 1},
				{Key: "Bee", Val: 5},
				{Key: "Butterfly", Val: 1},
			},
			expectedSize:    50,
			expectedIsEmpty: false,
			expectedStrings: []string{
				"<Ant:1>",
				"<Bald Eagle:20>",
				"<Bat:30>",
				"<Bee:5>",
				"<Blue Whale:80>",
				"<Brown Bear:30>",
				"<Butterfly:1>",
				"<Cat:15>",
				"<Cheetah:12>",
				"<Chimpanzee:40>",
				"<Crab:10>",
				"<Crocodile:70>",
				"<Dog:13>",
				"<Dolphin:40>",
				"<Elephant:70>",
				"<Ferret:7>",
				"<Frog:10>",
				"<Galapagos Tortoise:100>",
				"<Giraffe:26>",
				"<Goldfish:10>",
				"<Guinea Pig:6>",
				"<Hamster:3>",
				"<Hedgehog:5>",
				"<Horse:25>",
				"<Jellyfish:1>",
				"<Kangaroo:23>",
				"<Koala:20>",
				"<Lion:14>",
				"<Lizard:10>",
				"<Lobster:50>",
				"<Macaw:60>",
				"<Octopus:5>",
				"<Owl:25>",
				"<Parrot:50>",
				"<Peacock:20>",
				"<Penguin:20>",
				"<Platypus:17>",
				"<Polar Bear:25>",
				"<Rabbit:9>",
				"<Raven:15>",
				"<Salamander:20>",
				"<Sea Turtle:100>",
				"<Shark:30>",
				"<Snake:20>",
				"<Starfish:35>",
				"<Swan:30>",
				"<Tarantula:20>",
				"<Tiger:16>",
				"<Wolf:14>",
				"<Zebra:25>",
			},
			expectedAll: []KeyValue[string, int]{
				{Key: "Ant", Val: 1},
				{Key: "Bald Eagle", Val: 20},
				{Key: "Bat", Val: 30},
				{Key: "Bee", Val: 5},
				{Key: "Blue Whale", Val: 80},
				{Key: "Brown Bear", Val: 30},
				{Key: "Butterfly", Val: 1},
				{Key: "Cat", Val: 15},
				{Key: "Cheetah", Val: 12},
				{Key: "Chimpanzee", Val: 40},
				{Key: "Crab", Val: 10},
				{Key: "Crocodile", Val: 70},
				{Key: "Dog", Val: 13},
				{Key: "Dolphin", Val: 40},
				{Key: "Elephant", Val: 70},
				{Key: "Ferret", Val: 7},
				{Key: "Frog", Val: 10},
				{Key: "Galapagos Tortoise", Val: 100},
				{Key: "Giraffe", Val: 26},
				{Key: "Goldfish", Val: 10},
				{Key: "Guinea Pig", Val: 6},
				{Key: "Hamster", Val: 3},
				{Key: "Hedgehog", Val: 5},
				{Key: "Horse", Val: 25},
				{Key: "Jellyfish", Val: 1},
				{Key: "Kangaroo", Val: 23},
				{Key: "Koala", Val: 20},
				{Key: "Lion", Val: 14},
				{Key: "Lizard", Val: 10},
				{Key: "Lobster", Val: 50},
				{Key: "Macaw", Val: 60},
				{Key: "Octopus", Val: 5},
				{Key: "Owl", Val: 25},
				{Key: "Parrot", Val: 50},
				{Key: "Peacock", Val: 20},
				{Key: "Penguin", Val: 20},
				{Key: "Platypus", Val: 17},
				{Key: "Polar Bear", Val: 25},
				{Key: "Rabbit", Val: 9},
				{Key: "Raven", Val: 15},
				{Key: "Salamander", Val: 20},
				{Key: "Sea Turtle", Val: 100},
				{Key: "Shark", Val: 30},
				{Key: "Snake", Val: 20},
				{Key: "Starfish", Val: 35},
				{Key: "Swan", Val: 30},
				{Key: "Tarantula", Val: 20},
				{Key: "Tiger", Val: 16},
				{Key: "Wolf", Val: 14},
				{Key: "Zebra", Val: 25},
			},
			anyMatchPredicate: func(k string, v int) bool { return k == "Platypus" },
			expectedAnyMatch:  true,
			allMatchPredicate: func(k string, v int) bool { return v >= 1 },
			expectedAllMatch:  true,
		},
	}
}

func getOrderedSymbolTableTests() []orderedSymbolTableTest[string, int] {
	cmpKey := NewCompareFunc[string]()
	eqVal := NewEqualFunc[int]()

	return []orderedSymbolTableTest[string, int]{
		{
			name:   "ABC",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{Key: "B", Val: 2},
				{Key: "A", Val: 1},
				{Key: "C", Val: 3},
			},
			expectedSize:       3,
			expectedIsEmpty:    false,
			expectedMinKey:     "A",
			expectedMinVal:     1,
			expectedMinOK:      true,
			expectedMaxKey:     "C",
			expectedMaxVal:     3,
			expectedMaxOK:      true,
			floorKey:           "A",
			expectedFloorKey:   "A",
			expectedFloorVal:   1,
			expectedFloorOK:    true,
			ceilingKey:         "C",
			expectedCeilingKey: "C",
			expectedCeilingVal: 3,
			expectedCeilingOK:  true,
			selectRank:         1,
			expectedSelectKey:  "B",
			expectedSelectVal:  2,
			expectedSelectOK:   true,
			rankKey:            "C",
			expectedRank:       2,
			rangeKeyLo:         "A",
			rangeKeyHi:         "C",
			expectedRange: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
			},
			expectedRangeSize: 3,
			expectedString:    "{<A:1> <B:2> <C:3>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
			},
			anyMatchPredicate: func(k string, v int) bool { return v < 0 },
			expectedAnyMatch:  false,
			allMatchPredicate: func(k string, v int) bool { return v%2 == 0 },
			expectedAllMatch:  false,
		},
		{
			name:   "ABCDE",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{Key: "B", Val: 2},
				{Key: "A", Val: 1},
				{Key: "C", Val: 3},
				{Key: "E", Val: 5},
				{Key: "D", Val: 4},
			},
			expectedSize:       5,
			expectedIsEmpty:    false,
			expectedMinKey:     "A",
			expectedMinVal:     1,
			expectedMinOK:      true,
			expectedMaxKey:     "E",
			expectedMaxVal:     5,
			expectedMaxOK:      true,
			floorKey:           "B",
			expectedFloorKey:   "B",
			expectedFloorVal:   2,
			expectedFloorOK:    true,
			ceilingKey:         "D",
			expectedCeilingKey: "D",
			expectedCeilingVal: 4,
			expectedCeilingOK:  true,
			selectRank:         2,
			expectedSelectKey:  "C",
			expectedSelectVal:  3,
			expectedSelectOK:   true,
			rankKey:            "E",
			expectedRank:       4,
			rangeKeyLo:         "B",
			rangeKeyHi:         "D",
			expectedRange: []KeyValue[string, int]{
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
				{Key: "D", Val: 4},
			},
			expectedRangeSize: 3,
			expectedString:    "{<A:1> <B:2> <C:3> <D:4> <E:5>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
				{Key: "D", Val: 4},
				{Key: "E", Val: 5},
			},
			anyMatchPredicate: func(k string, v int) bool { return v == 0 },
			expectedAnyMatch:  false,
			allMatchPredicate: func(k string, v int) bool { return v > 0 },
			expectedAllMatch:  true,
		},
		{
			name:   "ADGJMPS",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{Key: "J", Val: 10},
				{Key: "A", Val: 1},
				{Key: "D", Val: 4},
				{Key: "S", Val: 19},
				{Key: "P", Val: 16},
				{Key: "M", Val: 13},
				{Key: "G", Val: 7},
			},
			expectedSize:       7,
			expectedIsEmpty:    false,
			expectedMinKey:     "A",
			expectedMinVal:     1,
			expectedMinOK:      true,
			expectedMaxKey:     "S",
			expectedMaxVal:     19,
			expectedMaxOK:      true,
			floorKey:           "C",
			expectedFloorKey:   "A",
			expectedFloorVal:   1,
			expectedFloorOK:    true,
			ceilingKey:         "R",
			expectedCeilingKey: "S",
			expectedCeilingVal: 19,
			expectedCeilingOK:  true,
			selectRank:         3,
			expectedSelectKey:  "J",
			expectedSelectVal:  10,
			expectedSelectOK:   true,
			rankKey:            "S",
			expectedRank:       6,
			rangeKeyLo:         "B",
			rangeKeyHi:         "R",
			expectedRange: []KeyValue[string, int]{
				{Key: "D", Val: 4},
				{Key: "G", Val: 7},
				{Key: "J", Val: 10},
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
			},
			expectedRangeSize: 5,
			expectedString:    "{<A:1> <D:4> <G:7> <J:10> <M:13> <P:16> <S:19>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "D", Val: 4},
				{Key: "G", Val: 7},
				{Key: "J", Val: 10},
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
				{Key: "S", Val: 19},
			},
			anyMatchPredicate: func(k string, v int) bool { return v%5 == 0 },
			expectedAnyMatch:  true,
			allMatchPredicate: func(k string, v int) bool { return v < 10 },
			expectedAllMatch:  false,
		},
		{
			name:   "Words",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{Key: "box", Val: 2},
				{Key: "dad", Val: 3},
				{Key: "baby", Val: 5},
				{Key: "dome", Val: 7},
				{Key: "band", Val: 11},
				{Key: "dance", Val: 13},
				{Key: "balloon", Val: 17},
			},
			expectedSize:       7,
			expectedIsEmpty:    false,
			expectedMinKey:     "baby",
			expectedMinVal:     5,
			expectedMinOK:      true,
			expectedMaxKey:     "dome",
			expectedMaxVal:     7,
			expectedMaxOK:      true,
			floorKey:           "bold",
			expectedFloorKey:   "band",
			expectedFloorVal:   11,
			expectedFloorOK:    true,
			ceilingKey:         "breeze",
			expectedCeilingKey: "dad",
			expectedCeilingVal: 3,
			expectedCeilingOK:  true,
			selectRank:         3,
			expectedSelectKey:  "box",
			expectedSelectVal:  2,
			expectedSelectOK:   true,
			rankKey:            "dance",
			expectedRank:       5,
			rangeKeyLo:         "a",
			rangeKeyHi:         "c",
			expectedRange: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
				{Key: "box", Val: 2},
			},
			expectedRangeSize: 4,
			expectedString:    "{<baby:5> <balloon:17> <band:11> <box:2> <dad:3> <dance:13> <dome:7>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
				{Key: "box", Val: 2},
				{Key: "dad", Val: 3},
				{Key: "dance", Val: 13},
				{Key: "dome", Val: 7},
			},
			anyMatchPredicate: func(k string, v int) bool { return strings.HasSuffix(k, "x") },
			expectedAnyMatch:  true,
			allMatchPredicate: func(k string, v int) bool { return k == strings.ToLower(k) },
			expectedAllMatch:  true,
		},
	}
}

func runSymbolTableTest(t *testing.T, st SymbolTable[string, int], test symbolTableTest[string, int]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("BeforePut", func(t *testing.T) {
			assert.True(t, st.verify())
			assert.Zero(t, st.Size())
			assert.True(t, st.IsEmpty())
		})

		t.Run("AfterPut", func(t *testing.T) {
			// Put
			for _, kv := range test.keyVals {
				st.Put(kv.Key, kv.Val)
				st.Put(kv.Key, kv.Val) // Update existing key-value
				assert.True(t, st.verify())
			}

			// Get
			for _, expected := range test.keyVals {
				val, ok := st.Get(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
			}

			assert.Equal(t, test.expectedSize, st.Size())
			assert.Equal(t, test.expectedIsEmpty, st.IsEmpty())

			// The key-value pairs are unordered, so we need to compare the strings pair-wise.
			str := st.String()
			for _, s := range test.expectedStrings {
				assert.Contains(t, str, s)
			}

			equals := st.Equals(test.equals)
			assert.Equal(t, test.expectedEquals, equals)

			// The key-value pairs are unordered, so we need to compare the lists pair-wise.
			all := Collect(st.All())
			for _, kv := range test.expectedAll {
				assert.Contains(t, all, kv)
			}
			for _, kv := range all {
				assert.Contains(t, test.expectedAll, kv)
			}

			anyMatch := st.AnyMatch(test.anyMatchPredicate)
			assert.Equal(t, test.expectedAnyMatch, anyMatch)

			allMatch := st.AllMatch(test.allMatchPredicate)
			assert.Equal(t, test.expectedAllMatch, allMatch)

			// Delete a a non-existent key
			val, ok := st.Delete("NonExistentKey")
			assert.False(t, ok)
			assert.Zero(t, val)

			for _, expected := range test.keyVals {
				val, ok := st.Delete(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
				assert.True(t, st.verify())
			}
		})

		t.Run("AfterDelete", func(t *testing.T) {
			assert.True(t, st.verify())
			assert.Zero(t, st.Size())
			assert.True(t, st.IsEmpty())
		})
	})
}

func runOrderedSymbolTableTest(t *testing.T, ost OrderedSymbolTable[string, int], test orderedSymbolTableTest[string, int]) {
	t.Run(test.name, func(t *testing.T) {
		var kvs []KeyValue[string, int]
		var minKey, maxKey, floorKey, ceilingKey, selectKey string
		var minVal, maxVal, floorVal, ceilingVal, selectVal int
		var minOK, maxOK, floorOK, ceilingOK, selectOK bool

		t.Run("BeforePut", func(t *testing.T) {
			assert.True(t, ost.verify())
			assert.Zero(t, ost.Size())
			assert.Zero(t, ost.Height())
			assert.True(t, ost.IsEmpty())

			minKey, minVal, minOK = ost.Min()
			assert.Empty(t, minKey)
			assert.Zero(t, minVal)
			assert.False(t, minOK)

			maxKey, maxVal, maxOK = ost.Max()
			assert.Empty(t, maxKey)
			assert.Zero(t, maxVal)
			assert.False(t, maxOK)

			floorKey, floorVal, floorOK = ost.Floor("")
			assert.Empty(t, floorKey)
			assert.Zero(t, floorVal)
			assert.False(t, floorOK)

			ceilingKey, ceilingVal, ceilingOK = ost.Ceiling("")
			assert.Empty(t, ceilingKey)
			assert.Zero(t, ceilingVal)
			assert.False(t, ceilingOK)

			selectKey, selectVal, selectOK = ost.Select(0)
			assert.Empty(t, selectKey)
			assert.Zero(t, selectVal)
			assert.False(t, selectOK)

			assert.Zero(t, ost.Rank(""))
			assert.Len(t, ost.Range("", ""), 0)
			assert.Zero(t, ost.RangeSize("", ""))
		})

		t.Run("AfterPut", func(t *testing.T) {
			// Put
			for _, kv := range test.keyVals {
				ost.Put(kv.Key, kv.Val)
				ost.Put(kv.Key, kv.Val) // Update existing key-value
				assert.True(t, ost.verify())
			}

			// Get
			for _, expected := range test.keyVals {
				val, ok := ost.Get(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
			}

			assert.Equal(t, test.expectedSize, ost.Size())
			assert.Equal(t, test.expectedHeight, ost.Height())
			assert.Equal(t, test.expectedIsEmpty, ost.IsEmpty())

			minKey, minVal, minOK = ost.Min()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)

			maxKey, maxVal, maxOK = ost.Max()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)

			floorKey, floorVal, floorOK = ost.Floor(test.floorKey)
			assert.Equal(t, test.expectedFloorKey, floorKey)
			assert.Equal(t, test.expectedFloorVal, floorVal)
			assert.Equal(t, test.expectedFloorOK, floorOK)

			ceilingKey, ceilingVal, ceilingOK = ost.Ceiling(test.ceilingKey)
			assert.Equal(t, test.expectedCeilingKey, ceilingKey)
			assert.Equal(t, test.expectedCeilingVal, ceilingVal)
			assert.Equal(t, test.expectedCeilingOK, ceilingOK)

			minKey, minVal, minOK = ost.DeleteMin()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)
			assert.True(t, ost.verify())
			ost.Put(minKey, minVal)

			maxKey, maxVal, maxOK = ost.DeleteMax()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)
			assert.True(t, ost.verify())
			ost.Put(maxKey, maxVal)

			selectKey, selectVal, selectOK = ost.Select(test.selectRank)
			assert.Equal(t, test.expectedSelectKey, selectKey)
			assert.Equal(t, test.expectedSelectVal, selectVal)
			assert.Equal(t, test.expectedSelectOK, selectOK)

			rank := ost.Rank(test.rankKey)
			assert.Equal(t, test.expectedRank, rank)

			kvs = ost.Range(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRange, kvs)

			rangeSize := ost.RangeSize(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRangeSize, rangeSize)

			// The key-value pairs are ordered, so we can directly compare the strings.
			assert.Equal(t, test.expectedString, ost.String())

			equals := ost.Equals(test.equals)
			assert.Equal(t, test.expectedEquals, equals)

			// The key-value pairs are ordered, so we can directly compare the lists.
			all := Collect(ost.All())
			assert.Equal(t, test.expectedAll, all)

			anyMatch := ost.AnyMatch(test.anyMatchPredicate)
			assert.Equal(t, test.expectedAnyMatch, anyMatch)

			allMatch := ost.AllMatch(test.allMatchPredicate)
			assert.Equal(t, test.expectedAllMatch, allMatch)

			// VLR Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedVRLTraverse, kvs)

			// LVR Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(LVR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedLVRTraverse, kvs)

			// RVL Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(RVL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedRVLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedDescendingTraverse, kvs)

			graphviz := ost.Graphviz()
			assert.Equal(t, test.expectedGraphviz, graphviz)

			// Delete a a non-existent key
			val, ok := ost.Delete("NonExistentKey")
			assert.False(t, ok)
			assert.Zero(t, val)

			for _, expected := range test.keyVals {
				val, ok := ost.Delete(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
				assert.True(t, ost.verify())
			}
		})

		t.Run("AfterDelete", func(t *testing.T) {
			assert.True(t, ost.verify())
			assert.Zero(t, ost.Size())
			assert.Zero(t, ost.Height())
			assert.True(t, ost.IsEmpty())

			minKey, minVal, minOK = ost.Min()
			assert.Empty(t, minKey)
			assert.Zero(t, minVal)
			assert.False(t, minOK)

			maxKey, maxVal, maxOK = ost.Max()
			assert.Empty(t, maxKey)
			assert.Zero(t, maxVal)
			assert.False(t, maxOK)

			floorKey, floorVal, floorOK = ost.Floor("")
			assert.Empty(t, floorKey)
			assert.Zero(t, floorVal)
			assert.False(t, floorOK)

			ceilingKey, ceilingVal, ceilingOK = ost.Ceiling("")
			assert.Empty(t, ceilingKey)
			assert.Zero(t, ceilingVal)
			assert.False(t, ceilingOK)

			selectKey, selectVal, selectOK = ost.Select(0)
			assert.Empty(t, selectKey)
			assert.Zero(t, selectVal)
			assert.False(t, selectOK)

			assert.Zero(t, ost.Rank(""))
			assert.Len(t, ost.Range("", ""), 0)
			assert.Zero(t, ost.RangeSize("", ""))
		})
	})
}
