package algorithm

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type naiveUnionFind struct {
	vn   int
	gn   int
	flag [][]bool
}

func newNaiveUnionFind(vn int) *naiveUnionFind {
	flag := make([][]bool, vn)
	for i := range flag {
		flag[i] = make([]bool, vn)
	}
	for u := 0; u < vn; u++ {
		flag[u][u] = true
	}
	return &naiveUnionFind{
		vn:   vn,
		gn:   vn,
		flag: flag,
	}
}

func (uf *naiveUnionFind) VN() int {
	return uf.vn
}

func (uf *naiveUnionFind) GN() int {
	return uf.gn
}

func (uf *naiveUnionFind) Size(x int) int {
	res := 0
	for _, elem := range uf.flag[x] {
		if elem {
			res++
		}
	}
	return res
}

func (uf *naiveUnionFind) IsSame(x, y int) bool {
	return uf.flag[x][y]
}

func (uf *naiveUnionFind) Unite(x, y int) bool {
	if uf.flag[x][y] {
		return false
	}
	uf.flag[x][y], uf.flag[y][x] = true, true
	for u := 0; u < uf.VN(); u++ {
		if !uf.flag[u][x] {
			continue
		}
		for v := 0; v < uf.VN(); v++ {
			if !uf.flag[y][v] {
				continue
			}
			uf.flag[u][y], uf.flag[y][u] = true, true
			uf.flag[x][v], uf.flag[v][x] = true, true
			uf.flag[u][v], uf.flag[v][u] = true, true
		}
	}
	uf.gn--
	return true
}

func (uf *naiveUnionFind) Reset() {
	for u := 0; u < uf.VN(); u++ {
		for v := u + 1; v < uf.VN(); v++ {
			uf.flag[u][v], uf.flag[v][u] = false, false
		}
	}
	uf.gn = uf.VN()
}

func TestUnionFind(t *testing.T) {
	vn := 100
	nuf, uf := newNaiveUnionFind(vn), NewUnionFind(vn)

	check := func() {
		assert.Equal(t, nuf.VN(), uf.VN())
		assert.Equal(t, nuf.GN(), uf.GN())

		for u := 0; u < vn; u++ {
			assert.Equal(t, nuf.Size(u), uf.Size(u))

			for v := 0; v < vn; v++ {
				assert.Equal(t, nuf.IsSame(u, v), uf.IsSame(u, v))
			}
		}
	}

	check()

	loop := 100
	for i := 0; i < loop; i++ {
		u, v := rand.Intn(vn), rand.Intn(vn)
		assert.Equal(t, nuf.Unite(u, v), uf.Unite(u, v))
		check()
	}

	nuf.Reset()
	uf.Reset()
	check()
}

func BenchmarkUnionFind(b *testing.B) {
	vn := int(1e7)
	uf := NewUnionFind(vn)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		uf.Unite(rand.Intn(vn), rand.Intn(vn))
	}

	for i := 0; i < b.N; i++ {
		uf.IsSame(rand.Intn(vn), rand.Intn(vn))
	}

	b.ReportMetric(float64(b.N)/float64(b.Elapsed().Milliseconds()), "op/ms")
}
