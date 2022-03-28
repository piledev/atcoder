package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readi()
	S := reads()
	t := make([]int, 1, 500000)
	m := make([]int, 0, 500000)
	m = append(m, N)

	for i := N - 1; i >= 0; i-- {
		if S[i] == 'R' {
			t[0] = i
			m = append(t, m...)
		} else {
			m = append(m, i)
		}
	}

	ans := fmt.Sprint(m)
	ans = strings.Trim(ans, "[")
	ans = strings.Trim(ans, "]")
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
