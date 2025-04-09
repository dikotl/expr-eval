package main

import (
	"fmt"
	"slices"
	"strings"
)

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{items: make([]T, 0)}
}

func NewQueueFrom[T any](items []T) Queue[T] {
	return Queue[T]{items: items}
}
func (queue *Queue[T]) Push(item T) {
	queue.items = slices.Insert(queue.items, 0, item)
}

func (queue *Queue[T]) Pop() (item T, present bool) {
	if len(queue.items) == 0 {
		return
	}
	item = queue.items[0]
	queue.items = queue.items[1 : len(queue.items)-1]
	return item, true
}

func (queue *Queue[T]) MustPop() T {
	item, ok := queue.Pop()
	if !ok {
		panic("can't pop empty queue")
	}
	return item
}

func (queue *Queue[T]) Peek() *T {
	if len(queue.items) == 0 {
		return nil
	}
	return &queue.items[0]
}

func (queue *Queue[T]) Len() int {
	return len(queue.items)
}

func (queue *Queue[T]) String() string {
	buf := strings.Builder{}
	for i, entry := range queue.items {
		if i > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprint(&buf, entry)
	}
	return buf.String()
}
