package input

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/list"
)

func newStack(n int, vs ...int) list.Stack[int] {
	s := list.NewStack[int](n, nil)
	for _, v := range vs {
		s.Push(v)
	}

	return s
}

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		src           io.Reader
		expectedError string
	}{
		{
			name:          "Success",
			n:             4096,
			src:           strings.NewReader("Lorem ipsum"),
			expectedError: "",
		},
		{
			name:          "Failure",
			n:             4096,
			src:           iotest.ErrReader(errors.New("io error")),
			expectedError: "io error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in, err := New(tc.n, tc.src)

			if tc.expectedError == "" {
				assert.NotNil(t, in)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, in)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestInput_loadFirst(t *testing.T) {
	tests := []struct {
		name          string
		in            *Input
		expectedError string
	}{
		{
			name: "Success",
			in: &Input{
				src:  strings.NewReader("Lorem ipsum"),
				buff: make([]byte, 2048),
			},
			expectedError: "",
		},
		{
			name: "Failure",
			in: &Input{
				src:  iotest.ErrReader(errors.New("io error")),
				buff: make([]byte, 2048),
			},
			expectedError: "io error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.loadFirst()

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestInput_loadSecond(t *testing.T) {
	tests := []struct {
		name          string
		in            *Input
		expectedError string
	}{
		{
			name: "Success",
			in: &Input{
				src:  strings.NewReader("Lorem ipsum"),
				buff: make([]byte, 2048),
			},
			expectedError: "",
		},
		{
			name: "Failure",
			in: &Input{
				src:  iotest.ErrReader(errors.New("io error")),
				buff: make([]byte, 2048),
			},
			expectedError: "io error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.loadSecond()

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestInput_next(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		file          string
		expectedCount int
	}{
		{
			name:          "Success",
			n:             1024,
			file:          "./fixture/lorem_ipsum",
			expectedCount: 3422,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(tc.file)
			assert.NoError(t, err)
			defer f.Close()

			in, err := New(tc.n, f)
			assert.NoError(t, err)

			var b byte
			var count int

			for b, err = in.next(); err == nil; b, err = in.next() {
				count++
				assert.NotZero(t, b)
			}

			assert.Equal(t, io.EOF, err)
			assert.Equal(t, tc.expectedCount, count)
		})
	}
}

func TestInput_Next(t *testing.T) {
	// By putting 10 elements in the buff,
	// we ensure that we won't need to load the second half of the buffer.

	tests := []struct {
		name          string
		in            *Input
		expectedError string
		expectedRune  rune
		expectedSize  int
	}{
		{
			name: "FirstByte_EOF",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         io.EOF,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FirstByte_Invalid",
			in: &Input{
				src:         nil,
				buff:        []byte{0x80, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "invalid utf-8 character at 0",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FirstByte_Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0x69, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'i',
			expectedSize:  1,
		},
		{
			name: "SecondByte_EOF",
			in: &Input{
				src:         nil,
				buff:        []byte{0xC6, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "SecondByte_Invalid",
			in: &Input{
				src:         nil,
				buff:        []byte{0xC6, 0x40, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "invalid utf-8 character at 0",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "SecondByte_Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0xC6, 0xA9, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'Ʃ',
			expectedSize:  2,
		},
		{
			name: "ThirdByte_EOF",
			in: &Input{
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "ThirdByte_Invalid",
			in: &Input{
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x40, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "invalid utf-8 character at 0",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "ThirdByte_Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x80, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'ꩀ',
			expectedSize:  3,
		},
		{
			name: "FourthByte_EOF",
			in: &Input{
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FourthByte_Invalid",
			in: &Input{
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x40, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "invalid utf-8 character at 0",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FourthByte_Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x80, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  '𐀀',
			expectedSize:  4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, err := tc.in.Next()

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
				assert.Equal(t, tc.expectedRune, r)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRune, r)

				size, ok := tc.in.runeSizes.Peek()
				assert.True(t, ok)
				assert.Equal(t, tc.expectedSize, size)
			}
		})
	}
}

func TestInput_Retract(t *testing.T) {
	tests := []struct {
		name            string
		in              *Input
		retractCount    int
		expectedForward int
	}{
		{
			name: "Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     4,
				runeCount:   1,
				runeSizes:   newStack(4, 1, 2),
				err:         nil,
			},
			retractCount:    2,
			expectedForward: 1,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				runeCount:   2,
				runeSizes:   newStack(4, 2, 2),
				err:         nil,
			},
			retractCount:    2,
			expectedForward: 4,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 8,
				forward:     2,
				runeCount:   2,
				runeSizes:   newStack(4, 4, 2),
				err:         nil,
			},
			retractCount:    2,
			expectedForward: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.retractCount; i++ {
				tc.in.Retract()
			}

			assert.Equal(t, tc.expectedForward, tc.in.forward)
		})
	}
}

func TestInput_Peek(t *testing.T) {
	tests := []struct {
		name          string
		in            *Input
		expectedRune  rune
		expectedError string
	}{
		{
			name: "EOF",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         io.EOF,
			},
			expectedRune:  0,
			expectedError: "EOF",
		},
		{
			name: "Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0x69, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedRune:  'i',
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, err := tc.in.Peek()

			if tc.expectedError == "" {
				assert.Equal(t, tc.expectedRune, r)
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.expectedRune, r)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestInput_Lexeme(t *testing.T) {
	tests := []struct {
		name              string
		in                *Input
		expectedLexeme    string
		expectedPos       int
		expectedRuneCount int
	}{
		{
			name: "Empty",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedLexeme:    "",
			expectedPos:       0,
			expectedRuneCount: 0,
		},
		{
			name: "Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0x40, 0x68, 0x65, 0x72, 0x65, 0x20 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     5,
				runeCount:   1,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedLexeme:    "here",
			expectedPos:       1,
			expectedRuneCount: 5,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x40, 0x68, 0x65 /**/, 0x72, 0x65, 0x20, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				runeCount:   4,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedLexeme:    "here",
			expectedPos:       4,
			expectedRuneCount: 8,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x72, 0x65, 0x20, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x40, 0x68, 0x65},
				lexemeBegin: 10,
				forward:     2,
				runeCount:   10,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedLexeme:    "here",
			expectedPos:       10,
			expectedRuneCount: 14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lexeme, pos := tc.in.Lexeme()

			assert.Equal(t, tc.expectedLexeme, lexeme)
			assert.Equal(t, tc.expectedPos, pos)
			assert.Equal(t, tc.expectedRuneCount, tc.in.runeCount)
		})
	}
}

func TestInput_Skip(t *testing.T) {
	tests := []struct {
		name                string
		in                  *Input
		expectedPos         int
		expectedLexemeBegin int
		expectedRuneCount   int
	}{
		{
			name: "Empty",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				runeCount:   0,
				runeSizes:   newStack(4),
				err:         nil,
			},
			expectedPos:         0,
			expectedLexemeBegin: 0,
			expectedRuneCount:   0,
		},
		{
			name: "Success",
			in: &Input{
				src:         nil,
				buff:        []byte{0x40, 0x68, 0x65, 0x72, 0x65, 0x20 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     5,
				runeCount:   1,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedPos:         1,
			expectedLexemeBegin: 5,
			expectedRuneCount:   5,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x40, 0x68, 0x65 /**/, 0x72, 0x65, 0x20, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				runeCount:   4,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedPos:         4,
			expectedLexemeBegin: 8,
			expectedRuneCount:   8,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				src:         nil,
				buff:        []byte{0x72, 0x65, 0x20, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x40, 0x68, 0x65},
				lexemeBegin: 10,
				forward:     2,
				runeCount:   10,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				err:         nil,
			},
			expectedPos:         10,
			expectedLexemeBegin: 2,
			expectedRuneCount:   14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pos := tc.in.Skip()

			assert.Equal(t, tc.expectedPos, pos)
			assert.Equal(t, tc.expectedLexemeBegin, tc.in.lexemeBegin)
			assert.Equal(t, tc.expectedRuneCount, tc.in.runeCount)
		})
	}
}
