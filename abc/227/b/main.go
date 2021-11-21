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
	ans := 0
	for i := 0; i < N; i++ {
		s := readi()
		max := (s - 3) / 7
		ability := false
		for a := 1; a <= max; a++ {
			for b := 1; b <= max; b++ {
				if 4*a*b+3*(a+b) == s {
					ability = true
					// fmt.Println("a,b,s:", a, b, s)
				}
			}
			if ability {
				break
			}
		}
		if !ability {
			ans++
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
