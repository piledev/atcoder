package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var h []int
var dp []int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	h = make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
	}

	dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}

	fmt.Println(rec(n - 1))
}

func rec(i int) int {
	if dp[i] < math.MaxInt32 {
		return dp[i]
	}
	if i == 0 {
		return 0
	}
	res := math.MaxInt32
	chmin(&res, rec(i-1)+abs(h[i-1], h[i]))
	if i > 1 {
		chmin(&res, rec(i-2)+abs(h[i-2], h[i]))

	}
	dp[i] = res
	return res
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
