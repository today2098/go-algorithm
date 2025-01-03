//go:build go1.18

package algorithm

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	n, m := 100, 300

	dijkstra := NewDefaultDijkstra(n)
	assert.Equal(t, n, dijkstra.Order())

	g := make([][]int, n)
	for u := 0; u < n; u++ {
		g[u] = make([]int, n)
		for v := 0; v < n; v++ {
			if u == v {
				g[u][v] = 0
			} else {
				g[u][v] = dijkstra.Infinity()
			}
		}
	}

	for i := 0; i < m; i++ {
		u, v := rand.Intn(n), rand.Intn(n)
		cost := rand.Intn(1e4)

		dijkstra.AddEdge(u, v, cost)
		g[u][v] = min(g[u][v], cost)
	}

	for k := 0; k < n; k++ {
		for u := 0; u < n; u++ {
			if g[u][k] == dijkstra.Infinity() {
				continue
			}
			for v := 0; v < n; v++ {
				if g[k][v] == dijkstra.Infinity() {
					continue
				}
				g[u][v] = min(g[u][v], g[u][k]+g[k][v])
			}
		}
	}

	for s := 0; s < n; s++ {
		dijkstra.Dijkstra(s)

		for v := 0; v < n; v++ {
			want, got := g[s][v], dijkstra.Distance(v)
			assert.Equal(t, want, got)

			path := dijkstra.ShortestPath(v)
			t.Log(s, v, path)
			if want == dijkstra.Infinity() {
				assert.Len(t, path, 0)
			} else {
				sum := 0
				pre := s
				for i := range path {
					sum += g[pre][path[i]]
					pre = path[i]
				}

				assert.Equal(t, want, sum)
			}
		}
	}
}
