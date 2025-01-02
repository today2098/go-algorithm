//go:build go1.18

package algorithm

// Stack represents FIFO data structure.
type Queue[T any] struct {
	Data []T
}

// NewQueue returns an initialized Queue.
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

// Front returns the front element of queue.
func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("Queue: empty")
	}
	return q.Data[0]
}

// Push adds the element x to the back of queue.
func (q *Queue[T]) Push(x T) {
	q.Data = append(q.Data, x)
}

// Pop removes and returns the front element of queue.
func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("Queue: empty")
	}
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}
