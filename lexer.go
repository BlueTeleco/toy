// lexer
package color

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

// Lexer interface to build an Interpreter.
type Lexer interface {
	Lex() Token
}

// The Lexer will give a secuence of Tokens,
// a siple Type-Value pair to represent different
// posible tokens.
// Eg. Token{Type:"INT", Value:"3"}
type Token struct {
	Type  string
	Value string
}

// Simple implementantion of the Lexer interface.
type SimpleLexer struct {
	Reader *bufio.Scanner
	Text   string
	Pos    int
}

// Scans the next line in the input and assigns it to
// the Text string to be tokenized.
func (sl *SimpleLexer) ScanLine() error {
	sl.Reader.Scan()
	sl.Text = sl.Reader.Text()

	if err := sl.Reader.Err(); err != nil && err != io.EOF {
		return err
	}
	return nil
}

// Advances the position in the Text string.
func (sl *SimpleLexer) advance() error {
	if sl.Pos < len(sl.Text) {
		sl.Pos++
		return nil
	}
	return errors.New("Posicion fuera de rango")
}

// Gets a multiple digit int form the Text string.
func (sl *SimpleLexer) getInt() string {
	var str string
	for sl.Pos != len(sl.Text) && unicode.IsDigit(rune(sl.Text[sl.Pos])) {
		str += string(sl.Text[sl.Pos])
		if err := sl.advance(); err != nil {
			break
		}

	}
	return str
}

// Skips white spaces as defined in the unicode package.
// Includes, but is not limited to, spaces and tabs
func (sl *SimpleLexer) skipSpaces() {
	for unicode.IsSpace(rune(sl.Text[sl.Pos])) {
		if err := sl.advance(); err != nil {
			break
		}
	}
}

// Returns the next Token in the Text string.
// Posible Types: "EOF", "INT", "LPAR", "RPAR" and "OPR"
func (sl *SimpleLexer) Lex() Token {
	if sl.Pos == len(sl.Text) {
		return Token{"EOF", "EOF"}
	}

	if c := rune(sl.Text[sl.Pos]); unicode.IsSpace(c) {
		sl.skipSpaces()
		return sl.Lex()
	} else if unicode.IsDigit(c) {
		return Token{"INT", sl.getInt()}
	} else if c == '(' {
		sl.advance()
		return Token{"LPAR", string(c)}
	} else if c == ')' {
		sl.advance()
		return Token{"RPAR", string(c)}
	} else {
		sl.advance()
		return Token{"OPR", string(c)}
	}
}
