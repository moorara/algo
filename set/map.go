package set

import "github.com/moorara/algo/generic"

// Mapper is a function for converting a member of a set from one type to another type.
type Mapper[T, U any] func(T) U

// Map converts a set from one type to another type.
// You need to provide a compare function for the new type.
func (f Mapper[T, U]) Map(s Set[T], equal generic.EqualFunc[U]) Set[U] {
	members := make([]U, 0)
	for _, m := range s.Members() {
		members = append(members, f(m))
	}

	return &set[U]{
		equal:   equal,
		members: members,
	}
}
