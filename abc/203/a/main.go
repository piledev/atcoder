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
	a := readInt()
	b := readInt()
	c := readInt()

	var ans int
	switch {
	case a == b:
		ans = c
	case a == c:
		ans = b
	case b == c:
		ans = a
	default:
		ans = 0
	}
	fmt.Println(ans)
}
