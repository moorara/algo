package set

import "github.com/moorara/algo/generic"

// Mapper
type Mapper[T, U any] func(T) U

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
