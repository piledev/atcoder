package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	h := make([]int, N)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = readi()
	}
	for i := 0; i < N; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	fmt.Println(run_morau(N, h, dp))
	fmt.Println(run_kubaru(N, h, dp))
	fmt.Println(run_recursive(N, h, dp))
}

func run_morau(N int, h, dp []int) int {
	for i := 1; i < N; i++ {
		dp[i] = min(dp[i], dp[i-1]+abs(h[i-1]-h[i]))
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+abs(h[i-2]-h[i]))
		}
	}
	return dp[N-1]
}

func run_kubaru(N int, h, dp []int) int {
	for i := 0; i < N-1; i++ {
		dp[i+1] = min(dp[i+1], dp[i]+abs(h[i]-h[i+1]))
		if i+2 < N {
			dp[i+2] = min(dp[i+2], dp[i]+abs(h[i]-h[i+2]))
		}
	}
	return dp[N-1]
}

func run_recursive(N int, h, dp []int) int {

	return rec(h, dp, N, N-1)

}

func rec(h, dp []int, N, i int) int {
	if i == 0 {
		return 0
	}
	if dp[i] < math.MaxInt64 {
		return dp[i]
	}
	dp[i] = min(dp[i], rec(h, dp, N, i-1)+abs(h[i]-h[i-1]))
	if i-2 < N {
		dp[i] = min(dp[i], rec(h, dp, N, i-2)+abs(h[i]-h[i-2]))
	}
	fmt.Println(dp)
	return dp[i]
}
