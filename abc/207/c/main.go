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

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		t := readi()
		l := readi()
		r := readi()
		switch t {
		case 1:
			L[i] = l * 10
			R[i] = r * 10
		case 2:
			L[i] = l * 10
			R[i] = r*10 - 1
		case 3:
			L[i] = l*10 + 1
			R[i] = r * 10
		case 4:
			L[i] = l*10 + 1
			R[i] = r*10 - 1
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if (L[i] <= L[j] && L[j] <= R[i]) || (L[i] <= R[j] && R[j] <= R[i]) || (L[j] <= L[i] && R[i] <= R[j]) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
