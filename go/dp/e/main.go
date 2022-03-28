// d 問題との違いは,
// d: 1 <= W <= 10^5, 1 <= vi <= 10^9
// e: 1 <= W <= 10^9, 1 <= vi <= 10^3
// => ナップサックの大きさが違う。
// e で dp[i][j] の j を重さにしてしまうとめちゃくちゃ時間がかかる。
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()      // 荷物の数
	W := readInt()      // 限界重量
	w := make([]int, N) // 重さ
	v := make([]int, N) // 価値
	for i := 0; i < N; i++ {
		w[i] = readInt()
		v[i] = readInt()
	}

	ans := run(N, W, w, v)

	fmt.Println(ans)
}

func run(N, W int, w, v []int) int {
	// 1つあたりの価値の最大値は1000
	// それがN個与えられているので合計価値の最大値は 1000*N
	// vMax := 1000 * N
	vMax := 30
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, vMax+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	// 初期値。これがポイントだったりする。
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 1; j <= vMax; j++ {
			if j-v[i] >= 0 {
				dp[i+1][j] = min(dp[i][j], dp[i][j-v[i]]+w[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	res := 0
	for j := range dp[N] {
		if dp[N][j] <= W {
			res = j
		}
	}
	fmt.Println(dp)
	return res
}
