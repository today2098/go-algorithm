//go:build go1.18

package algorithm

// Stack represents FIFO data structure.
type Queue[T any] struct {
	Data []T
}

// NewStack returns an initialized Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Data: []T{},
	}
}

// Empty reports whether the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.Data)
}

// Top returns the top element of the queue.
func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("Queue: empty")
	}
	return q.Data[0]
}

// Push adds the element x to the top of the queue.
func (q *Queue[T]) Push(x T) {
	q.Data = append(q.Data, x)
}

// Pop removes and returns the top element of the queue.
func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("Queue: empty")
	}
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
