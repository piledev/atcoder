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
	M := readi()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readi()
	}
	memo := make([]bool, M)

	for _, v := range A {
		if v == 1 {
			continue
		}
		// if v-1 < len(memo) {
		// 	if memo[v-1] {
		// 		continue
		// 	}
		// }
		for i := 2; i < M; i++ {
			if v < i {
				break
			}
			if memo[i-1] {
				continue
			}
			if v%i == 0 {
				for j := 1; j <= M/i; j++ {
					memo[i*j-1] = true
				}
			}
		}
	}

	ans := make([]int, 0, M)
	for i := range memo {
		if memo[i] == false {
			ans = append(ans, i+1)
		}
	}
	sort.Ints(ans)

	fmt.Println(len(ans))
	for i := range ans {
		fmt.Println(ans[i])
	}
}
