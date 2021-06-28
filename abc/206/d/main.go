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
	A := make([]int, N)
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < N; i++ {
		A[i] = readi()
		m[A[i]] = A[i]
	}

	ans := 0
	for i := 0; i < N/2; i++ {
		l := A[i]
		r := A[N-1-i]
		for {
			if m[l] == l && m[r] == r {
				m[A[i]] = l
				m[A[N-1-i]] = r
				break
			} else {
				l = m[m[l]]
				r = m[m[r]]
			}
		}
		if l == r {
			continue
		}
		m[l] = m[r]
		ans++
	}
	fmt.Println(ans)
}
