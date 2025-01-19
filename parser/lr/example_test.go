package lr_test

import (
	"fmt"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

func ExampleStateMap() {
	start := grammar.NonTerminal("E′")

	p0 := &grammar.Production{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}}                                                 // E′ → E
	p1 := &grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}} // E → E + T
	p2 := &grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}}                                                  // E → T
	p3 := &grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}} // T → T * F
	p4 := &grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}}                                                  // T → F
	p5 := &grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}    // F → ( E )
	p6 := &grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}}                                                    // F → id

	I0 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p0, Start: &start, Dot: 0}, // E′ → •E
		// Non-Kernels
		lr.Item0{Production: p1, Start: &start, Dot: 0}, // E → •E + T
		lr.Item0{Production: p2, Start: &start, Dot: 0}, // E → •T
		lr.Item0{Production: p3, Start: &start, Dot: 0}, // T → •T * F
		lr.Item0{Production: p4, Start: &start, Dot: 0}, // T → •F
		lr.Item0{Production: p5, Start: &start, Dot: 0}, // F → •( E )
		lr.Item0{Production: p6, Start: &start, Dot: 0}, // F → •id
	)

	I1 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p0, Start: &start, Dot: 1}, // E′ → E•
		lr.Item0{Production: p1, Start: &start, Dot: 1}, // E → E•+ T
	)

	I2 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p2, Start: &start, Dot: 1}, // E → T•
		lr.Item0{Production: p3, Start: &start, Dot: 1}, // T → T•* F
	)

	I3 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p4, Start: &start, Dot: 1}, // T → F•
	)

	I4 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p5, Start: &start, Dot: 1}, // F → (•E )
		// Non-Kernels
		lr.Item0{Production: p1, Start: &start, Dot: 0}, // E → •E + T
		lr.Item0{Production: p2, Start: &start, Dot: 0}, // E → •T
		lr.Item0{Production: p3, Start: &start, Dot: 0}, // T → •T * F
		lr.Item0{Production: p4, Start: &start, Dot: 0}, // T → •F
		lr.Item0{Production: p5, Start: &start, Dot: 0}, // F → •( E )
		lr.Item0{Production: p6, Start: &start, Dot: 0}, // F → •id
	)

	I5 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p6, Start: &start, Dot: 1}, // F → id•
	)

	I6 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p1, Start: &start, Dot: 2}, // E → E +•T
		// Non-Kernels
		lr.Item0{Production: p3, Start: &start, Dot: 0}, // T → •T * F
		lr.Item0{Production: p4, Start: &start, Dot: 0}, // T → •F
		lr.Item0{Production: p5, Start: &start, Dot: 0}, // F → •( E )
		lr.Item0{Production: p6, Start: &start, Dot: 0}, // F → •id
	)

	I7 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p3, Start: &start, Dot: 2}, // T → T *•F
		// Non-Kernels
		lr.Item0{Production: p5, Start: &start, Dot: 0}, // F → •( E )
		lr.Item0{Production: p6, Start: &start, Dot: 0}, // F → •id
	)

	I8 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p1, Start: &start, Dot: 1}, // E → E• + T
		lr.Item0{Production: p5, Start: &start, Dot: 2}, // F → ( E•)
	)

	I9 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p1, Start: &start, Dot: 3}, // E → E + T•
		lr.Item0{Production: p3, Start: &start, Dot: 1}, // T → T•* F
	)

	I10 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p3, Start: &start, Dot: 3}, // T → T * F•
	)

	I11 := lr.NewItemSet(
		// Kernels
		lr.Item0{Production: p5, Start: &start, Dot: 3}, // F → ( E )•
	)

	C := lr.NewItemSetCollection(I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11)
	states := lr.BuildStateMap(C)
	fmt.Println(states)
}
