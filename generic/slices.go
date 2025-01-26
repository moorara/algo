package generic

import "iter"

// KeyValue is a generic struct that holds a key-value pair.
// K and V represent the types of the key and value, respectively
type KeyValue[K, V any] struct {
	Key K
	Val V
}

// Collect1 collects items in a collection from seq into a new slice and returns it.
func Collect1[T any](seq iter.Seq[T]) []T {
	items := make([]T, 0)
	for v := range seq {
		items = append(items, v)
	}

	return items
}

// Collect2 collects key-values in a collection from seq2 into a new slice and returns it.
func Collect2[K, V any](seq2 iter.Seq2[K, V]) []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0)
	for k, v := range seq2 {
		kvs = append(kvs, KeyValue[K, V]{k, v})
	}

	return kvs
}

// AnyMatch returns true if at least one item in a slice satisfies the provided predicate function.
// If no items satisfy the predicate or a slice is empty, it returns false.
func AnyMatch[T any](s []T, p Predicate1[T]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}

	return false
}

// AllMatch returns true if all items in a slice satisfy the provided predicate function.
// If a slice is empty, it returns true.
func AllMatch[T any](s []T, p Predicate1[T]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}

	return true
}

// SelectMatch selects a subset of items from a slice that satisfy the given predicate.
// It returns a new slice containing the matching items, of the same type as the original slice.
func SelectMatch[T any](s []T, p Predicate1[T]) []T {
	ss := make([]T, 0)

	for _, v := range s {
		if p(v) {
			ss = append(ss, v)
		}
	}

	return ss
}
