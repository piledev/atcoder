package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	ans := 0
	if N > 1 {
		ans = N - 1
	}
	fmt.Println(ans)

}
