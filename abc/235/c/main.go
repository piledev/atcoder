package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readi()
	Q := readi()
	a := make(map[int][]int, N)
	for i := 0; i < N; i++ {
		v := readi()
		a[v] = append(a[v], i+1)
	}

	for i := 0; i < Q; i++ {
		x := readi()
		k := readi()
		ans := 0

		if a[x] == nil {
			ans = -1
		} else if len(a[x]) < k {
			ans = -1
		} else {
			ans = a[x][k-1]
		}

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
