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
	S := reads()
	K := readi()

	if strings.Count(S, ".") <= K {
		fmt.Println(len(S))
		return
	}
	ans := 0
	for i := range S {
		fmt.Println("-----------------")
		fmt.Println("ans, len(S)-i", ans, len(S)-i)

		if ans >= len(S)-i {
			break
		}
		tlen := 0
		zan := K
		for j := i; j < len(S); j++ {
			// fmt.Println("S[", j, "]:", string(S[j]), zan)
			if S[j] == 'X' {
				tlen++
				ans = max(ans, tlen)
			} else {
				zan--
				if zan < 0 {
					ans = max(ans, tlen)
					break
				}
				tlen++
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
