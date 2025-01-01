//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/1/DSL_1_A

package main

import (
	"github.com/today2098/go-algorithm/algorithm"
	"github.com/today2098/go-algorithm/utils"
)

func main() {
	io := utils.NewIO()
	defer io.Wr.Flush()

	n, q := io.GetInt(), io.GetInt()

	uf := algorithm.NewUnionFind(n)
	for i := 0; i < q; i++ {
		com, x, y := io.GetInt(), io.GetInt(), io.GetInt()

		if com == 0 {
			uf.Unite(x, y)
		} else {
			if uf.IsSame(x, y) {
				io.Out(1)
			} else {
				io.Out(0)
			}
		}
	}
}
