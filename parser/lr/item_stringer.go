package lr

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/sort"
)

// itemSetStringer builds a string representation of an item set.
// It provides a human-readable visualization of a set of items.
type itemSetStringer struct {
	state *State
	items []Item

	b     bytes.Buffer
	width int
}

func (s *itemSetStringer) String() string {
	s.b.Reset()

	// Calculate the width of item set.
	s.calculateWidth()

	sort.Quick(s.items, cmpItem)

	s.printTopLine() // ┌────────[s]─────────┐

	sepPrinted := false
	for _, item := range s.items {
		if !item.IsKernel() && !sepPrinted {
			s.printMiddleLine()
			sepPrinted = true
		}
		s.printItem(item)
	}

	s.printBottomLine() // └────────────────────┘

	return s.b.String()
}

func (s *itemSetStringer) calculateWidth() {
	// Already set or calculated.
	if s.width != 0 {
		return
	}

	// Minimum width
	s.width = 10

	for _, item := range s.items {
		if l := utf8.RuneCountInString(item.String()); l > s.width {
			s.width = l
		}
	}

	// Padding
	s.width += 2
}

func (s *itemSetStringer) printTopLine() {
	var state string
	if s.state != nil {
		state = fmt.Sprintf("[%d]", *s.state)
	}

	pad := s.width - utf8.RuneCountInString(state)
	lpad := pad / 2
	rpad := pad - lpad

	s.b.WriteRune('┌')
	s.b.WriteString(strings.Repeat("─", lpad))
	s.b.WriteString(state)
	s.b.WriteString(strings.Repeat("─", rpad))
	s.b.WriteRune('┐')
	s.b.WriteRune('\n')
}

func (s *itemSetStringer) printItem(i Item) {
	item := i.String()

	rpad := s.width - utf8.RuneCountInString(item) - 2

	s.b.WriteRune('│')
	s.b.WriteRune(' ')
	s.b.WriteString(item)
	s.b.WriteString(strings.Repeat(" ", rpad))
	s.b.WriteRune(' ')
	s.b.WriteRune('│')
	s.b.WriteRune('\n')
}

func (s *itemSetStringer) printMiddleLine() {
	s.b.WriteRune('├')
	s.b.WriteString(strings.Repeat("╌", s.width))
	s.b.WriteRune('┤')
	s.b.WriteRune('\n')
}

func (s *itemSetStringer) printBottomLine() {
	s.b.WriteRune('└')
	s.b.WriteString(strings.Repeat("─", s.width))
	s.b.WriteRune('┘')
	s.b.WriteRune('\n')
}

// itemSetCollectionStringer builds a string representation of a collection of item sets.
// It provides a human-readable visualization of a collection of item sets.
type itemSetCollectionStringer struct {
	sets []ItemSet
}

func (c *itemSetCollectionStringer) String() string {
	var b bytes.Buffer

	// Calculate the maximum length for all item sets.
	maxWidth := c.calculateMaxWidth()

	for i, set := range c.sets {
		is := &itemSetStringer{
			state: (*State)(&i),
			items: generic.Collect1(set.All()),
			width: maxWidth,
		}

		b.WriteString(is.String())
	}

	return b.String()
}

func (c *itemSetCollectionStringer) calculateMaxWidth() int {
	// Minimum width
	maxWidth := 10

	for _, set := range c.sets {
		for item := range set.All() {
			if l := utf8.RuneCountInString(item.String()); l > maxWidth {
				maxWidth = l
			}
		}
	}

	// Padding
	maxWidth += 2

	return maxWidth
}
