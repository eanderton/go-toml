package toml

import (
	"fmt"
	"strconv"
	"unicode"
)

// Define tokens
type tokenType int

const (
	eof = -(iota + 1)
)

const (
	tokenError tokenType = iota
	tokenEOF
	tokenComment
	tokenKey
	tokenString
	tokenInteger
	tokenTrue
	tokenFalse
	tokenFloat
	tokenEqual
	tokenLeftBracket
	tokenRightBracket
	tokenLeftParen
	tokenRightParen
	tokenDoubleLeftBracket
	tokenDoubleRightBracket
	tokenDate
	tokenKeyGroup
	tokenKeyGroupArray
	tokenComma
	tokenColon
	tokenDollar
	tokenStar
	tokenQuestion
	tokenDot
	tokenDotDot
	tokenEOL
)

var tokenTypeNames = []string{
	"EOF",
	"Comment",
	"Key",
	"String",
	"Integer",
	"True",
	"False",
	"Float",
	"=",
	"[",
	"[",
	"(",
	")",
	"]]",
	"[[",
	"Date",
	"KeyGroup",
	"KeyGroupArray",
	",",
	":",
	"$",
	"*",
	"?",
	".",
	"..",
	"EOL",
}

type token struct {
	Position
	typ tokenType
	val string
}

func (tt tokenType) String() string {
	idx := int(tt)
	if idx < len(tokenTypeNames) {
		return tokenTypeNames[idx]
	}
	return "Unknown"
}

func (t token) Int() int {
	if result, err := strconv.Atoi(t.val); err != nil {
		panic(err)
	} else {
		return result
	}
}

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}

	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isKeyChar(r rune) bool {
	// "Keys start with the first non-whitespace character and end with the last
	// non-whitespace character before the equals sign."
	return !(isSpace(r) || r == '\r' || r == '\n' || r == eof || r == '=')
}

func isDigit(r rune) bool {
	return unicode.IsNumber(r)
}

func isHexDigit(r rune) bool {
	return isDigit(r) ||
		r == 'A' || r == 'B' || r == 'C' || r == 'D' || r == 'E' || r == 'F'
}
