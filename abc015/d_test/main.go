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
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	W := readInt()
	N := readInt()
	K := readInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = readInt()
		b[i] = readInt()
	}
	dp := make([][][]int, N+1)
	for n := range dp {
		dp[n] = make([][]int, K+1)
		for k := range dp[n] {
			dp[n][k] = make([]int, W+1)
		}
	}

	for n := 0; n < N; n++ {
		for k := 1; k <= K; k++ {
			for w := 1; w <= W; w++ {
				if w-a[n] >= 0 {
					// スクショを追加できるときは、追加した最大値と追加しない最大値を比較して大きい方を採用する.
					// 追加した場合の最大値は、
					// スクショの数が一つ少なく、今のスクショの幅が入る余地があるdpの値に、
					// 今のスクショの優先度を足す形になる。
					dp[n+1][k][w] = max(dp[n][k][w], dp[n][k-1][w-a[n]]+b[n])
				} else {
					// スクショを追加できないときは、追加しない最大値を採用する。
					dp[n+1][k][w] = dp[n][k][w]
				}
			}
		}
	}
	fmt.Println(dp[N][K][W])
}
