// 47 minutes
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
	A := make([]int, N)
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < N; i++ {
		A[i] = readi()
		m[A[i]]++
	}

	ans := 0

	for v, t := range m {
		for vv, tt := range m {
			if v < vv {
				ans += (v - vv) * (v - vv) * t * tt
			}
		}
	}
	fmt.Println(ans)
}
