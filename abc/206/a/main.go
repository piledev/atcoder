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

	x := (108 * N) / 100
	var ans string

	switch {
	case x < 206:
		ans = "Yay!"
	case x == 206:
		ans = "so-so"
	default:
		ans = ":("
	}
	fmt.Println(ans)
}
