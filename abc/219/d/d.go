package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	X := readi()
	Y := readi()
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
		B[i] = readi()
	}
	ans := -1

	mincnt := 301
	for i := 0; i < (1 << N); i++ {
		x := 0
		y := 0
		cnt := 0
		for j := 0; j < N; j++ {
			// b := i&(1<<j) >= 1
			// fmt.Printf("i: %d, j: %d, b: %v\n", i, j, b)
			if (i & (1 << j)) > 0 {
				// fmt.Println("true")
				cnt++
				x += A[j]
				y += B[j]
			}
		}
		if x >= X && y >= Y {
			mincnt = min(mincnt, cnt)
			// fmt.Println("mincnt:", mincnt)
			// break
		}
	}
	if mincnt < 301 {
		ans = mincnt
	}
	fmt.Println(ans)
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int    { res, _ := strconv.Atoi(reads()); return res }
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
