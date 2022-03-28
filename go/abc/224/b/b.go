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
	H := readi()
	W := readi()
	A := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			A[i][j] = readi()
		}
	}

	for i1 := 0; i1 < H-1; i1++ {
		for j1 := 0; j1 < W-1; j1++ {
			for i2 := i1 + 1; i2 < H; i2++ {
				for j2 := j1 + 1; j2 < W; j2++ {
					if A[i1][j1]+A[i2][j2] > A[i2][j1]+A[i1][j2] {
						fmt.Println("No")
						return
					}
				}
			}
		}
	}
	fmt.Println("Yes")
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
