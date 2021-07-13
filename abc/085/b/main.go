package main

import (
	"fmt"
)

func main() {
	var n, count int
	m := make(map[int]int)
	fmt.Scanf("%d", &n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d[i])
		m[d[i]]++
	}

	count = len(m)
	fmt.Println(count)
}
