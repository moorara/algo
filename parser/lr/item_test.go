package lr

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/sort"
)

type mockItem string

func (i mockItem) String() string {
	return string(i)
}

func (i mockItem) Equals(rhs Item) bool {
	ii, ok := rhs.(mockItem)
	return ok && i == ii
}

func (i mockItem) Compare(rhs Item) int {
	ii, _ := rhs.(mockItem)

	if i.IsInitial() && !ii.IsInitial() {
		return -1
	} else if !i.IsInitial() && ii.IsInitial() {
		return 1
	}

	if i.IsKernel() && !ii.IsKernel() {
		return -1
	} else if !i.IsKernel() && ii.IsKernel() {
		return 1
	}

	s1, s2 := i.String(), ii.String()
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}

	return 0
}

func (i mockItem) IsInitial() bool {
	return strings.Contains(string(i), "′→•")
}

func (i mockItem) IsKernel() bool {
	return i.IsInitial() || !strings.Contains(string(i), "→•")
}

func (i mockItem) IsComplete() bool {
	r, _ := utf8.DecodeLastRuneInString(string(i))
	return r == '•'
}

func (i mockItem) DotSymbol() (grammar.Symbol, bool) {
	dot := strings.IndexRune(string(i), '•')
	_, s1 := utf8.DecodeRuneInString(string(i[dot:]))
	r, s2 := utf8.DecodeRuneInString(string(i[dot+s1:]))

	if s2 == 0 {
		return nil, false
	}

	if unicode.IsUpper(r) {
		return grammar.NonTerminal(r), true
	}
	return grammar.Terminal(r), true
}

func (i mockItem) Next() (Item, bool) {
	dot := strings.IndexRune(string(i), '•')
	_, s1 := utf8.DecodeRuneInString(string(i[dot:]))
	r, s2 := utf8.DecodeRuneInString(string(i[dot+s1:]))

	if s2 == 0 {
		return nil, false
	}

	return mockItem(string(i[:dot]) + string(r) + "•" + string(i[dot+s1+s2:])), true
}

func getTestItemSets() []ItemSet {
	I0 := NewItemSet(
		// Kernels
		mockItem("E′→•E"),
		// Non-Kernels
		mockItem("E→•E+T"),
		mockItem("E→•T"),
		mockItem("T→•T*F"),
		mockItem("T→•F"),
		mockItem("F→•(E)"),
		mockItem("F→•id"),
	)

	I1 := NewItemSet(
		// Kernels
		mockItem("E′→E•"),
		mockItem("E→E•+T"),
	)

	I2 := NewItemSet(
		// Kernels
		mockItem("E→T•"),
		mockItem("T→T•*F"),
	)

	I3 := NewItemSet(
		// Kernels
		mockItem("T→F•"),
	)

	I4 := NewItemSet(
		// Kernels
		mockItem("F→(•E)"),
		// Non-Kernels
		mockItem("E→•E+T"),
		mockItem("E→•T"),
		mockItem("T→•T*F"),
		mockItem("T→•F"),
		mockItem("F→•(E)"),
		mockItem("F→•id"),
	)

	I5 := NewItemSet(
		// Kernels
		mockItem("F→id•"),
	)

	I6 := NewItemSet(
		// Kernels
		mockItem("E→E+•T"),
		// Non-Kernels
		mockItem("T→•T*F"),
		mockItem("T→•F"),
		mockItem("F→•(E)"),
		mockItem("F→•id"),
	)

	I7 := NewItemSet(
		// Kernels
		mockItem("T→T*•F"),
		// Non-Kernels
		mockItem("F→•(E)"),
		mockItem("F→•id"),
	)

	I8 := NewItemSet(
		// Kernels
		mockItem("E→E•+T"),
		mockItem("F→(E•)"),
	)

	I9 := NewItemSet(
		// Kernels
		mockItem("E→E+T•"),
		mockItem("T→T•*F"),
	)

	I10 := NewItemSet(
		// Kernels
		mockItem("T→T*F•"),
	)

	I11 := NewItemSet(
		// Kernels
		mockItem("F→(E)•"),
	)

	return []ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func TestNewItemSetCollection(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name string
		sets []ItemSet
	}{
		{
			name: "OK",
			sets: []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSetCollection(tc.sets...)

			assert.NotNil(t, s)
			for _, expectedItemSet := range tc.sets {
				assert.True(t, s.Contains(expectedItemSet))
			}
		})
	}
}

func TestNewItemSet(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
	}{
		{
			name: "OK",
			items: []Item{
				mockItem("E′ → E•"),
				mockItem("E → E•+ T"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSet(tc.items...)

			assert.NotNil(t, s)
			for _, expectedItem := range tc.items {
				assert.True(t, s.Contains(expectedItem))
			}
		})
	}
}

func TestCmpItemSet(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name         string
		sets         []ItemSet
		expectedSets []ItemSet
	}{
		{
			name:         "OK",
			sets:         []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedSets: []ItemSet{s[0], s[1], s[9], s[6], s[8], s[2], s[11], s[4], s[5], s[3], s[10], s[7]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.sets, cmpItemSet)
			assert.Equal(t, tc.expectedSets, tc.sets)
		})
	}
}
