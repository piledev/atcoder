package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var a, b []int
var N int

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N = readi()
	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = readi()
	}
	b = make([]int, N)
	for i := 0; i < N; i++ {
		b[i] = readi()
	}

	ans := rec(0, a[0])

	fmt.Println(ans)
}

func rec(i, v int) int {
	if i > N-1 {
		return 1
	}

	ret := 0
	for t := max(a[i], v); t <= b[i]; t++ {
		next := rec(i+1, t)
		ret += next
	}

	return ret % 998244353
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int    { res, _ := strconv.Atoi(reads()); return res }
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
