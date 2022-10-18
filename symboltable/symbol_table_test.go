package symboltable

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

type (
	symbolTableTest[K, V any] struct {
		name            string
		symbolTable     string
		cmpKey          generic.CompareFunc[K]
		eqVal           generic.EqualFunc[V]
		keyVals         []KeyValue[K, V]
		expectedSize    int
		expectedHeight  int
		expectedIsEmpty bool
		expectedKeyVals []KeyValue[K, V]
		equals          SymbolTable[K, V]
		expectedEquals  bool
	}

	orderedSymbolTableTest[K, V any] struct {
		name                       string
		symbolTable                string
		cmpKey                     generic.CompareFunc[K]
		eqVal                      generic.EqualFunc[V]
		keyVals                    []KeyValue[K, V]
		expectedSize               int
		expectedHeight             int
		expectedIsEmpty            bool
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
		expectedRangeSize          int
		expectedRange              []KeyValue[K, V]
		expectedKeyVals            []KeyValue[K, V]
		expectedVLRTraverse        []KeyValue[K, V]
		expectedVRLTraverse        []KeyValue[K, V]
		expectedLVRTraverse        []KeyValue[K, V]
		expectedRVLTraverse        []KeyValue[K, V]
		expectedLRVTraverse        []KeyValue[K, V]
		expectedRLVTraverse        []KeyValue[K, V]
		expectedAscendingTraverse  []KeyValue[K, V]
		expectedDescendingTraverse []KeyValue[K, V]
		equals                     OrderedSymbolTable[K, V]
		expectedEquals             bool
		expectedGraphviz           string
	}
)

func getSymbolTableTests() []symbolTableTest[string, int] {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	return []symbolTableTest[string, int]{
		{
			name:            "",
			symbolTable:     "",
			cmpKey:          cmpKey,
			eqVal:           eqVal,
			keyVals:         []KeyValue[string, int]{},
			expectedSize:    0,
			expectedIsEmpty: true,
			expectedKeyVals: []KeyValue[string, int]{},
		},
	}
}

func getOrderedSymbolTableTests() []orderedSymbolTableTest[string, int] {
	cmpKey := generic.NewCompareFunc[string]()
	eqVal := generic.NewEqualFunc[int]()

	return []orderedSymbolTableTest[string, int]{
		{
			name:   "ABC",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{"B", 2},
				{"A", 1},
				{"C", 3},
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
			expectedRangeSize:  3,
			expectedRange: []KeyValue[string, int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
			},
			expectedKeyVals: []KeyValue[string, int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
			},
		},
		{
			name:   "ABCDE",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{"B", 2},
				{"A", 1},
				{"C", 3},
				{"E", 5},
				{"D", 4},
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
			expectedRangeSize:  3,
			expectedRange: []KeyValue[string, int]{
				{"B", 2},
				{"C", 3},
				{"D", 4},
			},
			expectedKeyVals: []KeyValue[string, int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
				{"D", 4},
				{"E", 5},
			},
		},
		{
			name:   "ADGJMPS",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{"J", 10},
				{"A", 1},
				{"D", 4},
				{"S", 19},
				{"P", 16},
				{"M", 13},
				{"G", 7},
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
			expectedRangeSize:  5,
			expectedRange: []KeyValue[string, int]{
				{"D", 4},
				{"G", 7},
				{"J", 10},
				{"M", 13},
				{"P", 16},
			},
			expectedKeyVals: []KeyValue[string, int]{
				{"A", 1},
				{"D", 4},
				{"G", 7},
				{"J", 10},
				{"M", 13},
				{"P", 16},
				{"S", 19},
			},
		},
		{
			name:   "Words",
			cmpKey: cmpKey,
			eqVal:  eqVal,
			keyVals: []KeyValue[string, int]{
				{"box", 2},
				{"dad", 3},
				{"baby", 5},
				{"dome", 7},
				{"band", 11},
				{"dance", 13},
				{"balloon", 17},
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
			expectedRangeSize:  4,
			expectedRange: []KeyValue[string, int]{
				{"baby", 5},
				{"balloon", 17},
				{"band", 11},
				{"box", 2},
			},
			expectedKeyVals: []KeyValue[string, int]{
				{"baby", 5},
				{"balloon", 17},
				{"band", 11},
				{"box", 2},
				{"dad", 3},
				{"dance", 13},
				{"dome", 7},
			},
		},
	}
}

func runSymbolTableTest(t *testing.T, st SymbolTable[string, int], test symbolTableTest[string, int]) {
	t.Run(test.name, func(t *testing.T) {
		// Tree initially should be empty
		assert.True(t, st.IsEmpty())
		assert.Zero(t, st.Size())

		// TODO: verify
		assert.NotEmpty(t, test.symbolTable)
		assert.NotEmpty(t, test.cmpKey)
		assert.NotEmpty(t, test.keyVals)
		assert.NotEmpty(t, test.expectedSize)
		assert.NotEmpty(t, test.expectedIsEmpty)

		// Tree should be empty at the end
		assert.Zero(t, st.Size())
		assert.True(t, st.IsEmpty())
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

			assert.Equal(t, test.expectedRank, ost.Rank(test.rankKey))
			assert.Equal(t, test.expectedRangeSize, ost.RangeSize(test.rangeKeyLo, test.rangeKeyHi))

			kvs = ost.Range(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRange, kvs)

			kvs = ost.KeyValues()
			assert.Equal(t, test.expectedKeyVals, kvs)

			// VLR Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedVRLTraverse, kvs)

			// LVR Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(LVR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedLVRTraverse, kvs)

			// RVL Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(RVL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedRVLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[string, int]{}
			ost.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedDescendingTraverse, kvs)

			assert.Equal(t, test.expectedEquals, ost.Equals(test.equals))
			assert.Equal(t, test.expectedGraphviz, ost.Graphviz())

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
