package main

import (
	"fmt"
	"strings"
)

type Queue[T any] []T

func (s *Queue[T]) Push(value T) {
	(*s) = append((*s), value)
}

func (s *Queue[T]) Pop() T {
	item := (*s)[0]
	(*s) = (*s)[1 : len(*s)-1]
	return item
}

func (s *Queue[T]) Peek() *T {
	return &(*s)[0]
}

func (s *Queue[T]) String() string {
	buf := strings.Builder{}

	for i, entry := range *s {
		if i > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprint(&buf, entry)
	}

	return buf.String()
}
