// lexer
package toy

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

// SimpleLexer is an implementantion of the Lexer
// interface.
type SimpleLexer struct {
	reader   *bufio.Scanner
	text     string
	pos      int
	keyWords map[string]bool
}

// ScanLine scans the next line in the input and
// assigns it to the text string to be tokenized.
func (sl *SimpleLexer) scanLine() error {
	if sl.reader.Scan() {
		sl.text = sl.reader.Text()
		sl.pos = 0
		return nil
	} else {
		if err := sl.reader.Err(); err != nil {
			return err
		} else {
			return io.EOF
		}
	}
}

// Advances the position in the text string.
func (sl *SimpleLexer) advance() error {
	if sl.pos < len(sl.text) {
		sl.pos++
		return nil
	}
	return errors.New("Posicion fuera de rango")
}

// getInt gets a multiple digit int form the text string.
func (sl *SimpleLexer) getInt() string {
	var str string
	for sl.pos != len(sl.text) && unicode.IsDigit(rune(sl.text[sl.pos])) {
		str += string(sl.text[sl.pos])
		if err := sl.advance(); err != nil {
			break
		}

	}
	return str
}

// getWord gets an alphanumeric word. It can represent
// a keyword or not.
func (sl *SimpleLexer) getWord() string {
	var str string
	for sl.pos != len(sl.text) && (unicode.IsDigit(rune(sl.text[sl.pos])) || unicode.IsLetter(rune(sl.text[sl.pos]))) {
		str += string(sl.text[sl.pos])
		if err := sl.advance(); err != nil {
			break
		}

	}
	return str
}

// skipSpaces skips white spaces as defined in the
// unicode package. Includes, but is not limited to,
// spaces and tabs
func (sl *SimpleLexer) skipSpaces() {
	for unicode.IsSpace(rune(sl.text[sl.pos])) {
		if err := sl.advance(); err != nil {
			break
		}
	}
}

// Lex returns the next Token in the text string.
// Posible Types: "EOF", "INT", "LPAR", "RPAR", "OPR",
// "KW", "NAME", and "ASGN". If a line starts with '#'
// it skips it (it is a commnent line)
func (sl *SimpleLexer) Lex() Token {
	if sl.pos == len(sl.text) {
		err := sl.scanLine()
		switch err {
		case nil:
			return sl.Lex()
		case io.EOF:
			return Token{"EOF", "EOF"}
		default:
			panic("syntax error")
		}
	}

	switch c := rune(sl.text[sl.pos]); {
	case unicode.IsSpace(c):
		sl.skipSpaces()
		return sl.Lex()
	case unicode.IsDigit(c):
		return Token{"INT", sl.getInt()}
	case unicode.IsLetter(c):
		w := sl.getWord()
		if sl.keyWords[w] {
			return Token{"KW", w}
		}
		return Token{"NAME", w}
	case c == '#' && sl.pos == 0:
		sl.scanLine()
		return sl.Lex()
	case c == '(':
		sl.advance()
		return Token{"LPAR", string(c)}
	case c == ')':
		sl.advance()
		return Token{"RPAR", string(c)}
	case c == '{':
		sl.advance()
		return Token{"LCOR", string(c)}
	case c == '}':
		sl.advance()
		return Token{"RCOR", string(c)}
	case c == '=':
		sl.advance()
		return Token{"ASGN", string(c)}
	case c == ';':
		sl.advance()
		return Token{"DCOMA", string(c)}
	default:
		sl.advance()
		return Token{"OPR", string(c)}
	}
}

// NewSimpleLexer constructs a new SimpleLexer struct.
func NewSimpleLexer(sc *bufio.Scanner) *SimpleLexer {
	m := make(map[string]bool)
	m["if"] = true
	return &SimpleLexer{sc, "", 0, m}
}
