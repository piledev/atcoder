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
	s := readi()
	t := readi()
	x := readi()
	if s <= x && x < t {
		fmt.Println("Yes")
		return
	}
	if t < s {
		if s <= x || x < t {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
