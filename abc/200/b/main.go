// 8 minutes
package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	for i := 0; i < k; i++ {
		n = f(n, k)
	}
	fmt.Println(n)
}

func f(n, k int) int {
	if n%200 == 0 {
		n /= 200
	} else {
		n = n*1000 + 200
	}
	return n
}
