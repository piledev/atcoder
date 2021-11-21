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
	K := readi()
	m := make(map[int]int, N)
	tmax := 0
	for i := 0; i < N; i++ {
		a := readi()
		m[a]++
		tmax = max(tmax, a)
	}

	for i := 0; i < K; i++ {
		added := make([]int, tmax+1)
		for j := tmax; j > 0; j-- {
			if m[j] > 0 {
				m[j]--
				m[j-1]++
				added[j-1]++
				break
			}

		}

	}

	ans := N

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
