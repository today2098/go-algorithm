//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_C

package main

import (
	"github.com/today2098/go-algorithm/algorithm"
	"github.com/today2098/go-algorithm/utils"
)

func main() {
	io := utils.NewIO()
	defer io.Wr.Flush()

	n := io.GetInt()

	dq := algorithm.NewDeque[int]()
	for i := 0; i < n; i++ {
		cmd := io.GetString()

		if cmd == "insert" {
			x := io.GetInt()

			dq.PushFront(x)
		} else if cmd == "delete" {
			x := io.GetInt()

			node := dq.Data.Front()
			for node != nil && node.Value.(int) != x {
				node = node.Next()
			}

			if node != nil {
				dq.Data.Remove(node)
			}
		} else if cmd == "deleteFirst" {
			dq.PopFront()
		} else {
			dq.PopBack()
		}
	}

	ans := []any{}
	node := dq.Data.Front()
	for node != nil {
		ans = append(ans, node.Value.(int))
		node = node.Next()
	}

	io.Out(ans...)
}
