package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	W := readl()
	A := make([]int, N)
	m := make(map[int64]int64, N)
	ans := int64(0)
	for i := 0; i < N; i++ {
		A[i] = readi()
		m[int64(A[i])] += readl()
	}
	sort.Ints(A)
	for i := len(A); i > 0; i-- {
		if i < len(A) {
			if A[i-1] == A[i] {
				continue
			}
		}
		if W >= m[int64(A[i-1])] {
			ans += int64(A[i-1]) * m[int64(A[i-1])]
			W -= m[int64(A[i-1])]
		} else {
			ans += int64(A[i-1]) * W
			W = 0
			break
		}
	}

	fmt.Println(ans)
}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readl() int64   { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func readf() float64 { res, _ := strconv.ParseFloat(reads(), 64); return res }
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
