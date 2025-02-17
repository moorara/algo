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
	if seq == nil {
		return nil
	}

	items := make([]T, 0)
	for v := range seq {
		items = append(items, v)
	}

	return items
}

// Collect2 collects key-values in a collection from seq2 into a new slice and returns it.
func Collect2[K, V any](seq2 iter.Seq2[K, V]) []KeyValue[K, V] {
	if seq2 == nil {
		return nil
	}

	kvs := make([]KeyValue[K, V], 0)
	for k, v := range seq2 {
		kvs = append(kvs, KeyValue[K, V]{k, v})
	}

	return kvs
}

// Find searches for the index of the first occurrence of a value in a slice.
// If the value exists in the slice, it returns its index; otherwise, it returns -1.
func Find[T any](s []T, eq EqualFunc[T], val T) int {
	for i, v := range s {
		if eq(v, val) {
			return i
		}
	}

	return -1
}

// Contains checks if a slice contains all the specified values.
// It returns true if all values are found in the slice, otherwise returns false.
func Contains[T any](s []T, eq EqualFunc[T], vals ...T) bool {
	for _, val := range vals {
		if Find(s, eq, val) == -1 {
			return false
		}
	}

	return true
}

// AnyMatch returns true if at least one element in a slice satisfies the provided predicate function.
// If no elements satisfy the predicate or a slice is empty, it returns false.
func AnyMatch[T any](s []T, p Predicate1[T]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}

	return false
}

// AllMatch returns true if all elements in a slice satisfy the provided predicate function.
// If a slice is empty, it returns true.
func AllMatch[T any](s []T, p Predicate1[T]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}

	return true
}

// FirstMatch returns the first element in a slice that satisfies the given predicate.
// If no match is found, it returns the zero value of T and false.
func FirstMatch[T any](s []T, p Predicate1[T]) (T, bool) {
	for _, v := range s {
		if p(v) {
			return v, true
		}
	}

	var zeroT T
	return zeroT, false
}

// SelectMatch selects a subset of elements from a slice that satisfy the given predicate.
// It returns a new slice containing the matching elements, of the same type as the original slice.
func SelectMatch[T any](s []T, p Predicate1[T]) []T {
	matched := make([]T, 0)

	for _, v := range s {
		if p(v) {
			matched = append(matched, v)
		}
	}

	return matched
}

// PartitionMatch partitions the elements in a slice
// into two separate slices based on the provided predicate.
// The first slice contains the elements that satisfy the predicate (matched elements),
// while the second slice contains those that do not satisfy the predicate (unmatched elements).
// Both slices are of the same type as the original slice.
func PartitionMatch[T any](s []T, p Predicate1[T]) ([]T, []T) {
	matched := make([]T, 0)
	unmatched := make([]T, 0)

	for _, v := range s {
		if p(v) {
			matched = append(matched, v)
		} else {
			unmatched = append(unmatched, v)
		}
	}

	return matched, unmatched
}

// Transform applies a given transformation function to each element of a slice,
// converting it from type T to type U, and returns a new slice.
func Transform[T, U any](s []T, f func(T) U) []U {
	u := make([]U, len(s))
	for i := range s {
		u[i] = f(s[i])
	}

	return u
}
