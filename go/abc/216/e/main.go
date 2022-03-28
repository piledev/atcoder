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
	N := readi()
	K := readi()
	// m := make([]int64, 2000000100)
	m := make(map[int64]int64)
	maxA := int64(0)
	for i := int64(0); i < N; i++ {
		A := readi()
		m[A]++
		maxA = max(maxA, A)
	}

	ans := int64(0)
	for i := int64(0); i < K; i++ {
		// fmt.Println(m)
		ans += maxA
		m[maxA]--
		m[maxA-1]++
		if m[maxA] == 0 && maxA > 0 {
			maxA--
		}
		if maxA == 0 {
			break
		}
	}
	fmt.Println(ans)
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int64  { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
