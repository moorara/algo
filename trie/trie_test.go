package trie

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	. "github.com/moorara/algo/generic"
)

type trieTest[V any] struct {
	name                       string
	trie                       string
	eqVal                      EqualFunc[V]
	keyVals                    []KeyValue[string, V]
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
	expectedRange              []KeyValue[string, V]
	expectedRangeSize          int
	matchPattern               string
	expectedMatch              []KeyValue[string, V]
	withPrefixKey              string
	expectedWithPrefix         []KeyValue[string, V]
	longestPrefixOfKey         string
	expectedLongestPrefixOfKey string
	expectedLongestPrefixOfVal V
	expectedLongestPrefixOfOK  bool
	expectedString             string
	equal                      Trie[V]
	expectedEqual              bool
	expectedAll                []KeyValue[string, V]
	anyMatchPredicate          Predicate2[string, V]
	expectedAnyMatch           bool
	allMatchPredicate          Predicate2[string, V]
	expectedAllMatch           bool
	firstMatchPredicate        Predicate2[string, V]
	expectedFirstMatchKey      string
	expectedFirstMatchVal      V
	expectedFirstMatchOK       bool
	selectMatchPredicate       Predicate2[string, V]
	expectedSelectMatch        []KeyValue[string, V]
	partitionMatchPredicate    Predicate2[string, V]
	expectedPartitionMatched   []KeyValue[string, V]
	expectedPartitionUnmatched []KeyValue[string, V]
	expectedVLRTraverse        []KeyValue[string, V]
	expectedVRLTraverse        []KeyValue[string, V]
	expectedLVRTraverse        []KeyValue[string, V]
	expectedRVLTraverse        []KeyValue[string, V]
	expectedLRVTraverse        []KeyValue[string, V]
	expectedRLVTraverse        []KeyValue[string, V]
	expectedAscendingTraverse  []KeyValue[string, V]
	expectedDescendingTraverse []KeyValue[string, V]
	expectedDOT                string
}

