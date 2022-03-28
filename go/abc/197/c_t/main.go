package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ri() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	N := ri()
	A := make([]int, N)
	for i := range A {
		A[i] = ri()
	}

	ans := math.MaxInt64
	xor := 0
	// 区切りの位置をbit全探索する
	for i := 0; i < (1 << (N - 1)); i++ {
		ors := make([]int, 0)
		or := 0
		for j := 0; j < N; j++ {
			or ^= A[j]
			if (1 & (i >> j)) == 1 {
				ors = append(ors, or)
				or = 0
			}
		}

		// xor の計算
		for _, v := range ors {
			xor ^= v
		}

		// ans の設定
		if ans > xor {
			ans = xor
		}

	}

	fmt.Println(ans)

}
