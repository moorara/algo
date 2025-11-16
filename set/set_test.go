package set

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/stretchr/testify/assert"
)

type (
	setTest[T any] struct {
		name                string
		addTests            [][]T
		removeTests         [][]T
		containTests        []containTest[T]
		equalTests          []equalTest[T]
		anyMatchTests       []anyMatchTest[T]
		allMatchTests       []allMatchTest[T]
		firstMatchTests     []firstMatchTest[T]
		selectMatchTests    []selectMatchTest[T]
		partitionMatchTests []partitionMatchTest[T]
		isSubsetTests       []isSubsetTest[T]
		isSupersetTests     []isSupersetTest[T]
		unionTests          []unionTest[T]
		intersectionTests   []intersectionTest[T]
		differenceTests     []differenceTest[T]
		expectedSize        int
		expectedEmpty       bool
		expectedSubstrings  []string
		expectedAll         []T
	}

	containTest[T any] struct {
		vals     []T
		expected bool
	}

	equalTest[T any] struct {
		rhs      Set[T]
		expected bool
	}

	anyMatchTest[T any] struct {
		p        generic.Predicate1[T]
		expected bool
	}

	allMatchTest[T any] struct {
		p        generic.Predicate1[T]
		expected bool
	}

	firstMatchTest[T any] struct {
		p             generic.Predicate1[T]
		expectedMatch T
		expectedFound bool
	}

	selectMatchTest[T any] struct {
		p               generic.Predicate1[T]
		expectedMatched Set[T]
	}

	partitionMatchTest[T any] struct {
		p                 generic.Predicate1[T]
		expectedMatched   Set[T]
		expectedUnmatched Set[T]
	}

	isSubsetTest[T any] struct {
		superset Set[T]
		expected bool
	}

	isSupersetTest[T any] struct {
		subset   Set[T]
		expected bool
	}

	unionTest[T any] struct {
		sets     []Set[T]
		expected Set[T]
	}

	intersectionTest[T any] struct {
		sets     []Set[T]
		expected Set[T]
	}

	differenceTest[T any] struct {
		sets     []Set[T]
		expected Set[T]
	}
)

