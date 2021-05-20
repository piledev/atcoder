// 59 minutes
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%v", &a)
		a %= 200
		m[a]++
	}
	count := 0
	for _, v := range m {
		if v == 0 {
			continue
		}
		count += (v * (v - 1)) / 2
	}
	fmt.Println(count)
}
