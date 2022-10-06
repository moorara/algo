package symboltable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getBTreeTests() []orderedSymbolTableTest[string, int] {
	tests := getOrderedSymbolTableTests()

	tests[0].symbolTable = "B-Tree"
	tests[0].expectedHeight = 0
	tests[0].expectedVLRTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedVRLTraverse = []KeyValue[string, int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedRLVTraverse = []KeyValue[string, int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[string, int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedDotCode = `strict digraph "B-Tree" {
  rankdir=TB;
  concentrate=false;
  node [];

  1 [label="{ A | 1 } | { B | 2 } | { C | 3 } | {  |  }", style=bold, shape=record];
}`

	tests[1].symbolTable = "B-Tree"
	tests[1].expectedHeight = 1
	tests[1].expectedVLRTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedVRLTraverse = []KeyValue[string, int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedRLVTraverse = []KeyValue[string, int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[string, int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedDotCode = `strict digraph "B-Tree" {
  rankdir=TB;
  concentrate=false;
  node [];

  1 [label="{ * | <sentinel> } | { C | <C> } | {  |  } | {  |  }", style=solid, shape=Mrecord];
  2 [label="{ A | 1 } | { B | 2 } | {  |  } | {  |  }", style=bold, shape=record];
  3 [label="{ C | 3 } | { D | 4 } | { E | 5 } | {  |  }", style=bold, shape=record];

  1:sentinel -> 2 [];
  1:C -> 3 [];
}`

	tests[2].symbolTable = "B-Tree"
	tests[2].expectedHeight = 1
	tests[2].expectedVLRTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedVRLTraverse = []KeyValue[string, int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedLRVTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedRLVTraverse = []KeyValue[string, int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedAscendingTraverse = []KeyValue[string, int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[string, int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedDotCode = `strict digraph "B-Tree" {
  rankdir=TB;
  concentrate=false;
  node [];

  1 [label="{ * | <sentinel> } | { J | <J> } | { P | <P> } | {  |  }", style=solid, shape=Mrecord];
  2 [label="{ A | 1 } | { D | 4 } | { G | 7 } | {  |  }", style=bold, shape=record];
  3 [label="{ J | 10 } | { M | 13 } | {  |  } | {  |  }", style=bold, shape=record];
  4 [label="{ P | 16 } | { S | 19 } | {  |  } | {  |  }", style=bold, shape=record];

  1:sentinel -> 2 [];
  1:J -> 3 [];
  1:P -> 4 [];
}`

	tests[3].symbolTable = "B-Tree"
	tests[3].expectedHeight = 1
	tests[3].expectedVLRTraverse = []KeyValue[string, int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedVRLTraverse = []KeyValue[string, int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedLRVTraverse = []KeyValue[string, int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedRLVTraverse = []KeyValue[string, int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedAscendingTraverse = []KeyValue[string, int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[string, int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedDotCode = `strict digraph "B-Tree" {
  rankdir=TB;
  concentrate=false;
  node [];

  1 [label="{ * | <sentinel> } | { band | <band> } | { dad | <dad> } | {  |  }", style=solid, shape=Mrecord];
  2 [label="{ baby | 5 } | { balloon | 17 } | {  |  } | {  |  }", style=bold, shape=record];
  3 [label="{ band | 11 } | { box | 2 } | {  |  } | {  |  }", style=bold, shape=record];
  4 [label="{ dad | 3 } | { dance | 13 } | { dome | 7 } | {  |  }", style=bold, shape=record];

  1:sentinel -> 2 [];
  1:band -> 3 [];
  1:dad -> 4 [];
}`

	return tests
}

func TestBTree(t *testing.T) {
	tests := getBTreeTests()

	for _, tc := range tests {
		bt := NewBTree[string, int](4, tc.cmpKey)

		t.Run(tc.name, func(t *testing.T) {
			var kvs []KeyValue[string, int]
			var minKey, maxKey, floorKey, ceilingKey string
			var minVal, maxVal, floorVal, ceilingVal int
			var minOK, maxOK, floorOK, ceilingOK bool

			for _, kv := range tc.keyVals {
				bt.Put(kv.key, kv.val)
				bt.Put(kv.key, kv.val) // Update existing key-value
			}

			for _, expected := range tc.keyVals {
				val, ok := bt.Get(expected.key)
				assert.True(t, ok)
				assert.Equal(t, expected.val, val)
			}

			assert.Equal(t, tc.expectedSize, bt.Size())
			assert.Equal(t, tc.expectedHeight, bt.Height())
			assert.Equal(t, tc.expectedIsEmpty, bt.IsEmpty())

			minKey, minVal, minOK = bt.Min()
			assert.Equal(t, tc.expectedMinKey, minKey)
			assert.Equal(t, tc.expectedMinVal, minVal)
			assert.Equal(t, tc.expectedMinOK, minOK)

			maxKey, maxVal, maxOK = bt.Max()
			assert.Equal(t, tc.expectedMaxKey, maxKey)
			assert.Equal(t, tc.expectedMaxVal, maxVal)
			assert.Equal(t, tc.expectedMaxOK, maxOK)

			floorKey, floorVal, floorOK = bt.Floor(tc.floorKey)
			assert.Equal(t, tc.expectedFloorKey, floorKey)
			assert.Equal(t, tc.expectedFloorVal, floorVal)
			assert.Equal(t, tc.expectedFloorOK, floorOK)

			ceilingKey, ceilingVal, ceilingOK = bt.Ceiling(tc.ceilingKey)
			assert.Equal(t, tc.expectedCeilingKey, ceilingKey)
			assert.Equal(t, tc.expectedCeilingVal, ceilingVal)
			assert.Equal(t, tc.expectedCeilingOK, ceilingOK)

			kvs = bt.KeyValues()
			for _, kv := range kvs { // Soundness
				assert.Contains(t, tc.keyVals, kv)
			}
			for _, kv := range tc.keyVals { // Completeness
				assert.Contains(t, kvs, kv)
			}
			for i := 0; i < len(kvs)-1; i++ { // Sorted Ascending
				assert.Equal(t, -1, tc.cmpKey(kvs[i].key, kvs[i+1].key))
			}

			// VLR Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedVRLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[string, int]{}
			bt.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[string, int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedDescendingTraverse, kvs)

			// Graphviz dot language code
			assert.Equal(t, tc.expectedDotCode, bt.Graphviz())
		})
	}
}
