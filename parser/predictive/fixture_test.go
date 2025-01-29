package predictive

import "github.com/moorara/algo/grammar"

var grammars = []*grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]grammar.NonTerminal{"S", "E"},
		[]*grammar.Production{
			{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                  // S → E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("-"), grammar.NonTerminal("E")}}, // E → E - E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("/"), grammar.NonTerminal("E")}}, // E → E / E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("-"), grammar.NonTerminal("E")}},                           // E → - E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
		},
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]grammar.NonTerminal{"S", "E", "T", "F"},
		[]*grammar.Production{
			{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                  // S → E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("-"), grammar.NonTerminal("T")}}, // E → E - T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("/"), grammar.NonTerminal("F")}}, // T → T / F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
		},
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]*grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},                         // E → T E′
			{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}}, // E′ → + T E′
			{Head: "E′", Body: grammar.E}, // E′ → ε
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}},                         // T → F T′
			{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}}, // T′ → * F T′
			{Head: "T′", Body: grammar.E}, // T′ → ε
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}, // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                 // F → id
		},
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		[]*grammar.Production{
			{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
			{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT")}},       // name → "grammar" IDENT
			{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}},   // decls → decls decl
			{Head: "decls", Body: grammar.E}, // decls → ε
			{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                  // decl → token
			{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                   // decl → rule
			{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING")}}, // token → TOKEN "=" STRING
			{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX")}},  // token → TOKEN "=" REGEX
			{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs")}}, // rule → lhs "=" rhs
			{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("=")}},                             // rule → lhs "="
			{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // lhs → nonterm
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                         // rhs → rhs rhs
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},  // rhs → rhs "|" rhs
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},       // rhs → "(" rhs ")"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},       // rhs → "[" rhs "]"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},       // rhs → "{" rhs "}"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},     // rhs → "{{" rhs "}}"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // rhs → nonterm
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                    // rhs → term
			{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                  // nonterm → IDENT
			{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                     // term → TOKEN
			{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                    // term → STRING
		},
		"grammar",
	),
}

func getTestParsingTables() []*ParsingTable {
	pt0 := NewParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
	)

	pt0.addProduction("E", "id", &grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.addProduction("E", "(", &grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.addProduction("E′", "+", &grammar.Production{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.addProduction("E′", ")", &grammar.Production{Head: "E′", Body: grammar.E})
	pt0.addProduction("E′", grammar.Endmarker, &grammar.Production{Head: "E′", Body: grammar.E})
	pt0.addProduction("T", "id", &grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.addProduction("T", "(", &grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.addProduction("T′", "+", &grammar.Production{Head: "T′", Body: grammar.E})
	pt0.addProduction("T′", "*", &grammar.Production{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.addProduction("T′", ")", &grammar.Production{Head: "T′", Body: grammar.E})
	pt0.addProduction("T′", grammar.Endmarker, &grammar.Production{Head: "T′", Body: grammar.E})
	pt0.addProduction("F", "id", &grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}})
	pt0.addProduction("F", "(", &grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}})

	pt0.setSync("E", ")", true)
	pt0.setSync("E", grammar.Endmarker, true)
	pt0.setSync("T", "+", true)
	pt0.setSync("T", ")", true)
	pt0.setSync("T", grammar.Endmarker, true)
	pt0.setSync("F", "+", true)
	pt0.setSync("F", "*", true)
	pt0.setSync("F", ")", true)
	pt0.setSync("F", grammar.Endmarker, true)

	pt1 := NewParsingTable(
		[]grammar.Terminal{"a", "b", "e", "i", "t"},
		[]grammar.NonTerminal{"S", "S′", "E"},
	)

	pt1.addProduction("S", "a", &grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("a")}})
	pt1.addProduction("S", "i", &grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("i"), grammar.NonTerminal("E"), grammar.Terminal("t"), grammar.NonTerminal("S"), grammar.NonTerminal("S′")}})
	pt1.addProduction("S′", "e", &grammar.Production{Head: "S′", Body: grammar.E})
	pt1.addProduction("S′", "e", &grammar.Production{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("e"), grammar.NonTerminal("S")}})
	pt1.addProduction("S′", grammar.Endmarker, &grammar.Production{Head: "S′", Body: grammar.E})
	pt1.addProduction("E", "b", &grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("b")}})

	pt2 := NewParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	return []*ParsingTable{pt0, pt1, pt2}
}
