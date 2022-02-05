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
	H := readi()
	W := readi()
	B := make([][]int, W)
	for i := 0; i < W; i++ {
		B[i] = make([]int, H)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			B[j][i] = readi()
		}
	}

	for i := range B {
		for j := range B[i] {
			if j == H-1 {
				fmt.Printf("%d\n", B[i][j])
			} else {
				fmt.Printf("%d ", B[i][j])
			}
		}
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
