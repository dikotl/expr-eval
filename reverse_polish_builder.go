package main

import (
	"strconv"
)

type ReversePolishBuilder struct {
	stack Stack[Token]
}

func AstToReversePolishNotation(ast AstNode, reTokenize bool) Stack[Token] {
	// Build sequence of actions from the AST using reverse polish notation.
	builder := ReversePolishBuilder{
		stack: NewStack[Token](), // We don't know how deep the AST is.
	}

	WalkAst(ast, &builder)
	actions := builder.stack

	// Re-tokenize the sequence of actions so token spans will point to
	// the newly generated actions.
	if reTokenize {
		// Setup tokenizer.
		t := Tokenizer{}
		t.input.Reset(actions.String())

		// Preallocate the buffer.
		reTokenizedActions := make([]Token, builder.stack.Len())

		// Push items in reversed order.
		for i := builder.stack.Len() - 1; i >= 0; i-- {
			reTokenizedActions[i] = t.NextToken()
		}

		// Replace old actions.
		actions = NewStackFrom(reTokenizedActions)
	}

	return actions
}

func (builder *ReversePolishBuilder) WalkNumber(node *NumberNode) {
	builder.stack.Push(Token{
		Kind: Number,
		Span: node.TokenSpan,
		Data: strconv.Itoa(node.Value),
	})
}

func (builder *ReversePolishBuilder) WalkVariable(node *VariableNode) {
	builder.stack.Push(Token{
		Kind: Variable,
		Data: node.Name,
		Span: node.TokenSpan,
	})
}

func (builder *ReversePolishBuilder) WalkBinaryOperation(node *BinaryOperation) {
	builder.stack.Push(Token{
		Kind: node.Kind,
		Span: node.TokenSpan,
	})

	WalkAst(node.X, builder)
	WalkAst(node.Y, builder)
}

func (builder *ReversePolishBuilder) WalkUnaryOperation(node *UnaryOperation) {
	builder.stack.Push(Token{
		Kind: Tilde,
		Span: node.TokenSpan,
	})

	WalkAst(node.X, builder)
}
