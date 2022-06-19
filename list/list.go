// Package list implements list data structures.
package list

type arrayNode[T any] struct {
	block []T
	next  *arrayNode[T]
}

func newArrayNode[T any](size int, next *arrayNode[T]) *arrayNode[T] {
	return &arrayNode[T]{
		block: make([]T, size),
		next:  next,
	}
}
