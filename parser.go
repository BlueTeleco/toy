// parser
package color

// Parser interface to build an Interpreter.
type Parser interface {
	Parse() Operator
}

// SimpleParser is one simple implementation of the
// Parser interface.
type SimpleParser struct {
	Lex       Lexer
	CurrToken Token
}

// eat consumes a Token of Type tokenType. If there is
// a sintax error it panics.
func (sp *SimpleParser) eat(tokenType string) {
	if sp.CurrToken.Type == tokenType {
		sp.CurrToken = sp.Lex.Lex()
	} else {
		panic("syntax error")
	}
}

// factor implements the factor rule:
//
// factor: INT | LPAR expr RPAR
//
func (sp *SimpleParser) factor() *Node {
	if t := sp.CurrToken; t.Type == "INT" {
		sp.eat("INT")
		return &Node{nil, nil, t.Value}
	} else {
		sp.eat("LPAR")
		node := sp.expr()
		sp.eat("RPAR")
		return node
	}
}

// term implements the term rule:
//
// term: factor((MUL|DIV) factor)*
//
func (sp *SimpleParser) term() *Node {
	node := sp.factor()
	for value := sp.CurrToken.Value; value == "*" || value == "/"; value = sp.CurrToken.Value {
		sp.eat("OPR")
		node = &Node{node, sp.factor(), value}
	}
	return node
}

// expr implements the expr rule:
//
// expr: term((SUM|SUBS) term)*
//
func (sp *SimpleParser) expr() *Node {
	node := sp.term()
	for value := sp.CurrToken.Value; value == "+" || value == "-"; value = sp.CurrToken.Value {
		sp.eat("OPR")
		node = &Node{node, sp.term(), value}
	}
	return node
}

// Parse parses the expresion into a tree.
// Return the root of the tree as an
// Operator
func (sp *SimpleParser) Parse() Operator {
	return sp.expr()
}
