package lr

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/moorara/algo/errors"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// NewParsingTable creates an empty parsing table for an LR parser.
func NewParsingTable(states []State, terminals []grammar.Terminal, nonTerminals []grammar.NonTerminal) *ParsingTable {
	opts := symboltable.HashOpts{}

	actions := symboltable.NewQuadraticHashTable(
		HashState,
		EqState,
		func(lhs, rhs symboltable.SymbolTable[grammar.Terminal, set.Set[*Action]]) bool {
			return lhs.Equal(rhs)
		},
		opts,
	)

	gotos := symboltable.NewQuadraticHashTable(
		HashState,
		EqState,
		func(lhs, rhs symboltable.SymbolTable[grammar.NonTerminal, State]) bool {
			return lhs.Equal(rhs)
		},
		opts,
	)

	return &ParsingTable{
		states:       states,
		terminals:    terminals,
		nonTerminals: nonTerminals,
		actions:      actions,
		gotos:        gotos,
	}
}

// ParsingTable represents an LR parsing table.
type ParsingTable struct {
	states       []State
	terminals    []grammar.Terminal
	nonTerminals []grammar.NonTerminal
	actions      symboltable.SymbolTable[State, symboltable.SymbolTable[grammar.Terminal, set.Set[*Action]]]
	gotos        symboltable.SymbolTable[State, symboltable.SymbolTable[grammar.NonTerminal, State]]
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

// String returns a human-readable string representation of the parsing table.
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

// Equal determines whether or not two parsing tables are the same.
func (t *ParsingTable) Equal(rhs *ParsingTable) bool {
	return t.actions.Equal(rhs.actions) &&
		t.gotos.Equal(rhs.gotos)
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
			Type:   MISSING_ACTION,
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
			Type:   MISSING_GOTO,
			State:  s,
			Symbol: A,
		}
	}

	state, ok := row.Get(A)
	if !ok || state == ErrState {
		return ErrState, &ParsingTableError{
			Type:   MISSING_GOTO,
			State:  s,
			Symbol: A,
		}
	}

	return state, nil
}

// ParsingTableErrorType represents the type of error associated with an LR parsing table.
type ParsingTableErrorType int

const (
	MISSING_ACTION ParsingTableErrorType = 1 + iota // No action for the given state s and terminal a.
	MISSING_GOTO                                    // No next state for the given state s and non-terminal A.
	CONFLICT                                        // Conflict (Shift/Reduce or Reduce/Reduce) for the given state s and terminal a.
)

// ParsingTableError represents an error encountered in an LR parsing table.
// This error occurs when there is ambiguity in the grammar or when the input is unacceptable.
type ParsingTableError struct {
	Type    ParsingTableErrorType
	State   State
	Symbol  grammar.Symbol
	Actions set.Set[*Action] // Only set for CONFLICT errors
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *ParsingTableError) Error() string {
	var b bytes.Buffer

	switch e.Type {
	case MISSING_ACTION:
		fmt.Fprintf(&b, "no action exists in the parsing table for ACTION[%d, %s]", e.State, e.Symbol)

	case MISSING_GOTO:
		fmt.Fprintf(&b, "no state exists in the parsing table for GOTO[%d, %s]", e.State, e.Symbol)

	case CONFLICT:
		if s, r, ok := e.isShiftReduceConflict(); ok {
			fmt.Fprintf(&b, "AMBIGUOUS Grammar: shift/reduce conflict in ACTION[%d, %s]\n", e.State, e.Symbol)
			// cannot decide whether to shift or reduce
			fmt.Fprintf(&b, "  %s\n", s)
			fmt.Fprintf(&b, "  %s\n", r)
		} else if r1, r2, ok := e.isReduceReduceConflict(); ok {
			fmt.Fprintf(&b, "AMBIGUOUS Grammar: reduce/reduce conflict in ACTION[%d, %s]\n", e.State, e.Symbol)
			// cannot decide between multiple reductions
			fmt.Fprintf(&b, "  %s\n", r1)
			fmt.Fprintf(&b, "  %s\n", r2)
		}

	default:
		fmt.Fprintf(&b, "invalid error: %d", e.Type)
	}

	return b.String()
}

