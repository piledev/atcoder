// 9 minutes
package main

import (
	"bufio"
	"fmt"
	"math"
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
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		a := readi()
		p := readi()
		x := readi()

		if x-a > 0 {
			ans = min(ans, p)
		}
	}
	if ans == math.MaxInt64 {
		ans = -1
	}
	fmt.Println(ans)
}
