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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	ans := 0
	m := 101
	for i := 0; i < 3; i++ {
		v := readi()
		ans += v
		m = min(m, v)
	}

	fmt.Println(ans - m)
}
