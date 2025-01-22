// Package input implements a two-buffer input reader.
//
// This package is particularly suited for implementing lexical analyzers and parsers,
// offering robust support for managing lexemes, handling multi-byte UTF-8 characters, and processing character streams.
//
// The two-buffer technique handles large input streams by splitting the buffer into two halves
// that are alternately reloaded, ensuring efficient processing without frequent I/O operations.
//
// For more information and details, see "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package input

import (
	"bytes"
	"fmt"
	"io"

	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/list"
)

const eof byte = 0x00

// Input implements the two-buffer scheme for reading the input characters.
//
// For more information and details, see "Compilers: Principles, Techniques, and Tools (2nd Edition)".
type Input struct {
	filename string
	src      io.Reader

	// The first and second halves of the buff are alternatively reloaded.
	// Each half is of the same size N. Usually, N should be the size of a disk block.
	buff []byte

	lexemeBegin int // Pointer lexemeBegin marks the beginning of the current lexeme.
	forward     int // Pointer forward scans ahead until a pattern match is found.

	offset     int // Tracks the offset (0-based), total number of runes, before lexemeBegin.
	line       int // Tracks the line number (1-based) before lexemeBegin.
	column     int // Tracks the column number (1-based) before lexemeBegin.
	nextColumn int // Tracks the column number (1-based) of the next rune to be read by forward.

	runeSizes   list.Stack[int] // Tracks the size of runes read between lexemeBegin and forward.
	lastColumns list.Stack[int] // Tracks the last column numbers for each line between lexemeBegin and forward.

	err error // Last error encountered.
}

// New creates a new input buffer of size N.
// N usually should be the size of a disk block.
func New(filename string, src io.Reader, n int) (*Input, error) {
	// buff is divided into two sub-buffers (first half and second half).
	buff := make([]byte, 2*n)

	in := &Input{
		filename:    filename,
		src:         src,
		buff:        buff,
		lexemeBegin: 0,
		forward:     0,
		offset:      0,
		line:        1,
		column:      1,
		nextColumn:  1,
		runeSizes:   list.NewStack[int](n, nil),
		lastColumns: list.NewStack[int](n, nil),
	}

	if err := in.loadFirst(); err != nil {
		return nil, err
	}

	return in, nil
}

// loadFirst reads the input and loads the first sub-buffer.
func (i *Input) loadFirst() error {
	high := len(i.buff) / 2

	n, err := i.src.Read(i.buff[:high])
	if err != nil {
		return err
	}

	if n < high {
		i.buff[n] = eof
	}

	return nil
}

// loadSecond reads the input and loads the second sub-buffer.
func (i *Input) loadSecond() error {
	low, high := len(i.buff)/2, len(i.buff)

	n, err := i.src.Read(i.buff[low:high])
	if err != nil {
		return err
	}

	if n < high-low {
		i.buff[low+n] = eof
	}

	return nil
}

// next returns the current byte at the forward pointer and advances the forward pointer to the next byte.
func (i *Input) next() (byte, error) {
	if i.err != nil {
		return 0, i.err
	}

	b := i.buff[i.forward]
	i.forward++

	// Determine whether or not the forward pointer has reached the end of any halves.
	// If so, it loads the other half and set the forward pointer to the beginning of it.
	// If the forward pointer has reached to the end of input, an io.EOF error will be returned.
	if i.forward == len(i.buff)/2 { // Is forward at the end of first half?
		i.err = i.loadSecond()
	} else if i.forward == len(i.buff) { // Is forward at the end of second half?
		if i.err = i.loadFirst(); i.err == nil {
			i.forward = 0 // beginning of the first half
		}
	} else if i.buff[i.forward] == eof {
		i.err = io.EOF
	}

	// The current read is fine, but the next one may return an error
	return b, nil
}

// pos returns the position before lexemeBegin.
func (i *Input) pos() lexer.Position {
	return lexer.Position{
		Filename: i.filename,
		Offset:   i.offset,
		Line:     i.line,
		Column:   i.column,
	}
}

// forwardPos returns the position of forward.
func (i *Input) forwardPos() lexer.Position {
	return lexer.Position{
		Filename: i.filename,
		Offset:   i.offset + i.runeSizes.Size(),
		Line:     i.line + i.lastColumns.Size(),
		Column:   i.nextColumn,
	}
}

