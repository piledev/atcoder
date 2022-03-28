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
	s := reads()

	if len(s) == 1 {
		fmt.Println("Yes")
		return
	}

	if len(s) == 2 {
		if s[0] == 'x' {
			fmt.Println("Yes")
			return
		} else {
			if s[1] == 'x' {
				fmt.Println("Yes")
				return
			} else {
				fmt.Println("No")
				return
			}
		}
	}

	s0 := s[0]
	s1 := s[1]
	s2 := s[2]
	if s0 == 'o' {
		if s1 == 'o' || s2 == 'o' {
			fmt.Println("No")
			return
		}
	} else {
		if !((s1 == 'x' && s2 == 'o') || (s1 == 'o' && s2 == 'x')) {
			fmt.Println("No")
			return
		}
	}

	for i := range s {
		switch i % 3 {
		case 0:
			if s[i] != s0 {
				fmt.Println("No")
				return
			}
		case 1:
			if s[i] != s1 {
				fmt.Println("No")
				return
			}
		case 2:
			if s[i] != s2 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
	return

}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readl() int64   { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func readf() float64 { res, _ := strconv.ParseFloat(reads(), 64); return res }
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
