package main

type (
	AstNode interface {
		node()
	}

	NumberNode struct {
		TokenSpan
		Value int
	}

	VariableNode struct {
		TokenSpan
		Name string
	}

	BinaryOperation struct {
		TokenSpan
		Kind TokenKind
		X    AstNode
		Y    AstNode
	}

	UnaryOperation struct {
		TokenSpan
		Kind TokenKind
		X    AstNode
	}
)

func (node *NumberNode) node()      {}
func (node *VariableNode) node()    {}
func (node *BinaryOperation) node() {}
func (node *UnaryOperation) node()  {}
