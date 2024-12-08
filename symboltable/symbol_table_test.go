package symboltable

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/moorara/algo/generic"
)

type (
	// nolint: unused
	symbolTableTest[K, V any] struct {
		name              string
		symbolTable       string
		cmpKey            CompareFunc[K]
		eqVal             EqualFunc[V]
		keyVals           []KeyValue[K, V]
		expectedSize      int
		expectedHeight    int
		expectedIsEmpty   bool
		expectedString    string
		equals            SymbolTable[K, V]
		expectedEquals    bool
		expectedAll       []KeyValue[K, V]
		anyMatchPredicate Predicate2[K, V]
		expectedAnyMatch  bool
		allMatchPredicate Predicate2[K, V]
		expectedAllMatch  bool
	}

	orderedSymbolTableTest[K, V any] struct {
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
)

// nolint: unused
func getSymbolTableTests() []symbolTableTest[string, int] {
	cmpKey := NewCompareFunc[string]()
	eqVal := NewEqualFunc[int]()

	return []symbolTableTest[string, int]{
		{
			name:              "TBD",
			symbolTable:       "",
			cmpKey:            cmpKey,
			eqVal:             eqVal,
			keyVals:           []KeyValue[string, int]{},
			expectedSize:      0,
			expectedIsEmpty:   true,
			expectedString:    "{}",
			expectedAll:       []KeyValue[string, int]{},
			anyMatchPredicate: func(k string, v int) bool { return false },
			expectedAnyMatch:  false,
			allMatchPredicate: func(k string, v int) bool { return false },
			expectedAllMatch:  false,
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

// nolint: unused
func runSymbolTableTest(t *testing.T, st SymbolTable[string, int], test symbolTableTest[string, int]) {
	t.Run(test.name, func(t *testing.T) {
		var kvs []KeyValue[string, int]

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

			assert.Equal(t, test.expectedString, st.String())

			equals := st.Equals(test.equals)
			assert.Equal(t, test.expectedEquals, equals)

			kvs = Collect(st.All())
			assert.Equal(t, test.expectedAll, kvs)

			any := st.AnyMatch(test.anyMatchPredicate)
			assert.Equal(t, test.expectedAnyMatch, any)

			all := st.AllMatch(test.allMatchPredicate)
			assert.Equal(t, test.expectedAllMatch, all)

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
			assert.Zero(t, ost.RangeSize("", ""))
			assert.Len(t, ost.Range("", ""), 0)
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

			assert.Equal(t, test.expectedString, ost.String())

			equals := ost.Equals(test.equals)
			assert.Equal(t, test.expectedEquals, equals)

			kvs = Collect(ost.All())
			assert.Equal(t, test.expectedAll, kvs)

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
			assert.Zero(t, ost.RangeSize("", ""))
			assert.Len(t, ost.Range("", ""), 0)
		})
	})
}
