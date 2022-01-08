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
	black := "#"
	white := "."
	n := readl()
	a := readl()
	b := readl()
	p := readl()
	q := readl()
	r := readl()
	s := readl()

	t11 := max(1-a, 1-b)
	t12 := min(n-a, n-b)
	t21 := max(1-a, b-n)
	t22 := min(n-a, b-1)
	fmt.Println(t11, t12, t21, t22)

	for i := p - 1; i < q; i++ {
		for j := r - 1; j < s; j++ {
			if (a+t11 <= i && i <= a+t12 && b+t11 <= j && j <= b+t12) ||
				(a+t21 <= i && i <= a+t22 && b-t22 <= j && j <= b-t21) {
				fmt.Print(black)
			} else {
				fmt.Print(white)
			}
		}
		fmt.Print("\n")
	}

}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readl() int64   { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func readf() float64 { res, _ := strconv.ParseFloat(reads(), 64); return res }
func min(x, y int64) int64 {
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
