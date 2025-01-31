package predictive

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
)

func getTestParsingTables() []*ParsingTable {
	pt0 := NewParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
	)

	pt0.addProduction("E", "(", parsertest.Prods[0][0])                // E → T E′
	pt0.addProduction("E", "id", parsertest.Prods[0][0])               // E → T E′
	pt0.addProduction("E′", ")", parsertest.Prods[0][2])               // E′ → ε
	pt0.addProduction("E′", "+", parsertest.Prods[0][1])               // E′ → + T E′
	pt0.addProduction("E′", grammar.Endmarker, parsertest.Prods[0][2]) // E′ → ε
	pt0.addProduction("T", "(", parsertest.Prods[0][3])                // T → F T′
	pt0.addProduction("T", "id", parsertest.Prods[0][3])               // T → F T′
	pt0.addProduction("T′", ")", parsertest.Prods[0][5])               // T′ → ε
	pt0.addProduction("T′", "*", parsertest.Prods[0][4])               // T′ → * F T′
	pt0.addProduction("T′", "+", parsertest.Prods[0][5])               // T′ → ε
	pt0.addProduction("T′", grammar.Endmarker, parsertest.Prods[0][5]) // T′ → ε
	pt0.addProduction("F", "(", parsertest.Prods[0][6])                // F → ( E )
	pt0.addProduction("F", "id", parsertest.Prods[0][7])               // F → id

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
