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
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	K := readi()
	a := make([]int, N)
	s := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = readi()
	}
	_ = copy(s, a)

	if K%N == 0 {
		tmp := K / N
		for i := 0; i < N; i++ {
			fmt.Println(tmp)
		}
		return
	}

	sort.Ints(s)
	d := K / N
	m := K % N
	t := s[m-1]
	// fmt.Println(d, m, t, a, s)
	for _, v := range a {
		if v <= t {
			fmt.Println(d + 1)
		} else {
			fmt.Println(d)
		}
	}
	return

}
