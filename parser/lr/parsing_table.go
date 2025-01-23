package lr

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/moorara/algo/errors"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// ParsingTable represents a parsing table for an LR parser.
type ParsingTable struct {
	states       []State
	terminals    []grammar.Terminal
	nonTerminals []grammar.NonTerminal
	actions      symboltable.SymbolTable[State, symboltable.SymbolTable[grammar.Terminal, set.Set[*Action]]]
	gotos        symboltable.SymbolTable[State, symboltable.SymbolTable[grammar.NonTerminal, State]]
}

// NewParsingTable creates an empty parsing table for an LR parser.
func NewParsingTable(states []State, terminals []grammar.Terminal, nonTerminals []grammar.NonTerminal) *ParsingTable {
	opts := symboltable.HashOpts{}

	actions := symboltable.NewQuadraticHashTable(
		HashState,
		EqState,
		func(lhs, rhs symboltable.SymbolTable[grammar.Terminal, set.Set[*Action]]) bool {
			return lhs.Equals(rhs)
		},
		opts,
	)

	gotos := symboltable.NewQuadraticHashTable(
		HashState,
		EqState,
		func(lhs, rhs symboltable.SymbolTable[grammar.NonTerminal, State]) bool {
			return lhs.Equals(rhs)
		},
		opts,
	)

	return &ParsingTable{
		states:       states,
		terminals:    append(terminals, grammar.Endmarker),
		nonTerminals: nonTerminals,
		actions:      actions,
		gotos:        gotos,
	}
}

func (t *ParsingTable) getActions(s State, a grammar.Terminal) (set.Set[*Action], bool) {
	if row, ok := t.actions.Get(s); ok {
		if actions, ok := row.Get(a); ok {
			if actions != nil {
				return actions, true
			}
		}
	}

	return nil, false
}

func (t *ParsingTable) getActionString(s State, a grammar.Terminal) string {
	set, ok := t.getActions(s, a)
	if !ok || set.Size() == 0 {
		return ""
	}

	actions := generic.Collect1(set.All())
	sort.Quick(actions, cmpAction)

	var b bytes.Buffer

	for _, a := range actions {
		fmt.Fprintf(&b, "%s ┆ ", a)
	}
	b.Truncate(b.Len() - 5)

	return b.String()
}

func (t *ParsingTable) getGotoString(s State, A grammar.NonTerminal) string {
	row, ok := t.gotos.Get(s)
	if !ok {
		return ""
	}

	state, ok := row.Get(A)
	if !ok || state == ErrState {
		return ""
	}

	return strconv.Itoa(int(state))
}

// String returns a string representation of a parsing table.
func (t *ParsingTable) String() string {
	ts := &tableStringer[State, grammar.Terminal, grammar.NonTerminal]{
		K1Title:  "STATE",
		K1Values: t.states,
		K2Title:  "ACTION",
		K2Values: t.terminals,
		K3Title:  "GOTO",
		K3Values: t.nonTerminals,
		GetK1K2:  t.getActionString,
		GetK1K3:  t.getGotoString,
	}

	return ts.String()
}

// Equals determines whether or not two parsing tables are the same.
func (t *ParsingTable) Equals(rhs *ParsingTable) bool {
	return t.actions.Equals(rhs.actions) &&
		t.gotos.Equals(rhs.gotos)
}

// AddACTION adds a new action for state s and terminal a to the parsing table.
// Multiple actions can be added for the same state s and terminal a.
// It returns false if the ACTION[s,a] contains more than one action, indicating a conflict.
func (t *ParsingTable) AddACTION(s State, a grammar.Terminal, action *Action) bool {
	if _, ok := t.actions.Get(s); !ok {
		t.actions.Put(s, symboltable.NewQuadraticHashTable(
			grammar.HashTerminal,
			grammar.EqTerminal,
			eqActionSet,
			symboltable.HashOpts{},
		))
	}

	row, _ := t.actions.Get(s)
	if _, ok := row.Get(a); !ok {
		row.Put(a, set.New[*Action](eqAction))
	}

	actions, _ := row.Get(a)
	actions.Add(action)

	return actions.Size() == 1
}

// SetGOTO updates the next state for state s and non-terminal A in the parsing table.
// If the next state is ErrState, it will not be added to the table.
func (t *ParsingTable) SetGOTO(s State, A grammar.NonTerminal, next State) {
	if next == ErrState {
		return
	}

	if _, ok := t.gotos.Get(s); !ok {
		t.gotos.Put(s, symboltable.NewQuadraticHashTable(
			grammar.HashNonTerminal,
			grammar.EqNonTerminal,
			EqState,
			symboltable.HashOpts{},
		))
	}

	row, _ := t.gotos.Get(s)
	row.Put(A, next)
}

