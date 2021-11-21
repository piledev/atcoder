package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	K := readi()
	A := make([]int, N)
	o := make([]int, N)
	for i := 0; i < N; i++ {
		p1 := readi()
		p2 := readi()
		p3 := readi()
		A[i] = p1 + p2 + p3
		o[i] = p1 + p2 + p3
	}
	sort.Ints(o)
	target := o[len(o)-K]

	for i := range A {
		if A[i]+300 >= target {
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
	}
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
