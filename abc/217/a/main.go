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
	s := reads()
	t := reads()
	st := make([]string, 2)
	st[0] = s
	st[1] = t

	ans := "No"
	sort.Strings(st)
	if st[0] == s {
		ans = "Yes"
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
