package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
	W := readInt()
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = readInt()
		v[i] = readInt()
	}

	// dp初期化
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		// i個目までの品物のうち重量j以下の重量の中での価値の最大値をメモしていく
		for j := 0; j <= W; j++ {
			if i == 0 {
				if j >= w[i] {
					dp[i][j] = v[i]
				}
			} else {
				if w[i] <= j {
					// i番目の品物の重量がj以下の場合は品物を入れられる
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
				} else {
					// i番目の品物の重量がjより大きい場合は品物を入れられないので最大価値は変わらない
					dp[i][j] = dp[i-1][j]
				}
			}
			// fmt.Println(i, j, dp)
		}
	}
	fmt.Println(dp[N-1][W])
}
