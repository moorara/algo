package parser

import (
	"fmt"

	"github.com/moorara/algo/generic"
)

// Table is the interface for a generic table.
// It abstracts away the string representation of any parsing table.
type Table interface {
	fmt.Stringer
	generic.Equaler[Table]

	// Rows
	getFirstKeyTitle() string
	getFirstKeyStrings() []string

	// Columns
	getSecondKeyTitle() string
	getSecondKeyStrings() []string

	// Cells
	getEntryString() string
}
