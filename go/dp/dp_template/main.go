// 29 minutes 9 minutes 9 minutes
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var N, K int
var h []int
var dp []int

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N = readi()
	K = readi()
	h = make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = readi()
	}
	dp = make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	// run_kubaru()
	// run_morau()
	run_recursive()
	fmt.Println(dp[N-1])
}

func run_kubaru() {
}

func run_morau() {
}

func run_recursive() {
}

func rec(i int) int {
	return 0
}
