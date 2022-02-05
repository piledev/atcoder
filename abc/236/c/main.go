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
	M := readi()
	S := make([]string, N)
	for i := 0; i < N; i++ {
		S[i] = reads()
	}
	T := make([]string, M)
	for i := 0; i < M; i++ {
		T[i] = reads()
	}

	curt := 0
	for s := 0; s < N; s++ {
		if S[s] == T[curt] {
			fmt.Println("Yes")
			curt++
		} else {
			fmt.Println("No")
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
