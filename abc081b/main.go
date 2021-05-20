package main

import (
	"fmt"
)

func main() {
	var n int
	answer := 0
	ok := true
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for ok == true {
		for i := 0; i < n; i++ {
			if a[i]%2 == 1 {
				ok = false
			} else {
				a[i] /= 2
			}
		}
		if ok {
			answer++
		}
	}
	fmt.Println(answer)
}
