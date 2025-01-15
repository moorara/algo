package input_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/moorara/algo/lexer/input"
)

func ExampleInput() {
	bufferSize := 4096

	file, err := os.Open("./fixture/lorem_ipsum")
	if err != nil {
		panic(err)
	}

	in, err := input.New("lorem_ipsum", file, bufferSize)
	if err != nil {
		panic(err)
	}

	// advanceDFA simulates a determinist finite automata for identifying words.
	advanceDFA := func(state int, r rune) int {
		switch state {
		case 0:
			switch r {
			case ' ', ',', '.', '\n':
				return 0
			default:
				return 1
			}

		case 1:
			switch r {
			case ' ', ',', '.', '\n':
				return 3
			default:
				return 2
			}

		case 2:
			switch r {
			case ' ', ',', '.', '\n':
				return 3
			default:
				return 2
			}

		case 3:
			switch r {
			case ' ', ',', '.', '\n':
				return 0
			default:
				return 1
			}

		default:
			return -1
		}
	}

	// Reads runes from the input and feeds them into the DFA.
	var state int
	var r rune
	for r, err = in.Next(); err == nil; r, err = in.Next() {
		state = advanceDFA(state, r)
		switch state {
		case 1:
			in.Retract()
			in.Skip()
		case 3:
			in.Retract()
			lexeme, pos := in.Lexeme()
			fmt.Printf("Lexeme %q at %s\n", lexeme, pos)
		}
	}

	fmt.Println(err)
}

func ExampleInput_Next() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New("lorem_ipsum", src, bufferSize)
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

	in, err := input.New("lorem_ipsum", src, bufferSize)
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

func ExampleInput_Lexeme() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New("lorem_ipsum", src, bufferSize)
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
	fmt.Printf("lexeme: %q  position: %s\n", lexeme, pos)
}

func ExampleInput_Skip() {
	bufferSize := 4096
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	in, err := input.New("lorem_ipsum", src, bufferSize)
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
	fmt.Printf("position of skipped lexeme: %s\n", pos)

	for range 5 {
		if _, err = in.Next(); err != nil {
			panic(err)
		}
	}

	// Reading the next lexeme.
	lexeme, pos := in.Lexeme()
	fmt.Printf("lexeme: %q  position: %s\n", lexeme, pos)
}
