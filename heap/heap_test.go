package heap

import (
	"testing"

	. "github.com/moorara/algo/generic"
	"github.com/stretchr/testify/assert"
)

var (
	eqVal  = NewEqualFunc[string]()
	cmpMin = NewCompareFunc[int]()
	cmpMax = NewReverseCompareFunc[int]()
)

type indexedKeyValue[K, V any] struct {
	index int
	key   K
	val   V
}

type heapTest[K, V any] struct {
	name             string
	heap             string
	size             int
	cmpKey           CompareFunc[K]
	eqVal            EqualFunc[V]
	inserts          []KeyValue[K, V]
	expectedSize     int
	expectedIsEmpty  bool
	expectedPeek     KeyValue[K, V]
	expectedContains []KeyValue[K, V]
	expectedDelete   []KeyValue[K, V]
	expectedDOT      string
}

type indexedHeapTest[K, V any] struct {
	name                string
	heap                string
	cap                 int
	cmpKey              CompareFunc[K]
	eqVal               EqualFunc[V]
	inserts             []indexedKeyValue[K, V]
	changeKeys          []indexedKeyValue[K, V]
	expectedSize        int
	expectedIsEmpty     bool
	expectedPeek        indexedKeyValue[K, V]
	expectedPeekIndex   []indexedKeyValue[K, V]
	expectedContains    []indexedKeyValue[K, V]
	expectedDelete      []indexedKeyValue[K, V]
	expectedDeleteIndex []indexedKeyValue[K, V]
	expectedDOT         string
}

type mergeableHeapTest[K, V any] struct {
	name             string
	heap             string
	cmpKey           CompareFunc[K]
	eqVal            EqualFunc[V]
	inserts          []KeyValue[K, V]
	merge            MergeableHeap[K, V]
	expectedSize     int
	expectedIsEmpty  bool
	expectedPeek     KeyValue[K, V]
	expectedContains []KeyValue[K, V]
	expectedDelete   []KeyValue[K, V]
	expectedDOT      string
}

func getHeapTests() []heapTest[int, string] {
	return []heapTest[int, string]{
		{
			name:             "MinHeap_Empty",
			size:             2,
			cmpKey:           cmpMin,
			eqVal:            eqVal,
			inserts:          []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{Key: 0, Val: ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:             "MaxHeap_Empty",
			size:             2,
			cmpKey:           cmpMax,
			eqVal:            eqVal,
			inserts:          []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{Key: 0, Val: ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:   "MinHeap_FewItems",
			size:   2,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
			},
		},
		{
			name:   "MaxHeap_FewItems",
			size:   2,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 30, Val: "Task#3"},
			expectedContains: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
		{
			name:   "MinHeap_SomeItems",
			size:   4,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
			},
		},
		{
			name:   "MaxHeap_SomeItems",
			size:   4,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 50, Val: "Task#5"},
			expectedContains: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
		{
			name:   "MinHeap_ManyItems",
			size:   4,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 30, Val: "Task#3"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 80, Val: "Task#8"},
				{Key: 90, Val: "Task#9"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 80, Val: "Task#8"},
				{Key: 90, Val: "Task#9"},
			},
		},
		{
			name:   "MaxHeap_ManyItems",
			size:   4,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 90, Val: "Task#9"},
			expectedContains: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 60, Val: "Task#6"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 60, Val: "Task#6"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
	}
}

