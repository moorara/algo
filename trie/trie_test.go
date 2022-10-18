package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

type trieTest[V any] struct {
	name                       string
	symbolTable                string
	eqVal                      generic.EqualFunc[V]
	keyVals                    []KeyValue[V]
	expectedSize               int
	expectedHeight             int
	expectedIsEmpty            bool
	expectedMinKey             string
	expectedMinVal             V
	expectedMinOK              bool
	expectedMaxKey             string
	expectedMaxVal             V
	expectedMaxOK              bool
	floorKey                   string
	expectedFloorKey           string
	expectedFloorVal           V
	expectedFloorOK            bool
	ceilingKey                 string
	expectedCeilingKey         string
	expectedCeilingVal         V
	expectedCeilingOK          bool
	selectRank                 int
	expectedSelectKey          string
	expectedSelectVal          V
	expectedSelectOK           bool
	rankKey                    string
	expectedRank               int
	rangeKeyLo                 string
	rangeKeyHi                 string
	expectedRangeSize          int
	expectedRange              []KeyValue[V]
	expectedKeyVals            []KeyValue[V]
	expectedVLRTraverse        []KeyValue[V]
	expectedVRLTraverse        []KeyValue[V]
	expectedLVRTraverse        []KeyValue[V]
	expectedRVLTraverse        []KeyValue[V]
	expectedLRVTraverse        []KeyValue[V]
	expectedRLVTraverse        []KeyValue[V]
	expectedAscendingTraverse  []KeyValue[V]
	expectedDescendingTraverse []KeyValue[V]
	equals                     Trie[V]
	expectedEquals             bool
	expectedGraphviz           string
	matchPattern               string
	expectedMatch              []KeyValue[V]
	withPrefixKey              string
	expectedWithPrefix         []KeyValue[V]
	longestPrefixOfKey         string
	expectedLongestPrefixOfKey string
	expectedLongestPrefixOfVal V
	expectedLongestPrefixOfOK  bool
}

