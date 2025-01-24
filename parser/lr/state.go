package lr

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/sort"
)

const ErrState = State(-1)

var (
	EqState   = generic.NewEqualFunc[State]()
	HashState = hash.HashFuncForInt[State](nil)

	CmpState = func(lhs, rhs State) int {
		return int(lhs) - int(rhs)
	}
)

// State represents a state in the LR parsing table or automaton.
type State int

// StateMap is a generic map that associates an state (index) with an item set.
type StateMap []ItemSet

// BuildStateMap constructs a deterministic mapping of states to item sets.
// It creates a StateMap that associates a state (index) with each item set in the collection.
func BuildStateMap(C ItemSetCollection) StateMap {
	states := generic.Collect1(C.All())
	sort.Quick(states, cmpItemSet)

	return states
}

// Find finds the state corresponding to a given item set.
// Returns the state if found, or ErrState if no match exists.
func (m StateMap) Find(I ItemSet) State {
	for s := range m {
		if m[s].Equal(I) {
			return State(s)
		}
	}

	return ErrState
}

// All returns all states in the map.
func (m StateMap) All() []State {
	states := make([]State, len(m))
	for i := range m {
		states[i] = State(i)
	}

	return states
}

// String returns a string representation of all states in the map.
func (m StateMap) String() string {
	var b bytes.Buffer

	// Calculate the maximum width
	width := 10 // minimum width
	for _, I := range m {
		for i := range I.All() {
			if l := utf8.RuneCountInString(i.String()); l > width {
				width = l
			}
		}
	}
	width += 2 // padding

	for s, I := range m {
		items := generic.Collect1(I.All())
		sort.Quick(items, cmpItem)

		printTopLine(&b, width, s) // ┌────────[i]─────────┐

		sepPrinted := false
		for _, item := range items {
			if !item.IsKernel() && !sepPrinted {
				printMiddleLine(&b, width)
				sepPrinted = true
			}
			printItem(&b, width, item)
		}

		printBottomLine(&b, width) // └────────────────────┘
	}

	return b.String()
}

func printTopLine(b *bytes.Buffer, width int, state int) {
	s := strconv.Itoa(state)
	pad := width - utf8.RuneCountInString(s) - 2
	lpad := pad / 2
	rpad := pad - lpad
	b.WriteRune('┌')
	b.WriteString(strings.Repeat("─", lpad))
	b.WriteRune('[')
	b.WriteString(s)
	b.WriteRune(']')
	b.WriteString(strings.Repeat("─", rpad))
	b.WriteRune('┐')
	b.WriteRune('\n')
}

func printItem(b *bytes.Buffer, width int, item Item) {
	s := item.String()
	rpad := width - utf8.RuneCountInString(s) - 2
	b.WriteRune('│')
	b.WriteRune(' ')
	b.WriteString(s)
	b.WriteString(strings.Repeat(" ", rpad))
	b.WriteRune(' ')
	b.WriteRune('│')
	b.WriteRune('\n')
}

func printMiddleLine(b *bytes.Buffer, width int) {
	b.WriteRune('├')
	b.WriteString(strings.Repeat("╌", width))
	b.WriteRune('┤')
	b.WriteRune('\n')
}

func printBottomLine(b *bytes.Buffer, width int) {
	b.WriteRune('└')
	b.WriteString(strings.Repeat("─", width))
	b.WriteRune('┘')
	b.WriteRune('\n')
}
