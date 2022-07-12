package trie

import "testing"

func getPatriciaTests() []trieTest[int] {
	tests := getTrieTests()

	tests[0].symbolTable = "Patricia"
	tests[0].expectedHeight = 2
	tests[0].expectedVLRTraverse = []KeyValue[int]{{"B", 2}, {"A", 1}, {"C", 3}}
	tests[0].expectedVRLTraverse = []KeyValue[int]{{"B", 2}, {"A", 1}, {"C", 3}}
	tests[0].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"C", 3}, {"B", 2}}
	tests[0].expectedRVLTraverse = []KeyValue[int]{{"B", 2}, {"C", 3}, {"A", 1}}
	tests[0].expectedLRVTraverse = []KeyValue[int]{{"C", 3}, {"A", 1}, {"B", 2}}
	tests[0].expectedRLVTraverse = []KeyValue[int]{{"C", 3}, {"A", 1}, {"B", 2}}
	tests[0].expectedAscendingTraverse = []KeyValue[int]{{"A", 1}, {"B", 2}, {"C", 3}}
	tests[0].expectedDescendingTraverse = []KeyValue[int]{{"C", 3}, {"B", 2}, {"A", 1}}
	tests[0].expectedDotCode = `strict digraph "Patricia" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ B,2 | { <l>• | 0 | 01000010 | <r>• } }"];
  2 [label="{ A,1 | { <l>• | 7 | 01000001 | <r>• } }"];
  3 [label="{ C,3 | { <l>• | 8 | 01000011 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 2 [color=red, style=dashed];
  2:r -> 3 [color=blue];
  3:l -> 1 [color=red, style=dashed];
  3:r -> 3 [color=red, style=dashed];
}`

	tests[1].symbolTable = "Patricia"
	tests[1].expectedHeight = 3
	tests[1].expectedVLRTraverse = []KeyValue[int]{{"B", 2}, {"D", 4}, {"A", 1}, {"C", 3}, {"E", 5}}
	tests[1].expectedVRLTraverse = []KeyValue[int]{{"B", 2}, {"D", 4}, {"E", 5}, {"A", 1}, {"C", 3}}
	tests[1].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"C", 3}, {"D", 4}, {"E", 5}, {"B", 2}}
	tests[1].expectedRVLTraverse = []KeyValue[int]{{"B", 2}, {"E", 5}, {"D", 4}, {"C", 3}, {"A", 1}}
	tests[1].expectedLRVTraverse = []KeyValue[int]{{"C", 3}, {"A", 1}, {"E", 5}, {"D", 4}, {"B", 2}}
	tests[1].expectedRLVTraverse = []KeyValue[int]{{"E", 5}, {"C", 3}, {"A", 1}, {"D", 4}, {"B", 2}}
	tests[1].expectedAscendingTraverse = []KeyValue[int]{{"A", 1}, {"B", 2}, {"C", 3}, {"D", 4}, {"E", 5}}
	tests[1].expectedDescendingTraverse = []KeyValue[int]{{"E", 5}, {"D", 4}, {"C", 3}, {"B", 2}, {"A", 1}}
	tests[1].expectedDotCode = `strict digraph "Patricia" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ B,2 | { <l>• | 0 | 01000010 | <r>• } }"];
  2 [label="{ D,4 | { <l>• | 6 | 01000100 | <r>• } }"];
  3 [label="{ A,1 | { <l>• | 7 | 01000001 | <r>• } }"];
  4 [label="{ C,3 | { <l>• | 8 | 01000011 | <r>• } }"];
  5 [label="{ E,5 | { <l>• | 8 | 01000101 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 5 [color=blue];
  3:l -> 3 [color=red, style=dashed];
  3:r -> 4 [color=blue];
  4:l -> 1 [color=red, style=dashed];
  4:r -> 4 [color=red, style=dashed];
  5:l -> 2 [color=red, style=dashed];
  5:r -> 5 [color=red, style=dashed];
}`

	tests[2].symbolTable = "Patricia"
	tests[2].expectedHeight = 4
	tests[2].expectedVLRTraverse = []KeyValue[int]{{"J", 10}, {"P", 16}, {"D", 4}, {"A", 1}, {"G", 7}, {"M", 13}, {"S", 19}}
	tests[2].expectedVRLTraverse = []KeyValue[int]{{"J", 10}, {"P", 16}, {"S", 19}, {"D", 4}, {"M", 13}, {"A", 1}, {"G", 7}}
	tests[2].expectedLVRTraverse = []KeyValue[int]{{"A", 1}, {"G", 7}, {"D", 4}, {"M", 13}, {"P", 16}, {"S", 19}, {"J", 10}}
	tests[2].expectedRVLTraverse = []KeyValue[int]{{"J", 10}, {"S", 19}, {"P", 16}, {"M", 13}, {"D", 4}, {"G", 7}, {"A", 1}}
	tests[2].expectedLRVTraverse = []KeyValue[int]{{"G", 7}, {"A", 1}, {"M", 13}, {"D", 4}, {"S", 19}, {"P", 16}, {"J", 10}}
	tests[2].expectedRLVTraverse = []KeyValue[int]{{"S", 19}, {"M", 13}, {"G", 7}, {"A", 1}, {"D", 4}, {"P", 16}, {"J", 10}}
	tests[2].expectedAscendingTraverse = []KeyValue[int]{{"A", 1}, {"D", 4}, {"G", 7}, {"J", 10}, {"M", 13}, {"P", 16}, {"S", 19}}
	tests[2].expectedDescendingTraverse = []KeyValue[int]{{"S", 19}, {"P", 16}, {"M", 13}, {"J", 10}, {"G", 7}, {"D", 4}, {"A", 1}}
	tests[2].expectedDotCode = `strict digraph "Patricia" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ J,10 | { <l>• | 0 | 01001010 | <r>• } }"];
  2 [label="{ P,16 | { <l>• | 4 | 01010000 | <r>• } }"];
  3 [label="{ D,4 | { <l>• | 5 | 01000100 | <r>• } }"];
  4 [label="{ A,1 | { <l>• | 6 | 01000001 | <r>• } }"];
  5 [label="{ G,7 | { <l>• | 7 | 01000111 | <r>• } }"];
  6 [label="{ M,13 | { <l>• | 6 | 01001101 | <r>• } }"];
  7 [label="{ S,19 | { <l>• | 7 | 01010011 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 7 [color=blue];
  3:l -> 4 [color=blue];
  3:r -> 6 [color=blue];
  4:l -> 4 [color=red, style=dashed];
  4:r -> 5 [color=blue];
  5:l -> 3 [color=red, style=dashed];
  5:r -> 5 [color=red, style=dashed];
  6:l -> 1 [color=red, style=dashed];
  6:r -> 6 [color=red, style=dashed];
  7:l -> 2 [color=red, style=dashed];
  7:r -> 7 [color=red, style=dashed];
}`

	tests[3].symbolTable = "Patricia"
	tests[3].expectedHeight = 4
	tests[3].expectedVLRTraverse = []KeyValue[int]{{"box", 2}, {"dad", 3}, {"band", 11}, {"baby", 5}, {"balloon", 17}, {"dome", 7}, {"dance", 13}}
	tests[3].expectedVRLTraverse = []KeyValue[int]{{"box", 2}, {"dad", 3}, {"dome", 7}, {"dance", 13}, {"band", 11}, {"baby", 5}, {"balloon", 17}}
	tests[3].expectedLVRTraverse = []KeyValue[int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"dad", 3}, {"dance", 13}, {"dome", 7}, {"box", 2}}
	tests[3].expectedRVLTraverse = []KeyValue[int]{{"box", 2}, {"dome", 7}, {"dance", 13}, {"dad", 3}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedLRVTraverse = []KeyValue[int]{{"balloon", 17}, {"baby", 5}, {"band", 11}, {"dance", 13}, {"dome", 7}, {"dad", 3}, {"box", 2}}
	tests[3].expectedRLVTraverse = []KeyValue[int]{{"dance", 13}, {"dome", 7}, {"balloon", 17}, {"baby", 5}, {"band", 11}, {"dad", 3}, {"box", 2}}
	tests[3].expectedAscendingTraverse = []KeyValue[int]{{"baby", 5}, {"balloon", 17}, {"band", 11}, {"box", 2}, {"dad", 3}, {"dance", 13}, {"dome", 7}}
	tests[3].expectedDescendingTraverse = []KeyValue[int]{{"dome", 7}, {"dance", 13}, {"dad", 3}, {"box", 2}, {"band", 11}, {"balloon", 17}, {"baby", 5}}
	tests[3].expectedDotCode = `strict digraph "Patricia" {
  rankdir=TB;
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ box,2 | { <l>• | 0 | 01100010 01101111 01111000 | <r>• } }"];
  2 [label="{ dad,3 | { <l>• | 6 | 01100100 01100001 01100100 | <r>• } }"];
  3 [label="{ band,11 | { <l>• | 13 | 01100010 01100001 01101110 01100100 | <r>• } }"];
  4 [label="{ baby,5 | { <l>• | 21 | 01100010 01100001 01100010 01111001 | <r>• } }"];
  5 [label="{ balloon,17 | { <l>• | 23 | 01100010 01100001 01101100 01101100 01101111 01101111 01101110 | <r>• } }"];
  6 [label="{ dome,7 | { <l>• | 13 | 01100100 01101111 01101101 01100101 | <r>• } }"];
  7 [label="{ dance,13 | { <l>• | 21 | 01100100 01100001 01101110 01100011 01100101 | <r>• } }"];

  1:l -> 2 [color=blue];
  2:l -> 3 [color=blue];
  2:r -> 6 [color=blue];
  3:l -> 4 [color=blue];
  3:r -> 1 [color=red, style=dashed];
  4:l -> 4 [color=red, style=dashed];
  4:r -> 5 [color=blue];
  5:l -> 5 [color=red, style=dashed];
  5:r -> 3 [color=red, style=dashed];
  6:l -> 7 [color=blue];
  6:r -> 6 [color=red, style=dashed];
  7:l -> 2 [color=red, style=dashed];
  7:r -> 7 [color=red, style=dashed];
}`

	return tests
}

