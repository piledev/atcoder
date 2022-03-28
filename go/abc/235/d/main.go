package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var a int
var ans int = -1
var mdep []int

func main() {
	sc.Split(bufio.ScanWords)
	a = readi()
	N := readi()

	// if a == N {
	// 	fmt.Println(0)
	// 	return
	// }
	mdep = make([]int, 1000000) // default 0

	rec(N, 0)

	fmt.Println(ans)
}

func rec(val, dep int) {
	// fmt.Println("val,dep: ", val, dep)
	if val == 1 {
		ans = min(ans, dep)
		return
	}

	if ans > 0 && ans <= dep {
		return
	}

	if mdep[val] <= dep && mdep[val] > 0 {
		return
	} else {
		mdep[val] = dep
	}

	if val%a == 0 {
		rec(val/a, dep+1)
	}

	if val > 10 {
		sval := strconv.Itoa(val)
		head := string(sval[0])
		seco := string(sval[1])
		if seco == "0" {
			return
		}
		sval = sval[1:] + head
		val, _ = strconv.Atoi(sval)
		rec(val, dep+1)
	}

}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readl() int64   { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func readf() float64 { res, _ := strconv.ParseFloat(reads(), 64); return res }
func min(x, y int) int {
	if x > y || x == -1 {
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
