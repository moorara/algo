package predictive

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/errors"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

var (
	eqParsingTableRow = func(lhs, rhs symboltable.SymbolTable[grammar.Terminal, *parsingTableEntry]) bool {
		return lhs.Equals(rhs)
	}

	eqParsingTableEntry = func(lhs, rhs *parsingTableEntry) bool {
		return lhs.Equals(rhs)
	}
)

// ParsingTable represents the interface for a parsing table of a predictive parser.
type ParsingTable interface {
	fmt.Stringer
	generic.Equaler[ParsingTable]

	// Error checks for conflicts in the parsing table.
	// If there are multiple productions for at least one combination of non-terminal A and terminal a,
	// the method returns an error containing details about the conflicting productions.
	// If no conflicts are found, it returns nil.
	Error() error

	// AddProduction adds a new production to the parsing table.
	// Multiple productions can be added for the same non-terminal A and terminal a.
	// It returns false if the M[A,a] entry is marked as a synchronization token.
	AddProduction(grammar.NonTerminal, grammar.Terminal, grammar.Production) bool

	// SetSync updates the parsing table to mark or unmark
	// terminal a as a synchronization symbol for non-terminal A.
	// It returns false if the M[A,a] entry contains any productions.
	SetSync(grammar.NonTerminal, grammar.Terminal, bool) bool

	// IsEmpty returns true if there are no productions in the M[A,a] entry.
	IsEmpty(grammar.NonTerminal, grammar.Terminal) bool

	// IsSync returns true if the M[A,a] entry is marked as a synchronization symbol for non-terminal A.
	IsSync(grammar.NonTerminal, grammar.Terminal) bool

	// GetProduction returns the single production from the M[A,a] entry if exactly one production exists.
	// It returns the production and true if successful, or a default value and false otherwise.
	GetProduction(grammar.NonTerminal, grammar.Terminal) (grammar.Production, bool)
}

// BuildParsingTable constructs a parsing table for a predictive parser.
//
// Predictive parsers are recursive-descent parsers that do not require backtracking.
// They can be constructed for a specific class of grammars called LL(1).
// An LL(1) grammar must not be left-recursive or ambiguous.
//
// If a grammar is left-recursive or ambiguous,
// the resulting parsing table will contain one or more multiply defined entries.
//
// This method constructs a parsing table for any context-free grammar.
// To identify errors in the table, use the CheckErrors method.
// Some errors may be resolved by eliminating left recursion and applying left factoring to the grammar.
// However, certain grammars cannot be transformed into LL(1) even after these transformations.
// Some languages may have no LL(1) grammar at all.
func BuildParsingTable(G grammar.CFG) ParsingTable {
	/*
	 * For each production A → α of the grammar:
	 *
	 *   1. For each terminal a in FIRST(α), add A → α to M[A,a].
	 *   2. If ε is in FIRST(α), then for each terminal b in FOLLOW(A), add A → α to M[A,b].
	 *      If ε is in FIRST(α) and $ is in FOLLOW(A), add A → α to M[A,$] as well.
	 *   3. If, after performing the above, there is no production at all in M[A,a],
	 *      then set M[A,a] to error (can be represented by an empty entry in the table).
	 */

	terminals := G.OrderTerminals()
	_, _, nonTerminals := G.OrderNonTerminals()
	table := newParsingTable(terminals, nonTerminals)

	first := G.ComputeFIRST()
	follow := G.ComputeFOLLOW(first)

	// For each production A → α
	for p := range G.Productions.All() {
		A := p.Head
		FIRSTα := first(p.Body)

		// For each terminal a ∈ FIRST(α), add A → α to M[A,a].
		for a := range FIRSTα.Terminals.All() {
			table.AddProduction(A, a, p)
		}

		// If ε ∈ FIRST(α)
		if FIRSTα.IncludesEmpty {
			FOLLOWA := follow(A)

			// For each terminal b ∈ FOLLOW(A), add A → α to M[A,b].
			for b := range FOLLOWA.Terminals.All() {
				table.AddProduction(A, b, p)
			}

			// If ε ∈ FIRST(α) and $ ∈ FOLLOW(A), add A → α to M[A,$] as well.
			if FOLLOWA.IncludesEndmarker {
				table.AddProduction(A, grammar.Endmarker, p)
			}
		}
	}

	/*
	 * Panic-mode error recovery is based on the idea of skipping over symbols on the input
	 * until a token in a selected set of synchronizing tokens appears.
	 * Its effectiveness depends on the choice of synchronizing set.
	 * The sets should be chosen so that the parser recovers quickly from errors that are likely to occur.
	 */

	// As a starting point, add all terminals in FOLLOW(A) to the synchronization set for non-terminal A.
	for _, A := range nonTerminals {
		FOLLOWA := follow(A)

		for a := range FOLLOWA.Terminals.All() {
			// If M[A,a] has any productions, the sync flag will not be set.
			table.SetSync(A, a, true)
		}

		if FOLLOWA.IncludesEndmarker {
			// If M[A,$] has any productions, the sync flag will not be set.
			table.SetSync(A, grammar.Endmarker, true)
		}
	}

	return table
}

// parsingTable represents a parsing table for a predictive parser.
type parsingTable struct {
	nonTerminals []grammar.NonTerminal
	terminals    []grammar.Terminal
	table        symboltable.SymbolTable[
		grammar.NonTerminal,
		symboltable.SymbolTable[
			grammar.Terminal,
			*parsingTableEntry,
		],
	]
}

