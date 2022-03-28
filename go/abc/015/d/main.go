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

	// ここもうちょっと良く見てみること。具体的な値を入れてどういうことか考えたい。
	for n := 0; n < N; n++ {
		for k := 1; k <= K; k++ {
			for w := 1; w <= W; w++ {
				// wよりa[n]小さい値が0以上のとき
				if w-a[n] >= 0 {
					// dp[n][k-1][w-a[n]]にn個目のスクショを追加すると、
					// [k-1]個が1個増えて[k]になり
					// [w-a[n]]の幅がa[n]だけ増えて[w]になり
					// 優先度にはb[n]が加算される。
					dp[n+1][k][w] = dp[n][k-1][w-a[n]] + b[n]
				}
				// dp[n+1][k][w]の現在の値と、1つ前のスクショを比較し、大きい方を採用する。
				dp[n+1][k][w] = max(dp[n+1][k][w], dp[n][k][w])

			}
		}
	}
	fmt.Println(dp[N][K][W])
}
