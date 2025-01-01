package algorithm

// UnionFind represents disjoint-set data structure.
type UnionFind struct {
	vn  int
	gn  int
	par []int
}

// NewUnionFind returns an initialized UnionFind.
func NewUnionFind(vn int) *UnionFind {
	par := make([]int, vn)
	for i := 0; i < vn; i++ {
		par[i] = -1
	}
	return &UnionFind{
		vn:  vn,
		gn:  vn,
		par: par,
	}
}

// VN returns the number of elements.
func (uf *UnionFind) VN() int {
	return uf.vn
}

// GN returns the number of sets.
func (uf *UnionFind) GN() int {
	return uf.gn
}

// Root returns the root of the set that contains element x.
func (uf *UnionFind) Root(x int) int {
	if !uf.isValidIdx(x) {
		panic("UnionFind: index out of range")
	}
	if uf.par[x] < 0 {
		return x
	}
	// NOTE: recursion
	return uf.Root(uf.par[x])
}

// Size returns the size of the set that contains element x.
func (uf *UnionFind) Size(x int) int {
	if !uf.isValidIdx(x) {
		panic("UnionFind: index out of range")
	}
	return -uf.par[uf.Root(x)]
}

// IsSame reports whether the elements x and y belong to the same set.
func (uf *UnionFind) IsSame(x, y int) bool {
	if !(uf.isValidIdx(x) && uf.isValidIdx(y)) {
		panic("UnionFind: index out of range")
	}
	return uf.Root(x) == uf.Root(y)
}

// Unite merges the set that contains element x and the set that contains element y
func (uf *UnionFind) Unite(x, y int) bool {
	if !(uf.isValidIdx(x) && uf.isValidIdx(y)) {
		panic("UnionFind: index out of range")
	}
	x, y = uf.Root(x), uf.Root(y)
	if x == y {
		return false
	}
	// NOTE: Merge with a technique called by "union by size".
	if uf.par[x] > uf.par[y] {
		x, y = y, x
	}
	uf.par[x], uf.par[y] = uf.par[x]+uf.par[y], x
	uf.gn--
	return true
}

// Reset resets uf.
func (uf *UnionFind) Reset() {
	for i := 0; i < uf.VN(); i++ {
		uf.par[i] = -1
	}
	uf.gn = uf.VN()
}

func (uf *UnionFind) isValidIdx(x int) bool {
	return 0 <= x && x < uf.VN()
}