func newParsingTable(terminals []grammar.Terminal, nonTerminals []grammar.NonTerminal) *parsingTable {
	return &parsingTable{
		nonTerminals: nonTerminals,
		terminals:    append(terminals, grammar.Endmarker),
		table: symboltable.NewQuadraticHashTable(
			grammar.HashNonTerminal,
			grammar.EqNonTerminal,
			eqParsingTableRow,
			symboltable.HashOpts{},
		),
	}
}

func (t *parsingTable) getEntry(A grammar.NonTerminal, a grammar.Terminal) (*parsingTableEntry, bool) {
	if row, ok := t.table.Get(A); ok {
		if entry, ok := row.Get(a); ok {
			return entry, true
		}
	}

	return nil, false
}

func (t *parsingTable) ensureEntry(A grammar.NonTerminal, a grammar.Terminal) *parsingTableEntry {
	if _, ok := t.table.Get(A); !ok {
		t.table.Put(A, symboltable.NewQuadraticHashTable(
			grammar.HashTerminal,
			grammar.EqTerminal,
			eqParsingTableEntry,
			symboltable.HashOpts{},
		))
	}

	row, _ := t.table.Get(A)

	if _, ok := row.Get(a); !ok {
		row.Put(a, &parsingTableEntry{
			Productions: set.New(grammar.EqProduction),
			Sync:        false,
		})
	}

	entry, _ := row.Get(a)

	return entry
}

func (t *parsingTable) String() string {
	ts := &parser.TableStringer[grammar.NonTerminal, grammar.Terminal]{
		K1Title:  "Non-Terminal",
		K1Values: t.nonTerminals,
		K2Title:  "Terminal",
		K2Values: t.terminals,
		GetEntry: func(A grammar.NonTerminal, a grammar.Terminal) string {
			if e, ok := t.getEntry(A, a); ok {
				return e.String()
			}

			return ""
		},
	}

	return ts.String()
}

func (t *parsingTable) Equals(rhs ParsingTable) bool {
	u, ok := rhs.(*parsingTable)
	return ok && t.table.Equals(u.table)
}

func (t *parsingTable) Error() error {
	var err = &errors.MultiError{
		Format: errors.BulletErrorFormat,
	}

	for _, A := range t.nonTerminals {
		for _, a := range t.terminals {
			if e, ok := t.getEntry(A, a); ok {
				if e.Productions.Size() > 1 {
					err = errors.Append(err, &parsingTableError{
						NonTerminal: A,
						Terminal:    a,
						Productions: e.Productions,
					})
				}
			}
		}
	}

	return err.ErrorOrNil()
}

func (t *parsingTable) AddProduction(A grammar.NonTerminal, a grammar.Terminal, prod grammar.Production) bool {
	e := t.ensureEntry(A, a)
	if !e.Sync {
		e.Productions.Add(prod)
		return true
	}

	return false
}

func (t *parsingTable) SetSync(A grammar.NonTerminal, a grammar.Terminal, sync bool) bool {
	e := t.ensureEntry(A, a)
	if e.Productions.Size() == 0 {
		e.Sync = sync
		return true
	}

	return false
}

func (t *parsingTable) IsEmpty(A grammar.NonTerminal, a grammar.Terminal) bool {
	if e, ok := t.getEntry(A, a); ok {
		return e.Productions.Size() == 0
	}

	return true
}

func (t *parsingTable) IsSync(A grammar.NonTerminal, a grammar.Terminal) bool {
	if e, ok := t.getEntry(A, a); ok {
		return e.Productions.Size() == 0 && e.Sync
	}

	return false
}

func (t *parsingTable) GetProduction(A grammar.NonTerminal, a grammar.Terminal) (grammar.Production, bool) {
	if e, ok := t.getEntry(A, a); ok {
		if prods := e.Productions; prods.Size() == 1 {
			for p := range prods.All() {
				return p, true
			}
		}
	}

	return grammar.Production{}, false
}

// parsingTableEntry defines the type of each entry in the parsing table.
type parsingTableEntry struct {
	Productions set.Set[grammar.Production]
	Sync        bool
}

func (e *parsingTableEntry) String() string {
	if e.Sync {
		return "sync"
	}

	if e.Productions.Size() == 0 {
		return ""
	}

	prods := grammar.OrderProductionSet(e.Productions)
	ss := make([]string, len(prods))
	for i, p := range prods {
		ss[i] = p.String()
	}

	return strings.Join(ss, " ┆ ")
}

func (e *parsingTableEntry) Equals(rhs *parsingTableEntry) bool {
	return e.Productions.Equals(rhs.Productions) && e.Sync == rhs.Sync
}

// parsingTableError represents an error for a parsing table of a predictive parser.
// This error occurs due to the presence of left recursion or ambiguity in the grammar.
type parsingTableError struct {
	NonTerminal grammar.NonTerminal
	Terminal    grammar.Terminal
	Productions set.Set[grammar.Production]
}

func (e *parsingTableError) Error() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "multiple productions in parsing table at M[%s, %s]:\n", e.NonTerminal, e.Terminal)
	for _, p := range grammar.OrderProductionSet(e.Productions) {
		fmt.Fprintf(b, "  %s\n", p)
	}

	return b.String()
}
