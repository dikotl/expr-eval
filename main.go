package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Input any math expression, then press ← or → to move " +
		"through the expression. Ctrl+C to exit.")

	stdin := bufio.NewReader(os.Stdin)
	expr := readExpr(stdin)

	p := NewParser(expr)
	ast := p.Expr()

	// fmt.Print("Normalized expression: ")
	// WalkAst(ast, AstPrinter{})
	// fmt.Println()

	polish := ReversePolishBuilder{Stack: NewStack[Token]()}
	WalkAst(ast, &polish)

	inputStr := polish.Stack.String()
	fmt.Printf("Parsed expression in reverse polish notation:\n%s\n", inputStr)
	t := Tokenizer{}
	t.input.Reset(inputStr)
	newInput := make([]Token, 0, len(polish.items))
	for range len(polish.items) {
		newInput = append(newInput, t.NextToken())
	}
	slices.Reverse(newInput)

	first := true
	items := newInput
	itemIndex := len(items) - 1

	setRawMode(true)
	defer setRawMode(false)
	// defer runInterruptNotifyMonitor()()

	buf := make([]byte, 3)

	for {
		trimmedInput := NewStackFrom(items[itemIndex:])
		stack, err := Eval(trimmedInput)
		Display(items[itemIndex].Span, stack, err, first)
		first = false

		n, err := os.Stdin.Read(buf)

		if err != nil {
			panic(err)
		}

		switch {
		case n == 3 && buf[0] == 27 && buf[1] == 91: // Esc
			switch buf[2] {
			case 'C':
				itemIndex = max(itemIndex-1, 0)
			case 'D':
				itemIndex = min(itemIndex+1, len(items)-1)
			default:
				// Ignore other escape sequences
			}

		case n == 1 && buf[0] == 3: // SIGINT
			return
		}
	}
}

func readExpr(r *bufio.Reader) string {
	for {
		s, err := r.ReadString('\n')

		if err == nil {
			return strings.TrimSpace(s)
		}

		fmt.Printf("error: %s\n", err)
	}
}
