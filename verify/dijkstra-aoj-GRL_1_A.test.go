//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/library/5/GRL/1/GRL_1_A

package main

import (
	"github.com/today2098/go-algorithm/algorithm"
	"github.com/today2098/go-algorithm/utils"
)

func main() {
	io := utils.NewIO()
	defer io.Wr.Flush()

	n, m, r := io.GetInt(), io.GetInt(), io.GetInt()

	dijkstra := algorithm.NewDefaultDijkstra(n)
	for i := 0; i < m; i++ {
		s, t, d := io.GetInt(), io.GetInt(), io.GetInt()
		dijkstra.AddEdge(s, t, d)
	}
	dijkstra.Dijkstra(r)

	for i := 0; i < n; i++ {
		ans := dijkstra.Distance(i)
		if ans == dijkstra.Infinity() {
			io.Out("INF")
		} else {
			io.Out(ans)
		}
	}
}
