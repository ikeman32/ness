// Copyright (c) 2023 David H Isakson II
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	COMMENT   = "COMMENT"
	MULTILINE = "MULTILINE"

	// Identifiers and literals
	IDENT  = "IDENT"  // add, foo, x, y ...
	INT    = "INT"    // 1,2,3 ...
	DOUBLE = "DOUBLE" // 1.5067...
	STRING = "STRING" // "This is a string"
	BOOL   = "BOOL"   // true, false

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA = ","
	SEMI  = ";"
	COLON = ":"
	DOT   = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	CREATE   = "CREATE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	AS       = "AS"
	INTEGER  = "INTEGER"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"Create": CREATE,
	"true":   TRUE,
	"false":  FALSE,
	"If":     IF,
	"Else":   ELSE,
	"Return": RETURN,
	"as":     AS,
	"int":    INTEGER,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
