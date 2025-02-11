package lr

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/moorara/algo/errors"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
)

// Associativity represents the associativity property of a terminal or a production rule.
type Associativity int

const (
	NONE  Associativity = iota // Not associative
	LEFT                       // Left-associative
	RIGHT                      // Right-associative
)

// String returns a string representation of the associativity.
func (a Associativity) String() string {
	switch a {
	case NONE:
		return "NONE"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return fmt.Sprintf("Invalid Associativity(%d)", a)
	}
}

// Precedence represents the precedence of a terminal or a production rule.
type Precedence struct {
	Order int
	Associativity
}

// String returns a string representation of the predencces.
func (p *Precedence) String() string {
	return fmt.Sprintf("%d:%s", p.Order, p.Associativity)
}

// Equal determines whether or not two predencces are the same.
func (p *Precedence) Equal(rhs *Precedence) bool {
	return p.Associativity == rhs.Associativity && p.Order == rhs.Order
}

// ActionHandlePair associates an action with a precedence handle.
// It is used to resolve precedence conflicts when multiple handles have the same precedence order.
// The action is selected based on the associativity of the handles and the types of actions involved.
type ActionHandlePair struct {
	Action *Action
	Handle *PrecedenceHandle
}

// String returns a string representation of the pair of action-handle.
func (p *ActionHandlePair) String() string {
	return fmt.Sprintf("<%s, %s>", p.Action, p.Handle)
}

// Equal determines whether or not two pairs of action-handle are the same.
func (p *ActionHandlePair) Equal(rhs *ActionHandlePair) bool {
	return p.Action.Equal(rhs.Action) && p.Handle.Equal(rhs.Handle)
}

// PrecedenceLevels represents an ordered list of precedence levels defined
// for specific terminals or production rules of a context-free grammar.
// The order of these levels is crucial for resolving conflicts.
type PrecedenceLevels []*PrecedenceLevel

// String returns a string representation of the list of precedence levels.
func (p PrecedenceLevels) String() string {
	levels := make([]string, len(p))
	for i, l := range p {
		levels[i] = l.String()
	}

	return strings.Join(levels, "\n")
}

// Equal determines whether or not two ordered list of precedence levels are the same.
func (p PrecedenceLevels) Equal(rhs PrecedenceLevels) bool {
	if len(p) != len(rhs) {
		return false
	}

	for i := range len(p) {
		if !p[i].Equal(rhs[i]) {
			return false
		}
	}

	return true
}

// Validate checks whether a list of precedence levels is valid.
// A precedence handle must not appear in multiple levels.
func (p PrecedenceLevels) Validate() error {
	var err *errors.MultiError

	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			if h := p[i].Handles.Intersection(p[j].Handles); h.Size() > 0 {
				err = errors.Append(err, fmt.Errorf("%s appeared in more than one precedence level", h))
			}
		}
	}

	return err.ErrorOrNil()
}

// Precedence searches for a precedence handle in the list of precedence levels.
// If found, it returns the handle's associativity and its order in the list.
// If not found, it returns nil and false.
func (p PrecedenceLevels) Precedence(h *PrecedenceHandle) (*Precedence, bool) {
	for i, level := range p {
		if level.Handles.Contains(h) {
			return &Precedence{
				Order:         i,
				Associativity: level.Associativity,
			}, true
		}
	}

	return nil, false
}

// Compare compares two action-handle pairs to determine their precedence.
// A handle that appears earlier in the list has higher precedence and is considered larger.
// If both handles have the same precedence level, they are compared based on the associativity of their level.
// In this case, the associativity determines which actions types precede.
//
// If either handle is not found in the list, the function returns -1 and false.
func (p PrecedenceLevels) Compare(lhs, rhs *ActionHandlePair) (int, error) {
	const invalid = 6765

	if lhs.Equal(rhs) {
		return 0, nil
	}

	lp, lok := p.Precedence(lhs.Handle)
	rp, rok := p.Precedence(rhs.Handle)

	switch {
	case !lok && !rok:
		return invalid, fmt.Errorf("no associativity and precedence specified: %s, %s", lhs.Handle, rhs.Handle)
	case !lok:
		return invalid, fmt.Errorf("no associativity and precedence specified: %s", lhs.Handle)
	case !rok:
		return invalid, fmt.Errorf("no associativity and precedence specified: %s", rhs.Handle)
	}

	// A lower index in the list indicates higher precedence.
	if lp.Order < rp.Order {
		return 1, nil
	} else if lp.Order > rp.Order {
		return -1, nil
	}

	// If the two handles have the same order, they also have the same associativity,
	// and the caller should choose between competing actions based on their action types.

	switch lp.Associativity {
	case NONE:
		if lhs.Handle.Equal(rhs.Handle) {
			return invalid, fmt.Errorf("not associative: %s", lhs.Handle)
		} else {
			return invalid, fmt.Errorf("not associative: %s and %s", lhs.Handle, rhs.Handle)
		}

	case LEFT:
		// For left-associative handles, reduces precede shifts.
		if lhs.Action.Type == REDUCE && rhs.Action.Type == SHIFT {
			return 1, nil
		} else if lhs.Action.Type == SHIFT && rhs.Action.Type == REDUCE {
			return -1, nil
		}

	case RIGHT:
		// For right-associative handles, shifts precede reduces.
		if lhs.Action.Type == SHIFT && rhs.Action.Type == REDUCE {
			return 1, nil
		} else if lhs.Action.Type == REDUCE && rhs.Action.Type == SHIFT {
			return -1, nil
		}
	}

	// At this point, both handles have the same precedence order and associativity,
	// and both actions are of the REDUCE type.
	// In this case, no precedence can be determined and the
	// conflicting handles should be assigned to separate precedence levels.

	if lhs.Handle.Equal(rhs.Handle) {
		return invalid, fmt.Errorf("assign separate precedences: %s and %s", lhs.Action.Production, rhs.Action.Production)
	}

	return invalid, fmt.Errorf("assign separate precedences: %s and %s", lhs.Handle, rhs.Handle)
}

