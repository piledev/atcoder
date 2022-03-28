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
	Q := readi()
	A := make([]int, 0, Q)
	adds := make([]int, 0, Q)
	issorted := false
	for i := 0; i < Q; i++ {
		q := readi()
		switch q {

		case 1:
			a := readi()
			adds = append(adds, a)

		case 2:
			if issorted {
				sort.Ints(A)
				issorted = false
			}
			A = append(A, adds...)
			adds = nil

			fmt.Println(A[0])
			A = A[1:]

		case 3:
			issorted = true
			A = append(A, adds...)
			adds = nil
		}
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
