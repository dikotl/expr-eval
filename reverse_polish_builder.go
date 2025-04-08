package main

import "strconv"

type ReversePolishBuilder struct {
	Stack[Token]
}

func (b *ReversePolishBuilder) WalkNumber(node *NumberNode) {
	b.Push(Token{
		Kind: Number,
		Span: node.TokenSpan,
		Data: strconv.Itoa(node.Value),
	})
}

func (b *ReversePolishBuilder) WalkVariable(node *VariableNode) {
	b.Push(Token{
		Kind: Variable,
		Data: node.Name,
		Span: node.TokenSpan,
	})
}

func (b *ReversePolishBuilder) WalkBinaryOperation(node *BinaryOperation) {
	b.Push(Token{
		Kind: node.Kind,
		Span: node.TokenSpan,
	})

	WalkAst(node.X, b)
	WalkAst(node.Y, b)
}

func (b *ReversePolishBuilder) WalkUnaryOperation(node *UnaryOperation) {
	b.Push(Token{
		Kind: node.Kind,
		Span: node.TokenSpan,
	})

	WalkAst(node.X, b)
}
