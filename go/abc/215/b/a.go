package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int64 {
	sc.Scan()
	res, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()

	k := 0
	temp := int64(1)
	for {
		temp *= 2
		if temp <= N {
			k++
		} else {
			break
		}
	}

	fmt.Println(k)
}
