// interpreter
package toy

import (
	"strconv"
)

// Te parser will return an Interpreter interface.
// This interface provides a method that the
// Interpreter can use to interpret the parsed
// expresion.
type Interpreter interface {
	Interprete() int
}

// OprNode is one implementation of the Interpreter
// interface. A node of a tree structure, that
// contains an Operation to apply to it children.
type OprNode struct {
	Left      Interpreter
	Right     Interpreter
	Operation string
}

// OprNode Interprete implementation.
func (on *OprNode) Interprete() int {
	switch on.Operation {
	case "+", "&":
		return on.Left.Interprete() + on.Right.Interprete()
	case "-":
		return on.Left.Interprete() - on.Right.Interprete()
	case "*", "|":
		return on.Left.Interprete() * on.Right.Interprete()
	case "/":
		return on.Left.Interprete() / on.Right.Interprete()
	case ">":
		if on.Left.Interprete() > on.Right.Interprete() {
			return 1
		}
		return 0
	}
	i, _ := strconv.Atoi(on.Operation)
	return i
}

// Variables contains the map where the Variables are stored.
var Variables map[string]int

// VarNode is one implementation of the Interpreter
// interface. A node of a tree structure, that assigns
// or retrieves a variable's value.
type VarNode struct {
	Value Interpreter
	Name  string
}

// VarNode Interprete implementation.
func (vn *VarNode) Interprete() int {
	if vn.Value != nil {
		Variables[vn.Name] = vn.Value.Interprete()
		return 0
	}
	return Variables[vn.Name]
}

// BlockNode is one implementation of the Interpreter
// interface. A node of a tree structure, that contains
// multiple sons.
type BlockNode struct {
	Sons []Interpreter
}

// BlockNode Interprete implementation.
func (bn *BlockNode) Interprete() int {
	for _, son := range bn.Sons {
		son.Interprete()
	}
	return 0
}

// IfNode is one implementation of the Interpreter
// interface. A node of a tree structure that implements
// an if control structure.
type IfNode struct {
	Expresion Interpreter
	Block     Interpreter
}

// IfNode Interprete implementation.
func (in *IfNode) Interprete() int {
	if in.Expresion.Interprete() > 0 {
		return in.Block.Interprete()
	}
	return 1
}
