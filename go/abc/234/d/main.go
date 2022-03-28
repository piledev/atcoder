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
	n := readi()
	k := readi()
	m := make([]bool, n)
	kth := n + 1
	for i := 0; i < n; i++ {
		p := readi()
		m[p-1] = true
		if i < k {
			kth = min(kth, p)
			if i == k-1 {
				fmt.Println(kth)
			}
		} else {
			if kth < p {
				j := kth
				for m[j] == false {
					j++
				}
				kth = j + 1
			}
			fmt.Println(kth)
		}
	}
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
