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
	A := make([]int64, N)
	Asum := int64(0)
	for i := int64(0); i < N; i++ {
		A[i] = readi()
		Asum += A[i]
	}
	X := readi()

	cnt := (X / Asum) * N
	t := Asum * (X / Asum)
	for i := range A {
		t += A[i]
		cnt++
		if t > X {
			break
		}
	}

	fmt.Println(cnt)
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int64  { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
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
