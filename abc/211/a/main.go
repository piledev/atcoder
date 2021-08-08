package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() float32 {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return float32(res)
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	a := readi()
	b := readi()
	c := (a-b)/3 + b

	fmt.Println(c)
}
