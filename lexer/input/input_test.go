package input

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/lexer"
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
		filename      string
		src           io.Reader
		n             int
		expectedError string
	}{
		{
			name:          "Success",
			filename:      "lorem_ipsum",
			src:           strings.NewReader("Lorem ipsum"),
			n:             4096,
			expectedError: "",
		},
		{
			name:          "Failure",
			filename:      "lorem_ipsum",
			src:           iotest.ErrReader(errors.New("io error")),
			n:             4096,
			expectedError: "io error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in, err := New(tc.filename, tc.src, tc.n)

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

			defer func() {
				if err := f.Close(); err != nil {
					assert.Failf(t, "error on closing %s: %s", f.Name(), err)
				}
			}()

			in, err := New(tc.file, f, tc.n)
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
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         io.EOF,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FirstByte_Invalid",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x80, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "test:1:1: invalid utf-8 character",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FirstByte_Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x69, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'i',
			expectedSize:  1,
		},
		{
			name: "SecondByte_EOF",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xC6, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "SecondByte_Invalid",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xC6, 0x40, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "test:1:1: invalid utf-8 character",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "SecondByte_Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xC6, 0xA9, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'Ʃ',
			expectedSize:  2,
		},
		{
			name: "ThirdByte_EOF",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "ThirdByte_Invalid",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x40, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "test:1:1: invalid utf-8 character",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "ThirdByte_Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xEA, 0xA9, 0x80, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  'ꩀ',
			expectedSize:  3,
		},
		{
			name: "FourthByte_EOF",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "EOF",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FourthByte_Invalid",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x40, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "test:1:1: invalid utf-8 character",
			expectedRune:  0,
			expectedSize:  0,
		},
		{
			name: "FourthByte_Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0xF0, 0x90, 0x80, 0x80, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  '𐀀',
			expectedSize:  4,
		},
		{
			name: "Newline_Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x4C, 0x6F, 0x72, 0x65, 0x6D, 0x0A /*newline*/, 0x69, 0x70, 0x73, 0x75, 0x6D},
				lexemeBegin: 5,
				forward:     5,
				offset:      5,
				line:        1,
				column:      6,
				nextColumn:  6,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedError: "",
			expectedRune:  '\n',
			expectedSize:  1,
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
		name               string
		in                 *Input
		retractCount       int
		expectedForward    int
		expectedNextColumn int
	}{
		{
			name: "Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     4,
				offset:      1,
				line:        1,
				column:      2,
				nextColumn:  4,
				runeSizes:   newStack(4, 1, 2),
				lastColumns: newStack(4),
				err:         nil,
			},
			retractCount:       2,
			expectedForward:    1,
			expectedNextColumn: 2,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				offset:      2,
				line:        1,
				column:      3,
				nextColumn:  5,
				runeSizes:   newStack(4, 2, 2),
				lastColumns: newStack(4),
				err:         nil,
			},
			retractCount:       2,
			expectedForward:    4,
			expectedNextColumn: 3,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 8,
				forward:     2,
				offset:      2,
				line:        1,
				column:      3,
				nextColumn:  5,
				runeSizes:   newStack(4, 4, 2),
				lastColumns: newStack(4),
				err:         nil,
			},
			retractCount:       2,
			expectedForward:    8,
			expectedNextColumn: 3,
		},
		{
			name: "Success_Newline",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x4C, 0x6F, 0x72, 0x65, 0x6D, 0x0A /*newline*/, 0x69, 0x70, 0x73, 0x75, 0x6D},
				lexemeBegin: 5,
				forward:     7,
				offset:      5,
				line:        1,
				column:      6,
				nextColumn:  2,
				runeSizes:   newStack(4, 1, 1),
				lastColumns: newStack(4, 6),
				err:         nil,
			},
			retractCount:       2,
			expectedForward:    5,
			expectedNextColumn: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for range tc.retractCount {
				tc.in.Retract()
			}

			assert.Equal(t, tc.expectedForward, tc.in.forward)
			assert.Equal(t, tc.expectedNextColumn, tc.in.nextColumn)
		})
	}
}

func TestInput_Lexeme(t *testing.T) {
	tests := []struct {
		name           string
		in             *Input
		expectedLexeme string
		expectedPos    lexer.Position
		expectedOffset int
		expectedLine   int
		expectedColumn int
	}{
		{
			name: "Empty",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedLexeme: "",
			expectedPos:    lexer.Position{Filename: "test", Offset: 0, Line: 1, Column: 1},
			expectedOffset: 0,
			expectedLine:   1,
			expectedColumn: 1,
		},
		{
			name: "Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x40, 0x68, 0x65, 0x72, 0x65, 0x20 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     5,
				offset:      1,
				line:        1,
				column:      2,
				nextColumn:  6,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedLexeme: "here",
			expectedPos:    lexer.Position{Filename: "test", Offset: 1, Line: 1, Column: 2},
			expectedOffset: 5,
			expectedLine:   1,
			expectedColumn: 6,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x40, 0x68, 0x65 /**/, 0x72, 0x65, 0x20, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				offset:      4,
				line:        1,
				column:      5,
				nextColumn:  9,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedLexeme: "here",
			expectedPos:    lexer.Position{Filename: "test", Offset: 4, Line: 1, Column: 5},
			expectedOffset: 8,
			expectedLine:   1,
			expectedColumn: 9,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x72, 0x65, 0x20, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x40, 0x68, 0x65},
				lexemeBegin: 10,
				forward:     2,
				offset:      10,
				line:        1,
				column:      11,
				nextColumn:  15,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedLexeme: "here",
			expectedPos:    lexer.Position{Filename: "test", Offset: 10, Line: 1, Column: 11},
			expectedOffset: 14,
			expectedLine:   1,
			expectedColumn: 15,
		},
		{
			name: "Success_Newline",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x4C, 0x6F, 0x72, 0x65, 0x6D, 0x0A /*newline*/, 0x69, 0x70, 0x73, 0x75, 0x6D, 0x20 /*space*/, 0x64, 0x6F, 0x6C, 0x6F, 0x72},
				lexemeBegin: 0,
				forward:     11,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  12,
				runeSizes:   newStack(4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
				lastColumns: newStack(4, 6),
				err:         nil,
			},
			expectedLexeme: "Lorem\nipsum",
			expectedPos:    lexer.Position{Filename: "test", Offset: 0, Line: 1, Column: 1},
			expectedOffset: 11,
			expectedLine:   2,
			expectedColumn: 12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lexeme, pos := tc.in.Lexeme()

			assert.Equal(t, tc.expectedLexeme, lexeme)
			assert.Equal(t, tc.expectedPos, pos)
			assert.Equal(t, tc.expectedOffset, tc.in.offset)
			assert.Equal(t, tc.expectedLine, tc.in.line)
			assert.Equal(t, tc.expectedColumn, tc.in.column)
		})
	}
}

