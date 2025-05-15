// Package trie implements prefix tree data structures.
package trie

import (
	"fmt"

	"github.com/moorara/algo/generic"
)

// Trie represents a trie (prefix tree) abstract data type.
type Trie[V any] interface {
	fmt.Stringer
	generic.Equaler[Trie[V]]
	generic.Collection2[string, V]
	generic.Tree2[string, V]

	verify() bool

	Height() int
	Min() (string, V, bool)
	Max() (string, V, bool)
	Floor(string) (string, V, bool)
	Ceiling(string) (string, V, bool)
	DeleteMin() (string, V, bool)
	DeleteMax() (string, V, bool)
	Select(int) (string, V, bool)
	Rank(string) int
	Range(string, string) []generic.KeyValue[string, V]
	RangeSize(string, string) int
	Match(string) []generic.KeyValue[string, V]
	WithPrefix(string) []generic.KeyValue[string, V]
	LongestPrefixOf(string) (string, V, bool)
}
