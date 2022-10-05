package heap

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/stretchr/testify/assert"
)

type (
	KeyValue[K, V any] struct {
		key K
		val V
	}

	IndexedKeyValue[K, V any] struct {
		idx int
		key K
		val V
	}

	heapTest[K, V any] struct {
		name             string
		heap             string
		size             int
		cmpKey           generic.CompareFunc[K]
		eqVal            generic.EqualFunc[V]
		insertKVs        []KeyValue[K, V]
		expectedSize     int
		expectedIsEmpty  bool
		expectedPeek     KeyValue[K, V]
		expectedContains []KeyValue[K, V]
		expectedDelete   []KeyValue[K, V]
		expectedGraphviz string
	}

	indexedHeapTest[K, V any] struct {
		name                string
		heap                string
		cap                 int
		cmpKey              generic.CompareFunc[K]
		eqVal               generic.EqualFunc[V]
		insertKVs           []IndexedKeyValue[K, V]
		changeKeyKVs        []IndexedKeyValue[K, V]
		expectedSize        int
		expectedIsEmpty     bool
		expectedPeek        IndexedKeyValue[K, V]
		expectedPeekIndex   []IndexedKeyValue[K, V]
		expectedContains    []IndexedKeyValue[K, V]
		expectedDelete      []IndexedKeyValue[K, V]
		expectedDeleteIndex []IndexedKeyValue[K, V]
		expectedGraphviz    string
	}
)

