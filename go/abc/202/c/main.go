package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	a := make(map[int]int)
	b := make([]int, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		a[readInt()]++
	}
	for i := 0; i < n; i++ {
		b[i] = readInt()
	}
	for i := 0; i < n; i++ {
		c[i] = readInt()
	}

	fmt.Println(run(n, a, b, c))
}

func run(n int, a map[int]int, b, c []int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans += a[b[c[i]-1]]
	}
	return ans
}
