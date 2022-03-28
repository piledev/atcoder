package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		p := readi()
		Q[p-1] = i + 1
	}

	ans := make([]byte, 0, math.MaxInt32)
	for i := range Q {
		str := strconv.Itoa(Q[i])
		ans = append(ans, str...)
		ans = append(ans, ' ')
	}

	fmt.Println(strings.Trim(string(ans), " "))
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
