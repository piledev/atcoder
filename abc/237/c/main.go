package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	S := reads()
	taila := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == 'a' {
			taila++
		} else {
			break
		}
	}

	T := S[:len(S)-taila]

	if len(T) == 0 {
		fmt.Println("Yes")
		return
	}

	heada := 0
	for i := 0; i < taila; i++ {
		if T[i] == 'a' {
			heada++
		} else {
			break
		}
	}

	T = T[heada:]

	// fmt.Println(T)

	r := len(T) - 1
	for l := 0; l < r; l++ {
		// fmt.Println(l, S[l], r, S[r])
		if T[l] == T[r] {
			r--
			continue
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
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
