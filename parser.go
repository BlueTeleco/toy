// parser
package toy

// Parser interface to build an Interpreter.
type Parser interface {
	Parse() Interpreter
}

// SimpleParser is one simple implementation of the
// Parser interface.
type SimpleParser struct {
	lex       Lexer
	currToken Token
}

// eat consumes a Token of Type tokenType. If there is
// a sintax error it panics.
func (sp *SimpleParser) eat(tokenType string) {
	if sp.currToken.Type == tokenType {
		sp.currToken = sp.lex.Lex()
	} else {
		panic("syntax error")
	}
}

// factor implements the factor rule:
//
// factor: INT | NAME | LPAR expr RPAR
//
func (sp *SimpleParser) factor() Interpreter {
	if t := sp.currToken; t.Type == "INT" {
		sp.eat("INT")
		return &OprNode{nil, nil, t.Value}
	} else if t.Type == "NAME" {
		sp.eat("NAME")
		return &VarNode{nil, t.Value}
	} else {
		sp.eat("LPAR")
		node := sp.expr()
		sp.eat("RPAR")
		return node
	}
}

// term implements the term rule:
//
// term: factor((MUL|DIV|OR) factor)*
//
func (sp *SimpleParser) term() Interpreter {
	node := sp.factor()
	for value := sp.currToken.Value; value == "*" || value == "/" || value == "|"; value = sp.currToken.Value {
		sp.eat("OPR")
		node = &OprNode{node, sp.factor(), value}
	}
	return node
}

// expr implements the expr rule:
//
// expr: term((SUM|SUBS|AND) term)*
//
func (sp *SimpleParser) expr() Interpreter {
	node := sp.term()
	for value := sp.currToken.Value; value == "+" || value == "-" || value == "&"; value = sp.currToken.Value {
		sp.eat("OPR")
		node = &OprNode{node, sp.term(), value}
	}
	return node
}

// line implements the line rule:
//
// line: NAME ASGN expr
//
func (sp *SimpleParser) line() Interpreter {
	n := sp.currToken.Value
	sp.eat("NAME")
	sp.eat("ASGN")
	node := &VarNode{sp.expr(), n}
	sp.eat("DCOMA")
	return node
}

// ifblock implements the ifblock rule:
//
// ifblock: KW LPAR expr RPAR LCOR block RCOR
//
func (sp *SimpleParser) ifblock() Interpreter {
	sp.eat("KW")
	sp.eat("LPAR")
	enode := sp.expr()
	sp.eat("LCOR")
	bnode := sp.block()
	sp.eat("RCOR")
	return &IfNode{enode, bnode}
}

// block implements the block rule:
//
// block: (line | ifblock)*
//
func (sp *SimpleParser) block() Interpreter {
	lines := make([]Interpreter, 0, 20)
	for sp.currToken.Type != "EOF" {
		n := len(lines)
		if len(lines) == cap(lines) {
			temp := make([]Interpreter, n, n+20)
			copy(temp, lines)
			lines = temp
		}
		lines = lines[:n+1]
		if sp.currToken.Value == "if" {
			lines[n] = sp.ifblock()
		} else {
			lines[n] = sp.line()
		}
	}
	return &BlockNode{lines}
}

// Parse parses the expresion into a tree.
// Return the root of the tree as an
// Interpreter
func (sp *SimpleParser) Parse() Interpreter {
	return sp.block()
}

// NewSimpleParser constructs a new SimpleParser struct.
func NewSimpleParser(l Lexer) *SimpleParser {
	Variables = make(map[string]int)
	return &SimpleParser{l, l.Lex()}
}
