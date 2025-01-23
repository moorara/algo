package lr

import (
	"fmt"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

var (
	eqAction = func(lhs, rhs *Action) bool {
		return lhs.Equals(rhs)
	}

	eqActionSet = func(lhs, rhs set.Set[*Action]) bool {
		return lhs.Equals(rhs)
	}
)

// ActionType enumerates the possible types of actions in an LR parser.
type ActionType int

const (
	SHIFT  ActionType = 1 + iota // Advance to the next state by consuming input.
	REDUCE                       // Apply a production to reduce symbols on the stack.
	ACCEPT                       // Accept the input as successfully parsed.
	ERROR                        // Signal an error in parsing.
)

// Action represents an action in the LR parsing table or automaton.
type Action struct {
	Type       ActionType
	State      State               // Only set for SHIFT actions
	Production *grammar.Production // Only set for REDUCE actions
}

// String returns a string representation of an action.
func (a *Action) String() string {
	switch a.Type {
	case SHIFT:
		return fmt.Sprintf("SHIFT %d", a.State)
	case REDUCE:
		return fmt.Sprintf("REDUCE %s", a.Production)
	case ACCEPT:
		return "ACCEPT"
	case ERROR:
		return "ERROR"
	}

	return fmt.Sprintf("INVALID ACTION(%d)", a.Type)
}

// Equals determines whether or not two actions are the same.
func (a *Action) Equals(rhs *Action) bool {
	return a.Type == rhs.Type &&
		a.State == rhs.State &&
		equalProductions(a.Production, rhs.Production)
}

func equalProductions(lhs, rhs *grammar.Production) bool {
	if lhs == nil || rhs == nil {
		return lhs == rhs
	}
	return lhs.Equals(rhs)
}

func cmpAction(lhs, rhs *Action) int {
	if lhs.Type == SHIFT && rhs.Type == SHIFT {
		return int(lhs.State) - int(rhs.State)
	} else if lhs.Type == REDUCE && rhs.Type == REDUCE {
		return grammar.CmpProduction(lhs.Production, rhs.Production)
	}

	return int(lhs.Type) - int(rhs.Type)
}