func getTrieTests() []trieTest[int] {
	eqVal := NewEqualFunc[int]()

	return []trieTest[int]{
		{
			name:  "ABC",
			eqVal: eqVal,
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
			matchPattern:      "*",
			expectedMatch: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
			},
			withPrefixKey: "A",
			expectedWithPrefix: []KeyValue[string, int]{
				{Key: "A", Val: 1},
			},
			longestPrefixOfKey:         "F",
			expectedLongestPrefixOfKey: "",
			expectedLongestPrefixOfVal: 0,
			expectedLongestPrefixOfOK:  false,
			expectedString:             "{<A:1> <B:2> <C:3>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
			},
			anyMatchPredicate:        func(k string, v int) bool { return v < 0 },
			expectedAnyMatch:         false,
			allMatchPredicate:        func(k string, v int) bool { return v%2 == 0 },
			expectedAllMatch:         false,
			firstMatchPredicate:      func(k string, v int) bool { return v%5 == 0 },
			expectedFirstMatchKey:    "",
			expectedFirstMatchVal:    0,
			expectedFirstMatchOK:     false,
			selectMatchPredicate:     func(k string, v int) bool { return v%10 == 0 },
			expectedSelectMatch:      []KeyValue[string, int]{},
			partitionMatchPredicate:  func(k string, v int) bool { return v%10 == 0 },
			expectedPartitionMatched: []KeyValue[string, int]{},
			expectedPartitionUnmatched: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
			},
		},
		{
			name:  "ABCDE",
			eqVal: eqVal,
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
			matchPattern:      "*",
			expectedMatch: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
				{Key: "D", Val: 4},
				{Key: "E", Val: 5},
			},
			withPrefixKey: "C",
			expectedWithPrefix: []KeyValue[string, int]{
				{Key: "C", Val: 3},
			},
			longestPrefixOfKey:         "D",
			expectedLongestPrefixOfKey: "D",
			expectedLongestPrefixOfVal: 4,
			expectedLongestPrefixOfOK:  true,
			expectedString:             "{<A:1> <B:2> <C:3> <D:4> <E:5>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "B", Val: 2},
				{Key: "C", Val: 3},
				{Key: "D", Val: 4},
				{Key: "E", Val: 5},
			},
			anyMatchPredicate:     func(k string, v int) bool { return v == 0 },
			expectedAnyMatch:      false,
			allMatchPredicate:     func(k string, v int) bool { return v > 0 },
			expectedAllMatch:      true,
			firstMatchPredicate:   func(k string, v int) bool { return v%5 == 0 },
			expectedFirstMatchKey: "E",
			expectedFirstMatchVal: 5,
			expectedFirstMatchOK:  true,
			selectMatchPredicate:  func(k string, v int) bool { return v%2 == 0 },
			expectedSelectMatch: []KeyValue[string, int]{
				{Key: "B", Val: 2},
				{Key: "D", Val: 4},
			},
			partitionMatchPredicate: func(k string, v int) bool { return v%2 == 0 },
			expectedPartitionMatched: []KeyValue[string, int]{
				{Key: "B", Val: 2},
				{Key: "D", Val: 4},
			},
			expectedPartitionUnmatched: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "C", Val: 3},
				{Key: "E", Val: 5},
			},
		},
		{
			name:  "ADGJMPS",
			eqVal: eqVal,
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
			matchPattern:      "*",
			expectedMatch: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "D", Val: 4},
				{Key: "G", Val: 7},
				{Key: "J", Val: 10},
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
				{Key: "S", Val: 19},
			},
			withPrefixKey: "M",
			expectedWithPrefix: []KeyValue[string, int]{
				{Key: "M", Val: 13},
			},
			longestPrefixOfKey:         "P",
			expectedLongestPrefixOfKey: "P",
			expectedLongestPrefixOfVal: 16,
			expectedLongestPrefixOfOK:  true,
			expectedString:             "{<A:1> <D:4> <G:7> <J:10> <M:13> <P:16> <S:19>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "D", Val: 4},
				{Key: "G", Val: 7},
				{Key: "J", Val: 10},
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
				{Key: "S", Val: 19},
			},
			anyMatchPredicate:     func(k string, v int) bool { return v%5 == 0 },
			expectedAnyMatch:      true,
			allMatchPredicate:     func(k string, v int) bool { return v < 10 },
			expectedAllMatch:      false,
			firstMatchPredicate:   func(k string, v int) bool { return v == 13 },
			expectedFirstMatchKey: "M",
			expectedFirstMatchVal: 13,
			expectedFirstMatchOK:  true,
			selectMatchPredicate:  func(k string, v int) bool { return v > 10 },
			expectedSelectMatch: []KeyValue[string, int]{
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
				{Key: "S", Val: 19},
			},
			partitionMatchPredicate: func(k string, v int) bool { return v > 10 },
			expectedPartitionMatched: []KeyValue[string, int]{
				{Key: "M", Val: 13},
				{Key: "P", Val: 16},
				{Key: "S", Val: 19},
			},
			expectedPartitionUnmatched: []KeyValue[string, int]{
				{Key: "A", Val: 1},
				{Key: "D", Val: 4},
				{Key: "G", Val: 7},
				{Key: "J", Val: 10},
			},
		},
		{
			name:  "Words",
			eqVal: eqVal,
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
			matchPattern:      "d***e",
			expectedMatch: []KeyValue[string, int]{
				{Key: "dance", Val: 13},
			},
			withPrefixKey: "ba",
			expectedWithPrefix: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
			},
			longestPrefixOfKey:         "domestic",
			expectedLongestPrefixOfKey: "dome",
			expectedLongestPrefixOfVal: 7,
			expectedLongestPrefixOfOK:  true,
			expectedString:             "{<baby:5> <balloon:17> <band:11> <box:2> <dad:3> <dance:13> <dome:7>}",
			expectedAll: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
				{Key: "box", Val: 2},
				{Key: "dad", Val: 3},
				{Key: "dance", Val: 13},
				{Key: "dome", Val: 7},
			},
			anyMatchPredicate:     func(k string, v int) bool { return strings.HasSuffix(k, "x") },
			expectedAnyMatch:      true,
			allMatchPredicate:     func(k string, v int) bool { return k == strings.ToLower(k) },
			expectedAllMatch:      true,
			firstMatchPredicate:   func(k string, v int) bool { return strings.Contains(k, "alloo") },
			expectedFirstMatchKey: "balloon",
			expectedFirstMatchVal: 17,
			expectedFirstMatchOK:  true,
			selectMatchPredicate:  func(k string, v int) bool { return strings.HasPrefix(k, "ba") },
			expectedSelectMatch: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
			},
			partitionMatchPredicate: func(k string, v int) bool { return strings.HasPrefix(k, "ba") },
			expectedPartitionMatched: []KeyValue[string, int]{
				{Key: "baby", Val: 5},
				{Key: "balloon", Val: 17},
				{Key: "band", Val: 11},
			},
			expectedPartitionUnmatched: []KeyValue[string, int]{
				{Key: "box", Val: 2},
				{Key: "dad", Val: 3},
				{Key: "dance", Val: 13},
				{Key: "dome", Val: 7},
			},
		},
	}
}

