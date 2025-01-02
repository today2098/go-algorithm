//go:build go1.18

package algorithm

// Deque represents Double-ended queue data structure.
type Deque[T any] struct {
	FrontStack []T
	BackStack  []T
}

// NewDeque returns an initialized Deque.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		FrontStack: make([]T, 0),
		BackStack:  make([]T, 0),
	}
}

// Empty reports whether the deque is empty.
func (dq *Deque[T]) Empty() bool {
	return dq.Size() == 0
}

// Size returns the number of elements in the deque.
func (dq *Deque[T]) Size() int {
	return len(dq.FrontStack) + len(dq.BackStack)
}

// Front returns the front element of deque.
func (dq *Deque[T]) Front() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.FrontStack) == 0 {
		return dq.BackStack[0]
	}
	return dq.FrontStack[len(dq.FrontStack)-1]
}

// Back returns the back element of deque.
func (dq *Deque[T]) Back() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.BackStack) == 0 {
		return dq.FrontStack[0]
	}
	return dq.BackStack[len(dq.BackStack)-1]
}

// At returns the i th element of deque.
func (dq *Deque[T]) At(i int) T {
	if !dq.isValidIdx(i) {
		panic("Deque: index out of range")
	}
	if i < len(dq.FrontStack) {
		return dq.FrontStack[len(dq.FrontStack)-1-i]
	}
	return dq.BackStack[i-len(dq.FrontStack)]
}

// PushFront adds the element x to the front of deque.
func (dq *Deque[T]) PushFront(x T) {
	dq.FrontStack = append(dq.FrontStack, x)
}

// PushBack adds the element x to the back of deque.
func (dq *Deque[T]) PushBack(x T) {
	dq.BackStack = append(dq.BackStack, x)
}

// PopFront removes and returns the front element of deque.
func (dq *Deque[T]) PopFront() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.FrontStack) == 0 {
		res := dq.BackStack[0]
		dq.BackStack = dq.BackStack[1:]
		return res
	}
	res := dq.FrontStack[len(dq.FrontStack)-1]
	dq.FrontStack = dq.FrontStack[:len(dq.FrontStack)-1]
	return res
}

// PopBack removes and returns the back element of deque.
func (dq *Deque[T]) PopBack() T {
	if dq.Empty() {
		panic("Deque: empty")
	}
	if len(dq.BackStack) == 0 {
		res := dq.FrontStack[0]
		dq.FrontStack = dq.FrontStack[1:]
		return res
	}
	res := dq.BackStack[len(dq.BackStack)-1]
	dq.BackStack = dq.BackStack[:len(dq.BackStack)-1]
	return res
}

func (dq *Deque[T]) isValidIdx(i int) bool {
	return 0 <= i && i < dq.Size()
}
