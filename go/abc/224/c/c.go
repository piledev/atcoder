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
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = readi()
		y[i] = readi()
	}

	count := 0
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				if (x[i]-x[j])*(y[i]-y[k])-(y[i]-y[j])*(x[i]-x[k]) == 0 {
					continue
				}
				count++
			}
		}
	}
	fmt.Println(count)
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

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
