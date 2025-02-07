package parsertest

import "github.com/moorara/algo/grammar"

var Prods = [][]*grammar.Production{
	{ // G0
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},                         // E → T E′
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}}, // E′ → + T E′
		{Head: "E′", Body: grammar.E}, // E′ → ε
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}},                         // T → F T′
		{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}}, // T′ → * F T′
		{Head: "T′", Body: grammar.E}, // T′ → ε
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}, // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                 // F → id
	},
	{ // G1
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                          // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")}}, // S → CC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")}},    // C → cC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("d")}},                              // C → d
	},
	{ // G2
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                                                 // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L"), grammar.Terminal("="), grammar.NonTerminal("R")}}, // S → L = R
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("R")}},                                                  // S → R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("R")}},                           // L → *R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // L → id
		{Head: "R", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L")}},                                                  // R → L
	},
	{ // G3
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
	{ // G4
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
	},
	{ // G5
		{Head: "grammar′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("grammar")}},                                                      // grammar′ → grammar
		{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}},                            // grammar → name decls
		{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT"), grammar.NonTerminal("semi_opt")}}, // name → "grammar" IDENT semi_opt
		{Head: "semi_opt", Body: grammar.String[grammar.Symbol]{grammar.Terminal(";")}},                                                               // semi_opt → ";"
		{Head: "semi_opt", Body: grammar.E}, // semi_opt → ε
		{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}}, // decls → decls decl
		{Head: "decls", Body: grammar.E}, // decls → ε
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                                                   // decl → token
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("directive")}},                                                                               // decl → directive
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                                                    // decl → rule
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING"), grammar.NonTerminal("semi_opt")}}, // token → TOKEN "=" STRING semi_opt
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX"), grammar.NonTerminal("semi_opt")}},  // token → TOKEN "=" REGEX semi_opt
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("PREDEF"), grammar.NonTerminal("semi_opt")}}, // token → TOKEN "=" PREDEF semi_opt
		{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@left"), grammar.NonTerminal("handles"), grammar.NonTerminal("semi_opt")}},                // directive → "@left" handles semi_opt
		{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@right"), grammar.NonTerminal("handles"), grammar.NonTerminal("semi_opt")}},               // directive → "@right" handles semi_opt
		{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@none"), grammar.NonTerminal("handles"), grammar.NonTerminal("semi_opt")}},                // directive → "@none" handles semi_opt
		{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("handles"), grammar.NonTerminal("term")}},                                                 // handles → handles term
		{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("handles"), grammar.NonTerminal("rule_handle")}},                                          // handles → handles rule_handle
		{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                                                 // handles → term
		{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule_handle")}},                                                                          // handles → rule_handle
		{Head: "rule_handle", Body: grammar.String[grammar.Symbol]{grammar.Terminal("<"), grammar.NonTerminal("rule"), grammar.Terminal(">")}},                               // rule_handle → "<" rule ">"
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs"), grammar.Terminal(";")}},           // rule → lhs "=" rhs ";"
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.Terminal(";")}},                                       // rule → lhs "=" ";"
		{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                                                  // lhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                                                          // rhs → rhs rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},                                        // rhs → "(" rhs ")"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},                                        // rhs → "[" rhs "]"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},                                        // rhs → "{" rhs "}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},                                      // rhs → "{{" rhs "}}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},                                   // rhs → rhs "|" rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|")}},                                                               // rhs → rhs "|"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                                                  // rhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                                                     // rhs → term
		{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                                                   // nonterm → IDENT
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                                                      // term → TOKEN
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                                                     // term → STRING
	},
}

var Grammars = []*grammar.CFG{
	// G0
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		Prods[0],
		"E",
	),
	// G1
	grammar.NewCFG(
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
		Prods[1][1:],
		"S",
	),
	// G2
	grammar.NewCFG(
		[]grammar.Terminal{"=", "*", "id"},
		[]grammar.NonTerminal{"S", "L", "R"},
		Prods[2][1:],
		"S",
	),
	// G3
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		Prods[3][1:],
		"E",
	),
	// G4
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		Prods[4][1:],
		"E",
	),
	// G5
	grammar.NewCFG(
		[]grammar.Terminal{
			"=", ";", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "<", ">",
			"grammar", "@left", "@right", "@none",
			"IDENT", "TOKEN", "STRING", "REGEX", "PREDEF",
		},
		[]grammar.NonTerminal{
			"grammar", "name", "decls", "decl", "semi_opt", "token",
			"directive", "handles", "rule_handle", "rule", "lhs", "rhs", "nonterm", "term",
		},
		Prods[5][1:],
		"grammar",
	),
}
