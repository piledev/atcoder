// 31 minutes
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N int
var s, dp [][]int

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

func main() {
	sc.Split(bufio.ScanWords)
	N = readi()
	s = make([][]int, 3)
	for i := range s {
		s[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			s[j][i] = readi()
		}
	}
	dp = make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	// run_morau()
	// run_kubaru()
	run_recursive()
	fmt.Println(ans())
}

func ans() int {
	res := 0
	for i := range dp {
		res = max(res, dp[i][N-1])
	}
	return res
}

func log(a [][]int) {
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Println("----------------")
}

func run_kubaru() {
	for i := 0; i < N-1; i++ {
		if i == 0 {
			for j := 0; j < 3; j++ {
				dp[j][i] = s[j][i]
			}
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if k != j {
					dp[j][i+1] = max(dp[j][i+1], dp[k][i]+s[j][i+1])
					// log(dp)
				}
			}
		}
	}
}

func run_morau() {
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			if i == 0 {
				dp[j][i] = s[j][i]
				continue
			}
			for k := 0; k < 3; k++ {
				if k != j {
					dp[j][i] = max(dp[j][i], dp[k][i-1]+s[j][i])
				}
			}
		}
	}
}

func run_recursive() {
	for i := 0; i < 3; i++ {
		rec(i, N-1)
	}
	// log(dp)
}

func rec(x, y int) int {
	if dp[x][y] != 0 {
		return dp[x][y]
	}
	if y == 0 {
		dp[x][y] = s[x][y]
		return dp[x][y]
	}
	for i := 0; i < 3; i++ {
		if i == x {
			continue
		}
		dp[x][y] = max(dp[x][y], rec(i, y-1)+s[x][y])
	}

	return dp[x][y]
}
