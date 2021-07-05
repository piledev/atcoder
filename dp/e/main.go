package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ri() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	N := ri()
	_ = ri()
	w := make([]int, N)
	v := make([]int, N)
	vmax := 0
	for i := 0; i < N; i++ {
		w[i] = ri()
		v[i] = ri()
		vmax += v[i]
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, vmax+1)
		for j := range dp[i] {
			// dp[i][j] = math.MaxInt32
			dp[i][j] = 30
		}
	}
	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j <= vmax; j++ {
			if j >= v[i] {
				dp[i+1][j] = min(dp[i][j], dp[i][j-v[i]]+w[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

}
