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
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	X := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
		if i%2 == 1 {
			A[i] -= 1
		}
	}
	ans := "Yes"
	for _, v := range A {
		X -= v
		if X < 0 {
			ans = "No"
		}
	}

	fmt.Println(ans)
}
