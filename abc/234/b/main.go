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
	X := make([]float64, N)
	Y := make([]float64, N)
	for i := 0; i < N; i++ {
		X[i] = readf()
		Y[i] = readf()
	}

	ans := float64(0)
	for i := range X {
		for j := i + 1; j < N; j++ {
			ans = max(ans, distance(X[i], X[j], Y[i], Y[j]))

		}

	}

	fmt.Println(ans)
}

func distance(x1, x2, y1, y2 float64) float64 {
	ret := float64(0)
	ret = math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	return ret

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
func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
