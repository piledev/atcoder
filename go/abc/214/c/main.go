package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	s := make([]int, N)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = readi()
	}
	for i := 0; i < N; i++ {
		dp[i] = readi()
	}

	count := 0
	i := 0
	for count < N {
		temp := dp[i]

		if i == 0 {
			dp[i] = min(dp[i], dp[N-1]+s[N-1])
		} else {
			dp[i] = min(dp[i], dp[i-1]+s[i-1])
		}

		if dp[i] < temp {
			count = 0
		} else {
			count++
		}

		if i == N-1 {
			i = 0
		} else {
			i++
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}
}
