package lr

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
)

// ConflictError represents a conflict in an LR parsing table.
// A conflict occurs when the grammar is ambiguous, resulting in multiple actions
// being associated with a specific state s and terminal a in the ACTION table.
// A conflict is either a shift/reduce conflict or a reduce/reduce conflict.
type ConflictError struct {
	State    State
	Terminal grammar.Terminal
	Actions  set.Set[*Action]
}

// IsShiftReduce returns true if the conflict is a shift/reduce conflict.
// A shift/reduce conflict arises when there is at least one SHIFT action
// and one REDUCE action associated with the same state and terminal.
func (e *ConflictError) IsShiftReduce() bool {
	return e.Actions.AnyMatch(func(a *Action) bool {
		return a.Type == SHIFT
	}) && e.Actions.AnyMatch(func(a *Action) bool {
		return a.Type == REDUCE
	})
}

// IsReduceReduce returns true if the conflict is a reduce/reduce conflict.
// A reduce/reduce conflict occurs when all actions associated with a state
// and terminal are REDUCE actions, and there is more than one such action.
func (e *ConflictError) IsReduceReduce() bool {
	return e.Actions.Size() > 1 &&
		e.Actions.AllMatch(func(a *Action) bool {
			return a.Type == REDUCE
		})
}

// Handles generates a set of precedence handles derived from the conflict actions.
// A precedence handle is either a terminal symbol or a production rule that is used
// to define associativity and precedence for resolving conflicts in the parsing table.
// By specifying the associativity and precedence for these handles,
// shift/reduce and reduce/reduce conflicts can be resolved.
func (e *ConflictError) Handles() set.Set[*PrecedenceHandle] {
	dedup := set.New(eqPrecedenceHandle)

	for a := range e.Actions.All() {
		switch a.Type {
		case SHIFT:
			dedup.Add(&PrecedenceHandle{
				Terminal: &e.Terminal,
			})

		case REDUCE:
			first, ok := generic.FirstMatch(a.Production.Body, func(s grammar.Symbol) bool {
				return s.IsTerminal()
			})

			if ok {
				term := first.(grammar.Terminal)
				dedup.Add(&PrecedenceHandle{
					Terminal: &term,
				})
			} else {
				dedup.Add(&PrecedenceHandle{
					Production: a.Production,
				})
			}
		}
	}

	return dedup
}

// Error returns a detailed string representation of the conflict error.
func (e *ConflictError) Error() string {
	var b bytes.Buffer

	b.WriteString("Error:      Ambiguous Grammar\n")

	if e.IsShiftReduce() {
		fmt.Fprintf(&b, "Cause:      Shift/Reduce conflict in ACTION[%d, %s]\n", e.State, e.Terminal)
	} else if e.IsReduceReduce() {
		fmt.Fprintf(&b, "Cause:      Reduce/Reduce conflict in ACTION[%d, %s]\n", e.State, e.Terminal)
	}

	b.WriteString("Context:    The parser cannot decide whether to\n")

	actions := generic.Collect1(e.Actions.All())
	sort.Insertion(actions, cmpAction)

	for i, a := range actions {
		fmt.Fprintf(&b, "              %d. ", i+1)

		switch a.Type {
		case SHIFT:
			fmt.Fprintf(&b, "Shift the terminal %s", e.Terminal)
		case REDUCE:
			fmt.Fprintf(&b, "Reduce by production %s", a.Production)
		}

		if i < len(actions)-1 {
			b.WriteString(", or\n")
		} else {
			b.WriteString("\n")
		}
	}

	handles := generic.Collect1(e.Handles().All())
	sort.Insertion(handles, cmpPrecedenceHandle)

	if len(handles) == 1 {
		fmt.Fprintf(&b, "Resolution: Specify associativity for the %s in the grammar directives.\n", handles[0])
	} else {
		b.WriteString("Resolution: Specify precedence for the following in the grammar directives:\n")
		for _, handle := range handles {
			fmt.Fprintf(&b, "              • %s\n", handle)
		}
		b.WriteString("            Terminals or Productions listed earlier in the directives will have higher precedence.\n")
	}

	return b.String()
}

// AggregatedConflictError represents a collection of conflict errors.
// It is used to accumulate multiple conflict errors and generate a consolidated,
// more concise error message that is more useful for resolving all conflicts at once.
type AggregatedConflictError []*ConflictError

// ErrorOrNil returns an error if the AggregatedConflictError contains any errors, or nil if it has none.
// This method is useful for ensuring that a valid error value is returned after accumulating errors,
// indicating whether errors are present or not.
func (e AggregatedConflictError) ErrorOrNil() error {
	if len(e) == 0 {
		return nil
	}

	return e
}

// Error builds and returns a consolidated string representation of the collection of conflict errors.
// The resulting message is more concise and provides a clearer summary of all the conflicts.
func (e AggregatedConflictError) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	}

	var b bytes.Buffer

	b.WriteString("Error:      Ambiguous Grammar\n")
	b.WriteString("Cause:      Multiple conflicts in the parsing table:\n")

	for i, err := range e {
		fmt.Fprintf(&b, "              %d. ", i+1)

		if err.IsShiftReduce() {
			fmt.Fprintf(&b, "Shift/Reduce conflict in ACTION[%d, %s]\n", err.State, err.Terminal)
		} else if err.IsReduceReduce() {
			fmt.Fprintf(&b, "Reduce/Reduce conflict in ACTION[%d, %s]\n", err.State, err.Terminal)
		}
	}

	set := set.New(eqPrecedenceHandle)
	for _, err := range e {
		set = set.Union(err.Handles())
	}

	handles := generic.Collect1(set.All())
	sort.Insertion(handles, cmpPrecedenceHandle)

	if len(handles) == 1 {
		fmt.Fprintf(&b, "Resolution: Specify associativity for the %s in the grammar directives.\n", handles[0])
	} else {
		b.WriteString("Resolution: Specify precedence for the following in the grammar directives:\n")
		for _, handle := range handles {
			fmt.Fprintf(&b, "              • %s\n", handle)
		}
		b.WriteString("            Terminals or Productions listed earlier in the directives will have higher precedence.\n")
	}

	return b.String()
}

// Unwrap implements the unwrap interface for AggregatedConflictError.
// It returns the slice of accumulated errors wrapped in the AggregatedConflictError instance.
// If there are no errors, it returns nil, indicating that e does not wrap any error.
func (e AggregatedConflictError) Unwrap() []error {
	if len(e) == 0 {
		return nil
	}

	errs := make([]error, len(e))
	for i, err := range e {
		errs[i] = err
	}

	return errs
}
