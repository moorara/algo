package grammar

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/moorara/algo/errors"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

// ParsingTable represents the interface for a predictive parsing table used for LL(1) context-free grammars.
type ParsingTable interface {
	fmt.Stringer
	generic.Equaler[ParsingTable]

	Add(NonTerminal, Terminal, Production)
	Get(NonTerminal, Terminal) set.Set[Production]
	CheckErrors() error
}

// parsingTable implements the ParsingTable interface.
type parsingTable struct {
	nonTerminals []NonTerminal
	terminals    []Terminal
	table        symboltable.SymbolTable[NonTerminal, symboltable.SymbolTable[Terminal, set.Set[Production]]]
}

// NewParsingTable creates a new instance of the ParsingTable.
func NewParsingTable(terminals []Terminal, nonTerminals []NonTerminal) ParsingTable {
	t := &parsingTable{
		nonTerminals: nonTerminals,
		terminals:    append(terminals, endmarker),
		table: symboltable.NewQuadraticHashTable(
			hashNonTerminal,
			eqNonTerminal,
			func(lhs, rhs symboltable.SymbolTable[Terminal, set.Set[Production]]) bool {
				return lhs.Equals(rhs)
			},
			symboltable.HashOpts{},
		),
	}

	// Populate the parsing table with empty sets.
	for _, A := range t.nonTerminals {
		ARow := symboltable.NewQuadraticHashTable(
			hashTerminal,
			eqTerminal,
			eqProductionSet,
			symboltable.HashOpts{},
		)

		for _, a := range t.terminals {
			ARow.Put(a, set.New(eqProduction))
		}

		t.table.Put(A, ARow)
	}

	return t
}

// String returns a string representation of the parsing table.
func (t *parsingTable) String() string {
	return newParsingTablePrinter(t).String()
}

// Equals determines whether or not two parsing tables are the same.
func (t *parsingTable) Equals(rhs ParsingTable) bool {
	u, ok := rhs.(*parsingTable)
	return ok && t.table.Equals(u.table)
}

// Add adds a new entry to the parsing table.
// Multiple productions can be added for the same non-terminal A and terminal a.
func (t *parsingTable) Add(A NonTerminal, a Terminal, prod Production) {
	// The parsing table is pre-populated with empty sets.
	ARow, _ := t.table.Get(A)
	prods, _ := ARow.Get(a)
	prods.Add(prod)
}

// Get looks up an entry in the parsing table.
func (t *parsingTable) Get(A NonTerminal, a Terminal) set.Set[Production] {
	// The parsing table is pre-populated with empty sets.
	ARow, _ := t.table.Get(A)
	prods, _ := ARow.Get(a)
	return prods
}

// CheckErrors checks for conflicts in the parsing table.
// If there are multiple productions for at least one combination of non-terminal A and terminal a,
// the method returns an error containing details about the conflicting productions.
// If no conflicts are found, it returns nil.
func (t *parsingTable) CheckErrors() error {
	var err = &errors.MultiError{
		Format: errors.BulletErrorFormat,
	}

	for _, A := range t.nonTerminals {
		for _, a := range t.terminals {
			if prods := t.Get(A, a); prods.Size() > 1 {
				err = errors.Append(err, &ParsingTableError{
					NonTerminal: A,
					Terminal:    a,
					Productions: prods,
				})
			}
		}
	}

	return err.ErrorOrNil()
}

// parsingTablePrinter is used for building a string representation of a parsing table.
type parsingTablePrinter struct {
	t     *parsingTable
	b     *strings.Builder
	cLens []int
	tLen  int
}

func newParsingTablePrinter(t *parsingTable) *parsingTablePrinter {
	return &parsingTablePrinter{
		t: t,
		b: new(strings.Builder),
	}
}

func (p *parsingTablePrinter) String() string {
	p.b = new(strings.Builder)
	rowTitle, colTitle := "Non-Terminal", "Terminal"

	// Calculate the maximum length of each column and the total length of the table.
	p.calculateColumnLengths(rowTitle)

	p.printTopLine()               // ┌──────────────┬───────────────┐
	p.printCell01(colTitle)        // │              │   Terminal    │
	p.printSecondLine()            // ├──────────────┼───┬────┬───┬──┤
	p.printCell(0, rowTitle, true) // │ Non-Terminal │ ...

	//  ... a │ b │ c │ d │
	for i := 1; i < len(p.cLens); i++ {
		a := p.t.terminals[i-1]
		p.printCell(i, a.String(), false)
	}

	p.b.WriteRune('\n')

	for _, A := range p.t.nonTerminals {
		p.printMiddleLine()              // ├──────────────┼───────────────┤
		p.printCell(0, A.String(), true) // │      A       │ ...

		//  ... P1 │ P2 │ P3 │ P4 │
		for i := 1; i < len(p.cLens); i++ {
			a := p.t.terminals[i-1]
			s := p.flattenProductions(A, a)
			p.printCell(i, s, false)
		}

		p.b.WriteRune('\n')
	}

	p.printBottomLine() // └──────────────┴───────────────┘

	return p.b.String()
}

