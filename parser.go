package main

import (
	"fmt"
	"strconv"
	"strings"
)

var precedences = map[TokenKind]int{
	EOF:      -1,
	Plus:     2,
	Minus:    2,
	Asterisk: 3,
	Slash:    3,
	Percent:  3,
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
	return p.binary(0)
}

func (p *Parser) binary(precedence int) AstNode {
	x := p.prefixed()

	for precedences[p.Kind] >= precedence {
		operator := p.next()

		x = &BinaryOperation{
			X:         x,
			Y:         p.binary(precedences[operator.Kind] + 1),
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
				fmt.Printf(
					"%s^ expected closing parenthesis\n",
					strings.Repeat(" ", p.Span.A),
				)
				// os.Exit(2)
				panic("unreachable")
			}

			p.next()
			p.depth--
		}()

		return p.Expr()

	default:
		fmt.Printf(
			"%s^ expected operand\n",
			strings.Repeat(" ", p.Span.A),
		)
		// os.Exit(2)
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
