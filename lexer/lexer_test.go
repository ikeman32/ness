// Copyright (c) 2023 David H Isakson II
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package lexer

import (
	"testing"

	"github.com/ikeman32/ness/token"
)

func TestNextToken(t *testing.T) {
	input := `Create myInt as int; = 5.`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CREATE, "Create"},
		{token.IDENT, "myInt"},
		{token.AS, "as"},
		{token.INTEGER, "int"},
		{token.SEMI, ";"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
	}

	l := new(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// func NewLexer(input string) {
// 	panic("unimplemented")
// }
