package combinator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestE(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "OK",
			in:   newStringInput("abc"),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("abc"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := E(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestExpectRune(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		r             rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			r:             'a',
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			r:             'b',
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("a"),
			r:    'a',
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("ab"),
			r:    'a',
			expectedOut: &Output{
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
			out, err := ExpectRune(tc.r)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestNotExpectRune(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		r             rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			r:             'a',
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			r:             'a',
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("a"),
			r:    'b',
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("ab"),
			r:    'b',
			expectedOut: &Output{
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
			out, err := NotExpectRune(tc.r)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestExpectRuneIn(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		runes         []rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			runes:         []rune{'a', 'b'},
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			runes:         []rune{'0', '1'},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:  "Successful_WithoutRemaining",
			in:    newStringInput("a"),
			runes: []rune{'a', 'b'},
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name:  "Successful_WithRemaining",
			in:    newStringInput("ab"),
			runes: []rune{'a', 'b'},
			expectedOut: &Output{
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
			out, err := ExpectRuneIn(tc.runes...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestNotExpectRuneIn(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		runes         []rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			runes:         []rune{'A', 'B'},
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			runes:         []rune{'a', 'b'},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:  "Successful_WithoutRemaining",
			in:    newStringInput("a"),
			runes: []rune{'A', 'B'},
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name:  "Successful_WithRemaining",
			in:    newStringInput("ab"),
			runes: []rune{'A', 'B'},
			expectedOut: &Output{
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
			out, err := NotExpectRuneIn(tc.runes...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestExpectRuneInRange(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		lo, hi        rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			lo:            'a',
			hi:            'z',
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			lo:            '0',
			hi:            '9',
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:          "Invalid_Range",
			in:            newStringInput("a"),
			lo:            'z',
			hi:            'a',
			expectedError: "invalid range [z,a]",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("a"),
			lo:   'a',
			hi:   'z',
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("ab"),
			lo:   'a',
			hi:   'z',
			expectedOut: &Output{
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
			out, err := ExpectRuneInRange(tc.lo, tc.hi)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestNotExpectRuneInRange(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		lo, hi        rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			lo:            'A',
			hi:            'Z',
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("a"),
			lo:            'a',
			hi:            'z',
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:          "Invalid_Range",
			in:            newStringInput("a"),
			lo:            'Z',
			hi:            'A',
			expectedError: "invalid range [Z,A]",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("a"),
			lo:   'A',
			hi:   'Z',
			expectedOut: &Output{
				Result:    Result{'a', 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("ab"),
			lo:   'A',
			hi:   'Z',
			expectedOut: &Output{
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
			out, err := NotExpectRuneInRange(tc.lo, tc.hi)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestExpectRunes(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		runes         []rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			runes:         []rune{'a', 'b'},
			expectedError: "end of input",
		},
		{
			name:          "Input_NotEnough",
			in:            newStringInput("a"),
			runes:         []rune{'a', 'b'},
			expectedError: "end of input",
		},
		{
			name:          "Input_NotMatching",
			in:            newStringInput("ab"),
			runes:         []rune{'b', 'a'},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:  "Successful_EmptyRunes",
			in:    newStringInput("ab"),
			runes: []rune{},
			expectedOut: &Output{
				Result:    Result{[]rune{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:  "Successful_WithoutRemaining",
			in:    newStringInput("ab"),
			runes: []rune{'a', 'b'},
			expectedOut: &Output{
				Result:    Result{[]rune{'a', 'b'}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name:  "Successful_WithRemaining",
			in:    newStringInput("abcd"),
			runes: []rune{'a', 'b'},
			expectedOut: &Output{
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
			out, err := ExpectRunes(tc.runes...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestNotExpectRunes(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		runes         []rune
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			runes:         []rune{'b', 'a'},
			expectedError: "end of input",
		},
		{
			name:          "Input_NotEnough",
			in:            newStringInput("a"),
			runes:         []rune{'b', 'a'},
			expectedError: "end of input",
		},
		{
			name:          "Input_Matching",
			in:            newStringInput("ab"),
			runes:         []rune{'a', 'b'},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:  "Successful_EmptyRunes",
			in:    newStringInput("ab"),
			runes: []rune{},
			expectedOut: &Output{
				Result:    Result{[]rune{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name:  "Successful_WithoutRemaining",
			in:    newStringInput("ab"),
			runes: []rune{'b', 'a'},
			expectedOut: &Output{
				Result:    Result{[]rune{'a', 'b'}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name:  "Successful_WithRemaining",
			in:    newStringInput("abcd"),
			runes: []rune{'b', 'a'},
			expectedOut: &Output{
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
			out, err := NotExpectRunes(tc.runes...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestExpectString(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		s             string
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			s:             "ab",
			expectedError: "end of input",
		},
		{
			name:          "Input_NotEnough",
			in:            newStringInput("a"),
			s:             "ab",
			expectedError: "end of input",
		},
		{
			name:          "Input_NotMatching",
			in:            newStringInput("ab"),
			s:             "09",
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_EmptyString",
			in:   newStringInput("ab"),
			s:    "",
			expectedOut: &Output{
				Result:    Result{"", 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("ab"),
			s:    "ab",
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_With_Remaining",
			in:   newStringInput("abcd"),
			s:    "ab",
			expectedOut: &Output{
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
			out, err := ExpectString(tc.s)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestNotExpectString(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		s             string
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			s:             "ba",
			expectedError: "end of input",
		},
		{
			name:          "Input_NotEnough",
			in:            newStringInput("a"),
			s:             "ba",
			expectedError: "end of input",
		},
		{
			name:          "Input_Matching",
			in:            newStringInput("ab"),
			s:             "ab",
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_EmptyString",
			in:   newStringInput("ab"),
			s:    "",
			expectedOut: &Output{
				Result:    Result{"", 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("ab"),
			s:    "ba",
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			s:    "ba",
			expectedOut: &Output{
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
			out, err := NotExpectString(tc.s)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestALT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             []Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Parsers_Empty",
			in:   newStringInput("ab"),
			p:    []Parser{},
			expectedOut: &Output{
				Result: Result{Empty{}, 0, nil},
				Remaining: &stringInput{
					pos:   0,
					runes: []rune("ab"),
				},
			},
		},
		{
			name: "Input_Empty",
			in:   nil,
			p: []Parser{
				ExpectString("ab"),
				ExpectString("00"),
			},
			expectedError: "end of input",
		},
		{
			name: "Parser_Unsuccessful",
			in:   newStringInput("ab"),
			p: []Parser{
				ExpectString("00"),
				ExpectString("11"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_FirstParser",
			in:   newStringInput("ab"),
			p: []Parser{
				ExpectString("ab"),
				ExpectString("00"),
			},
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_SecondParser",
			in:   newStringInput("ab"),
			p: []Parser{
				ExpectString("00"),
				ExpectString("ab"),
			},
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("ab"),
				ExpectString("cd"),
			},
			expectedOut: &Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
		{
			name: "Unsuccessful_MultipleParsers",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("00"),
				ExpectString("11"),
				ExpectString("22"),
				ExpectString("33"),
				ExpectString("44"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_MultipleParsers",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("00"),
				ExpectString("11"),
				ExpectString("22"),
				ExpectString("33"),
				ExpectString("ab"),
			},
			expectedOut: &Output{
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
			out, err := ALT(tc.p...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestCONCAT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             []Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Parsers_Empty",
			in:   newStringInput("ab"),
			p:    []Parser{},
			expectedOut: &Output{
				Result: Result{Empty{}, 0, nil},
				Remaining: &stringInput{
					pos:   0,
					runes: []rune("ab"),
				},
			},
		},
		{
			name: "Input_Empty",
			in:   nil,
			p: []Parser{
				ExpectString("b"),
				ExpectString("a"),
			},
			expectedError: "end of input",
		},
		{
			name: "Input_NotEnough",
			in:   newStringInput("a"),
			p: []Parser{
				ExpectString("b"),
				ExpectString("a"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Unsuccessful_FirstParser",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("cd"),
				ExpectString("00"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Unsuccessful_SecondParser",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("00"),
				ExpectString("ab"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("abcd"),
			p: []Parser{
				ExpectString("ab"),
				ExpectString("cd"),
			},
			expectedOut: &Output{
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
			name: "Successful_WithRemaining",
			in:   newStringInput("abcdef"),
			p: []Parser{
				ExpectString("ab"),
				ExpectString("cd"),
			},
			expectedOut: &Output{
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
			name: "Unsuccessful_MultipleParsers",
			in:   newStringInput("abcdefghijklmn"),
			p: []Parser{
				ExpectString("cd"),
				ExpectString("ef"),
				ExpectString("ij"),
				ExpectString("ab"),
			},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_MultipleParsers",
			in:   newStringInput("abcdefghijklmn"),
			p: []Parser{
				ExpectString("ab"),
				ExpectString("cd"),
				ExpectString("ef"),
				ExpectString("gh"),
				ExpectString("ij"),
			},
			expectedOut: &Output{
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
			out, err := CONCAT(tc.p...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestOPT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Input_Empty",
			in:   nil,
			p:    ExpectString("ab"),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_EmptyResult",
			in:   newStringInput("ab"),
			p:    ExpectString("00"),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("ab"),
			p:    ExpectString("ab"),
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab"),
			expectedOut: &Output{
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
			out, err := OPT(tc.p)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestREP(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Input_Empty",
			in:   nil,
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_Zero",
			in:   newStringInput("ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_One",
			in:   newStringInput("1ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_Many",
			in:   newStringInput("1234ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_WithoutRemaining",
			in:   newStringInput("1234"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			out, err := REP(tc.p)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestREP1(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRuneInRange('0', '9'),
			expectedError: "end of input",
		},
		{
			name:          "Unsuccessful_Zero",
			in:            newStringInput("ab"),
			p:             ExpectRuneInRange('0', '9'),
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_One",
			in:   newStringInput("1ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_Many",
			in:   newStringInput("1234ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_WithoutRemaining",
			in:   newStringInput("1234"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			out, err := REP1(tc.p)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_ALT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		q             []Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectString("ab"),
			q:             []Parser{ExpectString("00")},
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("ab"),
			p:             ExpectString("00"),
			q:             []Parser{ExpectString("11")},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_FirstParser",
			in:   newStringInput("ab"),
			p:    ExpectString("ab"),
			q:    []Parser{ExpectString("00")},
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_SecondParser",
			in:   newStringInput("ab"),
			p:    ExpectString("00"),
			q:    []Parser{ExpectString("ab")},
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab"),
			q:    []Parser{ExpectString("cd")},
			expectedOut: &Output{
				Result: Result{"ab", 0, nil},
				Remaining: &stringInput{
					pos:   2,
					runes: []rune("cd"),
				},
			},
		},
		{
			name:          "Unsuccessful_MultipleParsers",
			in:            newStringInput("abcd"),
			p:             ExpectString("00"),
			q:             []Parser{ExpectString("11"), ExpectString("22"), ExpectString("33"), ExpectString("44")},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_MultipleParsers",
			in:   newStringInput("abcd"),
			p:    ExpectString("00"),
			q:    []Parser{ExpectString("11"), ExpectString("22"), ExpectString("33"), ExpectString("ab")},
			expectedOut: &Output{
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
			out, err := tc.p.ALT(tc.q...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_CONCAT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		q             []Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectString("a"),
			q:             []Parser{ExpectString("b")},
			expectedError: "end of input",
		},
		{
			name:          "Input_NotEnough",
			in:            newStringInput("a"),
			p:             ExpectString("a"),
			q:             []Parser{ExpectString("b")},
			expectedError: "end of input",
		},
		{
			name:          "Unsuccessful_FirstParser",
			in:            newStringInput("abcd"),
			p:             ExpectString("00"),
			q:             []Parser{ExpectString("cd")},
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name:          "Unsuccessful_SecondParser",
			in:            newStringInput("abcd"),
			p:             ExpectString("ab"),
			q:             []Parser{ExpectString("00")},
			expectedError: "unexpected rune 'c' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab"),
			q:    []Parser{ExpectString("cd")},
			expectedOut: &Output{
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
			name: "Successful_WithRemaining",
			in:   newStringInput("abcdef"),
			p:    ExpectString("ab"),
			q:    []Parser{ExpectString("cd")},
			expectedOut: &Output{
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
			name:          "Unsuccessful_MultipleParsers",
			in:            newStringInput("abcdefghijklmn"),
			p:             ExpectString("ab"),
			q:             []Parser{ExpectString("cd"), ExpectString("ef"), ExpectString("ij")},
			expectedError: "unexpected rune 'g' at position 0",
		},
		{
			name: "Successful_MultipleParsers",
			in:   newStringInput("abcdefghijklmn"),
			p:    ExpectString("ab"),
			q:    []Parser{ExpectString("cd"), ExpectString("ef"), ExpectString("gh"), ExpectString("ij")},
			expectedOut: &Output{
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
			out, err := tc.p.CONCAT(tc.q...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_OPT(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Input_Empty",
			in:   nil,
			p:    ExpectString("ab"),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_EmptyResult",
			in:   newStringInput("ab"),
			p:    ExpectString("00"),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("ab"),
			p:    ExpectString("ab"),
			expectedOut: &Output{
				Result:    Result{"ab", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab"),
			expectedOut: &Output{
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
			out, err := tc.p.OPT()(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_REP(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name: "Input_Empty",
			in:   nil,
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_Zero",
			in:   newStringInput("ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
				Result:    Result{Empty{}, 0, nil},
				Remaining: newStringInput("ab"),
			},
		},
		{
			name: "Successful_One",
			in:   newStringInput("1ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_Many",
			in:   newStringInput("1234ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_WithoutRemaining",
			in:   newStringInput("1234"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			out, err := tc.p.REP()(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_REP1(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRuneInRange('0', '9'),
			expectedError: "end of input",
		},
		{
			name:          "Unsuccessful_Zero",
			in:            newStringInput("ab"),
			p:             ExpectRuneInRange('0', '9'),
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_One",
			in:   newStringInput("1ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_Many",
			in:   newStringInput("1234ab"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			name: "Successful_WithoutRemaining",
			in:   newStringInput("1234"),
			p:    ExpectRuneInRange('0', '9'),
			expectedOut: &Output{
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
			out, err := tc.p.REP1()(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
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
		name          string
		in            Input
		p             Parser
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRune('!'),
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("{2,4}"),
			p:             ExpectRune('!'),
			expectedError: "unexpected rune '{' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("{2,4}"),
			p:    rangeParser,
			expectedOut: &Output{
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
			name: "Successful_WithRemaining",
			in:   newStringInput("{2,4}ab"),
			p:    rangeParser,
			expectedOut: &Output{
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
			out, err := tc.p.Flatten()(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
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
		name          string
		in            Input
		p             Parser
		pos           []int
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRune('!'),
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("{2,4}"),
			p:             ExpectRune('!'),
			expectedError: "unexpected rune '{' at position 0",
		},
		{
			name: "Result_NotList",
			in:   newStringInput("{2,4}"),
			p:    ExpectString("{2,4}"),
			expectedOut: &Output{
				Result:    Result{"{2,4}", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Indices_Invalid",
			in:   newStringInput("{2,4}"),
			p:    rangeParser,
			pos:  []int{-1, 5},
			expectedOut: &Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("{2,4}"),
			p:    rangeParser,
			pos:  []int{1, 3},
			expectedOut: &Output{
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
			name: "Successful_WithRemaining",
			in:   newStringInput("{2,4}ab"),
			p:    rangeParser,
			pos:  []int{1, 3},
			expectedOut: &Output{
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
			out, err := tc.p.Select(tc.pos...)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_Get(t *testing.T) {
	tests := []struct {
		name          string
		in            Input
		p             Parser
		i             int
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRune('!'),
			i:             0,
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("ab"),
			p:             ExpectRune('!'),
			i:             0,
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Result_NotList",
			in:   newStringInput("abcd"),
			p:    ExpectString("abcd"),
			i:    -1,
			expectedOut: &Output{
				Result:    Result{"abcd", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Index_LT_Zero",
			in:   newStringInput("abcd"),
			p:    ExpectRuneInRange('a', 'z').REP(),
			i:    -1,
			expectedOut: &Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name: "Index_GEQ_Len",
			in:   newStringInput("abcd"),
			p:    ExpectRuneInRange('a', 'z').REP(),
			i:    4,
			expectedOut: &Output{
				Result:    Result{Val: Empty{}},
				Remaining: nil,
			},
		},
		{
			name: "Successful_CONCAT",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab").CONCAT(ExpectString("cd")),
			i:    1,
			expectedOut: &Output{
				Result:    Result{"cd", 2, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_REP",
			in:   newStringInput("abcd"),
			p:    ExpectRuneIn('a', 'b', 'c', 'd').REP(),
			i:    2,
			expectedOut: &Output{
				Result:    Result{'c', 2, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_REP1",
			in:   newStringInput("abcd"),
			p:    ExpectRuneInRange('a', 'z').REP(),
			i:    3,
			expectedOut: &Output{
				Result:    Result{'d', 3, nil},
				Remaining: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := tc.p.Get(tc.i)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_Map(t *testing.T) {
	toUpper := func(r Result) (Result, error) {
		return Result{
			Val: strings.ToUpper(r.Val.(string)),
			Pos: r.Pos,
		}, nil
	}

	tests := []struct {
		name          string
		in            Input
		p             Parser
		f             MapFunc
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRune('!'),
			f:             toUpper,
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("ab"),
			p:             ExpectRune('!'),
			f:             toUpper,
			expectedError: "unexpected rune 'a' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("ab"),
			p:    ExpectString("ab"),
			f:    toUpper,
			expectedOut: &Output{
				Result:    Result{"AB", 0, nil},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("abcd"),
			p:    ExpectString("ab"),
			f:    toUpper,
			expectedOut: &Output{
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
			out, err := tc.p.Map(tc.f)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParser_Bind(t *testing.T) {
	// This syntax annotation is bound on a digit parser and expects that many 'a' runes to follow.
	// For example, input "4aaaa" is valid, while "4aaa" is not.
	annotate := func(r Result) Parser {
		count := int(r.Val.(rune) - '0')
		seq := make([]Parser, 0, count)
		for range count {
			seq = append(seq, ExpectRune('a'))
		}
		return CONCAT(seq...)
	}

	tests := []struct {
		name          string
		in            Input
		p             Parser
		f             BindFunc
		expectedOut   *Output
		expectedError string
	}{
		{
			name:          "Input_Empty",
			in:            nil,
			p:             ExpectRuneInRange('0', '9'),
			f:             annotate,
			expectedError: "end of input",
		},
		{
			name:          "Parser_Unsuccessful",
			in:            newStringInput("4aaaa"),
			p:             ExpectRuneInRange('a', 'f'),
			f:             annotate,
			expectedError: "unexpected rune '4' at position 0",
		},
		{
			name: "Successful_WithoutRemaining",
			in:   newStringInput("4aaaa"),
			p:    ExpectRuneInRange('0', '9'),
			f:    annotate,
			expectedOut: &Output{
				Result: Result{
					Val: List{
						{'a', 1, nil},
						{'a', 2, nil},
						{'a', 3, nil},
						{'a', 4, nil},
					},
					Pos: 1,
					Bag: nil,
				},
				Remaining: nil,
			},
		},
		{
			name: "Successful_WithRemaining",
			in:   newStringInput("4aaaa-tail"),
			p:    ExpectRuneInRange('0', '9'),
			f:    annotate,
			expectedOut: &Output{
				Result: Result{
					Val: List{
						{'a', 1, nil},
						{'a', 2, nil},
						{'a', 3, nil},
						{'a', 4, nil},
					},
					Pos: 1,
					Bag: nil,
				},
				Remaining: &stringInput{
					pos:   5,
					runes: []rune("-tail"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := tc.p.Bind(tc.f)(tc.in)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOut, out)
			} else {
				assert.Nil(t, out)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
