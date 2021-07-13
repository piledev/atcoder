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
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
	A := make([]int, N)
	for i := range A {
		A[i] = readInt()
	}

	ans := 0
	for i := range A {
		if A[i] < 10 {
			continue
		} else {
			ans += A[i] - 10
		}
	}
	fmt.Println(ans)
}
