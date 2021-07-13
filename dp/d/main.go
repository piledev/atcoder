package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N, W int
var w, v []int
var dp [][]int

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N = readi()
	W = readi()
	w = make([]int, N)
	v = make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = readi()
		v[i] = readi()
	}
	dp = make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	run_morau()
	fmt.Println(dp[N][W])
}

func run_morau() {
	for i := 0; i < N; i++ {
		for j := 1; j <= W; j++ {
			if w[i] <= j {
				dp[i+1][j] = max(dp[i][j], dp[i][j-w[i]]+v[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
}
