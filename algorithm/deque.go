//go:build go1.18

package algorithm

import "container/list"

type Deque[T any] struct {
	Data *list.List
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		Data: list.New(),
	}
}

func (dq *Deque[T]) Empty() bool {
	return dq.Size() == 0
}

func (dq *Deque[T]) Size() int {
	return dq.Data.Len()
}

func (dq *Deque[T]) Front() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	return dq.Data.Front().Value.(T)
}

func (dq *Deque[T]) Back() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	return dq.Data.Back().Value.(T)
}

func (dq *Deque[T]) PushFront(x T) {
	dq.Data.PushFront(x)
}

func (dq *Deque[T]) PushBack(x T) {
	dq.Data.PushBack(x)
}

func (dq *Deque[T]) PopFront() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	return dq.Data.Remove(dq.Data.Front()).(T)
}

func (dq *Deque[T]) PopBack() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	return dq.Data.Remove(dq.Data.Back()).(T)
}
