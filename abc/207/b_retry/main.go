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

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	a := readi()
	b := readi()
	c := readi()
	d := readi()
	if c*d-b <= 0 {
		fmt.Println(-1)
		return
	}
	ans := (a + (c*d - b) - 1) / (c*d - b)
	fmt.Println(ans)
}
