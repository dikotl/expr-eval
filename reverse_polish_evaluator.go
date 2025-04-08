package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type EvalState struct {
// 	input       []Token
// 	sp          int
// 	stack       Stack[int] //
// 	span        TokenSpan  // Top of the stack
// 	inputLength int
// }

// func NewEvalState(input Stack[Token]) EvalState {
// 	s := EvalState{inputLength: -1}
// 	return s
// }

func Eval(input Stack[Token]) (result Stack[int]) {
	result = NewStackFrom(make([]int, 0, len(input.Items())))

	for {
		item, ok := input.Pop()

		if !ok {
			return
		}

		switch item.Kind {
		case Number:
			value, _ := strconv.Atoi(item.Data)
			result.Push(value)

		case Variable:
			panic("unimplemented")

		case Plus:
			x := result.MustPop()
			y := result.MustPop()
			result.Push(x + y)

		case Minus:
			x := result.MustPop()
			y := result.MustPop()
			result.Push(x - y)

		case Asterisk:
			x := result.MustPop()
			y := result.MustPop()
			result.Push(x * y)

		case Slash:
			x := result.MustPop()
			y := result.MustPop()
			result.Push(x / y)

		case Percent:
			x := result.MustPop()
			y := result.MustPop()
			result.Push(x % y)

		default:
			panic("unimplemented")
		}
	}
}

func Display(span TokenSpan, stack Stack[int], inputLength int, first bool) {
	if !first {
		// Go up
		for range 3 {
			fmt.Print("\033[1A\033[K")
		}
	}

	fmt.Printf(
		"%s^%s%s\nStack\n%s\n",
		strings.Repeat(" ", span.A),
		strings.Repeat("~", span.B-span.A-1),
		strings.Repeat(" ", inputLength-span.B),
		stack.String(),
	)
}
