package predictive

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// tableStringer builds a string representation of a parsing table used during predictive parsing.
// It provides a human-readable visualization of the table's content.
type tableStringer[K1, K2 any] struct {
	K1Title  string
	K1Values []K1
	K2Title  string
	K2Values []K2
	GetK1K2  func(K1, K2) string

	b     bytes.Buffer
	cLens []int
	tLen  int
}

func (t *tableStringer[K1, K2]) String() string {
	t.b.Reset()

	// Calculate the maximum length of each column and the total length of the table.
	t.calculateColumnLengths()

	t.printTopLine()         // ┌──────────────┬───────────────┐
	t.printK2Title()         // │              │   Terminal    │
	t.printSecondLine()      // │ Non-Terminal ├───┬───┬───┬───┤
	t.printCell(0, "", true) // │              │ ...

	//  ... a │ b │ c │ d │
	for i, k2 := range t.K2Values {
		t.printCell(i+1, k2, false)
	}

	t.b.WriteRune('\n')

	for _, k1 := range t.K1Values {
		t.printMiddleLine()      // ├──────────────┼───┼───┼───┼───┤
		t.printCell(0, k1, true) // │      A       │ ...

		//  ... P1 │ P2 │ P3 │ P4 │
		for i, k2 := range t.K2Values {
			s := t.GetK1K2(k1, k2)
			t.printCell(i+1, s, false)
		}

		t.b.WriteRune('\n')
	}

	t.printBottomLine() // └──────────────┴───────────────┘

	return t.b.String()
}

func (t *tableStringer[K1, K2]) calculateColumnLengths() {
	t.cLens = make([]int, 1+len(t.K2Values))

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
		s := fmt.Sprintf("%v", k2)
		t.cLens[i+1] = utf8.RuneCountInString(s)

		for _, k1 := range t.K1Values {
			s := t.GetK1K2(k1, k2)

			if l := utf8.RuneCountInString(s); l > t.cLens[i+1] {
				t.cLens[i+1] = l
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

func (t *tableStringer[K1, K2]) printK2Title() {
	t.b.WriteRune('│')
	t.b.WriteString(strings.Repeat(" ", t.cLens[0]))

	lenI := t.tLen - t.cLens[0] - 3
	padI := lenI - utf8.RuneCountInString(t.K2Title)
	lpadI := padI / 2
	rpadI := padI - lpadI

	t.b.WriteRune('│')
	t.b.WriteString(strings.Repeat(" ", lpadI))
	t.b.WriteString(t.K2Title)
	t.b.WriteString(strings.Repeat(" ", rpadI))
	t.b.WriteRune('│')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2]) printCell(col int, v any, isFirstCell bool) {
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

func (t *tableStringer[K1, K2]) printTopLine() {
	t.b.WriteRune('┌')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))
	t.b.WriteRune('┬')
	t.b.WriteString(strings.Repeat("─", t.tLen-t.cLens[0]-3))
	t.b.WriteRune('┐')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2]) printSecondLine() {
	t.b.WriteRune('│')

	pad := t.cLens[0] - utf8.RuneCountInString(t.K1Title)
	lpad := pad / 2
	rpad := pad - lpad
	t.b.WriteString(strings.Repeat(" ", lpad))
	t.b.WriteString(t.K1Title)
	t.b.WriteString(strings.Repeat(" ", rpad))

	t.b.WriteRune('├')
	t.b.WriteString(strings.Repeat("─", t.cLens[1]))

	for i := 2; i < len(t.cLens); i++ {
		t.b.WriteRune('┬')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┤')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2]) printMiddleLine() {
	t.b.WriteRune('├')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))

	for i := 1; i < len(t.cLens); i++ {
		t.b.WriteRune('┼')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┤')
	t.b.WriteRune('\n')
}

func (t *tableStringer[K1, K2]) printBottomLine() {
	t.b.WriteRune('└')
	t.b.WriteString(strings.Repeat("─", t.cLens[0]))

	for i := 1; i < len(t.cLens); i++ {
		t.b.WriteRune('┴')
		t.b.WriteString(strings.Repeat("─", t.cLens[i]))
	}

	t.b.WriteRune('┘')
	t.b.WriteRune('\n')
}
