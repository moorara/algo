package lr

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

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
