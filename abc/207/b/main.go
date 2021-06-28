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

	ans := 0

	pre := -1 * a
	for i := 0; (a + b*i) > (c * i * d); i++ {
		if i > 0 {
			now := (c * i * d) - (a + b*i)
			if now <= pre {
				fmt.Println(-1)
				return
			}
		}
		ans = i
	}
	ans++

	fmt.Println(ans)
}