func (e *ParsingTableError) isShiftReduceConflict() (*Action, *Action, bool) {
	if e.Type != CONFLICT {
		return nil, nil, false
	}

	shift, ok := e.Actions.FirstMatch(func(action *Action) bool {
		return action.Type == SHIFT
	})

	if !ok {
		return nil, nil, false
	}

	reduce, ok := e.Actions.FirstMatch(func(action *Action) bool {
		return action.Type == REDUCE
	})

	if !ok {
		return nil, nil, false
	}

	return shift, reduce, true
}

func (e *ParsingTableError) isReduceReduceConflict() (*Action, *Action, bool) {
	if e.Type != CONFLICT {
		return nil, nil, false
	}

	actions := generic.Collect1(e.Actions.SelectMatch(func(action *Action) bool {
		return action.Type == REDUCE
	}).All())

	if len(actions) < 2 {
		return nil, nil, false
	}

	sort.Quick(actions, cmpAction)

	return actions[0], actions[1], true
}

// tableStringer builds a string representation of a parsing table used during LR parsing.
// It provides a human-readable visualization of the table's content.
type tableStringer[K1, K2, K3 any] struct {
	K1Title  string
	K1Values []K1
	K2Title  string
	K2Values []K2
	K3Title  string
	K3Values []K3
	GetK1K2  func(K1, K2) string
	GetK1K3  func(K1, K3) string

	b     bytes.Buffer
	cLens []int
	tLen  int
}

func (t *tableStringer[K1, K2, K3]) String() string {
	t.b.Reset()

	// Calculate the maximum length of each column and the total length of the table.
	t.calculateColumnLengths()

	t.printTopLine()         // ┌─────────┬───────────────┬───────────┐
	t.printK2Titles()        // │         │     ACTION    │   GOTO    │
	t.printSecondLine()      // │  STATE  ├───┬───┬───┬───┼───┬───┬───┤
	t.printCell(0, "", true) // │         │ ...

	//  ... a │ b │ c │ d │
	for i, k2 := range t.K2Values {
		t.printCell(i+1, k2, false)
	}

	//  ... A │ B │ C
	for i, k3 := range t.K3Values {
		t.printCell(i+1+len(t.K2Values), k3, false)
	}

	t.b.WriteRune('\n')

	for _, k1 := range t.K1Values {
		t.printMiddleLine()      // ├─────────┼───┼───┼───┼───┼───┼───┼───┤
		t.printCell(0, k1, true) // │    0    │ ...

		//  ... A1 │ A2 │ A3 │ A4 │
		for i, k2 := range t.K2Values {
			s := t.GetK1K2(k1, k2)
			t.printCell(i+1, s, false)
		}

		//  ... G1 │ G2 │ G3 │
		for i, k3 := range t.K3Values {
			s := t.GetK1K3(k1, k3)
			t.printCell(i+1+len(t.K2Values), s, false)
		}

		t.b.WriteRune('\n')
	}

	t.printBottomLine() // └─────────┴───────────────┴───────────┘

	return t.b.String()
}

func (t *tableStringer[K1, K2, K3]) calculateColumnLengths() {
	t.cLens = make([]int, 1+len(t.K2Values)+len(t.K3Values))

	// Find the maximum length of the first column, which belongs to K1.
	t.cLens[0] = utf8.RuneCountInString(t.K1Title)
	for _, k1 := range t.K1Values {
		s := fmt.Sprintf("%v", k1)

		if l := utf8.RuneCountInString(s); l > t.cLens[0] {
			t.cLens[0] = l
		}
	}

	// Find the maximum length of each K2 column.
	for i, k2 := range t.K2Values {
		c := i + 1
		s := fmt.Sprintf("%v", k2)
		t.cLens[c] = utf8.RuneCountInString(s)

		for _, k1 := range t.K1Values {
			s := t.GetK1K2(k1, k2)

			if l := utf8.RuneCountInString(s); l > t.cLens[c] {
				t.cLens[c] = l
			}
		}
	}

	// Find the maximum length of each K3 column.
	for i, k3 := range t.K3Values {
		c := i + 1 + len(t.K2Values)
		s := fmt.Sprintf("%v", k3)
		t.cLens[c] = utf8.RuneCountInString(s)

		for _, k1 := range t.K1Values {
			s := t.GetK1K3(k1, k3)

			if l := utf8.RuneCountInString(s); l > t.cLens[c] {
				t.cLens[c] = l
			}
		}
	}

	// Add padding to each column.
	for i := range t.cLens {
		t.cLens[i] += 2
	}

	// Calculate the total length of the table.
	t.tLen = len(t.cLens) + 1 // padding
	for _, l := range t.cLens {
		t.tLen += l
	}
}

