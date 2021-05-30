package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	ans := 0
	n := nextInt()
	a := make([][]int, n+1)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		if i < n {
			a[i] = []int{nextInt(), nextInt(), nextInt()}
		}
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+a[i][k])
			}
		}
	}

	for i := 0; i < 3; i++ {
		chmax(&ans, dp[n][i])
	}

	fmt.Println(ans)

}

func chmax(x *int, y int) bool {
	if *x < y {
		*x = y
		return true
	}
	return false
}
