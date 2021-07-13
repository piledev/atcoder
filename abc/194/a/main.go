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
	A := readi()
	B := readi()

	AB := A + B
	ans := 0

	if AB >= 15 && B >= 8 {
		//icecream
		ans = 1
	} else if AB >= 10 && B >= 3 {
		//icemilk
		ans = 2
	} else if AB >= 3 {
		// ractice
		ans = 3
	} else {
		// hyoka
		ans = 4
	}

	fmt.Println(ans)
}