func TestInput_Skip(t *testing.T) {
	tests := []struct {
		name                string
		in                  *Input
		expectedPos         lexer.Position
		expectedLexemeBegin int
		expectedOffset      int
		expectedLine        int
		expectedColumn      int
	}{
		{
			name: "Empty",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 0,
				forward:     0,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  1,
				runeSizes:   newStack(4),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedPos:         lexer.Position{Filename: "test", Offset: 0, Line: 1, Column: 1},
			expectedLexemeBegin: 0,
			expectedOffset:      0,
			expectedLine:        1,
			expectedColumn:      1,
		},
		{
			name: "Success",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x40, 0x68, 0x65, 0x72, 0x65, 0x20 /**/, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				lexemeBegin: 1,
				forward:     5,
				offset:      1,
				line:        1,
				column:      2,
				nextColumn:  6,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedPos:         lexer.Position{Filename: "test", Offset: 1, Line: 1, Column: 2},
			expectedLexemeBegin: 5,
			expectedOffset:      5,
			expectedLine:        1,
			expectedColumn:      6,
		},
		{
			name: "Success_FirstHalfToSecondHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x00, 0x00, 0x00, 0x40, 0x68, 0x65 /**/, 0x72, 0x65, 0x20, 0x00, 0x00, 0x00},
				lexemeBegin: 4,
				forward:     8,
				offset:      4,
				line:        1,
				column:      5,
				nextColumn:  9,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedPos:         lexer.Position{Filename: "test", Offset: 4, Line: 1, Column: 5},
			expectedLexemeBegin: 8,
			expectedOffset:      8,
			expectedLine:        1,
			expectedColumn:      9,
		},
		{
			name: "Success_SecondHalfToFirstHalf",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x72, 0x65, 0x20, 0x00, 0x00, 0x00 /**/, 0x00, 0x00, 0x00, 0x40, 0x68, 0x65},
				lexemeBegin: 10,
				forward:     2,
				offset:      10,
				line:        1,
				column:      11,
				nextColumn:  15,
				runeSizes:   newStack(4, 1, 1, 1, 1),
				lastColumns: newStack(4),
				err:         nil,
			},
			expectedPos:         lexer.Position{Filename: "test", Offset: 10, Line: 1, Column: 11},
			expectedLexemeBegin: 2,
			expectedOffset:      14,
			expectedLine:        1,
			expectedColumn:      15,
		},
		{
			name: "Success_Newline",
			in: &Input{
				filename:    "test",
				src:         nil,
				buff:        []byte{0x4C, 0x6F, 0x72, 0x65, 0x6D, 0x0A /*newline*/, 0x69, 0x70, 0x73, 0x75, 0x6D, 0x20 /*space*/, 0x64, 0x6F, 0x6C, 0x6F, 0x72},
				lexemeBegin: 0,
				forward:     11,
				offset:      0,
				line:        1,
				column:      1,
				nextColumn:  12,
				runeSizes:   newStack(4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
				lastColumns: newStack(4, 6),
				err:         nil,
			},
			expectedPos:         lexer.Position{Filename: "test", Offset: 0, Line: 1, Column: 1},
			expectedLexemeBegin: 11,
			expectedOffset:      11,
			expectedLine:        2,
			expectedColumn:      12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pos := tc.in.Skip()

			assert.Equal(t, tc.expectedPos, pos)
			assert.Equal(t, tc.expectedLexemeBegin, tc.in.lexemeBegin)
			assert.Equal(t, tc.expectedOffset, tc.in.offset)
			assert.Equal(t, tc.expectedLine, tc.in.line)
			assert.Equal(t, tc.expectedColumn, tc.in.column)
		})
	}
}

func TestInputError(t *testing.T) {
	tests := []struct {
		name          string
		e             *InputError
		expectedError string
	}{
		{
			name: "WithoutLineAndColumn",
			e: &InputError{
				Description: "invalid utf-8 rune",
				Pos: lexer.Position{
					Filename: "test_file",
					Offset:   69,
				},
			},
			expectedError: "test_file:69: invalid utf-8 rune",
		},
		{
			name: "WithLineAndColumn",
			e: &InputError{
				Description: "invalid utf-8 rune",
				Pos: lexer.Position{
					Filename: "test_file",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test_file:8:27: invalid utf-8 rune",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}
