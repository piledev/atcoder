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
	K := readl()
	A := make([]int64, N)
	for i := 0; i < N; i++ {
		A[i] = readl()
	}

	ans := 0
	t := 0
	sum := int64(0)
	for l := 0; l < N; l++ {
		for r := max(l, t); r < N; r++ {
			fmt.Println(l, r)
			sum += A[r]
			t = r
			if sum == K {
				ans++
				fmt.Println(sum, K, l, r)
			}
			if r == N-1 {
				break
			}
		}
		sum -= A[l]
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
