package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrNotEnoughItems = errors.New("not enough items in the stack")

var binaryOperations = map[TokenKind]func(x, y int) int{
	Plus:     func(x, y int) int { return x + y },
	Minus:    func(x, y int) int { return x - y },
	Asterisk: func(x, y int) int { return x * y },
	Slash: func(x, y int) int {
		if y == 0 {
			return 0
		}
		return x / y
	},
	Percent: func(x, y int) int {
		if y == 0 {
			return 0
		}
		return x % y
	},
	Caret: func(base, exp int) int {
		result := 1
		for exp != 0 {
			if exp%2 != 0 {
				result *= base
			}
			exp /= 2
			base *= base
		}
		return result
	},
}

func Eval(input Stack[Token]) (result Stack[int], _ error) {
	buffer := make([]int, 0, len(input.items))
	result = NewStackFrom(buffer)

	for item, ok := input.Pop(); ok; item, ok = input.Pop() {
		// TODO: negate operation
		switch item.Kind {
		case Number:
			value, _ := strconv.Atoi(item.Data)
			result.Push(value)

		case Variable:
			panic("Eval: variables are not implemented")

		case Tilde:
			x, ok := result.Pop()

			if !ok {
				return result, ErrNotEnoughItems
			}

			result.Push(-x)

		case Plus, Minus, Asterisk, Slash, Percent, Caret:
			x, y, err := pop2Items(&result)

			if err != nil {
				return result, err
			}

			result.Push(binaryOperations[item.Kind](x, y))

		default:
			panic("unreachable")
		}
	}

	return
}

func Display(span TokenSpan, stack Stack[int], err error, first bool) {
	const outputFormat = "%s^%s%s\n\rEvaluator Stack:\n\r%s\n\r"
	const outputLineCount = 3 // count or \n in format string

	if !first {
		for range outputLineCount {
			// Go up and clear the line.
			//
			//	\033[1A // Up by 1 line.
			//	\003[K  // Clear line to the end.
			fmt.Print("\033[1A\033[K")
		}
	}

	errStr := ""

	if err != nil {
		errStr = fmt.Sprintf(" error: %s", err.Error())
	}

	fmt.Printf(
		outputFormat,
		strings.Repeat(" ", span.A),          // Spaces before.
		strings.Repeat("~", span.B-span.A-1), // Underscore if span length > 1.
		errStr,                               // An error occurred in eval.
		stack.String(),                       // The stack of values.
	)
}

func pop2Items(stack *Stack[int]) (x, y int, err error) {
	var ok bool

	x, ok = stack.Pop()

	if !ok {
		err = ErrNotEnoughItems
		return
	}

	y, ok = stack.Pop()

	if !ok {
		err = ErrNotEnoughItems
		return
	}

	return x, y, nil
}
