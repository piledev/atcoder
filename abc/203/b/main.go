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
	N := readInt() // 荷物の数
	W := readInt() // 限界重量

	ans := run(N, W)

	fmt.Println(ans)
}

func run(N, W int) int {
	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			res += i*100 + j
		}
	}
	return res
}
