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
	// zero := "000000000000000000"
	N := readl()
	N = N % 998244353

	ans := int64(0)
	for i := int64(0); i < N; i++ {
		ans += f(i + 1)
	}
	fmt.Println(ans % 998244353)

}

func f(x int64) int64 {
	var ret int64
	if x < 10 {
		ret = x
	} else {
		s := strconv.FormatInt(x, 10)
		head := string(s[0])
		s = s[1:]
		if head != "1" {
			ihead, _ := strconv.ParseInt(string(head), 10, 64)
			ihead--
			head = strconv.FormatInt(ihead, 10)
			s = head + s
		}

		ret, _ = strconv.ParseInt(s, 10, 64)
		ret++
	}
	// fmt.Printf("f(%d) = %d\n", x, ret%998244353)
	return ret % 998244353
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
