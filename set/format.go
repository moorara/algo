package set

import (
	"fmt"
	"strings"
)

// StringFormat is a function type for formatting a set into a single string representation.
type StringFormat[T any] func([]T) string

func defaultStringFormat[T any](members []T) string {
	vals := make([]string, len(members))
	for i, m := range members {
		vals[i] = fmt.Sprintf("%v", m)
	}

	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}
