package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var A [][]int
var N int

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N = readi()
	A = make([][]int, 2*N)
	for i := 0; i < 2*N-1; i++ {
		A[i] = make([]int, 2*N)
		for j := range A[i] {
			if j > i {
				A[i][j] = readi()
			}
		}
	}
	used := make([]int, 2*N)
	ans := rec(0, used, 0)
	fmt.Println(ans)
}

func rec(couple int, used []int, point int) int {
	fmt.Println(used)
	if couple == N {
		return point
	}
	maxpoint := 0
	maxi := 0
	maxj := 0
	for i := range A {
		if used[i] == 1 {
			continue
		} else {
			used[i] = 1
			for j := i + 1; j < 2*N; j++ {
				if used[j] == 1 {
					continue
				} else {
					used[j] = 1
					point = point ^ A[i][j]
					t := rec(couple+1, used, point)
					if maxpoint < t {
						maxi = i
						maxj = j
						maxpoint = t
					}
					used[j] = 0
				}
			}
			used[i] = 0
		}
	}
	fmt.Println(maxi, maxj)
	used[maxi], used[maxj] = 1, 1
	return maxpoint
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
