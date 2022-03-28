package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func min(x, y int) int {
	if x == -1 {
		return y
	}
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	a := make([]int, N)
	p := make([]int, N)
	x := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = readi()
		p[i] = readi()
		x[i] = readi()
	}
	ans := -1
	for i := 0; i < N; i++ {
		if x[i]-a[i] > 0 {
			ans = min(ans, p[i])
		}
	}

	fmt.Println(ans)
}
