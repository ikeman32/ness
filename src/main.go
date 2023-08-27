// Copyright (c) 2023 duke
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/ikeman32/ness/lexer"
)

func main() {

	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}

	lexer := lexer.NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	}
}
