package main

import (
	"fmt"
)

type AstPrinter struct{}

func (AstPrinter) WalkNumber(node *NumberNode) {
	fmt.Printf("%d", node.Value)
}

func (AstPrinter) WalkVariable(node *VariableNode) {
	fmt.Print(node.Name)
}

func (AstPrinter) WalkBinaryOperation(node *BinaryOperation) {
	WalkAst(node.X, AstPrinter{})
	fmt.Printf(" %s ", node.Kind)
	WalkAst(node.Y, AstPrinter{})
}

func (AstPrinter) WalkUnaryOperation(node *UnaryOperation) {
	fmt.Printf("%s ", node.Kind)
	WalkAst(node.X, AstPrinter{})
}
