// 二次元DPを使用した場合
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "math"
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
  ans := math.MaxInt64
	total := 0
  T := make([]int,N)
	for i := range T {
		T[i] = readInt()
		total += T[i]
}

