package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	a := readInt()
	b := readInt()
	k := readInt()

	fmt.Println(run(a, b, k))
}

func run(a int, b int, k int) string {
	s := ""
	// 初期化
	for i := 0; i < a+b; i++ {
		if i < a {
			s += "a"
		} else {
			s += "b"
		}
	}

	counter := 1
	for counter < k {
		for i := 0; i < a+b; i++ {
			now := a + b - i - 1
			if s[now] == 'a' {
				// sa := strings.Split(s,"")
				// s[now] = byte('b')
			}

		}

		if strings.Count(s, "a") == a {
			counter++
		}

	}

	return s
}

func abs(x, y int) int {
	x = x - y
	if x < 0 {
		return x * -1
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
