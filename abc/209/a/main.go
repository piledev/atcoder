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

func main() {
	sc.Split(bufio.ScanWords)
	a := readi()
	b := readi()
	ans := b - a + 1
	if ans < 0 {
		ans = 0
	}

	fmt.Println(ans)
}
