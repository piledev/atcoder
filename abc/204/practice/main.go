package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ri() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := ri()
	W := ri()
	a := make([]int, N)
	for i := range a {
		a[i] = ri()
	}
	ans := false
	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			if (bit & (1 << i)) == 1 {
				sum += a[i]
				if sum == W {
					ans = true
				}
			}
		}
	}
	fmt.Println(ans)
}
