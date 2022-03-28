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
	Q := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}
	sort.Ints(A)

	memo := make([]bool, N)
	for i := 0; i < Q; i++ {
		x := readi()
		ans := 0
		l := 0
		r := N - 1
		m := (l + r) / 2
		for {
			memo[m] = true
			// fmt.Println("l,m,r: ", l, m, r)
			if A[m] == x {
				ans = N - m
				break
			}

			if x < A[m] {
				// fmt.Println("<- A[m]: ", A[m])
				r = m
				m = (l + r) / 2
				if l == m || r-l == 1 {
					ans = N - m
					break
				}
				continue
			}
			if A[m] < x {
				// fmt.Println("-> A[m]: ", A[m])
				l = m
				m = (l + r) / 2
				if r == m || r-l == 1 {
					ans = N - m
					break
				}
				continue
			}
		}
		// for j := range A {
		// 	if A[j] >= x {
		// 		ans = N - j
		// 		break
		// 	}
		// }
		fmt.Println(ans)
	}
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
