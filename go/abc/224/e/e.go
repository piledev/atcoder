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
	N := readi()
	A := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
	}
	for i := 0; i < N; i++ {
		r := readi()
		c := readi()
		a := readi()
		A[r][c] = a
	}

	ans := N

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
