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
	N := readi()
	M := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}

	b := make([]int, 1500001)
	for i := 0; i < M; i++ {
		b[A[i]]++
	}

	var ans int
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			ans = i
			break
		}
	}

	for i := 1; i < N-M+1; i++ {
		b[A[i-1]]--
		b[A[i+M-1]]++
		if b[A[i-1]] == 0 {
			if ans > A[i-1] {
				ans = A[i-1]
			}
		}
	}

	fmt.Println(ans)
}
