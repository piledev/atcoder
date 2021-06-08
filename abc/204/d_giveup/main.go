// 二分探索法らしいが、まだ理解できないので一旦諦める
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
	T := make([]int, N)
	memo := make([]int, 1<<N)
	total := 0
	for i := range T {
		T[i] = readInt()
		total += T[i]
	}

	for bit := 0; bit < (1 << N); bit++ {
		for i := 0; i < N; i++ {
			if (bit & (1 << i)) == 1 {
				memo[bit-1] += T[i]
			}
		}
	}
	// sort.Ints(memo)

	fmt.Println(memo)
}
