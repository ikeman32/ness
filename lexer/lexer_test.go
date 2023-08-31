// Copyright (c) 2023 David H Isakson II
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package lexer

import (
	"testing"

	"github.com/ikeman32/ness/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `Create myInt as int; = 5.
	Create ten as int; = 10. //Anothe comment
	Create ans as int.
	ans = myInt + ten.
	/*This is a multiline comment*/
	`
	// TODO: Fix EOF token so it is expected at the end.
	// skip comments after and before code lines
	expectedTokens := []token.Token{
		{Type: token.EOF, Literal: ""},
		{Type: token.CREATE, Literal: "Create"},
		{Type: token.IDENT, Literal: "myInt"},
		{Type: token.AS, Literal: "as"},
		{Type: token.INTEGER, Literal: "int"},
		{Type: token.SEMI, Literal: ";"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.DOT, Literal: "."},
		{Type: token.CREATE, Literal: "Create"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.AS, Literal: "as"},
		{Type: token.INTEGER, Literal: "int"},
		{Type: token.SEMI, Literal: ";"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.DOT, Literal: "."},
		{Type: token.COMMENT, Literal: "//"},
		{Type: token.IDENT, Literal: "Anothe"},
		{Type: token.IDENT, Literal: "comment"},
		{Type: token.CREATE, Literal: "Create"},
		{Type: token.IDENT, Literal: "ans"},
		{Type: token.AS, Literal: "as"},
		{Type: token.INTEGER, Literal: "int"},
		{Type: token.DOT, Literal: "."},
		{Type: token.IDENT, Literal: "ans"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENT, Literal: "myInt"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.DOT, Literal: "."},
		// {Type: token.EOF, Literal: ""},
	}

	l := Lexer{input: input}
	for i, expected := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expected.Type {
			t.Errorf("Test %d: Expected token type %s, got %s", i, expected.Type, tok.Type)
		}

		if tok.Literal != expected.Literal {
			t.Errorf("Test %d: Expected token literal %s, got %s", i, expected.Literal, tok.Literal)
		}
	}
}
