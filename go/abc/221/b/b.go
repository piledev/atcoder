package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	S := reads()
	T := reads()
	s := strings.Split(S, "")
	t := strings.Split(T, "")

	ans := "No"
	if S == T {
		fmt.Println("Yes")
		return
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			cnt++
			if cnt > 1 {
				ans = "No"
			}

			if i+1 < len(s) {
				if s[i+1] == t[i] && s[i] == t[i+1] {
					ans = "Yes"
					i++
				}
			} else {
				ans = "No"
				break
			}
		}
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
