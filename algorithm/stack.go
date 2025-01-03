//go:build go1.18

package algorithm

// Stack represents LIFO data structure.
type Stack[T any] struct {
	Data []T
}

// NewStack returns an initialized Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Data: []T{},
	}
}

// Empty reports whether the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.Size() == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.Data)
}

// Top returns the top element of the stack.
func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("Stack: empty")
	}
	return s.Data[len(s.Data)-1]
}

// Push adds the element x to the top of the stack.
func (s *Stack[T]) Push(x T) {
	s.Data = append(s.Data, x)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("Stack: empty")
	}
	res := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return res
}
