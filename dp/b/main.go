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
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j <= k; j++ {
			if j > i {
				break
			}
			dp[i] = min(dp[i], dp[i-j]+abs(h[i-j], h[i]))
		}
	}

	fmt.Println(dp[n-1])
}

func abs(x, y int) int {
	x = x - y
	if x < 0 {
		return x * -1
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