func (t *tableStringer[K1, K2, K3]) getTitleIAndIILens() (int, int) {
	lenI := len(t.K2Values) - 1
	for i := 1; i < 1+len(t.K2Values); i++ {
		lenI += t.cLens[i]
	}

	lenII := len(t.K3Values) - 1
	for i := 1 + len(t.K2Values); i < 1+len(t.K2Values)+len(t.K3Values); i++ {
		lenII += t.cLens[i]
	}

	return lenI, lenII
}

func (t *tableStringer[K1, K2, K3]) printK2Titles() {
	t.b.WriteRune('│')
	t.b.WriteString(strings.Repeat(" ", t.cLens[0]))

	lenI, lenII := t.getTitleIAndIILens()

	padI := lenI - utf8.RuneCountInString(t.K2Title)
	lpadI := padI / 2
	rpadI := padI - lpadI

	t.b.WriteRune('│')
	t.b.WriteString(strings.Repeat(" ", lpadI))
	t.b.WriteString(t.K2Title)
	t.b.WriteString(strings.Repeat(" ", rpadI))

	padII := lenII - utf8.RuneCountInString(t.K3Title)
	lpadII := padII / 2
	rpadII := padII - lpadII

	t.b.WriteRune('│')
	t.b.WriteString(strings.Repeat(" ", lpadII))
	t.b.WriteString(t.K3Title)
	t.b.WriteString(strings.Repeat(" ", rpadII))
	t.b.WriteRune('│')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2, K3]) printCell(col int, v any, isFirstCell bool) {
	if isFirstCell {
		t.b.WriteRune('│')
	}

	s := fmt.Sprintf("%v", v)

	pad := t.cLens[col] - utf8.RuneCountInString(s)
	lpad := pad / 2
	rpad := pad - lpad

	t.b.WriteString(strings.Repeat(" ", lpad))
	t.b.WriteString(s)
	t.b.WriteString(strings.Repeat(" ", rpad))
	t.b.WriteRune('│')
}

func (t *tableStringer[K1, K2, K3]) printTopLine() {
	lenI, lenII := t.getTitleIAndIILens()

	t.b.WriteRune('┌')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))
	t.b.WriteRune('┬')
	t.b.WriteString(strings.Repeat("─", lenI))
	t.b.WriteRune('┬')
	t.b.WriteString(strings.Repeat("─", lenII))
	t.b.WriteRune('┐')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2, K3]) printSecondLine() {
	t.b.WriteRune('│')

	pad := t.cLens[0] - utf8.RuneCountInString(t.K1Title)
	lpad := pad / 2
	rpad := pad - lpad

	t.b.WriteString(strings.Repeat(" ", lpad))
	t.b.WriteString(t.K1Title)
	t.b.WriteString(strings.Repeat(" ", rpad))

	i := 1
	t.b.WriteRune('├')
	t.b.WriteString(strings.Repeat("─", t.cLens[i]))

	for i++; i < 1+len(t.K2Values); i++ {
		t.b.WriteRune('┬')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┼')
	t.b.WriteString(strings.Repeat("─", t.cLens[i]))

	for i++; i < 1+len(t.K2Values)+len(t.K3Values); i++ {
		t.b.WriteRune('┬')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┤')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2, K3]) printMiddleLine() {
	t.b.WriteRune('├')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))

	for i := 1; i < len(t.cLens); i++ {
		t.b.WriteRune('┼')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┤')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2, K3]) printBottomLine() {
	t.b.WriteRune('└')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))

	for i := 1; i < len(t.cLens); i++ {
		t.b.WriteRune('┴')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┘')
	t.b.WriteRune('\n')
}
