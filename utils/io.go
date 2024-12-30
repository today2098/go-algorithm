package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type IO struct {
	Sc *bufio.Scanner
	Wr *bufio.Writer
}

func NewIO() *IO {
	io := &IO{
		Sc: bufio.NewScanner(os.Stdin),
		Wr: bufio.NewWriter(os.Stdout),
	}
	io.Sc.Buffer([]byte{}, math.MaxInt32)
	io.Sc.Split(bufio.ScanWords)
	return io
}

func (io *IO) GetInt() int {
	io.Sc.Scan()
	res, err := strconv.Atoi(io.Sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func (io *IO) GetInts(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = io.GetInt()
	}
	return res
}

func (io *IO) GetFloat64() float64 {
	io.Sc.Scan()
	res, err := strconv.ParseFloat(io.Sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (io *IO) GetString() string {
	io.Sc.Scan()
	return io.Sc.Text()
}

func (io *IO) GetStrings(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = io.GetString()
	}
	return res
}

func (io *IO) Out(x ...interface{}) {
	fmt.Fprintln(io.Wr, x...)
}

func (io *IO) Flush() {
	io.Wr.Flush()
}