func (p *parsingTablePrinter) flattenProductions(A NonTerminal, a Terminal) string {
	prods := orderProductionSet(p.t.Get(A, a))
	if len(prods) == 0 {
		return ""
	}

	ss := make([]string, len(prods))
	for i, p := range prods {
		ss[i] = p.String()
	}

	return strings.Join(ss, " ┆ ")
}

func (p *parsingTablePrinter) calculateColumnLengths(header0 string) {
	p.cLens = make([]int, len(p.t.terminals)+1)

	// Find the maximum length of the first column, which belongs to non-terminals.
	p.cLens[0] = utf8.RuneCountInString(header0)
	for _, A := range p.t.nonTerminals {
		if l := utf8.RuneCountInString(A.String()); l > p.cLens[0] {
			p.cLens[0] = l
		}
	}

	// Find the maximum length of each terminal column.
	for i, a := range p.t.terminals {
		p.cLens[i+1] = utf8.RuneCountInString(a.String())
		for _, A := range p.t.nonTerminals {
			s := p.flattenProductions(A, a)
			if l := utf8.RuneCountInString(s); l > p.cLens[i+1] {
				p.cLens[i+1] = l
			}
		}
	}

	// Add padding to each column.
	for i := range p.cLens {
		p.cLens[i] += 2
	}

	// Calculate the total length of the table.
	p.tLen = len(p.cLens) + 1
	for _, l := range p.cLens {
		p.tLen += l
	}
}

func (p *parsingTablePrinter) printCell01(header1 string) {
	len1 := p.tLen - p.cLens[0] - 3
	pad1 := len1 - len(header1)
	lpad1 := pad1 / 2
	rpad1 := pad1 - lpad1
	p.b.WriteRune('│')
	p.b.WriteString(strings.Repeat(" ", p.cLens[0]))
	p.b.WriteRune('│')
	p.b.WriteString(strings.Repeat(" ", lpad1))
	p.b.WriteString(header1)
	p.b.WriteString(strings.Repeat(" ", rpad1))
	p.b.WriteRune('│')
	p.b.WriteRune('\n')
}

func (p *parsingTablePrinter) printCell(col int, s string, isFirstCell bool) {
	if isFirstCell {
		p.b.WriteRune('│')
	}

	pad0 := p.cLens[col] - utf8.RuneCountInString(s)
	lpad0 := pad0 / 2
	rpad0 := pad0 - lpad0
	p.b.WriteString(strings.Repeat(" ", lpad0))
	p.b.WriteString(s)
	p.b.WriteString(strings.Repeat(" ", rpad0))
	p.b.WriteRune('│')
}

func (p *parsingTablePrinter) printTopLine() {
	p.b.WriteRune('┌')
	p.b.WriteString(strings.Repeat("─", p.cLens[0]))
	p.b.WriteRune('┬')
	p.b.WriteString(strings.Repeat("─", p.tLen-p.cLens[0]-3))
	p.b.WriteRune('┐')
	p.b.WriteRune('\n')
}

func (p *parsingTablePrinter) printSecondLine() {
	p.b.WriteRune('├')
	p.b.WriteString(strings.Repeat("─", p.cLens[0]))
	p.b.WriteRune('┼')
	p.b.WriteString(strings.Repeat("─", p.cLens[1]))
	for i := 2; i < len(p.cLens); i++ {
		p.b.WriteRune('┬')
		p.b.WriteString(strings.Repeat("─", p.cLens[i]))
	}
	p.b.WriteRune('┤')
	p.b.WriteRune('\n')
}

func (p *parsingTablePrinter) printMiddleLine() {
	p.b.WriteRune('├')
	p.b.WriteString(strings.Repeat("─", p.cLens[0]))
	for i := 1; i < len(p.cLens); i++ {
		p.b.WriteRune('┼')
		p.b.WriteString(strings.Repeat("─", p.cLens[i]))
	}
	p.b.WriteRune('┤')
	p.b.WriteRune('\n')
}

func (p *parsingTablePrinter) printBottomLine() {
	p.b.WriteRune('└')
	p.b.WriteString(strings.Repeat("─", p.cLens[0]))
	for i := 1; i < len(p.cLens); i++ {
		p.b.WriteRune('┴')
		p.b.WriteString(strings.Repeat("─", p.cLens[i]))
	}
	p.b.WriteRune('┘')
	p.b.WriteRune('\n')
}