// PrecedenceLevel defines a set of terminals and/or production rules (referred to as handles),
// each of which shares the same precedence level and associativity.
type PrecedenceLevel struct {
	Associativity Associativity
	Handles       PrecedenceHandles
}

// String returns a string representation of the precedence level.
func (p *PrecedenceLevel) String() string {
	return fmt.Sprintf("%s %s", p.Associativity, p.Handles)
}

// Equal determines whether or not two precedence levels are the same.
func (p *PrecedenceLevel) Equal(rhs *PrecedenceLevel) bool {
	return p.Associativity == rhs.Associativity &&
		p.Handles.Equal(rhs.Handles)
}

// PrecedenceHandles represents a set of terminals and/or production rules (referred to as handles).
type PrecedenceHandles set.Set[*PrecedenceHandle]

// NewPrecedenceHandles creates a new set of terminals and/or production rules (referred to as handles).
func NewPrecedenceHandles(handles ...*PrecedenceHandle) PrecedenceHandles {
	return set.NewWithFormat(
		eqPrecedenceHandle,
		func(h []*PrecedenceHandle) string {
			if len(h) == 0 {
				return ""
			}

			sort.Insertion(h, cmpPrecedenceHandle)

			var b bytes.Buffer
			for i := range len(h) {
				fmt.Fprintf(&b, "%s, ", h[i])
			}
			b.Truncate(b.Len() - 2)

			return b.String()
		},
		handles...,
	)
}

// cmpPrecedenceHandles compares two sets of handles and establishes an order between them.
func cmpPrecedenceHandles(lhs, rhs PrecedenceHandles) int {
	if lhs.Size() < rhs.Size() {
		return -1
	} else if lhs.Size() > rhs.Size() {
		return 1
	}

	ls := generic.Collect1(lhs.All())
	sort.Quick(ls, cmpPrecedenceHandle)

	rs := generic.Collect1(rhs.All())
	sort.Quick(rs, cmpPrecedenceHandle)

	for i := range len(ls) {
		if cmp := cmpPrecedenceHandle(ls[i], rs[i]); cmp != 0 {
			return cmp
		}
	}

	return 0
}

// PrecedenceHandle represents either a terminal symbol or a production rule
// in the context of determining precedence for conflict resolution.
type PrecedenceHandle struct {
	*grammar.Terminal
	*grammar.Production
}

// PrecedenceHandleForTerminal creates a new precedence handle represented by a terminal symbol.
func PrecedenceHandleForTerminal(t grammar.Terminal) *PrecedenceHandle {
	return &PrecedenceHandle{
		Terminal: &t,
	}
}

// PrecedenceHandleForProduction creates a new precedence handle for a production rule.
// If the production contains terminal symbols, the handle is represented by
// the first terminal in the right-hand side (body) of the production.
// Otherwise, the handle is represented by the production itself.
func PrecedenceHandleForProduction(p *grammar.Production) *PrecedenceHandle {
	first, ok := generic.FirstMatch(p.Body, func(s grammar.Symbol) bool {
		return s.IsTerminal()
	})

	if ok {
		t := first.(grammar.Terminal)
		return &PrecedenceHandle{
			Terminal: &t,
		}
	}

	return &PrecedenceHandle{
		Production: p,
	}
}

// IsTerminal returns true if the handle represents a terminal symbol.
func (h *PrecedenceHandle) IsTerminal() bool {
	return h.Terminal != nil && h.Production == nil
}

// IsProduction returns true if the handle represents a production rule.
func (h *PrecedenceHandle) IsProduction() bool {
	return h.Terminal == nil && h.Production != nil
}

// String returns a string representation of the handle.
func (h *PrecedenceHandle) String() string {
	if h.IsTerminal() {
		return h.Terminal.String()
	} else if h.IsProduction() {
		return fmt.Sprintf("%s = %s", h.Production.Head, h.Production.Body)
	}

	panic("PrecedenceHandle.String: invalid configuration")
}

// Equal determines whether or not two handles are the same.
func (h *PrecedenceHandle) Equal(rhs *PrecedenceHandle) bool {
	switch {
	case h.IsTerminal() && rhs.IsTerminal():
		return grammar.EqTerminal(*h.Terminal, *rhs.Terminal)
	case h.IsProduction() && rhs.IsProduction():
		return grammar.EqProduction(h.Production, rhs.Production)

	}

	return false
}

// eqPrecedenceHandle determines whether or not two handles are the same.
func eqPrecedenceHandle(lhs, rhs *PrecedenceHandle) bool {
	return lhs.Equal(rhs)
}

// cmpPrecedenceHandle compares two handles and establishes an order between them.
// Terminal handles come before Production handles.
func cmpPrecedenceHandle(lhs, rhs *PrecedenceHandle) int {
	switch {
	case lhs.IsTerminal() && rhs.IsProduction():
		return -1
	case lhs.IsProduction() && rhs.IsTerminal():
		return 1
	case lhs.IsTerminal() && rhs.IsTerminal():
		return grammar.CmpTerminal(*lhs.Terminal, *rhs.Terminal)
	case lhs.IsProduction() && rhs.IsProduction():
		return grammar.CmpProduction(lhs.Production, rhs.Production)
	}

	panic("cmpPrecedenceHandle: invalid configuration")
}
