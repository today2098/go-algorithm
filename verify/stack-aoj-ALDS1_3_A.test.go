//go:build ignore

// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/3/ALDS1_3_A

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/today2098/go-algorithm/algorithm"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt32)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	query := strings.Split(sc.Text(), " ")

	st := algorithm.NewStack[int]()
	for _, elem := range query {
		var tmp int
		switch elem {
		case "+":
			tmp = st.Pop() + st.Pop()
		case "-":
			tmp = -st.Pop() + st.Pop()
		case "*":
			tmp = st.Pop() * st.Pop()
		default:
			tmp, _ = strconv.Atoi(elem)
		}
		st.Push(tmp)
	}

	ans := st.Top()
	fmt.Println(ans)
}
