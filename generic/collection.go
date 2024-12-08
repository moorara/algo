package generic

import "iter"

type (
	// Predicate1 is a generic function type that takes a value
	//   and returns a boolean value, used for evaluating a condition.
	Predicate1[T any] func(T) bool

	// Predicate2 is a generic function type that takes a key-value pair
	//   and returns a boolean value, used for evaluating a condition.
	Predicate2[K, V any] func(K, V) bool
)

type (
	// Collection1 is a generic interface for a collection of items.
	Collection1[T any] interface {
		// All returns an iterator sequence containing all the items in the collection.
		// This allows for iterating over the entire collection using the range keyword.
		All() iter.Seq[T]

		// AnyMatch returns true if at least one item in the collection satisfies the provided predicate function.
		// If no items satisfy the predicate or the collection is empty,
		// it returns false.
		AnyMatch(Predicate1[T]) bool

		// AllMatch returns true if all items in the collection satisfy the provided predicate function.
		// If the collection is empty, it returns false.
		AllMatch(Predicate1[T]) bool
	}

	// Collection2 is a generic interface for a collection of key-value pairs.
	Collection2[K, V any] interface {
		// All returns an iterator sequence containing all the key-value pairs in the collection.
		// This allows for iterating over the entire collection using the range keyword.
		All() iter.Seq2[K, V]

		// AnyMatch returns true if at least one key-value pair in the collection satisfies the provided predicate function.
		// If no key-value pairs satisfy the predicate or the collection is empty, it returns false.
		AnyMatch(Predicate2[K, V]) bool

		// AllMatch returns true if all key-value pairs in the collection satisfy the provided predicate function.
		// If the collection is empty, it returns true.
		AllMatch(Predicate2[K, V]) bool
	}
)

// Collect collects key-value pairs in a collection from seq2 into a new slice and returns it.
func Collect[K, V any](seq2 iter.Seq2[K, V]) []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0)
	for k, v := range seq2 {
		kvs = append(kvs, KeyValue[K, V]{k, v})
	}

	return kvs
}