func getSetTests() []setTest[int] {
	eqInt := generic.NewEqualFunc[int]()
	cmpInt := generic.NewCompareFunc[int]()
	hashInt := hash.HashFuncForInt[int](nil)

	return []setTest[int]{
		{
			name: "OK",
			addTests: [][]int{
				{1, 2, 4, 8, 16, 32, 64},
				{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
				{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
			},
			removeTests: [][]int{
				{16, 32, 64},
				{13, 21, 34, 55},
				{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
				{101, 211, 307, 401, 503, 601, 701, 809, 907},
			},
			containTests: []containTest[int]{
				{vals: []int{1, 2, 4, 8}, expected: true},
				{vals: []int{3, 5, 7}, expected: true},
				{vals: []int{16, 32, 64}, expected: false},
				{vals: []int{13, 21, 34, 55}, expected: false},
				{vals: []int{31, 37, 41, 43, 47, 53, 59, 61, 67}, expected: false},
				{vals: []int{101, 211, 307, 401, 503, 601, 701}, expected: false},
			},
			equalTests: []equalTest[int]{
				{rhs: New(eqInt), expected: false},
				{rhs: New(eqInt, 2, 4, 8), expected: false},
				{rhs: New(eqInt, 1, 2, 3, 4, 5, 6, 7), expected: false},
				{rhs: New(eqInt, 1, 2, 3, 4, 5, 6, 7, 8), expected: false},
				{rhs: New(eqInt, 1, 2, 3, 4, 5, 7, 8), expected: true},
				{rhs: NewStableSet(eqInt, 1, 2, 3, 4, 5, 7, 8), expected: true},
				{rhs: NewSortedSet(cmpInt, 1, 2, 3, 4, 5, 7, 8), expected: true},
				{rhs: NewHashSet(hashInt, eqInt, HashSetOpts{}, 1, 2, 3, 4, 5, 7, 8), expected: true},
			},
			anyMatchTests: []anyMatchTest[int]{
				{p: func(v int) bool { return v%2 == 0 }, expected: true},
				{p: func(v int) bool { return v > 10 }, expected: false},
			},
			allMatchTests: []allMatchTest[int]{
				{p: func(v int) bool { return v < 10 }, expected: true},
				{p: func(v int) bool { return v%2 == 0 }, expected: false},
			},
			firstMatchTests: []firstMatchTest[int]{
				{p: func(v int) bool { return v%3 == 0 }, expectedMatch: 3, expectedFound: true},
				{p: func(v int) bool { return v > 10 }, expectedMatch: 0, expectedFound: false},
			},
			selectMatchTests: []selectMatchTest[int]{
				{
					p:               func(v int) bool { return v%2 == 0 },
					expectedMatched: New(eqInt, 2, 4, 8),
				},
				{
					p:               func(v int) bool { return v > 10 },
					expectedMatched: New(eqInt),
				},
			},
			partitionMatchTests: []partitionMatchTest[int]{
				{
					p:                 func(v int) bool { return v%2 == 0 },
					expectedMatched:   New(eqInt, 2, 4, 8),
					expectedUnmatched: New(eqInt, 1, 3, 5, 7),
				},
				{
					p:                 func(v int) bool { return v > 10 },
					expectedMatched:   New(eqInt),
					expectedUnmatched: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
			},
			isSubsetTests: []isSubsetTest[int]{
				{superset: New(eqInt, 2, 3, 4, 5, 7, 8), expected: false},
				{superset: New(eqInt, 1, 2, 3, 4, 5, 7, 8), expected: true},
				{superset: New(eqInt, 1, 2, 3, 4, 5, 7, 8, 9), expected: true},
			},
			isSupersetTests: []isSupersetTest[int]{
				{subset: New(eqInt, 2, 3, 4, 5, 7, 8), expected: true},
				{subset: New(eqInt, 1, 2, 3, 4, 5, 7, 8), expected: true},
				{subset: New(eqInt, 1, 2, 3, 4, 5, 7, 8, 9), expected: false},
			},
			unionTests: []unionTest[int]{
				{
					sets:     []Set[int]{},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
				{
					sets: []Set[int]{
						New(eqInt),
					},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
				{
					sets: []Set[int]{
						New(eqInt, 16, 32, 64),
						New(eqInt, 13, 21, 34, 55),
					},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8, 13, 16, 21, 32, 34, 55, 64),
				},
			},
			intersectionTests: []intersectionTest[int]{
				{
					sets:     []Set[int]{},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
				{
					sets: []Set[int]{
						New(eqInt),
					},
					expected: New(eqInt),
				},
				{
					sets: []Set[int]{
						New(eqInt, 2, 4, 8),
						New(eqInt, 2, 3, 5, 8),
					},
					expected: New(eqInt, 2, 8),
				},
			},
			differenceTests: []differenceTest[int]{
				{
					sets:     []Set[int]{},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
				{
					sets: []Set[int]{
						New(eqInt),
					},
					expected: New(eqInt, 1, 2, 3, 4, 5, 7, 8),
				},
				{
					sets: []Set[int]{
						New(eqInt, 2, 4, 8),
						New(eqInt, 2, 3, 5, 8),
					},
					expected: New(eqInt, 1, 7),
				},
			},
			expectedSize:       7,
			expectedEmpty:      false,
			expectedSubstrings: []string{"1", "2", "3", "4", "5", "7", "8"},
			expectedAll:        []int{1, 2, 3, 4, 5, 7, 8},
		},
	}
}

func runSetTest(t *testing.T, set Set[int], test setTest[int]) {
	t.Run(test.name, func(t *testing.T) {
		t.Run("Before", func(t *testing.T) {
			assert.Zero(t, set.Size())
			assert.True(t, set.IsEmpty())
			assert.False(t, set.Contains(-1))
			assert.Equal(t, "{}", set.String())
		})

		t.Run("Add", func(t *testing.T) {
			for _, vals := range test.addTests {
				set.Add(vals...)
			}
		})

		t.Run("Remove", func(t *testing.T) {
			for _, vals := range test.removeTests {
				set.Remove(vals...)
			}
		})

		t.Run("Contains", func(t *testing.T) {
			for _, test := range test.containTests {
				b := set.Contains(test.vals...)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("Equal", func(t *testing.T) {
			for _, test := range test.equalTests {
				b := set.Equal(test.rhs)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("AnyMatch", func(t *testing.T) {
			for _, test := range test.anyMatchTests {
				b := set.AnyMatch(test.p)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("AllMatch", func(t *testing.T) {
			for _, test := range test.allMatchTests {
				b := set.AllMatch(test.p)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("FirstMatch", func(t *testing.T) {
			for _, test := range test.firstMatchTests {
				match, found := set.FirstMatch(test.p)
				assert.Equal(t, test.expectedMatch, match)
				assert.Equal(t, test.expectedFound, found)
			}
		})

		t.Run("SelectMatch", func(t *testing.T) {
			for _, test := range test.selectMatchTests {
				matched := set.SelectMatch(test.p)
				assert.True(t, matched.(Set[int]).Equal(test.expectedMatched), "Expected:\n%s\nGot:\n%s", test.expectedMatched, matched)
			}
		})

		t.Run("PartitionMatch", func(t *testing.T) {
			for _, test := range test.partitionMatchTests {
				matched, unmatched := set.PartitionMatch(test.p)
				assert.True(t, matched.(Set[int]).Equal(test.expectedMatched), "Expected:\n%s\nGot:\n%s", test.expectedMatched, matched)
				assert.True(t, unmatched.(Set[int]).Equal(test.expectedUnmatched), "Expected:\n%s\nGot:\n%s", test.expectedUnmatched, unmatched)
			}
		})

		t.Run("IsSubset", func(t *testing.T) {
			for _, test := range test.isSubsetTests {
				b := set.IsSubset(test.superset)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("IsSuperset", func(t *testing.T) {
			for _, test := range test.isSupersetTests {
				b := set.IsSuperset(test.subset)
				assert.Equal(t, test.expected, b)
			}
		})

		t.Run("Union", func(t *testing.T) {
			for _, test := range test.unionTests {
				union := set.Union(test.sets...)
				assert.True(t, union.Equal(test.expected), "Expected:\n%s\nGot:\n%s", test.expected, union)
			}
		})

		t.Run("Intersection", func(t *testing.T) {
			for _, test := range test.intersectionTests {
				intersection := set.Intersection(test.sets...)
				assert.True(t, intersection.Equal(test.expected), "Expected:\n%s\nGot:\n%s", test.expected, intersection)
			}
		})

		t.Run("Difference", func(t *testing.T) {
			for _, test := range test.differenceTests {
				difference := set.Difference(test.sets...)
				assert.True(t, difference.Equal(test.expected), "Expected:\n%s\nGot:\n%s", test.expected, difference)
			}
		})

		t.Run("Clone", func(t *testing.T) {
			clone := set.Clone()
			assert.True(t, clone.Equal(set))
		})

		t.Run("CloneEmpty", func(t *testing.T) {
			clone := set.CloneEmpty()
			assert.Zero(t, clone.Size())
			assert.True(t, clone.IsEmpty())
		})

		t.Run("Size", func(t *testing.T) {
			assert.Equal(t, test.expectedSize, set.Size())
		})

		t.Run("Empty", func(t *testing.T) {
			assert.Equal(t, test.expectedEmpty, set.IsEmpty())
		})

		t.Run("String", func(t *testing.T) {
			str := set.String()
			for _, substr := range test.expectedSubstrings {
				assert.Contains(t, str, substr)
			}
		})

		t.Run("All", func(t *testing.T) {
			all := generic.Collect1(set.All())

			assert.Len(t, all, len(test.expectedAll))
			for _, v := range test.expectedAll {
				assert.Contains(t, all, v)
			}
		})

		t.Run("RemoveAll", func(t *testing.T) {
			set.RemoveAll()
		})

		t.Run("After", func(t *testing.T) {
			assert.Zero(t, set.Size())
			assert.True(t, set.IsEmpty())
			assert.False(t, set.Contains(-1))
			assert.Equal(t, "{}", set.String())
		})
	})
}

func TestPowerset(t *testing.T) {
	opts := HashSetOpts{}
	hashInt := hash.HashFuncForInt[int](nil)
	eqInt := generic.NewEqualFunc[int]()
	cmpInt := generic.NewCompareFunc[int]()
	eqSet := func(a, b Set[int]) bool { return a.Equal(b) }

	tests := []struct {
		name     string
		s        Set[int]
		expected Set[Set[int]]
	}{
		{
			name: "Empty_Set",
			s:    New(eqInt),
			expected: New(eqSet,
				New(eqInt),
			),
		},
		{
			name: "NonEmpty_Set",
			s:    New(eqInt, 3, 5, 7),
			expected: New(eqSet,
				New(eqInt),
				New(eqInt, 3),
				New(eqInt, 5),
				New(eqInt, 7),
				New(eqInt, 3, 5),
				New(eqInt, 3, 7),
				New(eqInt, 5, 7),
				New(eqInt, 3, 5, 7),
			),
		},
		{
			name: "NonEmpty_StableSet",
			s:    NewStableSet(eqInt, 3, 5, 7),
			expected: NewStableSet(eqSet,
				NewStableSet(eqInt),
				NewStableSet(eqInt, 3),
				NewStableSet(eqInt, 5),
				NewStableSet(eqInt, 7),
				NewStableSet(eqInt, 3, 5),
				NewStableSet(eqInt, 3, 7),
				NewStableSet(eqInt, 5, 7),
				NewStableSet(eqInt, 3, 5, 7),
			),
		},
		{
			name: "NonEmpty_SortedSet",
			s:    NewSortedSet(cmpInt, 3, 5, 7),
			expected: NewSortedSet(cmpSortedSet[int],
				NewSortedSet(cmpInt),
				NewSortedSet(cmpInt, 3),
				NewSortedSet(cmpInt, 5),
				NewSortedSet(cmpInt, 7),
				NewSortedSet(cmpInt, 3, 5),
				NewSortedSet(cmpInt, 3, 7),
				NewSortedSet(cmpInt, 5, 7),
				NewSortedSet(cmpInt, 3, 5, 7),
			),
		},
		{
			name: "NonEmpty_HashSet",
			s:    NewHashSet(hashInt, eqInt, opts, 3, 5, 7),
			expected: NewHashSet(hashHashSet[int], eqHashSet[int], opts,
				NewHashSet(hashInt, eqInt, opts),
				NewHashSet(hashInt, eqInt, opts, 3),
				NewHashSet(hashInt, eqInt, opts, 5),
				NewHashSet(hashInt, eqInt, opts, 7),
				NewHashSet(hashInt, eqInt, opts, 3, 5),
				NewHashSet(hashInt, eqInt, opts, 3, 7),
				NewHashSet(hashInt, eqInt, opts, 5, 7),
				NewHashSet(hashInt, eqInt, opts, 3, 5, 7),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			powerset := Powerset(tc.s)
			assert.True(t, powerset.Equal(tc.expected))
		})
	}
}

func TestPartitions(t *testing.T) {
	opts := HashSetOpts{}
	hashInt := hash.HashFuncForInt[int](nil)
	eqInt := generic.NewEqualFunc[int]()
	cmpInt := generic.NewCompareFunc[int]()
	eqSet := func(a, b Set[int]) bool { return a.Equal(b) }
	eqPartition := func(a, b Set[Set[int]]) bool { return a.Equal(b) }

	tests := []struct {
		name     string
		s        Set[int]
		expected Set[Set[Set[int]]]
	}{
		{
			name: "Empty_Set",
			s:    New(eqInt),
			expected: New(eqPartition,
				New(eqSet),
			),
		},
		{
			name: "NonEmpty_Set",
			s:    New(eqInt, 3, 5, 7),
			expected: New(eqPartition,
				New(eqSet,
					New(eqInt, 3),
					New(eqInt, 5),
					New(eqInt, 7),
				),
				New(eqSet,
					New(eqInt, 7),
					New(eqInt, 3, 5),
				),
				New(eqSet,
					New(eqInt, 5),
					New(eqInt, 3, 7),
				),
				New(eqSet,
					New(eqInt, 3),
					New(eqInt, 5, 7),
				),
				New(eqSet,
					New(eqInt, 3, 5, 7),
				),
			),
		},
		{
			name: "NonEmpty_StableSet",
			s:    NewStableSet(eqInt, 3, 5, 7),
			expected: NewStableSet(eqPartition,
				NewStableSet(eqSet,
					NewStableSet(eqInt, 3),
					NewStableSet(eqInt, 5),
					NewStableSet(eqInt, 7),
				),
				NewStableSet(eqSet,
					NewStableSet(eqInt, 7),
					NewStableSet(eqInt, 3, 5),
				),
				NewStableSet(eqSet,
					NewStableSet(eqInt, 5),
					NewStableSet(eqInt, 3, 7),
				),
				NewStableSet(eqSet,
					NewStableSet(eqInt, 3),
					NewStableSet(eqInt, 5, 7),
				),
				NewStableSet(eqSet,
					NewStableSet(eqInt, 3, 5, 7),
				),
			),
		},
		{
			name: "NonEmpty_SortedSet",
			s:    NewSortedSet(cmpInt, 3, 5, 7),
			expected: NewSortedSet(cmpSortedPartition[int],
				NewSortedSet(cmpSortedSet[int],
					NewSortedSet(cmpInt, 3),
					NewSortedSet(cmpInt, 5),
					NewSortedSet(cmpInt, 7),
				),
				NewSortedSet(cmpSortedSet[int],
					NewSortedSet(cmpInt, 7),
					NewSortedSet(cmpInt, 3, 5),
				),
				NewSortedSet(cmpSortedSet[int],
					NewSortedSet(cmpInt, 5),
					NewSortedSet(cmpInt, 3, 7),
				),
				NewSortedSet(cmpSortedSet[int],
					NewSortedSet(cmpInt, 3),
					NewSortedSet(cmpInt, 5, 7),
				),
				NewSortedSet(cmpSortedSet[int],
					NewSortedSet(cmpInt, 3, 5, 7),
				),
			),
		},
		{
			name: "NonEmpty_HashSet",
			s:    NewHashSet(hashInt, eqInt, opts, 3, 5, 7),
			expected: NewHashSet(hashHashPartition[int], eqHashPartition[int], opts,
				NewHashSet(hashHashSet[int], eqHashSet[int], opts,
					NewHashSet(hashInt, eqInt, opts, 3),
					NewHashSet(hashInt, eqInt, opts, 5),
					NewHashSet(hashInt, eqInt, opts, 7),
				),
				NewHashSet(hashHashSet[int], eqHashSet[int], opts,
					NewHashSet(hashInt, eqInt, opts, 7),
					NewHashSet(hashInt, eqInt, opts, 3, 5),
				),
				NewHashSet(hashHashSet[int], eqHashSet[int], opts,
					NewHashSet(hashInt, eqInt, opts, 5),
					NewHashSet(hashInt, eqInt, opts, 3, 7),
				),
				NewHashSet(hashHashSet[int], eqHashSet[int], opts,
					NewHashSet(hashInt, eqInt, opts, 3),
					NewHashSet(hashInt, eqInt, opts, 5, 7),
				),
				NewHashSet(hashHashSet[int], eqHashSet[int], opts,
					NewHashSet(hashInt, eqInt, opts, 3, 5, 7),
				),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			partitions := Partitions(tc.s)
			assert.True(t, partitions.Equal(tc.expected))
		})
	}
}
