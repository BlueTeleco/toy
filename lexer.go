package color

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

type Lexer interface {
	Lex() Token
}

type Token struct {
	Type  string
	Value string
}

type SimpleLexer struct {
	Reader *bufio.Scanner
	Text   string
	Pos    int
}

func (sl *SimpleLexer) ScanLine() error {
	sl.Reader.Scan()
	sl.Text = sl.Reader.Text()

	if err := sl.Reader.Err(); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (sl *SimpleLexer) Advance() error {
	if sl.Pos < len(sl.Text) {
		sl.Pos++
		return nil
	}
	return errors.New("Posicion fuera de rango")
}

func (sl *SimpleLexer) GetInt() string {
	var str string
	for unicode.IsDigit(rune(sl.Text[sl.Pos])) {
		str += string(sl.Text[sl.Pos])
		if err := sl.Advance(); err != nil {
			break
		}

	}
	return str
}

func (sl *SimpleLexer) SkipSpaces() {
	for unicode.IsSpace(rune(sl.Text[sl.Pos])) {
		if err := sl.Advance(); err != nil {
			break
		}
	}
}

func (sl *SimpleLexer) Lex() Token {
	if sl.Pos == len(sl.Text) {
		return Token{"EOF", "EOF"}
	}

	if c := rune(sl.Text[sl.Pos]); unicode.IsSpace(c) {
		sl.SkipSpaces()
		return sl.Lex()
	} else if unicode.IsDigit(c) {
		return Token{"INT", sl.GetInt()}
	} else if c == '(' {
		sl.Advance()
		return Token{"LPAR", string(c)}
	} else if c == ')' {
		sl.Advance()
		return Token{"RPAR", string(c)}
	} else {
		sl.Advance()
		return Token{"OPR", string(c)}
	}
}
