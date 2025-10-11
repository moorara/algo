package cont

import (
	"bytes"
	"fmt"
	"iter"
)

// FormatList is a function type for formatting a range list into a single string representation.
type FormatList[T Continuous] func(iter.Seq[Range[T]]) string

func defaultFormatList[T Continuous](all iter.Seq[Range[T]]) string {
	var b bytes.Buffer

	for r := range all {
		fmt.Fprintf(&b, "%s ", r)
	}

	// Remove the last space
	if b.Len() > 0 {
		b.Truncate(b.Len() - 1)
	}

	return b.String()
}

// FormatMap is a function type for formatting a range map into a single string representation.
type FormatMap[T Continuous, V any] func(iter.Seq2[Range[T], V]) string

func defaultFormatMap[T Continuous, V any](all iter.Seq2[Range[T], V]) string {
	var b bytes.Buffer

	for r, v := range all {
		fmt.Fprintf(&b, "%s:%v ", r, v)
	}

	// Remove the last space
	if b.Len() > 0 {
		b.Truncate(b.Len() - 1)
	}

	return b.String()
}
