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
	k := readi()
	a := readi()
	b := readi()

	if k == 10 {
		fmt.Println(a * b)
		return
	}
	a = change(a, k)
	b = change(b, k)

	fmt.Println(a * b)

}

func change(a, k int) int {
	sum := 0
	for i := 1; i <= 100; i++ {
		s := a % 10
		a /= 10
		tk := 1
		for j := 1; j < i; j++ {
			tk *= k
		}
		sum += s * tk
		if a == 0 {
			break
		}
	}
	return sum
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
