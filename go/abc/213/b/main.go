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
	N := readi()
	A := make([]int, N)
	m := make(map[int]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
		m[A[i]] = i + 1
	}

	sort.Ints(A)

	ans := m[A[N-2]]

	fmt.Println(ans)
}
