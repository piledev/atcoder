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
	X := readi()
	x := make([]int, 4)
	x[0] = X / 1000
	x[1] = (X % 1000) / 100
	x[2] = (X % 100) / 10
	x[3] = X % 10

	ans := "Strong"

	if X%1111 == 0 {
		ans = "Weak"
	}

	count := 0
	for i := 0; i < 3; i++ {
		if x[i+1]-x[i] == 1 {
			count++
		} else if x[i+1] == 0 && x[i] == 9 {
			count++
		}
	}
	if count == 3 {
		ans = "Weak"
	}

	fmt.Println(ans)
}
