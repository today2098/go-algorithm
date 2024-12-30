//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_B

package main

import (
	"github.com/today2098/go-algorithm/algorithm"
	"github.com/today2098/go-algorithm/utils"
)

func main() {
	io := utils.NewIO()
	defer io.Wr.Flush()

	n, q := io.GetInt(), io.GetInt()

	type task struct {
		name string
		time int
	}
	que := algorithm.NewQueue[task]()
	for range n {
		name, time := io.GetString(), io.GetInt()
		que.Push(task{
			name: name,
			time: time,
		})
	}

	now := 0
	for !que.Empty() {
		t := que.Pop()
		if t.time <= q {
			now += t.time
			io.Out(t.name, now)
		} else {
			now += q
			t.time -= q
			que.Push(t)
		}
	}
}
