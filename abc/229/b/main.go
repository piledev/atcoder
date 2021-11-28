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
	a := reads()
	b := reads()

	ans := "Easy"
	count := min(len(a), len(b))
	for i := 0; i < count; i++ {
		ia, _ := strconv.Atoi(string(a[len(a)-1-i]))
		ib, _ := strconv.Atoi(string(b[len(b)-1-i]))
		// fmt.Println(ia, ib)
		if ia+ib > 9 {
			ans = "Hard"
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