func getIndexedHeapTests() []indexedHeapTest[int, string] {
	return []indexedHeapTest[int, string]{
		{
			name:                "MinHeap_Empty",
			cap:                 10,
			cmpKey:              cmpMin,
			eqVal:               eqVal,
			inserts:             []indexedKeyValue[int, string]{},
			changeKeys:          []indexedKeyValue[int, string]{},
			expectedSize:        0,
			expectedIsEmpty:     true,
			expectedPeek:        indexedKeyValue[int, string]{},
			expectedPeekIndex:   []indexedKeyValue[int, string]{},
			expectedContains:    []indexedKeyValue[int, string]{},
			expectedDelete:      []indexedKeyValue[int, string]{},
			expectedDeleteIndex: []indexedKeyValue[int, string]{},
		},
		{
			name:                "MaxHeap_Empty",
			cap:                 10,
			cmpKey:              cmpMax,
			eqVal:               eqVal,
			inserts:             []indexedKeyValue[int, string]{},
			changeKeys:          []indexedKeyValue[int, string]{},
			expectedSize:        0,
			expectedIsEmpty:     true,
			expectedPeek:        indexedKeyValue[int, string]{},
			expectedPeekIndex:   []indexedKeyValue[int, string]{},
			expectedContains:    []indexedKeyValue[int, string]{},
			expectedDelete:      []indexedKeyValue[int, string]{},
			expectedDeleteIndex: []indexedKeyValue[int, string]{},
		},
		{
			name:   "MinHeap_FewItems",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 30, "Task#3"},
				{1, 1, "Task#1"},
				{2, 200, "Task#2"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 1, key: 10},
				{index: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{1, 10, "Task#1"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{1, 10, "Task#1"},
				{2, 20, "Task#2"},
				{0, 30, "Task#3"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{1, 10, "Task#1"},
				{2, 20, "Task#2"},
				{0, 30, "Task#3"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{1, 10, "Task#1"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 30, "Task#3"},
				{2, 20, "Task#2"},
			},
		},
		{
			name:   "MaxHeap_FewItems",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{1, 3, "Task#3"},
				{2, 200, "Task#2"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 1, key: 30},
				{index: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{1, 30, "Task#3"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{1, 30, "Task#3"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{2, 20, "Task#2"},
			},
		},
		{
			name:   "MinHeap_SomeItems",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 50, "Task#5"},
				{1, 30, "Task#3"},
				{2, 4, "Task#4"},
				{3, 10, "Task#1"},
				{4, 200, "Task#2"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 2, key: 40},
				{index: 3, key: 10},
				{index: 4, key: 20},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{3, 10, "Task#1"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{3, 10, "Task#1"},
				{4, 20, "Task#2"},
				{1, 30, "Task#3"},
				{2, 40, "Task#4"},
				{0, 50, "Task#5"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{3, 10, "Task#1"},
				{4, 20, "Task#2"},
				{1, 30, "Task#3"},
				{2, 40, "Task#4"},
				{0, 50, "Task#5"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{3, 10, "Task#1"},
				{4, 20, "Task#2"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 50, "Task#5"},
				{2, 40, "Task#4"},
				{1, 30, "Task#3"},
			},
		},
		{
			name:   "MaxHeap_SomeItems",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{1, 30, "Task#3"},
				{2, 2, "Task#2"},
				{3, 50, "Task#5"},
				{4, 400, "Task#4"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 2, key: 20},
				{index: 3, key: 50},
				{index: 4, key: 40},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{3, 50, "Task#5"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{3, 50, "Task#5"},
				{4, 40, "Task#4"},
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{3, 50, "Task#5"},
				{4, 40, "Task#4"},
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{3, 50, "Task#5"},
				{4, 40, "Task#4"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{2, 20, "Task#2"},
				{1, 30, "Task#3"},
			},
		},
		{
			name:   "MinHeap_ManyItems",
			cap:    10,
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 90, "Task#9"},
				{1, 80, "Task#8"},
				{2, 70, "Task#7"},
				{3, 40, "Task#4"},
				{4, 5, "Task#5"},
				{5, 6, "Task#6"},
				{6, 30, "Task#3"},
				{7, 100, "Task#1"},
				{8, 200, "Task#2"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 4, key: 50},
				{index: 5, key: 60},
				{index: 6, key: 30},
				{index: 7, key: 10},
				{index: 8, key: 20},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{7, 10, "Task#1"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{7, 10, "Task#1"},
				{8, 20, "Task#2"},
				{6, 30, "Task#3"},
				{3, 40, "Task#4"},
				{4, 50, "Task#5"},
				{5, 60, "Task#6"},
				{2, 70, "Task#7"},
				{1, 80, "Task#8"},
				{0, 90, "Task#9"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{7, 10, "Task#1"},
				{8, 20, "Task#2"},
				{6, 30, "Task#3"},
				{3, 40, "Task#4"},
				{4, 50, "Task#5"},
				{5, 60, "Task#6"},
				{2, 70, "Task#7"},
				{1, 80, "Task#8"},
				{0, 90, "Task#9"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{7, 10, "Task#1"},
				{8, 20, "Task#2"},
				{6, 30, "Task#3"},
				{3, 40, "Task#4"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 90, "Task#9"},
				{1, 80, "Task#8"},
				{2, 70, "Task#7"},
				{5, 60, "Task#6"},
				{4, 50, "Task#5"},
			},
		},
		{
			name:   "MaxHeap_ManyItems",
			cap:    10,
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{3, 50, "Task#5"},
				{4, 4, "Task#4"},
				{5, 6, "Task#6"},
				{6, 70, "Task#7"},
				{7, 900, "Task#9"},
				{8, 800, "Task#8"},
			},
			changeKeys: []indexedKeyValue[int, string]{
				{index: 4, key: 40},
				{index: 5, key: 60},
				{index: 6, key: 70},
				{index: 7, key: 90},
				{index: 8, key: 80},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    indexedKeyValue[int, string]{7, 90, "Task#9"},
			expectedPeekIndex: []indexedKeyValue[int, string]{
				{7, 90, "Task#9"},
				{8, 80, "Task#8"},
				{6, 70, "Task#7"},
				{5, 60, "Task#6"},
				{3, 50, "Task#5"},
				{4, 40, "Task#4"},
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedContains: []indexedKeyValue[int, string]{
				{7, 90, "Task#9"},
				{8, 80, "Task#8"},
				{6, 70, "Task#7"},
				{5, 60, "Task#6"},
				{3, 50, "Task#5"},
				{4, 40, "Task#4"},
				{1, 30, "Task#3"},
				{2, 20, "Task#2"},
				{0, 10, "Task#1"},
			},
			expectedDelete: []indexedKeyValue[int, string]{
				{7, 90, "Task#9"},
				{8, 80, "Task#8"},
				{6, 70, "Task#7"},
				{5, 60, "Task#6"},
			},
			expectedDeleteIndex: []indexedKeyValue[int, string]{
				{0, 10, "Task#1"},
				{2, 20, "Task#2"},
				{1, 30, "Task#3"},
				{4, 40, "Task#4"},
				{3, 50, "Task#5"},
			},
		},
	}
}

func getMergeableHeapTests() []mergeableHeapTest[int, string] {
	return []mergeableHeapTest[int, string]{
		{
			name:             "MinHeap_Empty",
			cmpKey:           cmpMin,
			eqVal:            eqVal,
			inserts:          []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{Key: 0, Val: ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:             "MaxHeap_Empty",
			cmpKey:           cmpMax,
			eqVal:            eqVal,
			inserts:          []KeyValue[int, string]{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     KeyValue[int, string]{Key: 0, Val: ""},
			expectedContains: []KeyValue[int, string]{},
			expectedDelete:   []KeyValue[int, string]{},
		},
		{
			name:   "MinHeap_FewItems",
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
			},
		},
		{
			name:   "MaxHeap_FewItems",
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 30, Val: "Task#3"},
			expectedContains: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
		{
			name:   "MinHeap_SomeItems",
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
			},
		},
		{
			name:   "MaxHeap_SomeItems",
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 50, Val: "Task#5"},
			expectedContains: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
		{
			name:   "MinHeap_ManyItems",
			cmpKey: cmpMin,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 30, Val: "Task#3"},
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 10, Val: "Task#1"},
			expectedContains: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 80, Val: "Task#8"},
				{Key: 90, Val: "Task#9"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 20, Val: "Task#2"},
				{Key: 30, Val: "Task#3"},
				{Key: 40, Val: "Task#4"},
				{Key: 50, Val: "Task#5"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 80, Val: "Task#8"},
				{Key: 90, Val: "Task#9"},
			},
		},
		{
			name:   "MaxHeap_ManyItems",
			cmpKey: cmpMax,
			eqVal:  eqVal,
			inserts: []KeyValue[int, string]{
				{Key: 10, Val: "Task#1"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 60, Val: "Task#6"},
				{Key: 70, Val: "Task#7"},
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    KeyValue[int, string]{Key: 90, Val: "Task#9"},
			expectedContains: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 60, Val: "Task#6"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
			expectedDelete: []KeyValue[int, string]{
				{Key: 90, Val: "Task#9"},
				{Key: 80, Val: "Task#8"},
				{Key: 70, Val: "Task#7"},
				{Key: 60, Val: "Task#6"},
				{Key: 50, Val: "Task#5"},
				{Key: 40, Val: "Task#4"},
				{Key: 30, Val: "Task#3"},
				{Key: 20, Val: "Task#2"},
				{Key: 10, Val: "Task#1"},
			},
		},
	}
}

func runHeapTest(t *testing.T, heap Heap[int, string], test heapTest[int, string]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("Before", func(t *testing.T) {
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

		t.Run("Insert", func(t *testing.T) {
			for _, kv := range test.inserts {
				heap.Insert(kv.Key, kv.Val)
			}
		})

		t.Run("Size", func(t *testing.T) {
			assert.Equal(t, test.expectedSize, heap.Size())
		})

		t.Run("IsEmpty", func(t *testing.T) {
			assert.Equal(t, test.expectedIsEmpty, heap.IsEmpty())
		})

		t.Run("Peek", func(t *testing.T) {
			peekKey, peekVal, peekOK := heap.Peek()
			if test.expectedSize == 0 {
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, test.expectedPeek.Key, peekKey)
				assert.Equal(t, test.expectedPeek.Val, peekVal)
				assert.True(t, peekOK)
			}
		})

		t.Run("Contains", func(t *testing.T) {
			for _, kv := range test.expectedContains {
				assert.True(t, heap.ContainsKey(kv.Key))
				assert.True(t, heap.ContainsValue(kv.Val))
			}
		})

		t.Run("DOT", func(t *testing.T) {
			assert.Equal(t, test.expectedDOT, heap.DOT())
		})

		t.Run("Delete", func(t *testing.T) {
			for _, kv := range test.expectedDelete {
				deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, kv.Key, deleteKey)
				assert.Equal(t, kv.Val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("DeleteAll", func(t *testing.T) {
			for _, kv := range test.inserts {
				heap.Insert(kv.Key, kv.Val)
			}

			heap.DeleteAll()
		})

		t.Run("After", func(t *testing.T) {
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
		t.Run("Before", func(t *testing.T) {
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

		t.Run("Insert", func(t *testing.T) {
			for _, kv := range test.inserts {
				ok := heap.Insert(kv.index, kv.key, kv.val)
				assert.True(t, ok)
			}

			// Try inserting an index already on the heap.
			if len(test.inserts) > 0 {
				kv := test.inserts[0]
				ok := heap.Insert(kv.index, kv.key, kv.val)
				assert.False(t, ok)
			}
		})

		t.Run("ChangeKey", func(t *testing.T) {
			for _, kv := range test.changeKeys {
				ok := heap.ChangeKey(kv.index, kv.key)
				assert.True(t, ok)
			}

			// Try changing the key for an index not on the heap.
			ok := heap.ChangeKey(-1, 69)
			assert.False(t, ok)
		})

		t.Run("Size", func(t *testing.T) {
			assert.Equal(t, test.expectedSize, heap.Size())
		})

		t.Run("IsEmpty", func(t *testing.T) {
			assert.Equal(t, test.expectedIsEmpty, heap.IsEmpty())
		})

		t.Run("Peek", func(t *testing.T) {
			peekIndex, peekKey, peekVal, peekOK := heap.Peek()
			if test.expectedSize == 0 {
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, test.expectedPeek.index, peekIndex)
				assert.Equal(t, test.expectedPeek.key, peekKey)
				assert.Equal(t, test.expectedPeek.val, peekVal)
				assert.True(t, peekOK)
			}
		})

		t.Run("PeekIndex", func(t *testing.T) {
			for _, kv := range test.expectedPeekIndex {
				peekKey, peekVal, peekOK := heap.PeekIndex(kv.index)
				assert.Equal(t, kv.key, peekKey)
				assert.Equal(t, kv.val, peekVal)
				assert.True(t, peekOK)
			}
		})

		t.Run("Contains", func(t *testing.T) {
			for _, kv := range test.expectedContains {
				assert.True(t, heap.ContainsIndex(kv.index))
				assert.True(t, heap.ContainsKey(kv.key))
				assert.True(t, heap.ContainsValue(kv.val))
			}
		})

		t.Run("DOT", func(t *testing.T) {
			assert.Equal(t, test.expectedDOT, heap.DOT())
		})

		t.Run("Delete", func(t *testing.T) {
			for _, kv := range test.expectedDelete {
				deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, kv.index, deleteIndex)
				assert.Equal(t, kv.key, deleteKey)
				assert.Equal(t, kv.val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("DeleteIndex", func(t *testing.T) {
			for _, kv := range test.expectedDeleteIndex {
				deleteKey, deleteVal, deleteOK := heap.DeleteIndex(kv.index)
				assert.Equal(t, kv.key, deleteKey)
				assert.Equal(t, kv.val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("DeleteAll", func(t *testing.T) {
			for _, kv := range test.inserts {
				heap.Insert(kv.index, kv.key, kv.val)
			}

			heap.DeleteAll()
		})

		t.Run("After", func(t *testing.T) {
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

func runMergeableHeapTest(t *testing.T, heap MergeableHeap[int, string], test mergeableHeapTest[int, string]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("Before", func(t *testing.T) {
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

		t.Run("Insert", func(t *testing.T) {
			for _, kv := range test.inserts {
				heap.Insert(kv.Key, kv.Val)
			}
		})

		t.Run("Merge", func(t *testing.T) {
			heap.Merge(test.merge)
		})

		t.Run("Size", func(t *testing.T) {
			assert.Equal(t, test.expectedSize, heap.Size())
		})

		t.Run("IsEmpty", func(t *testing.T) {
			assert.Equal(t, test.expectedIsEmpty, heap.IsEmpty())
		})

		t.Run("Peek", func(t *testing.T) {
			peekKey, peekVal, peekOK := heap.Peek()
			if test.expectedSize == 0 {
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, test.expectedPeek.Key, peekKey)
				assert.Equal(t, test.expectedPeek.Val, peekVal)
				assert.True(t, peekOK)
			}
		})

		t.Run("Contains", func(t *testing.T) {
			for _, kv := range test.expectedContains {
				assert.True(t, heap.ContainsKey(kv.Key))
				assert.True(t, heap.ContainsValue(kv.Val))
			}
		})

		t.Run("DOT", func(t *testing.T) {
			assert.Equal(t, test.expectedDOT, heap.DOT())
		})

		t.Run("Delete", func(t *testing.T) {
			for _, kv := range test.expectedDelete {
				deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, kv.Key, deleteKey)
				assert.Equal(t, kv.Val, deleteVal)
				assert.True(t, deleteOK)
			}
		})

		t.Run("DeleteAll", func(t *testing.T) {
			for _, kv := range test.inserts {
				heap.Insert(kv.Key, kv.Val)
			}

			heap.DeleteAll()
		})

		t.Run("After", func(t *testing.T) {
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
