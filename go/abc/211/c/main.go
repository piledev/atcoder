package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	S := reads()

	slen := len(S)
	T := "chokudai"
	dp := make([][]int, slen+1)
	for i := range dp {
		dp[i] = make([]int, 9)
	}

	fmt.Println(slen)

	for i := 0; i < slen+1; i++ {
		for j := 0; j < 9; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else {
				if S[i] != T[j] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = getAns(dp[i-1][j] + dp[i-1][j-1])
				}
			}
		}
	}

	ans := dp[slen][8]

	fmt.Println(ans)
}

func getAns(i int) int {
	return i % 1_000_000_007
}
