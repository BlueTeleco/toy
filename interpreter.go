// interpreter
package color

import (
	"strconv"
)

// Te parser will return an Operator interface.
// This interface provides a method that the
// Interpreter can use to interpret the parsed
// expresion.
type Operator interface {
	Operate() int
}

// Node is one implementation of the Operator
// interface. A Node of a tree structure, that
// contains an Operation to apply to it children.
type Node struct {
	Left      *Node
	Right     *Node
	Operation string
}

// Node Operate implementation.
func (n *Node) Operate() int {
	switch n.Operation {
	case "+":
		return n.Left.Operate() + n.Right.Operate()
	case "-":
		return n.Left.Operate() - n.Right.Operate()
	case "*":
		return n.Left.Operate() * n.Right.Operate()
	case "/":
		return n.Left.Operate() / n.Right.Operate()
	}
	i, _ := strconv.Atoi(n.Operation)
	return i
}
