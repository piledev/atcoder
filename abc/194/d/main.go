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
	fN := float64(N)
	ans := float64(0)

	for i := 1; i < N; i++ {
		ans += 1.0 * fN / (fN - float64(i))
	}

	fmt.Println(ans)
}
