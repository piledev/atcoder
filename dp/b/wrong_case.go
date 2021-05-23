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
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	k := readInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 110000
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-j]+abs(h[i], h[i-j]))

		}
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
	if x < y {
		return x
	}
	return y
}
