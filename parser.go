package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var precedences = map[TokenKind]int{
	Plus:     2,
	Minus:    2,
	Asterisk: 3,
	Slash:    3,
	Percent:  3,
	Caret:    4,
}

type Parser struct {
	Tokenizer
	Token
	depth int
}

func NewParser(expr string) Parser {
	p := Parser{}
	p.input.Reset(expr)
	p.next()
	return p
}

func (p *Parser) Expr() AstNode {
	return p.binary(1)
}

func (p *Parser) binary(precedence int) AstNode {
	x := p.prefixed()

	for {
		prec, present := precedences[p.Kind]

		if !present || prec < precedence {
			break
		}

		operator := p.next()

		x = &BinaryOperation{
			X:         x,
			Y:         p.binary(prec + 1),
			Kind:      operator.Kind,
			TokenSpan: operator.Span,
		}
	}

	return x
}

func (p *Parser) prefixed() AstNode {
	switch p.Kind {
	case Plus, Minus:
		operator := p.next()

		return &UnaryOperation{
			TokenSpan: operator.Span,
			Kind:      operator.Kind,
			X:         p.prefixed(),
		}

	default:
		return p.operand()
	}
}

func (p *Parser) operand() AstNode {
	switch p.Kind {
	case Variable:
		tok := p.next()

		return &VariableNode{
			TokenSpan: tok.Span,
			Name:      tok.Data,
		}

	case Number:
		tok := p.next()
		value, _ := strconv.Atoi(tok.Data)

		return &NumberNode{
			TokenSpan: tok.Span,
			Value:     value,
		}

	case LeParen:
		p.next()
		p.depth++

		defer func() {
			if p.Kind != RiParen {
				p.error("expected closing parenthesis")
			}

			p.next()
			p.depth--
		}()

		return p.Expr()

	default:
		p.error("expected operand")
		panic("unreachable")
	}
}

func (p *Parser) next() (previous Token) {
	previous = p.Token

	if p.Kind != EOF {
		p.Token = p.NextToken()
	}

	return
}

func (p *Parser) error(message string) {
	fmt.Printf("%s^ %s\n", strings.Repeat(" ", p.Span.A), message)
	os.Exit(2)
}
