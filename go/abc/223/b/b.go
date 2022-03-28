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
	S := reads()
	if len(S) == 1 {
		fmt.Println(S)
		fmt.Println(S)
		return
	}
	s := make([]string, len(S))
	t := S
	for i := 0; i < len(S); i++ {
		t = string(t[1:]) + string(t[0])
		s[i] = t
	}
	sort.Strings(s)
	fmt.Println(s[0])
	fmt.Println(s[len(S)-1])

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
