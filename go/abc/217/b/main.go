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
	m := make(map[string]bool, 4)
	m["ABC"] = false
	m["ARC"] = false
	m["AGC"] = false
	m["AHC"] = false
	for i := 0; i < 3; i++ {
		a := reads()
		m[a] = true
	}

	for k, v := range m {
		if !v {
			fmt.Println(k)
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
