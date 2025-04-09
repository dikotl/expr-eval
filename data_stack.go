package main

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{items: make([]T, 0)}
}

func NewStackFrom[T any](items []T) Stack[T] {
	return Stack[T]{items: items}
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) Pop() (_ T, ok bool) {
	if len(stack.items) == 0 {
		return
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return item, true
}

func (stack *Stack[T]) MustPop() T {
	item, ok := stack.Pop()
	if !ok {
		panic("can't pop empty stack")
	}
	return item
}

func (stack Stack[T]) Peek() *T {
	if len(stack.items) == 0 {
		return nil
	}
	return &stack.items[len(stack.items)-1]
}

func (stack Stack[T]) String() string {
	buf := strings.Builder{}
	for i := len(stack.items) - 1; i >= 0; i-- {
		if i != len(stack.items)-1 {
			buf.WriteByte(' ')
		}
		fmt.Fprint(&buf, stack.items[i])
	}
	return buf.String()
}
