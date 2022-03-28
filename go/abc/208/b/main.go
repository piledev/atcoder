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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	p := readi()
	ans := 0

	coins := make([]int, 10)
	for i := 0; i < 10; i++ {
		yen := 1
		for j := 1; j <= i+1; j++ {
			yen *= j
		}
		coins[i] = yen
	}

	for p > 0 {
		p = run(p, coins)
		ans++
		// fmt.Println(p, ans)
	}

	fmt.Println(ans)
}

func run(p int, coins []int) int {
	for i := 10; i > 0; i-- {
		if p >= coins[i-1] {
			p -= coins[i-1]
			break
		}
	}
	return p

}
