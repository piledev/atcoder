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
	N := readi()
	A := make(map[int][]int, N)
	for i := 0; i < N-1; i++ {
		a := readi()
		b := readi()
		A[a] = append(A[a], b)
		A[b] = append(A[b], a)
	}
	ans := "No"
	for i := range A {
		if len(A[i]) == N-1 {
			ans = "Yes"
		}
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
