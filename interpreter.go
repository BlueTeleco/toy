// interpreter
package color

// Te parser will return an Interpreter interface.
// This interface provides a method that the
// Interpreter can use to interpret the parsed
// expresion.
type Interpreter interface {
	Interprete() int
}

// OprNode is one implementation of the Interpreter
// interface. A OprNode of a tree structure, that
// contains an Operation to apply to it children.
type OprNode struct {
	Left      Interpreter
	Right     Interpreter
	Operation string
}

// OprNode Interprete implementation.
func (n *OprNode) Interprete() int {
	switch n.Operation {
	case "+":
		return n.Left.Interprete() + n.Right.Interprete()
	case "-":
		return n.Left.Interprete() - n.Right.Interprete()
	case "*":
		return n.Left.Interprete() * n.Right.Interprete()
	case "/":
		return n.Left.Interprete() / n.Right.Interprete()
	}
	i, _ := strconv.Atoi(n.Operation)
	return i
}
