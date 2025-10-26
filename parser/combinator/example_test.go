package combinator_test

import (
	"fmt"

	"github.com/moorara/algo/parser/combinator"
)

func Example() {
	digit := combinator.ExpectRuneInRange('0', '9').Map(toDigit)      // digit → "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
	num := digit.REP1().Map(toNum)                                    // num → digit+
	expr := num.CONCAT(combinator.ExpectRune('+'), num).Map(evalExpr) // expr → num "+" num

	in := newStringInput("27 + 69")
	if out, ok := expr(in); ok {
		n := out.Result.Val.(int)
		fmt.Println(n)
	}
}

func ExampleParser_Map() {
	// Production rule: num → [0-9]+
	num := combinator.ExpectRuneInRange('0', '9').REP1().Map(func(r combinator.Result) (combinator.Result, bool) {
		l := r.Val.(combinator.List)

		num := 0
		for _, r := range l {
			num = num*10 + int(r.Val.(rune)-'0')
		}

		return combinator.Result{Val: num, Pos: l[0].Pos}, true
	})

	in := newStringInput("1024")
	if out, ok := num(in); ok {
		d := out.Result.Val.(int)
		fmt.Println(d)
	}
}

func ExampleParser_Bind() {
	// Production rule: copy → letter* "+" letter*
	// Context-sensitive constraint: the two letter* must be identical.

	letters := combinator.ExpectRuneInRange('a', 'z').REP().Map(toString)
	copy := combinator.CONCAT(
		letters,
		combinator.ExpectRune('+'),
		letters,
	).Bind(func(r combinator.Result) combinator.Parser {
		return func(in combinator.Input) (combinator.Output, bool) {
			r0, _ := r.Get(0)
			r2, _ := r.Get(2)

			s0 := r0.Val.(string)
			s1 := r2.Val.(string)

			if s0 == s1 {
				return combinator.Output{
					Result: combinator.Result{
						Val: s0 + s1,
						Pos: r0.Pos,
					},
					Remaining: in,
				}, true
			}

			return combinator.Output{}, false
		}
	})

	for _, s := range []string{"foo + bar", "baz + baz"} {
		in := newStringInput(s)
		if out, ok := copy(in); ok {
			fmt.Println("Successfully parsed:", out.Result.Val)
		} else {
			fmt.Println("Failed to parse:", s)
		}
	}
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

func toString(r combinator.Result) (combinator.Result, bool) {
	l := r.Val.(combinator.List)

	runes := make([]rune, len(l))
	for i, r := range l {
		runes[i] = r.Val.(rune)
	}

	return combinator.Result{
		Val: string(runes),
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
	// Ignore whitespaces
	if s.runes[0] == ' ' {
		s.pos, s.runes = s.pos+1, s.runes[1:]
		return s.Current()
	}

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
