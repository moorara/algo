package set

import (
	"fmt"
	"strings"
)

// Format is a function type for formatting a set into a single string representation.
type Format[T any] func([]T) string

func defaultFormat[T any](members []T) string {
	vals := make([]string, len(members))
	for i, m := range members {
		vals[i] = fmt.Sprintf("%v", m)
	}

	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}
