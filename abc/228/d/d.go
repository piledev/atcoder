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
	N := int64(2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2)
	A := make([]int64, N)
	for i := range A {
		A[i] = -1
	}

	Q := readi()
	for i := 0; i < Q; i++ {
		t := readi()
		x := readi64()
		switch t {
		case 1:
			h := x
			m := h % N
			for {
				if A[m] > -1 {
					m = (m + 1) % N
				} else {
					break
				}
			}
			A[m] = x
		case 2:
			m := x % N
			fmt.Println(A[m])
		}
	}
}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readi64() int64 { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
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