// Error checks the parsing table for any conflicts between actions.
// A conflict occurs when multiple actions are assigned to the same state and terminal symbol.
// Conflicts arise when the grammar is ambiguous.
// If any conflicts are found, it returns an error with detailed descriptions of the conflicts.
func (t *ParsingTable) Error() error {
	var err = &errors.MultiError{
		Format: errors.BulletErrorFormat,
	}

	// Check for ACTION conflicts.
	for _, s := range t.states {
		for _, a := range t.terminals {
			if actions, ok := t.getActions(s, a); ok {
				if actions.Size() > 1 {
					err = errors.Append(err, &ParsingTableError{
						Type:    CONFLICT,
						State:   s,
						Symbol:  a,
						Actions: actions,
					})
				}
			}
		}
	}

	return err.ErrorOrNil()
}

// ACTION looks up and returns the action for state s and terminal a.
// If the ACTION[s,a] contains more than one action,
// it returns an erroneous ACTION and an error, indicating a conflict.
func (t *ParsingTable) ACTION(s State, a grammar.Terminal) (*Action, error) {
	actions, ok := t.getActions(s, a)
	if !ok || actions.Size() == 0 {
		return &Action{Type: ERROR}, &ParsingTableError{
			Type:   NO_ACTION,
			State:  s,
			Symbol: a,
		}
	}

	if actions.Size() == 1 {
		for action := range actions.All() {
			return action, nil
		}
	}

	// Conflict
	return &Action{Type: ERROR}, &ParsingTableError{
		Type:    CONFLICT,
		State:   s,
		Symbol:  a,
		Actions: actions,
	}
}

// GOTO looks up and returns the next state for state s and non-terminal A.
// If the GOTO[s,A] contains more than one state, it returns an error.
func (t *ParsingTable) GOTO(s State, A grammar.NonTerminal) (State, error) {
	row, ok := t.gotos.Get(s)
	if !ok {
		return ErrState, &ParsingTableError{
			Type:   NO_GOTO,
			State:  s,
			Symbol: A,
		}
	}

	state, ok := row.Get(A)
	if !ok || state == ErrState {
		return ErrState, &ParsingTableError{
			Type:   NO_GOTO,
			State:  s,
			Symbol: A,
		}
	}

	return state, nil
}

type ParsingTableErrorType int

const (
	NO_ACTION ParsingTableErrorType = 1 + iota // No action for the given state s and terminal a.
	NO_GOTO                                    // No next state for the given state s and non-terminal A.
	CONFLICT                                   // Conflict (Shift/Reduce or Reduce/Reduce) for the given state s and terminal a.
)

// ParsingTableError represents an error encountered in an LR parsing table.
// This error occurs when there is ambiguity in the grammar or when the input is unacceptable.
type ParsingTableError struct {
	Type    ParsingTableErrorType
	State   State
	Symbol  grammar.Symbol
	Actions set.Set[*Action]
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *ParsingTableError) Error() string {
	var b bytes.Buffer

	if e.Type == NO_ACTION {
		fmt.Fprintf(&b, "no action for ACTION[%d, %s]", e.State, e.Symbol)
	} else if e.Type == NO_GOTO {
		fmt.Fprintf(&b, "no state for GOTO[%d, %s]", e.State, e.Symbol)
	} else if e.isSRConflict() {
		fmt.Fprintf(&b, "shift/reduce conflict at ACTION[%d, %s]\n", e.State, e.Symbol)
	} else if e.isRRConflict() {
		fmt.Fprintf(&b, "reduce/reduce conflict at ACTION[%d, %s]\n", e.State, e.Symbol)
	} else {
		fmt.Fprintf(&b, "invalid error: %d", e.Type)
	}

	if e.Type == CONFLICT {
		actions := generic.Collect1(e.Actions.All())
		sort.Quick(actions, cmpAction)

		for _, action := range actions {
			fmt.Fprintf(&b, "  %s\n", action)
		}
	}

	return b.String()
}

// isSRConflict determines whether or not the error is a Shift/Reduce conflict.
func (e *ParsingTableError) isSRConflict() bool {
	return e.Type == CONFLICT &&
		e.Actions.AnyMatch(func(action *Action) bool {
			return action.Type == SHIFT
		}) && e.Actions.AnyMatch(func(action *Action) bool {
		return action.Type == REDUCE
	})
}

// isRRConflict determines whether or not the error is a Reduce/Reduce conflict.
func (e *ParsingTableError) isRRConflict() bool {
	return e.Type == CONFLICT &&
		e.Actions.AllMatch(func(action *Action) bool {
			return action.Type == REDUCE
		})
}
