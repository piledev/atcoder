package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	n := readi()
	m := readi()
	a := make([]int, n)
	b := make([]int, m)
	for i := range a {
		a[i] = readi()
	}
	for i := range b {
		b[i] = readi()
	}
	sort.Ints(a)
	sort.Ints(b)

	ans := 1000_000_000_0
	i, j := 0, 0
	for i < n && j < m {
		ans = min(ans, abs(a[i]-b[j]))
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
