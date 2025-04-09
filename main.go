package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const Prompt = "Input simple math expression, then press ← or → to move " +
	"through the expression. Ctrl+C to exit."

// CLI flag.
var (
	origFlag = flag.Bool(
		"orig",
		false,
		"Uses original expression in step-through process",
	)

	normalizeFlag = flag.Bool(
		"normalize",
		false,
		"Normalizes the expression, print it and quit",
	)
)

func main() {
	flag.Parse()
	fmt.Println(Prompt)

	// Read the expression.
	stdin := bufio.NewReader(os.Stdin)
	expr := readExpr(stdin)

	// Parse it into an AST.
	p := NewParser(expr)
	ast := p.Expr()

	if *normalizeFlag {
		WalkAst(ast, AstPrinter{})
		fmt.Println()
		return
	}

	actions := AstToReversePolishNotation(ast, !*origFlag)
	actionIndex := len(actions.items) - 1
	first := true

	if !*origFlag {
		fmt.Println(actions.String())
	}

	// Set terminal into raw mode so we can handle arrow keys gracefully.
	setRawMode(true)
	defer setRawMode(false)
	// defer runInterruptNotifyMonitor()() // Ctrl+C hook.

	var buf [3]byte

	for {
		// Cut the input so we evaluate items before the pointer.
		trimmedInput := NewStackFrom(actions.items[actionIndex:])
		stack, err := Eval(trimmedInput)
		Display(actions.items[actionIndex].Span, stack, err, first)
		first = false

		// Read the input, we need exactly 3 bytes.
		n, err := os.Stdin.Read(buf[:])
		if err != nil {
			panic(err)
		}

		// Test the input for left\right arrow or Ctrl+C
		switch {
		case n == 3 && buf[0] == 27 && buf[1] == 91: // Esc
			switch buf[2] {
			case 'C': // left arrow
				actionIndex = max(actionIndex-1, 0)
			case 'D': // right arrow
				actionIndex = min(actionIndex+1, actions.Len()-1)
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
