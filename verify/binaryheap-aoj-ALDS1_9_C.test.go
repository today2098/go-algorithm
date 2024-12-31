//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/9/ALDS1_9_C

package main

import (
	"github.com/today2098/go-algorithm/algorithm"
	"github.com/today2098/go-algorithm/utils"
)

func main() {
	io := utils.NewIO()
	defer io.Wr.Flush()

	pque := algorithm.NewDefaultBinaryHeap[int]()
	for {
		query := io.GetString()
		if query == "end" {
			break
		}

		if query == "insert" {
			x := io.GetInt()
			pque.Push(x)
		} else {
			io.Out(pque.Pop())
		}
	}
}
