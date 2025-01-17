package combinator_test

import (
	"fmt"

	"github.com/moorara/algo/parser/combinator"
)

// stringInput implements the combinator.Input interface for strings.
type stringInput struct {
	pos   int
	runes []rune
}

func newStringInput(s string) combinator.Input {
	return &stringInput{
		pos:   0,
		runes: []rune(s),
	}
}

func (s *stringInput) Current() (rune, int) {
	return s.runes[0], s.pos
}

func (s *stringInput) Remaining() combinator.Input {
	if len(s.runes) == 1 {
		return nil
	}

	return &stringInput{
		pos:   s.pos + 1,
		runes: s.runes[1:],
	}
}

func Example() {
	add := combinator.ExpectRune('+').Map(toAdd)                 // add → "+"
	digit := combinator.ExpectRuneInRange('0', '9').Map(toDigit) // digit → "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
	num := digit.REP1().Map(toNum)                               // num → digit+
	expr := num.CONCAT(add, num).Map(evalExpr)                   // expr → num "+" num

	in := newStringInput("27+69")
	if out, ok := expr(in); ok {
		n := out.Result.Val.(int)
		fmt.Println(n)
	}
}

func toAdd(r combinator.Result) (combinator.Result, bool) {
	return combinator.Result{
		Val: nil,
		Pos: r.Pos,
	}, true
}

func toDigit(r combinator.Result) (combinator.Result, bool) {
	v := r.Val.(rune)
	digit := int(v - '0')

	return combinator.Result{
		Val: digit,
		Pos: r.Pos,
	}, true
}

func toNum(r combinator.Result) (combinator.Result, bool) {
	l := r.Val.(combinator.List)

	var num int
	for _, r := range l {
		num = num*10 + r.Val.(int)
	}

	return combinator.Result{
		Val: num,
		Pos: l[0].Pos,
	}, true
}

func evalExpr(r combinator.Result) (combinator.Result, bool) {
	r0, _ := r.Get(0)
	r2, _ := r.Get(2)

	n0 := r0.Val.(int)
	n2 := r2.Val.(int)

	return combinator.Result{
		Val: int(n0 + n2),
		Pos: r0.Pos,
	}, true
}
