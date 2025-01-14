package lexer_test

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
)

func ExampleLexer() {
	src := strings.NewReader(`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)

	var s scanner.Scanner
	s.Init(src)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		token := lexer.Token{
			Terminal: grammar.Terminal("WORD"),
			Lexeme:   s.TokenText(),
			Pos: lexer.Position{
				Filename: "lorem_ipsum",
				Offset:   s.Position.Offset,
				Line:     s.Position.Line,
				Column:   s.Position.Column,
			},
		}

		fmt.Println(token)
	}
}