func runTrieTest(t *testing.T, trie Trie[int], test trieTest[int]) {
	t.Run(test.name, func(t *testing.T) {
		var kvs []KeyValue[string, int]
		var minKey, maxKey, floorKey, ceilingKey, selectKey string
		var minVal, maxVal, floorVal, ceilingVal, selectVal int
		var minOK, maxOK, floorOK, ceilingOK, selectOK bool

		t.Run("Before", func(t *testing.T) {
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

		t.Run("Put", func(t *testing.T) {
			for _, kv := range test.keyVals {
				trie.Put(kv.Key, kv.Val)
				trie.Put(kv.Key, kv.Val) // Update existing key-value
				assert.True(t, trie.verify())
			}
		})

		t.Run("Get", func(t *testing.T) {
			// Get a non-existent key
			val, ok := trie.Get("NonExistentKey")
			assert.False(t, ok)
			assert.Zero(t, val)

			for _, expected := range test.keyVals {
				val, ok := trie.Get(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
			}
		})

		t.Run("Size", func(t *testing.T) {
			assert.Equal(t, test.expectedSize, trie.Size())
		})

		t.Run("Height", func(t *testing.T) {
			assert.Equal(t, test.expectedHeight, trie.Height())
		})

		t.Run("IsEmpty", func(t *testing.T) {
			assert.Equal(t, test.expectedIsEmpty, trie.IsEmpty())
		})

		t.Run("Min", func(t *testing.T) {
			minKey, minVal, minOK = trie.Min()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)
		})

		t.Run("Max", func(t *testing.T) {
			maxKey, maxVal, maxOK = trie.Max()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)
		})

		t.Run("Floor", func(t *testing.T) {
			floorKey, floorVal, floorOK = trie.Floor(test.floorKey)
			assert.Equal(t, test.expectedFloorKey, floorKey)
			assert.Equal(t, test.expectedFloorVal, floorVal)
			assert.Equal(t, test.expectedFloorOK, floorOK)
		})

		t.Run("Ceiling", func(t *testing.T) {
			ceilingKey, ceilingVal, ceilingOK = trie.Ceiling(test.ceilingKey)
			assert.Equal(t, test.expectedCeilingKey, ceilingKey)
			assert.Equal(t, test.expectedCeilingVal, ceilingVal)
			assert.Equal(t, test.expectedCeilingOK, ceilingOK)
		})

		t.Run("DeleteMin", func(t *testing.T) {
			minKey, minVal, minOK = trie.DeleteMin()
			assert.Equal(t, test.expectedMinKey, minKey)
			assert.Equal(t, test.expectedMinVal, minVal)
			assert.Equal(t, test.expectedMinOK, minOK)
			assert.True(t, trie.verify())
			trie.Put(minKey, minVal)
		})

		t.Run("DeleteMax", func(t *testing.T) {
			maxKey, maxVal, maxOK = trie.DeleteMax()
			assert.Equal(t, test.expectedMaxKey, maxKey)
			assert.Equal(t, test.expectedMaxVal, maxVal)
			assert.Equal(t, test.expectedMaxOK, maxOK)
			assert.True(t, trie.verify())
			trie.Put(maxKey, maxVal)
		})

		t.Run("Select", func(t *testing.T) {
			selectKey, selectVal, selectOK = trie.Select(test.selectRank)
			assert.Equal(t, test.expectedSelectKey, selectKey)
			assert.Equal(t, test.expectedSelectVal, selectVal)
			assert.Equal(t, test.expectedSelectOK, selectOK)
		})

		t.Run("Rank", func(t *testing.T) {
			rank := trie.Rank(test.rankKey)
			assert.Equal(t, test.expectedRank, rank)
		})

		t.Run("Range", func(t *testing.T) {
			kvs = trie.Range(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRange, kvs)
		})

		t.Run("RangeSize", func(t *testing.T) {
			rangeSize := trie.RangeSize(test.rangeKeyLo, test.rangeKeyHi)
			assert.Equal(t, test.expectedRangeSize, rangeSize)
		})

		t.Run("String", func(t *testing.T) {
			assert.Equal(t, test.expectedString, trie.String())
		})

		t.Run("Equal", func(t *testing.T) {
			equal := trie.Equal(test.equal)
			assert.Equal(t, test.expectedEqual, equal)
		})

		t.Run("All", func(t *testing.T) {
			kvs = Collect2(trie.All())
			assert.Equal(t, test.expectedAll, kvs)
		})

		t.Run("AnyMatch", func(t *testing.T) {
			anyMatch := trie.AnyMatch(test.anyMatchPredicate)
			assert.Equal(t, test.expectedAnyMatch, anyMatch)
		})

		t.Run("AllMatch", func(t *testing.T) {
			allMatch := trie.AllMatch(test.allMatchPredicate)
			assert.Equal(t, test.expectedAllMatch, allMatch)
		})

		t.Run("FirstMatch", func(t *testing.T) {
			key, val, ok := trie.FirstMatch(test.firstMatchPredicate)
			assert.Equal(t, test.expectedFirstMatchKey, key)
			assert.Equal(t, test.expectedFirstMatchVal, val)
			assert.Equal(t, test.expectedFirstMatchOK, ok)
		})

		t.Run("SelectMatch", func(t *testing.T) {
			selected := trie.SelectMatch(test.selectMatchPredicate)
			assert.Equal(t, test.expectedSelectMatch, Collect2(selected.All()))
		})

		t.Run("PartitionMatch", func(t *testing.T) {
			matched, unmatched := trie.PartitionMatch(test.partitionMatchPredicate)
			assert.Equal(t, test.expectedPartitionMatched, generic.Collect2(matched.All()))
			assert.Equal(t, test.expectedPartitionUnmatched, generic.Collect2(unmatched.All()))
		})

		t.Run("Traverse", func(t *testing.T) {
			// VLR Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedVRLTraverse, kvs)

			// LVR Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(LVR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedLVRTraverse, kvs)

			// RVL Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(RVL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedRVLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[string, int]{}
			trie.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{Key: key, Val: val})
				return true
			})
			assert.Equal(t, test.expectedDescendingTraverse, kvs)
		})

		t.Run("DOT", func(t *testing.T) {
			assert.Equal(t, test.expectedDOT, trie.DOT())
		})

		t.Run("Match", func(t *testing.T) {
			kvs = trie.Match(test.matchPattern)
			assert.Equal(t, test.expectedMatch, kvs)
		})

		t.Run("WithPrefix", func(t *testing.T) {
			kvs = trie.WithPrefix(test.withPrefixKey)
			assert.Equal(t, test.expectedWithPrefix, kvs)
		})

		t.Run("LongestPrefixOf", func(t *testing.T) {
			longestPrefixOfKey, longestPrefixVal, longestPrefixOK := trie.LongestPrefixOf(test.longestPrefixOfKey)
			assert.Equal(t, test.expectedLongestPrefixOfKey, longestPrefixOfKey)
			assert.Equal(t, test.expectedLongestPrefixOfVal, longestPrefixVal)
			assert.Equal(t, test.expectedLongestPrefixOfOK, longestPrefixOK)
		})

		t.Run("Delete", func(t *testing.T) {
			// Delete a non-existent key
			val, ok := trie.Delete("NonExistentKey")
			assert.False(t, ok)
			assert.Zero(t, val)

			for _, expected := range test.keyVals {
				val, ok := trie.Delete(expected.Key)
				assert.True(t, ok)
				assert.Equal(t, expected.Val, val)
				assert.True(t, trie.verify())
			}
		})

		t.Run("DeleteAll", func(t *testing.T) {
			for _, kv := range test.keyVals {
				trie.Put(kv.Key, kv.Val)
			}

			trie.DeleteAll()
			assert.True(t, trie.verify())
		})

		t.Run("After", func(t *testing.T) {
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
