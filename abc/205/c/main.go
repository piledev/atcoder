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

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer(math.MaxInt64)
	a := readi()
	b := readi()
	c := readi()

	ans := ""
	if c%2 == 0 {
		switch {
		case abs(a) > abs(b):
			ans = ">"
		case abs(a) < abs(b):
			ans = "<"
		default:
			ans = "="
		}
	} else {
		switch {
		case a > b:
			ans = ">"
		case a < b:
			ans = "<"
		default:
			ans = "="
		}
	}

	fmt.Println(ans)
}
