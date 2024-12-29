package combinator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// empty --> ε
var _empty = func(in Input) (Output, bool) {
	return Output{
		Result: Result{
			Val: Empty{},
		},
		Remaining: in,
	}, true
}

// stringInput implements the input interface for strings.
type stringInput struct {
	pos   int
	runes []rune
}

func newStringInput(s string) Input {
	return &stringInput{
		pos:   0,
		runes: []rune(s),
	}
}

func (s *stringInput) Current() (rune, int) {
	return s.runes[0], s.pos
}

func (s *stringInput) Remaining() Input {
	if len(s.runes) == 1 {
		return nil
	}

	return &stringInput{
		pos:   s.pos + 1,
		runes: s.runes[1:],
	}
}

func TestResult_Get(t *testing.T) {
	tests := []struct {
		name           string
		r              Result
		i              int
		expectedOK     bool
		expectedResult Result
	}{
		{
			name: "Not_List",
			r: Result{
				Val: 'a',
			},
			i:          0,
			expectedOK: false,
		},
		{
			name: "Index_LT_Zero",
			r: Result{
				Val: List{
					Result{'a', 0, nil},
					Result{'b', 1, nil},
					Result{'c', 2, nil},
					Result{'d', 3, nil},
				},
			},
			i:          -1,
			expectedOK: false,
		},
		{
			name: "Index_GEQ_Len",
			r: Result{
				Val: List{
					Result{'a', 0, nil},
					Result{'b', 1, nil},
					Result{'c', 2, nil},
					Result{'d', 3, nil},
				},
			},
			i:          4,
			expectedOK: false,
		},
		{
			name: "Successful",
			r: Result{
				Val: List{
					Result{'a', 0, nil},
					Result{'b', 1, nil},
					Result{'c', 2, nil},
					Result{'d', 3, nil},
				},
			},
			i:              2,
			expectedOK:     true,
			expectedResult: Result{'c', 2, nil},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, ok := tc.r.Get(tc.i)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}

func TestExpectRune(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		r           rune
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			r:          'a',
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("a"),
			r:          'b',
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("a"),
			r:          'a',
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("ab"),
			r:          'a',
			expectedOK: true,
			expectedOut: Output{
				Result: Result{'a', 0, nil},
				Remaining: &stringInput{
					pos:   1,
					runes: []rune("b"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := ExpectRune(tc.r)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestExpectRuneIn(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		runes       []rune
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			runes:      []rune{'a', 'b'},
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("a"),
			runes:      []rune{'0', '1'},
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("a"),
			runes:      []rune{'a', 'b'},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("ab"),
			runes:      []rune{'a', 'b'},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{'a', 0, nil},
				Remaining: &stringInput{
					pos:   1,
					runes: []rune("b"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := ExpectRuneIn(tc.runes...)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestExpectRuneInRange(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		low, up     rune
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			low:        'a',
			up:         'z',
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("a"),
			low:        '0',
			up:         '9',
			expectedOK: false,
		},
		{
			name:       "Invalid_Range",
			in:         newStringInput("a"),
			low:        'z',
			up:         'a',
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("a"),
			low:        'a',
			up:         'z',
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("ab"),
			low:        'a',
			up:         'z',
			expectedOK: true,
			expectedOut: Output{
				Result: Result{'a', 0, nil},
				Remaining: &stringInput{
					pos:   1,
					runes: []rune("b"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := ExpectRuneInRange(tc.low, tc.up)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestExpectRunes(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		runes       []rune
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			runes:      []rune{'a', 'b'},
			expectedOK: false,
		},
		{
			name:       "Input_Not_Enough",
			in:         newStringInput("a"),
			runes:      []rune{'a', 'b'},
			expectedOK: false,
		},
		{
			name:       "Input_Not_Matching",
			in:         newStringInput("ab"),
			runes:      []rune{'0', '9'},
			expectedOK: false,
		},
		{
			name:       "Successful_Empty_Runes",
			in:         newStringInput("ab"),
			runes:      []rune{},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{[]rune{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:       "Successful_Witouth_Remaining",
			in:         newStringInput("ab"),
			runes:      []rune{'a', 'b'},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{[]rune{'a', 'b'}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcd"),
			runes:      []rune{'a', 'b'},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{[]rune{'a', 'b'}, 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := ExpectRunes(tc.runes...)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestExpectString(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		s           string
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			s:          "ab",
			expectedOK: false,
		},
		{
			name:       "Input_Not_Enough",
			in:         newStringInput("a"),
			s:          "ab",
			expectedOK: false,
		},
		{
			name:       "Input_Not_Matching",
			in:         newStringInput("ab"),
			s:          "09",
			expectedOK: false,
		},
		{
			name:       "Successful_Empty_String",
			in:         newStringInput("ab"),
			s:          "",
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"", 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("ab"),
			s:          "ab",
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcd"),
			s:          "ab",
			expectedOK: true,
			expectedOut: Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := ExpectString(tc.s)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_CONCAT(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		q           []Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectString("a"),
			q:          []Parser{ExpectString("b")},
			expectedOK: false,
		},
		{
			name:       "Input_Not_Enough",
			in:         newStringInput("a"),
			p:          ExpectString("a"),
			q:          []Parser{ExpectString("b")},
			expectedOK: false,
		},
		{
			name:       "First_Parser_Unsuccessful",
			in:         newStringInput("abcd"),
			p:          ExpectString("00"),
			q:          []Parser{ExpectString("cd")},
			expectedOK: false,
		},
		{
			name:       "Second_Parser_Unsuccessful",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("00")},
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("cd")},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{"ab", 0, nil},
						Result{"cd", 2, nil},
					},
					Pos: 0,
				},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcdef"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("cd")},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{"ab", 0, nil},
						Result{"cd", 2, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   4,
					runes: []rune("ef"),
				},
			},
		},
		{
			name:       "Unuccessful_Multiple_Parsers",
			in:         newStringInput("abcdefghijklmn"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("cd"), ExpectString("ef"), ExpectString("ij")},
			expectedOK: false,
		},
		{
			name:       "Successful_Multiple_Parsers",
			in:         newStringInput("abcdefghijklmn"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("cd"), ExpectString("ef"), ExpectString("gh"), ExpectString("ij")},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{"ab", 0, nil},
						Result{"cd", 2, nil},
						Result{"ef", 4, nil},
						Result{"gh", 6, nil},
						Result{"ij", 8, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   10,
					runes: []rune("klmn"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.CONCAT(tc.q...)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_ALT(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		q           []Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("00")},
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("ab"),
			p:          ExpectString("00"),
			q:          []Parser{ExpectString("11")},
			expectedOK: false,
		},
		{
			name:       "First_Parser_Successful",
			in:         newStringInput("ab"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("00")},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Second_Parser_Successful",
			in:         newStringInput("ab"),
			p:          ExpectString("00"),
			q:          []Parser{ExpectString("ab")},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab"),
			q:          []Parser{ExpectString("cd")},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
		{
			name:       "Unsuccessful_Multiple_Parsers",
			in:         newStringInput("abcd"),
			p:          ExpectString("00"),
			q:          []Parser{ExpectString("11"), ExpectString("22"), ExpectString("33"), ExpectString("44")},
			expectedOK: false,
		},
		{
			name:       "Successful_Multiple_Parsers",
			in:         newStringInput("abcd"),
			p:          ExpectString("00"),
			q:          []Parser{ExpectString("11"), ExpectString("22"), ExpectString("33"), ExpectString("ab")},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.ALT(tc.q...)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_OPT(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectString("ab"),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_Empty_Result",
			in:         newStringInput("ab"),
			p:          ExpectString("00"),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("ab"),
			p:          ExpectString("ab"),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab"),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.OPT()(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_REP(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_Zero",
			in:         newStringInput("ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:       "Successful_One",
			in:         newStringInput("1ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   1,
					runes: []rune("ab"),
				},
			},
		},
		{
			name:       "Successful_Many",
			in:         newStringInput("1234ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
						Result{'2', 1, nil},
						Result{'3', 2, nil},
						Result{'4', 3, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   4,
					runes: []rune("ab"),
				},
			},
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("1234"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
						Result{'2', 1, nil},
						Result{'3', 2, nil},
						Result{'4', 3, nil},
					},
					Pos: 0,
				},
				Remaining: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.REP()(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_REP1(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: false,
		},
		{
			name:       "Unsuccessful_Zero",
			in:         newStringInput("ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: false,
		},
		{
			name:       "Successful_One",
			in:         newStringInput("1ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   1,
					runes: []rune("ab"),
				},
			},
		},
		{
			name:       "Successful_Many",
			in:         newStringInput("1234ab"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
						Result{'2', 1, nil},
						Result{'3', 2, nil},
						Result{'4', 3, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   4,
					runes: []rune("ab"),
				},
			},
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("1234"),
			p:          ExpectRuneInRange('0', '9'),
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'1', 0, nil},
						Result{'2', 1, nil},
						Result{'3', 2, nil},
						Result{'4', 3, nil},
					},
					Pos: 0,
				},
				Remaining: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.REP1()(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_Flatten(t *testing.T) {
	rangeParser := ExpectRune('{').CONCAT(
		ExpectRuneInRange('0', '9'),
		ExpectRune(','),
		ExpectRune(' ').OPT(),
		ExpectRuneInRange('0', '9'),
		ExpectRune('}'),
	)

	tests := []struct {
		name        string
		in          Input
		p           Parser
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRune('!'),
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("{2,4}"),
			p:          ExpectRune('!'),
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("{2,4}"),
			p:          rangeParser,
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'{', 0, nil},
						Result{'2', 1, nil},
						Result{',', 2, nil},
						Result{'4', 3, nil},
						Result{'}', 4, nil},
					},
					Pos: 0,
				},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("{2,4}ab"),
			p:          rangeParser,
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'{', 0, nil},
						Result{'2', 1, nil},
						Result{',', 2, nil},
						Result{'4', 3, nil},
						Result{'}', 4, nil},
					},
					Pos: 0,
				},
				Remaining: &stringInput{
					pos:   5,
					runes: []rune("ab"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.Flatten()(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_Select(t *testing.T) {
	rangeParser := ExpectRune('{').CONCAT(
		ExpectRuneInRange('0', '9'),
		ExpectRune(','),
		ExpectRuneInRange('0', '9'),
		ExpectRune('}'),
	)

	tests := []struct {
		name        string
		in          Input
		p           Parser
		pos         []int
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRune('!'),
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("{2,4}"),
			p:          ExpectRune('!'),
			expectedOK: false,
		},
		{
			name:       "Result_Not_List",
			in:         newStringInput("{2,4}"),
			p:          ExpectString("{2,4}"),
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"{2,4}", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Indices_Invalid",
			in:         newStringInput("{2,4}"),
			p:          rangeParser,
			pos:        []int{-1, 5},
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("{2,4}"),
			p:          rangeParser,
			pos:        []int{1, 3},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'2', 1, nil},
						Result{'4', 3, nil},
					},
					Pos: 1,
				},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("{2,4}ab"),
			p:          rangeParser,
			pos:        []int{1, 3},
			expectedOK: true,
			expectedOut: Output{
				Result: Result{
					Val: List{
						Result{'2', 1, nil},
						Result{'4', 3, nil},
					},
					Pos: 1,
				},
				Remaining: &stringInput{
					pos:   5,
					runes: []rune("ab"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.Select(tc.pos...)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_Get(t *testing.T) {
	tests := []struct {
		name        string
		in          Input
		p           Parser
		i           int
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRune('!'),
			i:          0,
			expectedOK: false,
		},
		{
			name:       "Parser_Unuccessful",
			in:         newStringInput("ab"),
			p:          ExpectRune('!'),
			i:          0,
			expectedOK: false,
		},
		{
			name:       "Result_Not_List",
			in:         newStringInput("abcd"),
			p:          ExpectString("abcd"),
			i:          -1,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"abcd", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Index_LT_Zero",
			in:         newStringInput("abcd"),
			p:          ExpectRuneInRange('a', 'z').REP(),
			i:          -1,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name:       "Index_GEQ_Len",
			in:         newStringInput("abcd"),
			p:          ExpectRuneInRange('a', 'z').REP(),
			i:          4,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_CONCAT",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab").CONCAT(ExpectString("cd")),
			i:          1,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"cd", 2, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_REP",
			in:         newStringInput("abcd"),
			p:          ExpectRuneIn('a', 'b', 'c', 'd').REP(),
			i:          2,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{'c', 2, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_REP1",
			in:         newStringInput("abcd"),
			p:          ExpectRuneInRange('a', 'z').REP(),
			i:          3,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{'d', 3, nil},
				Remaining: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.Get(tc.i)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_Map(t *testing.T) {
	toUpper := func(r Result) (Result, bool) {
		return Result{
			Val: strings.ToUpper(r.Val.(string)),
			Pos: r.Pos,
		}, true
	}

	tests := []struct {
		name        string
		in          Input
		p           Parser
		f           Mapper
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRune('!'),
			f:          toUpper,
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("ab"),
			p:          ExpectRune('!'),
			f:          toUpper,
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("ab"),
			p:          ExpectString("ab"),
			f:          toUpper,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{"AB", 0, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("abcd"),
			p:          ExpectString("ab"),
			f:          toUpper,
			expectedOK: true,
			expectedOut: Output{
				Result: Result{"AB", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.Map(tc.f)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func TestParser_Bind(t *testing.T) {
	annotate := func(r Result) Parser {
		if r.Val.(rune) == '(' {
			return ExpectRune(' ').REP().CONCAT(ExpectRune(')')).Get(1)
		}
		return _empty
	}

	tests := []struct {
		name        string
		in          Input
		p           Parser
		f           Binder
		expectedOK  bool
		expectedOut Output
	}{
		{
			name:       "Input_Empty",
			in:         nil,
			p:          ExpectRune('('),
			f:          annotate,
			expectedOK: false,
		},
		{
			name:       "Parser_Unsuccessful",
			in:         newStringInput("(  )"),
			p:          ExpectRune('['),
			f:          annotate,
			expectedOK: false,
		},
		{
			name:       "Successful_Without_Remaining",
			in:         newStringInput("(  )"),
			p:          ExpectRune('('),
			f:          annotate,
			expectedOK: true,
			expectedOut: Output{
				Result:    Result{')', 3, nil},
				Remaining: nil,
			},
		},
		{
			name:       "Successful_With_Remaining",
			in:         newStringInput("(  )tail"),
			p:          ExpectRune('('),
			f:          annotate,
			expectedOK: true,
			expectedOut: Output{
				Result: Result{')', 3, nil},
				Remaining: &stringInput{
					pos:   4,
					runes: []rune("tail"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, ok := tc.p.Bind(tc.f)(tc.in)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}
