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
	A := make([]float64, N)
	B := make([]float64, N)
	C := make([]float64, N)
	totaltime := float64(0)
	for i := 0; i < N; i++ {
		A[i] = float64(readi())
		B[i] = float64(readi())
		C[i] = A[i] / B[i]
		totaltime += C[i]
	}
	halftime := totaltime / 2

	temptime := float64(0)
	ans := float64(0)
	for i := 0; i < N; i++ {
		if temptime+C[i] >= halftime {
			ans += (halftime - temptime) * B[i]
			break
		}

		ans += A[i]
		temptime += C[i]
	}

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
