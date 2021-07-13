// 29 minutes 9 minutes 9 minutes
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N int
var a []int
var dp []int

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

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N = readi()
	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = readi()
	}
	dp = make([]int, N+1)

	for i := range a {
		dp[i+1] = max(dp[i], dp[i]+a[i])
	}

	fmt.Println(dp[N])

}
