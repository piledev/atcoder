// 18:15
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readi()
	K := readi()
	fx := f(N)
	for i := 1; i < K; i++ {
		fx = f(fx)
		if fx == 0 {
			break
		}
	}

	fmt.Println(fx)
}

func f(input int) int {
	s := strconv.Itoa(input)
	a := make([]int, len(s))
	for i := range a {
		a[i], _ = strconv.Atoi(string(s[i]))
		fmt.Println("a:", a)
	}
	sort.Ints(a)
	g1, g2 := 0, 0
	for i := range a {
		g1 = g1*10 + a[len(a)-i-1]
		g2 = g2*10 + a[i]
	}
	fmt.Printf("f(%d) = %d - %d = %d\n", input, g1, g2, g1-g2)
	return g1 - g2
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