func TestPatricia(t *testing.T) {
	tests := getPatriciaTests()

	for _, tc := range tests {
		pat := NewPatricia[int]()
		runTrieTest(t, pat, tc)

		/* t.Run(tc.name, func(t *testing.T) {
			var kvs []KeyValue[int]
			var minKey, maxKey, floorKey, ceilingKey, selectKey string
			var minVal, maxVal, floorVal, ceilingVal, selectVal int
			var minOK, maxOK, floorOK, ceilingOK, selectOK bool

			for _, kv := range tc.keyVals {
				pat.Put(kv.key, kv.val)
				pat.Put(kv.key, kv.val) // Update existing key-value
			}

			for _, expected := range tc.keyVals {
				val, ok := pat.Get(expected.key)
				assert.True(t, ok)
				assert.Equal(t, expected.val, val)
			}

			assert.Equal(t, tc.expectedSize, pat.Size())
			assert.Equal(t, tc.expectedHeight, pat.Height())
			assert.Equal(t, tc.expectedIsEmpty, pat.IsEmpty())

			minKey, minVal, minOK = pat.Min()
			assert.Equal(t, tc.expectedMinKey, minKey)
			assert.Equal(t, tc.expectedMinVal, minVal)
			assert.Equal(t, tc.expectedMinOK, minOK)

			maxKey, maxVal, maxOK = pat.Max()
			assert.Equal(t, tc.expectedMaxKey, maxKey)
			assert.Equal(t, tc.expectedMaxVal, maxVal)
			assert.Equal(t, tc.expectedMaxOK, maxOK)

			floorKey, floorVal, floorOK = pat.Floor(tc.floorKey)
			assert.Equal(t, tc.expectedFloorKey, floorKey)
			assert.Equal(t, tc.expectedFloorVal, floorVal)
			assert.Equal(t, tc.expectedFloorOK, floorOK)

			ceilingKey, ceilingVal, ceilingOK = pat.Ceiling(tc.ceilingKey)
			assert.Equal(t, tc.expectedCeilingKey, ceilingKey)
			assert.Equal(t, tc.expectedCeilingVal, ceilingVal)
			assert.Equal(t, tc.expectedCeilingOK, ceilingOK)

			minKey, minVal, minOK = pat.DeleteMin()
			assert.Equal(t, tc.expectedMinKey, minKey)
			assert.Equal(t, tc.expectedMinVal, minVal)
			assert.Equal(t, tc.expectedMinOK, minOK)
			pat.Put(minKey, minVal)

			maxKey, maxVal, maxOK = pat.DeleteMax()
			assert.Equal(t, tc.expectedMaxKey, maxKey)
			assert.Equal(t, tc.expectedMaxVal, maxVal)
			assert.Equal(t, tc.expectedMaxOK, maxOK)
			pat.Put(maxKey, maxVal)

			selectKey, selectVal, selectOK = pat.Select(tc.selectRank)
			assert.Equal(t, tc.expectedSelectKey, selectKey)
			assert.Equal(t, tc.expectedSelectVal, selectVal)
			assert.Equal(t, tc.expectedSelectOK, selectOK)

			assert.Equal(t, tc.expectedRank, pat.Rank(tc.rankKey))
			assert.Equal(t, tc.expectedRangeSize, pat.RangeSize(tc.rangeKeyLo, tc.rangeKeyHi))

			kvs = pat.Range(tc.rangeKeyLo, tc.rangeKeyHi)
			for _, kv := range kvs { // Soundness
				assert.Contains(t, tc.expectedRange, kv)
			}
			for _, kv := range tc.expectedRange { // Completeness
				assert.Contains(t, kvs, kv)
			}
			for i := 0; i < len(kvs)-1; i++ { // Sorted Ascending
				assert.Equal(t, -1, tc.cmpKey(kvs[i].key, kvs[i+1].key))
			}

			kvs = pat.KeyValues()
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
			kvs = []KeyValue[int]{}
			pat.Traverse(VLR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedVLRTraverse, kvs)

			// VRL Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(VRL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedVRLTraverse, kvs)

			// LVR Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(LVR, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedLVRTraverse, kvs)

			// RVL Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(RVL, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedRVLTraverse, kvs)

			// LRV Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(LRV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedLRVTraverse, kvs)

			// RLV Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(RLV, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedRLVTraverse, kvs)

			// Ascending Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(Ascending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedAscendingTraverse, kvs)

			// Descending Traversal
			kvs = []KeyValue[int]{}
			pat.Traverse(Descending, func(key string, val int) bool {
				kvs = append(kvs, KeyValue[int]{key, val})
				return true
			})
			assert.Equal(t, tc.expectedDescendingTraverse, kvs)

			// Graphviz dot language code
			assert.Equal(t, tc.expectedDotCode, pat.Graphviz())

			for _, expected := range tc.keyVals {
				val, ok := pat.Delete(expected.key)
				assert.True(t, ok)
				assert.Equal(t, expected.val, val)
			}
		}) */
	}
}
