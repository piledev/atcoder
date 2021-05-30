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
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	h := make([]int, n+2)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}

	dp := make([]int, n+2)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		chmin(&dp[i+1], dp[i]+abs(h[i], h[i+1]))
		chmin(&dp[i+2], dp[i]+abs(h[i], h[i+2]))
	}
	fmt.Println(dp[n-1])
}

func abs(x, y int) int {
	ret := x - y
	if ret < 0 {
		return ret * -1
	}
	return ret
}

func chmin(x *int, y int) bool {
	if *x > y {
		*x = y
		return true
	}
	return false
}
