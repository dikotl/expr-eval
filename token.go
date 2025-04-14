package main

//go:generate go tool stringer -type=TokenKind -output=token_string.go -linecomment
type TokenKind byte

const (
	InvalidToken TokenKind = iota
	EOF
	Number
	Variable
	Plus     // +
	Minus    // -
	Asterisk // *
	Slash    // /
	Percent  // %
	Caret    // ^
	LeParen  // (
	RiParen  // )
	Tilde    // ~
)

type TokenSpan struct{ A, B int }

type Token struct {
	Kind TokenKind
	Span TokenSpan
	Data string
}

func (t Token) String() string {
	switch t.Kind {
	case Number, Variable:
		return t.Data

	case EOF, Plus, Minus, Asterisk, Slash, Percent, Caret, LeParen, RiParen, Tilde:
		return t.Kind.String()

	default:
		panic("unreachable")
	}
}