// Next advances to the next rune in the input and returns it.
// If the end of the input is reached, it returns the io.EOF error.
func (i *Input) Next() (rune, error) {
	// First byte
	b0, err := i.next()
	if err != nil {
		return 0, err
	}

	x := first[b0]

	if x >= as {
		if x == xx {
			return 0, &InputError{
				Description: "invalid utf-8 character",
				Pos:         i.forwardPos(),
			}
		}

		// Check for new line
		if b0 == '\n' {
			i.lastColumns.Push(i.nextColumn)
			i.nextColumn = 1
		} else {
			i.nextColumn++
		}

		i.runeSizes.Push(1)
		return rune(b0), nil
	}

	size := int(x & 0b0111)

	// Second byte
	b1, err := i.next()
	if err != nil {
		return 0, err
	}

	accept := acceptRanges[x>>4]
	if b1 < accept.lo || accept.hi < b1 {
		return 0, &InputError{
			Description: "invalid utf-8 character",
			Pos:         i.forwardPos(),
		}
	}

	if size == 2 {
		i.runeSizes.Push(size)
		i.nextColumn++
		return rune(b0&mask2)<<6 | rune(b1&maskx), nil
	}

	// Third byte
	b2, err := i.next()
	if err != nil {
		return 0, err
	}

	if b2 < locb || hicb < b2 {
		return 0, &InputError{
			Description: "invalid utf-8 character",
			Pos:         i.forwardPos(),
		}
	}

	if size == 3 {
		i.runeSizes.Push(size)
		i.nextColumn++
		return rune(b0&mask3)<<12 | rune(b1&maskx)<<6 | rune(b2&maskx), nil
	}

	// Fourth byte
	b3, err := i.next()
	if err != nil {
		return 0, err
	}

	if b3 < locb || hicb < b3 {
		return 0, &InputError{
			Description: "invalid utf-8 character",
			Pos:         i.forwardPos(),
		}
	}

	i.runeSizes.Push(size)
	i.nextColumn++
	return rune(b0&mask4)<<18 | rune(b1&maskx)<<12 | rune(b2&maskx)<<6 | rune(b3&maskx), nil
}

// Retract recedes to the last rune in the input.
func (i *Input) Retract() {
	if size, ok := i.runeSizes.Pop(); ok {
		i.forward -= size
		if i.forward < 0 { // adjust the forward pointer if needed
			i.forward += len(i.buff)
		}

		// Check for new line
		if i.buff[i.forward] == '\n' {
			if lastColumn, ok := i.lastColumns.Pop(); ok {
				i.nextColumn = lastColumn
			}
		} else {
			i.nextColumn--
		}
	}
}

// Lexeme returns the current lexeme alongside its position.
func (i *Input) Lexeme() (string, lexer.Position) {
	pos := i.pos()

	var lexeme bytes.Buffer
	for i.lexemeBegin != i.forward {
		lexeme.WriteByte(i.buff[i.lexemeBegin])
		i.lexemeBegin++
		if i.lexemeBegin == len(i.buff) { // Is lexemeBegin at the end of second half?
			i.lexemeBegin = 0 // beginning of the first half
		}
	}

	for !i.runeSizes.IsEmpty() {
		i.runeSizes.Pop()
		i.offset++
	}

	for !i.lastColumns.IsEmpty() {
		i.lastColumns.Pop()
		i.line++
	}

	i.column = i.nextColumn

	return lexeme.String(), pos
}

// Skip skips over the pending lexeme in the input.
func (i *Input) Skip() lexer.Position {
	pos := i.pos()

	i.lexemeBegin = i.forward

	for !i.runeSizes.IsEmpty() {
		i.runeSizes.Pop()
		i.offset++
	}

	for !i.lastColumns.IsEmpty() {
		i.lastColumns.Pop()
		i.line++
	}

	i.column = i.nextColumn

	return pos
}

// InputError represents an error encountered when reading from an input source.
type InputError struct {
	Description string
	Pos         lexer.Position
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *InputError) Error() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s: %s", e.Pos, e.Description)
	return b.String()
}
