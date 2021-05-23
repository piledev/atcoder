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

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(h[0], h[1])
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-2]+abs(h[i-2], h[i]), dp[i-1]+abs(h[i-1], h[i]))
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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
