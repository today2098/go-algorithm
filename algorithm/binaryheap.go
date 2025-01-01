//go:build go1.18

package algorithm

// BinaryHeapCompFunc represents the function used by BinaryHeap to compare two elements.
type BinaryHeapCompFunc[T any] func(a, b T) bool

// BinaryHeap represents priority-queue data structure.
type BinaryHeap[T any] struct {
	comp BinaryHeapCompFunc[T]
	tree []T
}

// NewBinaryHeap returns an initialized BinaryHeap.
func NewBinaryHeap[T any](f BinaryHeapCompFunc[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		comp: f,
		tree: make([]T, 1),
	}
}

// NewDefaultBinaryHeap returns an initialized BinaryHeap.
func NewDefaultBinaryHeap[T int | float64 | byte | rune | string]() *BinaryHeap[T] {
	return NewBinaryHeap(func(a, b T) bool {
		return a > b
	})
}

// Empty reports whether the priority-queue is empty.
func (h *BinaryHeap[T]) Empty() bool {
	return h.Size() == 0
}

// Size returns the number of elements in the priority-queue.
func (h *BinaryHeap[T]) Size() int {
	return len(h.tree) - 1
}

// Top returns the top element of the priority-queue.
func (h *BinaryHeap[T]) Top() T {
	if h.Empty() {
		panic("BinaryHeap: empty")
	}
	return h.tree[1]
}

// Push adds the element x to the top of the priority-queue.
func (h *BinaryHeap[T]) Push(x T) {
	h.tree = append(h.tree, x)
	h.shiftUp(h.Size())
}

// Pop removes and returns the top element of the priority-queue.
func (h *BinaryHeap[T]) Pop() T {
	if h.Empty() {
		panic("BinaryHeap: empty")
	}
	res := h.tree[1]
	h.tree[1] = h.tree[h.Size()]
	h.tree = h.tree[:len(h.tree)-1]
	if !h.Empty() {
		h.shiftDown(1)
	}
	return res
}

func (h *BinaryHeap[T]) shiftUp(i int) {
	p := i >> 1
	for 1 <= p {
		if h.comp(h.tree[p], h.tree[i]) {
			break
		}
		h.tree[p], h.tree[i] = h.tree[i], h.tree[p]
		i = p
		p >>= 1
	}
}

func (h *BinaryHeap[T]) shiftDown(i int) {
	l, r := i<<1, i<<1|1
	for l <= h.Size() {
		if h.Size() < r || h.comp(h.tree[l], h.tree[r]) {
			if h.comp(h.tree[i], h.tree[l]) {
				break
			}
			h.tree[i], h.tree[l] = h.tree[l], h.tree[i]
			i = l
		} else {
			if h.comp(h.tree[i], h.tree[r]) {
				break
			}
			h.tree[i], h.tree[r] = h.tree[r], h.tree[i]
			i = r
		}
		l, r = i<<1, i<<1|1
	}
}
