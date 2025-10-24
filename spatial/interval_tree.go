package spatial

import "cmp"

// Bound represents a mathematical bound, which can be either open or closed.
type Bound[T cmp.Ordered] struct {
	Value T
	Open  bool
}

// Interval represents a mathematical interval.
type Interval[T cmp.Ordered] struct {
	Start Bound[T]
	End   Bound[T]
}

// intervalTreeNode represents a node in the interval tree.
type intervalTreeNode[K cmp.Ordered, V any] struct {
	in          Interval[K]
	max         Bound[K]
	val         V
	left, right *intervalTreeNode[K, V]
}

// IntervalTree represents an interval tree data structure.
// An interval tree is used to store a set of intervals and allows efficient queries to
// find all intervals that overlap with a given point or another interval.
//
// The interval tree is a balanced binary search tree (BST) where each node stores an interval.
// The tree uses the start point of each interval as the key, and each node is augmented with the maximum end point in its subtree.
// Building the tree from n intervals takes O(nlogn) time.
// Inserting or deleting an interval takes O(logn) time, assuming the tree remains balanced.
// To find all intervals overlapping a given interval or point, queries take O(logn + k) time, where k is the number of results returned.
type IntervalTree[K cmp.Ordered, V any] struct {
	root *intervalTreeNode[K, V]
}

func NewIntervalTree[K cmp.Ordered, V any]() *IntervalTree[K, V] {
	return &IntervalTree[K, V]{
		root: nil,
	}
}

func (t *IntervalTree[K, V]) Add(in Interval[K], val V) {

}

func (t *IntervalTree[K, V]) Remove(in Interval[K]) {

}

func (t *IntervalTree[K, V]) SearchPoint(pt K) {

}

func (t *IntervalTree[K, V]) SearchInterval(in Interval[K]) {

}
