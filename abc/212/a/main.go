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
	// sc.Buffer([]byte{},math.MaxInt64)
	a := readi()
	b := readi()
	ans := ""

	if 0 < a && b == 0 {
		ans = "Gold"
	} else if a == 0 && 0 < b {
		ans = "Silver"
	} else if 0 < a && 0 < b {
		ans = "Alloy"
	}

	fmt.Println(ans)
}
