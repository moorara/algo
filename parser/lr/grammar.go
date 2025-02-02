package lr

import "github.com/moorara/algo/grammar"

var primeSuffixes = []string{
	"′", // Prime (U+2032)
	"″", // Double Prime (U+2033)
	"‴", // Triple Prime (U+2034)
	"⁗", // Quadruple Prime (U+2057)
}

// augment augments a context-free grammar G by creating a new start symbol S′
// and adding a production "S′ → S", where S is the original start symbol of G.
// This transformation is used to prepare grammars for LR parsing.
// The function clones the input grammar G, ensuring the original grammar remains unmodified.
func augment(G *grammar.CFG) *grammar.CFG {
	augG := G.Clone()

	// A special symbol used to indicate the end of a string.
	augG.Terminals.Add(grammar.Endmarker)

	newS := augG.AddNewNonTerminal(G.Start, primeSuffixes...)
	augG.Start = newS
	augG.Productions.Add(&grammar.Production{
		Head: newS,
		Body: grammar.String[grammar.Symbol]{G.Start},
	})

	return augG
}

// Grammar represents a context-free grammar with additional features
// tailored for LR parsing and constructing an LR parser.
// The grammar is augmented with a new start symbol S′ and a new production "S′ → S"
// to facilitate the construction of the LR parser.
// It extends a regular context-free grammar by incorporating precedence and associativity information
// to handle ambiguities in the grammar and resolve conflicts.
// It also provides essential functionalities for building an LR automaton and the corresponding LR parsing table.
//
// Always use one of the provided functions to create a new instance of this type.
// Direct instantiation of the struct is discouraged.
type Grammar struct {
	*grammar.CFG
	Automaton
}

// NewGrammarWithLR0 creates a new augmented grammar based on the complete sets of LR(0) items.
func NewGrammarWithLR0(G *grammar.CFG) *Grammar {
	G = augment(G)

	return &Grammar{
		CFG: G,
		Automaton: &automaton{
			G: G,
			calculator: &calculator0{
				G: G,
			},
		},
	}
}

// NewGrammarWithLR1 creates a new augmented grammar based on the complete sets of LR(1) items.
func NewGrammarWithLR1(G *grammar.CFG) *Grammar {
	G = augment(G)
	FIRST := G.ComputeFIRST()

	return &Grammar{
		CFG: G,
		Automaton: &automaton{
			G: G,
			calculator: &calculator1{
				G:     G,
				FIRST: FIRST,
			},
		},
	}
}

// NewGrammarWithLR0Kernel creates a new augmented grammar based on the kernel sets of LR(0) items.
func NewGrammarWithLR0Kernel(G *grammar.CFG) *Grammar {
	G = augment(G)

	return &Grammar{
		CFG: G,
		Automaton: &kernelAutomaton{
			G: G,
			calculator: &calculator0{
				G: G,
			},
		},
	}
}

// NewGrammarWithLR1Kernel creates a new augmented grammar based on the kernel sets of LR(1) items.
func NewGrammarWithLR1Kernel(G *grammar.CFG) *Grammar {
	G = augment(G)
	FIRST := G.ComputeFIRST()

	return &Grammar{
		CFG: G,
		Automaton: &kernelAutomaton{
			G: G,
			calculator: &calculator1{
				G:     G,
				FIRST: FIRST,
			},
		},
	}
}
