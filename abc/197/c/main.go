package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}
	ans := math.MaxInt64

	for i := 1; i < N; i++ {
		a := A[:i]
		b := A[i:]
		asum := 0
		bsum := 0
		for _, v := range a {
			asum |= v
		}
		for _, v := range b {
			bsum |= v
		}
		// fmt.Println(ans, asum^bsum)
		fmt.Println(len(a), len(b))
		ans = min(ans, asum^bsum)
	}
	fmt.Println(ans)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
