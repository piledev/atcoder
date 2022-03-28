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
	p := make([]bool, 360)
	p[0] = true
	cur := 0
	N := readi()
	for i := 0; i < N; i++ {
		t := readi()
		cur = (cur + t) % 360
		p[cur] = true
		// fmt.Println(t, cur)
	}

	ans := 0
	cnt := 0
	for i := range p {
		cnt++
		if p[i] {
			// fmt.Println(ans, cnt)
			ans = max(ans, cnt)
			cnt = 0
		}
	}
	ans = max(ans, cnt+1)

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
