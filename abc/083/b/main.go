package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b, total int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	for i := 1; i <= n; i++ {
		istr := strconv.Itoa(i)
		sum := 0
		for _, v := range istr {
			vint, _ := strconv.Atoi(string(v))
			sum += vint
		}
		if a <= sum && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}
