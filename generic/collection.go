package generic

import "iter"

type (
	// Predicate1 is a generic function type that takes a value
	//   and returns a boolean value, used for evaluating a condition.
	Predicate1[T any] func(T) bool

	// Predicate2 is a generic function type that takes a key-value
	//   and returns a boolean value, used for evaluating a condition.
	Predicate2[K, V any] func(K, V) bool
)

type (
	// Collection1 is a generic interface for a collection of items.
	Collection1[T any] interface {
		// Size returns the number of items in the collection.
		Size() int

		// IsEmpty returns true if the collection contains no items.
		IsEmpty() bool

		// Add adds the specified items to the collection.
		// The behavior when adding an existing item depends on the implementation.
		Add(...T)

		// Remove removes the specified items from the collection.
		// If an item does not exist, it is ignored.
		Remove(...T)

		// RemoveAll removes all items from the collection, leaving it empty.
		RemoveAll()

		// Contains returns true if the collection includes all the specified items.
		Contains(...T) bool

		// All returns an iterator sequence containing all the items in the collection.
		// This allows for iterating over the entire collection using the range keyword.
		All() iter.Seq[T]

		// AnyMatch returns true if at least one item in the collection satisfies the provided predicate function.
		// If no items satisfy the predicate or the collection is empty, it returns false.
		AnyMatch(Predicate1[T]) bool

		// AllMatch returns true if all items in the collection satisfy the provided predicate function.
		// If the collection is empty, it returns true.
		AllMatch(Predicate1[T]) bool

		// SelectMatch selects a subset of items from the collection that satisfy the given predicate.
		// It returns a new collection containing the matching items, of the same type as the original collection.
		SelectMatch(Predicate1[T]) Collection1[T]
	}

	// Collection2 is a generic interface for a collection of key-values.
	Collection2[K, V any] interface {
		// Size returns the number of key-values in the collection.
		Size() int

		// IsEmpty returns true if the collection contains no key-values.
		IsEmpty() bool

		// Put adds a new key-value to the collection.
		// If the key already exists, its value is updated.
		Put(K, V)

		// Get returns the value associated with the given key in the collection.
		// If the key does not exist, the second return value will be false.
		Get(K) (V, bool)

		// Delete deletes the key-value specified by the given key from the collection.
		// If the key does not exist, the second return value will be false.
		Delete(K) (V, bool)

		// DeleteAll deletes all key-values from the collection, leaving it empty.
		DeleteAll()

		// All returns an iterator sequence containing all the key-values in the collection.
		// This allows for iterating over the entire collection using the range keyword.
		All() iter.Seq2[K, V]

		// AnyMatch returns true if at least one key-value in the collection satisfies the provided predicate function.
		// If no key-values satisfy the predicate or the collection is empty, it returns false.
		AnyMatch(Predicate2[K, V]) bool

		// AllMatch returns true if all key-values in the collection satisfy the provided predicate function.
		// If the collection is empty, it returns true.
		AllMatch(Predicate2[K, V]) bool

		// SelectMatch selects a subset of key-values from the collection that satisfy the given predicate.
		// It returns a new collection containing the matching key-values, of the same type as the original collection.
		SelectMatch(Predicate2[K, V]) Collection2[K, V]
	}
)
