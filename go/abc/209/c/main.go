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
	C := make([]int, N)
	for i := 0; i < N; i++ {
		C[i] = readi()
	}
	ans := 1
	sort.Ints(C)

	for i, v := range C {
		ans *= v - i
		ans = forans(ans)
	}

	fmt.Println(forans(ans))
}

func forans(i int) int {
	div := int(1000_000_007)
	return i % div
}
