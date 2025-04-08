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

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (_ T, ok bool) {
	if len(s.items) == 0 {
		return
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) MustPop() T {
	item, ok := s.Pop()
	if !ok {
		panic("can't pop empty stack")
	}
	return item
}

func (s Stack[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}

func (s Stack[T]) String() string {
	buf := strings.Builder{}

	for i := len(s.items) - 1; i >= 0; i-- {
		if i != len(s.items)-1 {
			buf.WriteByte(' ')
		}
		fmt.Fprint(&buf, s.items[i])
	}

	return buf.String()
}

func (s Stack[T]) Items() []T {
	return s.items
}
