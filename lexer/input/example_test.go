package input_test

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/lexer/input"
)

func ExampleInput_Next() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New(bufferSize, src)
	if err != nil {
		panic(err)
	}

	// Reading the next rune.
	r, err := in.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)
}

func ExampleInput_Retract() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New(bufferSize, src)
	if err != nil {
		panic(err)
	}

	r, err := in.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)

	r, err = in.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)

	// Undoing the last rune read.
	in.Retract()

	r, err = in.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)
}

func ExampleInput_Peek() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New(bufferSize, src)
	if err != nil {
		panic(err)
	}

	r, err := in.Peek()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)

	r, err = in.Next()
	if err != nil {
		panic(err)
	}

	fmt.Printf("rune: %c\n", r)
}

func ExampleInput_Lexeme() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New(bufferSize, src)
	if err != nil {
		panic(err)
	}

	for range 5 {
		r, err := in.Next()
		if err != nil {
			panic(err)
		}

		fmt.Printf("rune: %c\n", r)
	}

	// Reading the current lexeme.
	lexeme, pos := in.Lexeme()
	fmt.Printf("lexeme: %q  position: %d\n", lexeme, pos)
}

func ExampleInput_Skip() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New(bufferSize, src)
	if err != nil {
		panic(err)
	}

	for range 6 {
		if _, err = in.Next(); err != nil {
			panic(err)
		}
	}

	// Skiping the current lexeme.
	pos := in.Skip()
	fmt.Printf("position of skipped lexeme: %d\n", pos)

	for range 5 {
		if _, err = in.Next(); err != nil {
			panic(err)
		}
	}

	// Reading the next lexeme.
	lexeme, pos := in.Lexeme()
	fmt.Printf("lexeme: %q  position: %d\n", lexeme, pos)
}
