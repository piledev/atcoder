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
	S := reads()
	K := readi()

	a := make([]int, 0, len(S))
	totalc := 0
	countx := 0

	for i := 0; i < len(S); i++ {
		if S[i] == '.' {
			if countx > 0 {
				a = append(a, countx)
				countx = 0
			}
			totalc++
			a = append(a, 0)
		} else {
			countx++
		}
	}
	if countx > 0 {
		a = append(a, countx)
	}

	// fmt.Println(a, totalc)

	ans := 0
	if totalc <= K {
		ans = len(S)
		fmt.Println(ans)
		return
	}

	for i := range a {
		if len(a)-i < ans {
			break
		}
		tempxlen := 0
		tempccount := 0
		for j := i; j < len(a); j++ {
			// fmt.Println("a[", j, "]: ", a[j])
			// a[i] が X のとき
			if a[j] > 0 {
				tempxlen += a[j]

				// . のとき
			} else {
				if tempccount < K {
					tempccount++
					tempxlen++
				} else {
					ans = max(ans, tempxlen)
					break
				}
			}
		}
	}
	fmt.Println(ans)
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
