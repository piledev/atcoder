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

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer(math.MaxInt64)
	N := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}

	B := make([]bool, N)
	for _, v := range A {
		B[v-1] = true
	}

	ans := "Yes"
	for _, b := range B {
		if !b {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}
