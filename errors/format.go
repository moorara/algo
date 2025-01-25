package errors

import (
	"bytes"
	"fmt"
	"strings"
)

// ErrorFormat is a function type for formatting a slice of errors into a single string representation.
type ErrorFormat func([]error) string

var (
	// DefaultErrorFormat formats a slice of errors into a single string,
	// with each error on a new line.
	DefaultErrorFormat = defaultErrorFormat

	// BulletErrorFormat formats a slice of errors into a single string,
	// with each error indented, prefixed by a bullet point, and properly spaced.
	BulletErrorFormat = bulletErrorFormat
)

func defaultErrorFormat(errs []error) string {
	if len(errs) == 0 {
		return ""
	}

	var b bytes.Buffer

	for _, err := range errs {
		fmt.Fprintln(&b, err)
	}

	return b.String()
}

func bulletErrorFormat(errs []error) string {
	if len(errs) == 0 {
		return ""
	}

	var b bytes.Buffer

	if len(errs) == 1 {
		b.WriteString("1 error occurred:\n\n")
	} else {
		fmt.Fprintf(&b, "%d errors occurred:\n\n", len(errs))
	}

	for _, err := range errs {
		for i, line := range strings.Split(err.Error(), "\n") {
			if i == 0 {
				fmt.Fprintf(&b, "  • %s\n", line)
			} else {
				fmt.Fprintf(&b, "    %s\n", line)
			}
		}
	}

	return b.String()
}
