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
	m := make([]string, 120)
	temp := N
	for i := range m {
		if temp%2 == 1 {
			m[i] = "A"
			temp -= 1
		} else if temp > 0 {
			m[i] = "B"
			temp /= 2
		} else {
			break
		}
	}

	// ans := ""
	for i := 119; i >= 0; i-- {
		if m[i] != "" {
			fmt.Printf(m[i])
		}
	}
	fmt.Println("")
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
