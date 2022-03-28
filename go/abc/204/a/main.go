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
	x := readInt()
	y := readInt()

	var ans int
	switch {
	case x == y:
		ans = x
	case (x == 0) && (y == 1):
		ans = 2
	case (x == 0) && (y == 2):
		ans = 1
	case (x == 1) && (y == 2):
		ans = 0
	case (x == 1) && (y == 0):
		ans = 2
	case (x == 2) && (y == 1):
		ans = 0
	case (x == 2) && (y == 0):
		ans = 1
	}
	fmt.Println(ans)
}
