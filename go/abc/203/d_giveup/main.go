// 二分探索法らしいが、まだ理解できないので一旦諦める
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
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
	K := readInt()
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			A[i][j] = readInt()
		}
	}
	fmt.Println(run(N, K, A))
}

func run(N, K int, A [][]int) int {
	res := 0
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
