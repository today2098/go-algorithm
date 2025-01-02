//go:build go1.18

package algorithm

type Deque[T any] struct {
	front []T
	back  []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		front: make([]T, 0),
		back:  make([]T, 0),
	}
}

func (dq *Deque[T]) Empty() bool {
	return dq.Size() == 0
}

func (dq *Deque[T]) Size() int {
	return len(dq.front) + len(dq.back)
}

func (dq *Deque[T]) Front() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.front) == 0 {
		return dq.back[0]
	}
	return dq.front[len(dq.front)-1]
}

func (dq *Deque[T]) Back() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.back) == 0 {
		return dq.front[0]
	}
	return dq.back[len(dq.back)-1]
}

func (dq *Deque[T]) At(i int) T {
	if !dq.isValidIdx(i) {
		panic("Deque: index out of range")
	}
	if i < len(dq.front) {
		return dq.front[len(dq.front)-1-i]
	}
	return dq.back[i-len(dq.front)]
}

func (dq *Deque[T]) PushFront(x T) {
	dq.front = append(dq.front, x)
}

func (dq *Deque[T]) PushBack(x T) {
	dq.back = append(dq.back, x)
}

func (dq *Deque[T]) PopFront() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.front) == 0 {
		res := dq.back[0]
		dq.back = dq.back[1:]
		return res
	}
	res := dq.front[len(dq.front)-1]
	dq.front = dq.front[:len(dq.front)-1]
	return res
}

func (dq *Deque[T]) PopBack() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.back) == 0 {
		res := dq.front[0]
		dq.front = dq.front[1:]
		return res
	}
	res := dq.back[len(dq.back)-1]
	dq.back = dq.back[:len(dq.back)-1]
	return res
}

func (dq *Deque[T]) isValidIdx(i int) bool {
	return 0 <= i && i < dq.Size()
}
