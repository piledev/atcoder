package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &B[i])
	}

	sort.Ints(A)
	sort.Ints(B)
	ans := B[0] - A[N-1] + 1
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)

}
