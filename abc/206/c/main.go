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
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readi()
	a := make([]int, N)
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < N; i++ {
		a[i] = readi()
		m[a[i]]++
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans += (N - 1 - i) - (m[a[i]] - 1)
		m[a[i]]--
	}
	fmt.Println(ans)
}
