// 44 minutes
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
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
		B[i] = readi()
	}

	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				ans = min(ans, A[i]+B[j])
			} else {
				ans = min(ans, max(A[i], B[j]))
			}
		}
	}

	fmt.Println(ans)
}
