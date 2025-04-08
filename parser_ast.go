package main

type (
	AstNode interface {
		Span() TokenSpan
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

func (n *NumberNode) Span() TokenSpan {
	return n.TokenSpan
}

func (n *VariableNode) Span() TokenSpan {
	return n.TokenSpan
}

func (n *BinaryOperation) Span() TokenSpan {
	return TokenSpan{
		A: n.X.Span().A,
		B: n.Y.Span().B,
	}
}

func (n *UnaryOperation) Span() TokenSpan {
	return TokenSpan{
		A: n.A,
		B: n.X.Span().B,
	}
}
