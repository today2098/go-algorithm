//go:build go1.18

package algorithm

import "math"

type dijkstraPair[T any] struct {
	first  T
	second int
}

// Dijkstra is a helper structure that solves shortest path problem by Dijkstra algorithm.
type Dijkstra[T int | int64 | float64] struct {
	vn  int
	g   [][]dijkstraPair[T] // pair of (cost, to)
	d   []T
	pre []int
	inf T
}

// NewDijkstra returns an initialized Dijkstra.
func NewDijkstra[T int | int64 | float64](vn int, inf T) *Dijkstra[T] {
	d := make([]T, vn)
	pre := make([]int, vn)
	for v := 0; v < vn; v++ {
		d[v] = inf
		pre[v] = -1
	}
	return &Dijkstra[T]{
		vn:  vn,
		g:   make([][]dijkstraPair[T], vn),
		d:   d,
		pre: pre,
		inf: inf,
	}
}

// NewDefaultDijkstra returns an initialized Dijkstra.
func NewDefaultDijkstra(vn int) *Dijkstra[int] {
	return NewDijkstra(vn, math.MaxInt)
}

// Infinity returns the inifinity value.
func (d *Dijkstra[T]) Infinity() T {
	return d.inf
}

// Order returns the number of nodes in the graph.
func (d *Dijkstra[T]) Order() int {
	return d.vn
}

// AddEdge adds a new edge that connects the two nodes with cost.
func (d *Dijkstra[T]) AddEdge(from, to int, cost T) {
	if !(d.isValidIdx(from) && d.isValidIdx(to)) {
		panic("Dijkstra: index out of range")
	}
	d.g[from] = append(d.g[from], dijkstraPair[T]{
		first:  cost,
		second: to,
	})
}

// Dijkstra calculates the shortest distances between s and each nodes.
func (d *Dijkstra[T]) Dijkstra(s int) {
	if !d.isValidIdx(s) {
		panic("Dijkstra: index out of range")
	}
	for v := 0; v < d.Order(); v++ {
		d.d[v] = d.Infinity()
		d.pre[v] = -1
	}
	d.d[s] = 0
	pq := NewBinaryHeap(func(a, b dijkstraPair[T]) bool {
		return a.first < b.first
	})
	pq.Push(dijkstraPair[T]{
		first:  0, // distance
		second: s, // node
	})
	for !pq.Empty() {
		dist, from := pq.Top().first, pq.Top().second
		pq.Pop()
		if d.d[from] < dist {
			continue
		}
		for i := range d.g[from] {
			cost, to := d.g[from][i].first, d.g[from][i].second
			if d.d[to] > d.d[from]+cost {
				d.d[to] = d.d[from] + cost
				d.pre[to] = from
				pq.Push(dijkstraPair[T]{
					first:  d.d[to],
					second: to,
				})
			}
		}
	}
}

// Distance returns the shortest distance between s and t.
func (d *Dijkstra[T]) Distance(t int) T {
	if !d.isValidIdx(t) {
		panic("Dijkstra: index out of range")
	}
	return d.d[t]
}

// ShortestPath returns the shortest path between s and t.
func (d *Dijkstra[T]) ShortestPath(t int) []int {
	if !d.isValidIdx(t) {
		panic("Dijkstra: index out of range")
	}
	path := make([]int, 0)
	if d.Distance(t) == d.Infinity() {
		return path
	}
	for t != -1 {
		path = append(path, t)
		t = d.pre[t]
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}
	return path
}

func (d *Dijkstra[T]) isValidIdx(v int) bool {
	return 0 <= v && v < d.Order()
}
