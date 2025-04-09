package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
	"unicode"
)

var operators = map[rune]TokenKind{
	'+': Plus,
	'-': Minus,
	'*': Asterisk,
	'/': Slash,
	'%': Percent,
	'^': Caret,
	'(': LeParen,
	')': RiParen,
}

const eof = '\000'

type Tokenizer struct {
	input  strings.Reader
	peeked rune
	index  int
	size   int
}

func (t *Tokenizer) NextToken() (token Token) {
	t.SkipWhile(isWhitespace)

	token.Span.A = t.index
	peeked := t.Peek()

	switch {
	case peeked == eof:
		token.Kind = EOF

	case unicode.IsDigit(peeked):
		token.Kind = Number
		token.Data = t.TakeWhile(unicode.IsDigit)

	case unicode.IsLetter(peeked):
		token.Kind = Variable
		token.Data = t.TakeWhile(unicode.IsLetter)

	case isOperator(peeked):
		token.Kind = operators[peeked]
		t.Advance()

	default:
		fmt.Printf("%s^ illegal character\n", strings.Repeat(" ", t.index))
		os.Exit(2)
	}

	token.Span.B = t.index
	return
}

func (t *Tokenizer) Next() (peeked rune) {
	char, size, err := t.input.ReadRune()

	if err != nil {
		if err == io.EOF {
			if t.peeked != eof {
				t.peeked = eof
				t.index += 1
				t.size = 0
			}
			return eof
		}
		// Unreachable because [strings.Reader.ReadRune] can't return other error.
		panic(err)
	}

	t.peeked = char
	t.index += t.size
	t.size = size
	return char
}

func (s *Tokenizer) Advance() (previous rune) {
	previous = s.peeked
	s.Next()
	return previous
}

func (t *Tokenizer) Peek() (peeked rune) {
	if t.peeked == eof {
		// Likely first read. Or last.
		t.Next()
	}

	return t.peeked
}

// Takes all characters while predicate returns true.
func (t *Tokenizer) TakeWhile(predicate func(rune) bool) string {
	buf := strings.Builder{}
	for char := range t.Chars() {
		if !predicate(char) {
			break
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Skips all characters while predicate returns true.
func (t *Tokenizer) SkipWhile(predicate func(rune) bool) (skipped int) {
	for char := range t.Chars() {
		if !predicate(char) {
			break
		}
		skipped++
	}
	return
}

func (t *Tokenizer) Chars() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for {
			if char := t.Peek(); !yield(char) || char == eof {
				return
			}
			t.Advance()
		}
	}
}

func isOperator(c rune) bool {
	_, ok := operators[c]
	return ok
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t'
}
