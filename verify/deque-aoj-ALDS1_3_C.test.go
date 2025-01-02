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

			for i := len(dq.FrontStack) - 1; i >= 0; i-- {
				if dq.FrontStack[i] == x {
					dq.FrontStack = append(dq.FrontStack[0:i], dq.FrontStack[i+1:len(dq.FrontStack)]...)
					goto Next
				}
			}
			for i := 0; i < len(dq.BackStack); i++ {
				if dq.BackStack[i] == x {
					dq.BackStack = append(dq.BackStack[0:i], dq.BackStack[i+1:len(dq.BackStack)]...)
					goto Next
				}
			}

		Next:
		} else if cmd == "deleteFirst" {
			dq.PopFront()
		} else {
			dq.PopBack()
		}
	}

	ans := []any{}
	for !dq.Empty() {
		ans = append(ans, dq.PopFront())
	}

	io.Out(ans...)
}