func getTrieTests() []trieTest[int] {
	eqVal := generic.NewEqualFunc[int]()

	return []trieTest[int]{
		{
			name:  "ABC",
			eqVal: eqVal,
			keyVals: []KeyValue[int]{
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
			expectedRange: []KeyValue[int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
			},
			expectedKeyVals: []KeyValue[int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
			},
			matchPattern: "*",
			expectedMatch: []KeyValue[int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
			},
			withPrefixKey: "A",
			expectedWithPrefix: []KeyValue[int]{
				{"A", 1},
			},
			longestPrefixOfKey:         "F",
			expectedLongestPrefixOfKey: "",
			expectedLongestPrefixOfVal: 0,
			expectedLongestPrefixOfOK:  false,
		},
		{
			name:  "ABCDE",
			eqVal: eqVal,
			keyVals: []KeyValue[int]{
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
			expectedRange: []KeyValue[int]{
				{"B", 2},
				{"C", 3},
				{"D", 4},
			},
			expectedKeyVals: []KeyValue[int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
				{"D", 4},
				{"E", 5},
			},
			matchPattern: "*",
			expectedMatch: []KeyValue[int]{
				{"A", 1},
				{"B", 2},
				{"C", 3},
				{"D", 4},
				{"E", 5},
			},
			withPrefixKey: "C",
			expectedWithPrefix: []KeyValue[int]{
				{"C", 3},
			},
			longestPrefixOfKey:         "D",
			expectedLongestPrefixOfKey: "D",
			expectedLongestPrefixOfVal: 4,
			expectedLongestPrefixOfOK:  true,
		},
		{
			name:  "ADGJMPS",
			eqVal: eqVal,
			keyVals: []KeyValue[int]{
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
			expectedRange: []KeyValue[int]{
				{"D", 4},
				{"G", 7},
				{"J", 10},
				{"M", 13},
				{"P", 16},
			},
			expectedKeyVals: []KeyValue[int]{
				{"A", 1},
				{"D", 4},
				{"G", 7},
				{"J", 10},
				{"M", 13},
				{"P", 16},
				{"S", 19},
			},
			matchPattern: "*",
			expectedMatch: []KeyValue[int]{
				{"A", 1},
				{"D", 4},
				{"G", 7},
				{"J", 10},
				{"M", 13},
				{"P", 16},
				{"S", 19},
			},
			withPrefixKey: "M",
			expectedWithPrefix: []KeyValue[int]{
				{"M", 13},
			},
			longestPrefixOfKey:         "P",
			expectedLongestPrefixOfKey: "P",
			expectedLongestPrefixOfVal: 16,
			expectedLongestPrefixOfOK:  true,
		},
		{
			name:  "Words",
			eqVal: eqVal,
			keyVals: []KeyValue[int]{
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
			expectedRange: []KeyValue[int]{
				{"baby", 5},
				{"balloon", 17},
				{"band", 11},
				{"box", 2},
			},
			expectedKeyVals: []KeyValue[int]{
				{"baby", 5},
				{"balloon", 17},
				{"band", 11},
				{"box", 2},
				{"dad", 3},
				{"dance", 13},
				{"dome", 7},
			},
			matchPattern: "d***e",
			expectedMatch: []KeyValue[int]{
				{"dance", 13},
			},
			withPrefixKey: "ba",
			expectedWithPrefix: []KeyValue[int]{
				{"baby", 5},
				{"balloon", 17},
				{"band", 11},
			},
			longestPrefixOfKey:         "domestic",
			expectedLongestPrefixOfKey: "dome",
			expectedLongestPrefixOfVal: 7,
			expectedLongestPrefixOfOK:  true,
		},
	}
}

func runTrieTest(t *testing.T, trie Trie[int], test trieTest[int]) {
	t.Run(test.name, func(t *testing.T) {
		var kvs []KeyValue[int]
		var minKey, maxKey, floorKey, ceilingKey, selectKey string
		var minVal, maxVal, floorVal, ceilingVal, selectVal int
		var minOK, maxOK, floorOK, ceilingOK, selectOK bool

		t.Run("BeforePut", func(t *testing.T) {
			assert.True(t, trie.verify())
			assert.Zero(t, trie.Size())
			assert.Zero(t, trie.Height())
			assert.True(t, trie.IsEmpty())

			minKey, minVal, minOK = trie.Min()
			assert.Empty(t, minKey)
			assert.Zero(t, minVal)
			assert.False(t, minOK)

			maxKey, maxVal, maxOK = trie.Max()
			assert.Empty(t, maxKey)
			assert.Zero(t, maxVal)
			assert.False(t, maxOK)

			floorKey, floorVal, floorOK = trie.Floor("")
			assert.Empty(t, floorKey)
			assert.Zero(t, floorVal)
			assert.False(t, floorOK)

			ceilingKey, ceilingVal, ceilingOK = trie.Ceiling("")
			assert.Empty(t, ceilingKey)
			assert.Zero(t, ceilingVal)
			assert.False(t, ceilingOK)

			selectKey, selectVal, selectOK = trie.Select(0)
			assert.Empty(t, selectKey)
			assert.Zero(t, selectVal)
			assert.False(t, selectOK)

			assert.Zero(t, trie.Rank(""))
			assert.Zero(t, trie.RangeSize("", ""))
			assert.Len(t, trie.Range("", ""), 0)
		})

		t.Run("AfterPut", func(t *testing.T) {
			// Put
			for _, kv := range test.keyVals {
				trie.Put(kv.Key, kv.Val)
				trie.Put(kv.Key, kv.Val) // Update existing key-value
				assert.True(t, trie.verify())
			}

			// Get
			for _, expected := range test.keyVals {
				val, ok := trie.Get(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
			}

			assert.Equal(t, test.expectedSize, trie.Size())
			assert.Equal(t, test.expectedHeight, trie.Height())
			assert.Equal(t, test.expectedIsEmpty, trie.IsEmpty())

			minKey, minVal, minOK = trie.Min()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)

			maxKey, maxVal, maxOK = trie.Max()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)

			floorKey, floorVal, floorOK = trie.Floor(test.floorKey)
			assert.Equal(t, test.expectedFloorKey, floorKey)
			assert.Equal(t, test.expectedFloorVal, floorVal)
			assert.Equal(t, test.expectedFloorOK, floorOK)

			ceilingKey, ceilingVal, ceilingOK = trie.Ceiling(test.ceilingKey)
			assert.Equal(t, test.expectedCeilingKey, ceilingKey)
			assert.Equal(t, test.expectedCeilingVal, ceilingVal)
			assert.Equal(t, test.expectedCeilingOK, ceilingOK)

			minKey, minVal, minOK = trie.DeleteMin()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)
			assert.True(t, trie.verify())
			trie.Put(minKey, minVal)

			maxKey, maxVal, maxOK = trie.DeleteMax()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)
			assert.True(t, trie.verify())
			trie.Put(maxKey, maxVal)

			selectKey, selectVal, selectOK = trie.Select(test.selectRank)
			assert.Equal(t, test.expectedSelectKey, selectKey)
			assert.Equal(t, test.expectedSelectVal, selectVal)
			assert.Equal(t, test.expectedSelectOK, selectOK)

			assert.Equal(t, test.expectedRank, trie.Rank(test.rankKey))
			assert.Equal(t, test.expectedRangeSize, trie.RangeSize(test.rangeKeyLo, test.rangeKeyHi))

			kvs = trie.Range(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRange, kvs)

			kvs = trie.KeyValues()
			assert.Equal(t, test.expectedKeyVals, kvs)

			// VLR Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedVRLTraverse, kvs)

			// LVR Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(LVR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedLVRTraverse, kvs)

			// RVL Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(RVL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedRVLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[int]{}
			trie.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, test.expectedDescendingTraverse, kvs)

			assert.Equal(t, test.expectedEquals, trie.Equals(test.equals))
			assert.Equal(t, test.expectedGraphviz, trie.Graphviz())

			kvs = trie.Match(test.matchPattern)
			assert.Equal(t, test.expectedMatch, kvs)

			kvs = trie.WithPrefix(test.withPrefixKey)
			assert.Equal(t, test.expectedWithPrefix, kvs)

			longestPrefixOfKey, longestPrefixVal, longestPrefixOK := trie.LongestPrefixOf(test.longestPrefixOfKey)
			assert.Equal(t, test.expectedLongestPrefixOfKey, longestPrefixOfKey)
			assert.Equal(t, test.expectedLongestPrefixOfVal, longestPrefixVal)
			assert.Equal(t, test.expectedLongestPrefixOfOK, longestPrefixOK)

			for _, expected := range test.keyVals {
				val, ok := trie.Delete(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
				assert.True(t, trie.verify())
			}
		})

		t.Run("AfterDelete", func(t *testing.T) {
			assert.True(t, trie.verify())
			assert.Zero(t, trie.Size())
			assert.Zero(t, trie.Height())
			assert.True(t, trie.IsEmpty())

			minKey, minVal, minOK = trie.Min()
			assert.Empty(t, minKey)
			assert.Zero(t, minVal)
			assert.False(t, minOK)

			maxKey, maxVal, maxOK = trie.Max()
			assert.Empty(t, maxKey)
			assert.Zero(t, maxVal)
			assert.False(t, maxOK)

			floorKey, floorVal, floorOK = trie.Floor("")
			assert.Empty(t, floorKey)
			assert.Zero(t, floorVal)
			assert.False(t, floorOK)

			ceilingKey, ceilingVal, ceilingOK = trie.Ceiling("")
			assert.Empty(t, ceilingKey)
			assert.Zero(t, ceilingVal)
			assert.False(t, ceilingOK)

			selectKey, selectVal, selectOK = trie.Select(0)
			assert.Empty(t, selectKey)
			assert.Zero(t, selectVal)
			assert.False(t, selectOK)

			assert.Zero(t, trie.Rank(""))
			assert.Zero(t, trie.RangeSize("", ""))
			assert.Len(t, trie.Range("", ""), 0)
		})
	})
}
