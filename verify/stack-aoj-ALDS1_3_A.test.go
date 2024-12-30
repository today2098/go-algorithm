//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/today2098/go-algorithm/algorithm"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt32)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	query := []rune(sc.Text())

	len := len(query)
	st := algorithm.NewStack[int]()
	for i := 0; i < len; i += 2 {
		c := query[i]
		if c == '+' {
			tmp := st.Pop() + st.Pop()
			st.Push(tmp)
		} else if c == '-' {
			tmp := -st.Pop() + st.Pop()
			st.Push(tmp)
		} else if c == '*' {
			tmp := st.Pop() * st.Pop()
			st.Push(tmp)
		} else {
			st.Push(int(c - '0'))
		}
	}

	ans := st.Pop()
	fmt.Println(ans)
}
