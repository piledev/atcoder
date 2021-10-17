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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}

	sum := 0
	for i := range A {
		t := 0
		for j := i + 1; j < N; j++ {
			if A[i] <= A[j] {
				for k := j - i; k > 1; k-- {
					sum *= 2
				}

			}

		}

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
