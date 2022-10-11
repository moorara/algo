package automata

import "github.com/moorara/algo/generic"

// markEntry associates a boolean field with each value.
type markEntry[T any] struct {
	val    T
	marked bool
}

// markList is an auxiliary data structure for the subset construction algorithm.
type markList[T any] struct {
	eq   generic.EqualFunc[T]
	list []markEntry[T]
}

// newMarkList creates and returns a new mark list.
func newMarkList[T any](eq generic.EqualFunc[T]) *markList[T] {
	return &markList[T]{
		eq:   eq,
		list: make([]markEntry[T], 0),
	}
}

// Values return the list of all values.
func (l *markList[T]) Values() []T {
	vals := make([]T, len(l.list))
	for i, e := range l.list {
		vals[i] = e.val
	}

	return vals
}

// AddUnmarked adds a new values to the list as unmarked and returns the index of the new value.
func (l *markList[T]) AddUnmarked(val T) int {
	l.list = append(l.list, markEntry[T]{
		val:    val,
		marked: false,
	})

	return len(l.list) - 1
}

// GetUnmarked returns the first unmarked value from the list.
func (l *markList[T]) GetUnmarked() (T, int) {
	for i, e := range l.list {
		if !e.marked {
			return e.val, i
		}
	}

	var zeroT T
	return zeroT, -1
}

// Contains determines whether or not the list contains a given value.
func (l *markList[T]) Contains(val T) int {
	for i, e := range l.list {
		if l.eq(e.val, val) {
			return i
		}
	}

	return -1
}

// MarkByIndex marks a value using its index in the list.
func (l *markList[T]) MarkByIndex(i int) {
	if 0 <= i && i < len(l.list) {
		l.list[i].marked = true
	}
}
