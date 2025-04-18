package main

import "fmt"

type AstWalker interface {
	WalkNumber(node *NumberNode)
	WalkVariable(node *VariableNode)
	WalkBinaryOperation(node *BinaryOperation)
	WalkUnaryOperation(node *UnaryOperation)
}

func WalkAst(ast AstNode, walker AstWalker) {
	switch node := ast.(type) {
	case nil:
		panic("WalkAst: got nil node")

	case *NumberNode:
		walker.WalkNumber(node)

	case *VariableNode:
		walker.WalkVariable(node)

	case *BinaryOperation:
		walker.WalkBinaryOperation(node)

	case *UnaryOperation:
		walker.WalkUnaryOperation(node)

	default:
		panic(fmt.Sprintf("unexpected ast node: %#v", node))
	}
}
