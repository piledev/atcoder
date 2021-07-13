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
	a := float64(readi())
	b := float64(readi())

	ans := (a - b) / a * 100

	fmt.Println(ans)
}