func getHeapTests() []heapTest[int, string] {
	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	return []heapTest[int, string]{
		{
			name:             "MinHeap_Empty",
			size:             2,
			cmpKey:           cmpMin,
			eqVal:            eqVal,
			insertKVs:        []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{0, ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:             "MaxHeap_Empty",
			size:             2,
			cmpKey:           cmpMax,
			eqVal:            eqVal,
			insertKVs:        []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{0, ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:   "MinHeap_FewEntries",
			size:   2,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{30, "thirty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{10, "ten"},
			expectedContains: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
			},
			expectedDelete: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
			},
		},
		{
			name:   "MaxHeap_FewEntries",
			size:   2,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{30, "thirty"},
			expectedContains: []KeyValue[int, string]{
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []KeyValue[int, string]{
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
		{
			name:   "MinHeap_SomeEntries",
			size:   4,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{50, "fifty"},
				{30, "thirty"},
				{40, "forty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{10, "ten"},
			expectedContains: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
			},
			expectedDelete: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
			},
		},
		{
			name:   "MaxHeap_SomeEntries",
			size:   4,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
				{50, "fifty"},
				{40, "forty"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{50, "fifty"},
			expectedContains: []KeyValue[int, string]{
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []KeyValue[int, string]{
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
		{
			name:   "MinHeap_ManyEntries",
			size:   4,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{30, "thirty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{10, "ten"},
			expectedContains: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{70, "seventy"},
				{80, "eighty"},
				{90, "ninety"},
			},
			expectedDelete: []KeyValue[int, string]{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{70, "seventy"},
				{80, "eighty"},
				{90, "ninety"},
			},
		},
		{
			name:   "MaxHeap_ManyEntries",
			size:   4,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []KeyValue[int, string]{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
				{50, "fifty"},
				{40, "forty"},
				{60, "sixty"},
				{70, "seventy"},
				{90, "ninety"},
				{80, "eighty"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{90, "ninety"},
			expectedContains: []KeyValue[int, string]{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{60, "sixty"},
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []KeyValue[int, string]{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{60, "sixty"},
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
	}
}

func getIndexedHeapTests() []indexedHeapTest[int, string] {
	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	return []indexedHeapTest[int, string]{
		{
			name:                "MinHeap_Empty",
			cap:                 10,
			cmpKey:              cmpMin,
			eqVal:               eqVal,
			insertKVs:           []IndexedKeyValue[int, string]{},
			changeKeyKVs:        []IndexedKeyValue[int, string]{},
			expectedSize:        0,
			expectedIsEmpty:     true,
			expectedPeek:        IndexedKeyValue[int, string]{},
			expectedPeekIndex:   []IndexedKeyValue[int, string]{},
			expectedContains:    []IndexedKeyValue[int, string]{},
			expectedDelete:      []IndexedKeyValue[int, string]{},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{},
		},
		{
			name:                "MaxHeap_Empty",
			cap:                 10,
			cmpKey:              cmpMax,
			eqVal:               eqVal,
			insertKVs:           []IndexedKeyValue[int, string]{},
			changeKeyKVs:        []IndexedKeyValue[int, string]{},
			expectedSize:        0,
			expectedIsEmpty:     true,
			expectedPeek:        IndexedKeyValue[int, string]{},
			expectedPeekIndex:   []IndexedKeyValue[int, string]{},
			expectedContains:    []IndexedKeyValue[int, string]{},
			expectedDelete:      []IndexedKeyValue[int, string]{},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{},
		},
		{
			name:   "MinHeap_FewEntries",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 30, "thirty"},
				{1, 1, "ten"},
				{2, 200, "twenty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 1, key: 10},
				{idx: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{1, 10, "ten"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{1, 10, "ten"},
				{2, 20, "twenty"},
				{0, 30, "thirty"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{1, 10, "ten"},
				{2, 20, "twenty"},
				{0, 30, "thirty"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{1, 10, "ten"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 30, "thirty"},
				{2, 20, "twenty"},
			},
		},
		{
			name:   "MaxHeap_FewEntries",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{1, 3, "thirty"},
				{2, 200, "twenty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 1, key: 30},
				{idx: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{1, 30, "thirty"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{1, 30, "thirty"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{2, 20, "twenty"},
			},
		},
		{
			name:   "MinHeap_SomeEntries",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 50, "fifty"},
				{1, 30, "thirty"},
				{2, 4, "forty"},
				{3, 10, "ten"},
				{4, 200, "twenty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 2, key: 40},
				{idx: 3, key: 10},
				{idx: 4, key: 20},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{3, 10, "ten"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{3, 10, "ten"},
				{4, 20, "twenty"},
				{1, 30, "thirty"},
				{2, 40, "forty"},
				{0, 50, "fifty"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{3, 10, "ten"},
				{4, 20, "twenty"},
				{1, 30, "thirty"},
				{2, 40, "forty"},
				{0, 50, "fifty"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{3, 10, "ten"},
				{4, 20, "twenty"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 50, "fifty"},
				{2, 40, "forty"},
				{1, 30, "thirty"},
			},
		},
		{
			name:   "MaxHeap_SomeEntries",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{1, 30, "thirty"},
				{2, 2, "twenty"},
				{3, 50, "fifty"},
				{4, 400, "forty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 2, key: 20},
				{idx: 3, key: 50},
				{idx: 4, key: 40},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{3, 50, "fifty"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{3, 50, "fifty"},
				{4, 40, "forty"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{2, 20, "twenty"},
				{1, 30, "thirty"},
			},
		},
		{
			name:   "MinHeap_ManyEntries",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 90, "ninety"},
				{1, 80, "eighty"},
				{2, 70, "seventy"},
				{3, 40, "forty"},
				{4, 5, "fifty"},
				{5, 6, "sixty"},
				{6, 30, "thirty"},
				{7, 100, "ten"},
				{8, 200, "twenty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 4, key: 50},
				{idx: 5, key: 60},
				{idx: 6, key: 30},
				{idx: 7, key: 10},
				{idx: 8, key: 20},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{7, 10, "ten"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
				{4, 50, "fifty"},
				{5, 60, "sixty"},
				{2, 70, "seventy"},
				{1, 80, "eighty"},
				{0, 90, "ninety"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
				{4, 50, "fifty"},
				{5, 60, "sixty"},
				{2, 70, "seventy"},
				{1, 80, "eighty"},
				{0, 90, "ninety"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 90, "ninety"},
				{1, 80, "eighty"},
				{2, 70, "seventy"},
				{5, 60, "sixty"},
				{4, 50, "fifty"},
			},
		},
		{
			name:   "MaxHeap_ManyEntries",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			insertKVs: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{3, 50, "fifty"},
				{4, 4, "forty"},
				{5, 6, "sixty"},
				{6, 70, "seventy"},
				{7, 900, "ninety"},
				{8, 800, "eighty"},
			},
			changeKeyKVs: []IndexedKeyValue[int, string]{
				{idx: 4, key: 40},
				{idx: 5, key: 60},
				{idx: 6, key: 70},
				{idx: 7, key: 90},
				{idx: 8, key: 80},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    IndexedKeyValue[int, string]{7, 90, "ninety"},
			expectedPeekIndex: []IndexedKeyValue[int, string]{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []IndexedKeyValue[int, string]{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []IndexedKeyValue[int, string]{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
			},
			expectedDeleteIndex: []IndexedKeyValue[int, string]{
				{0, 10, "ten"},
				{2, 20, "twenty"},
				{1, 30, "thirty"},
				{4, 40, "forty"},
				{3, 50, "fifty"},
			},
		},
	}
}

func runHeapTest(t *testing.T, heap Heap[int, string], test heapTest[int, string]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("BeforeInsert", func(t *testing.T) {
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(0))
			assert.False(t, heap.ContainsValue(""))

			peekKey, peekVal, peekOK := heap.Peek()
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			deleteKey, deleteVal, deleteOK := heap.Delete()
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, deleteOK)
		})

		t.Run("AfterInsert", func(t *testing.T) {
			for _, kv := range test.insertKVs {
				heap.Insert(kv.key, kv.val)
			}

			assert.Equal(t, test.expectedSize, heap.Size())
			assert.Equal(t, test.expectedIsEmpty, heap.IsEmpty())

			peekKey, peekVal, peekOK := heap.Peek()
			if test.expectedSize == 0 {
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, test.expectedPeek.key, peekKey)
				assert.Equal(t, test.expectedPeek.val, peekVal)
				assert.True(t, peekOK)
			}

			for _, kv := range test.expectedContains {
				assert.True(t, heap.ContainsKey(kv.key))
				assert.True(t, heap.ContainsValue(kv.val))
			}

			// Graphviz dot language code
			assert.Equal(t, test.expectedGraphviz, heap.Graphviz())

			for _, kv := range test.expectedDelete {
				deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, kv.key, deleteKey)
				assert.Equal(t, kv.val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("AfterDelete", func(t *testing.T) {
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(0))
			assert.False(t, heap.ContainsValue(""))

			peekKey, peekVal, peekOK := heap.Peek()
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			deleteKey, deleteVal, deleteOK := heap.Delete()
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, peekOK, deleteOK)
		})
	})
}

func runIndexedHeapTest(t *testing.T, heap IndexedHeap[int, string], test indexedHeapTest[int, string]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("BeforeInsert", func(t *testing.T) {
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsIndex(0))
			assert.False(t, heap.ContainsKey(0))
			assert.False(t, heap.ContainsValue(""))

			peekIndex, peekKey, peekVal, peekOK := heap.Peek()
			assert.Equal(t, -1, peekIndex)
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			peekKey, peekVal, peekOK = heap.PeekIndex(0)
			assert.Equal(t, -1, peekIndex)
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
			assert.Equal(t, -1, deleteIndex)
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, deleteOK)

			deleteKey, deleteVal, deleteOK = heap.DeleteIndex(0)
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, deleteOK)
		})

		t.Run("AfterInsert", func(t *testing.T) {
			for _, kv := range test.insertKVs {
				heap.Insert(kv.idx, kv.key, kv.val)
			}

			for _, kv := range test.changeKeyKVs {
				heap.ChangeKey(kv.idx, kv.key)
			}

			assert.Equal(t, test.expectedSize, heap.Size())
			assert.Equal(t, test.expectedIsEmpty, heap.IsEmpty())

			peekIndex, peekKey, peekVal, peekOK := heap.Peek()
			if test.expectedSize == 0 {
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, test.expectedPeek.idx, peekIndex)
				assert.Equal(t, test.expectedPeek.key, peekKey)
				assert.Equal(t, test.expectedPeek.val, peekVal)
				assert.True(t, peekOK)
			}

			for _, kv := range test.expectedPeekIndex {
				peekKey, peekVal, peekOK = heap.PeekIndex(kv.idx)
				assert.Equal(t, kv.key, peekKey)
				assert.Equal(t, kv.val, peekVal)
				assert.True(t, peekOK)
			}

			for _, kv := range test.expectedContains {
				assert.True(t, heap.ContainsIndex(kv.idx))
				assert.True(t, heap.ContainsKey(kv.key))
				assert.True(t, heap.ContainsValue(kv.val))
			}

			// Graphviz dot language code
			assert.Equal(t, test.expectedGraphviz, heap.Graphviz())

			for _, kv := range test.expectedDelete {
				deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, kv.idx, deleteIndex)
				assert.Equal(t, kv.key, deleteKey)
				assert.Equal(t, kv.val, deleteVal)
				assert.True(t, deleteOK)
			}

			for _, kv := range test.expectedDeleteIndex {
				deleteKey, deleteVal, deleteOK := heap.DeleteIndex(kv.idx)
				assert.Equal(t, kv.key, deleteKey)
				assert.Equal(t, kv.val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("AfterDelete", func(t *testing.T) {
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(0))
			assert.False(t, heap.ContainsValue(""))

			peekIndex, peekKey, peekVal, peekOK := heap.Peek()
			assert.Equal(t, -1, peekIndex)
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			peekKey, peekVal, peekOK = heap.PeekIndex(0)
			assert.Equal(t, -1, peekIndex)
			assert.Zero(t, peekKey)
			assert.Empty(t, peekVal)
			assert.False(t, peekOK)

			deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
			assert.Equal(t, -1, deleteIndex)
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, deleteOK)

			deleteKey, deleteVal, deleteOK = heap.DeleteIndex(0)
			assert.Zero(t, deleteKey)
			assert.Empty(t, deleteVal)
			assert.False(t, deleteOK)
		})
	})
}
