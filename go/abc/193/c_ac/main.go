package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	m := make(map[int]bool)
	ans := N

	for a := 2; a*a <= N; a++ {
		if m[a] {
			continue
		}
		for b := a * a; b <= N; b *= a {
			if m[b] {
				break
			}
			m[b] = true
			ans--
		}
	}
	fmt.Println(ans)
}
