package main

import (
	"fmt"
)

type Travel struct {
	t int
	x int
	y int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]Travel, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &a[i].t, &a[i].x, &a[i].y)
	}

	ok := true
	now := Travel{0, 0, 0}
	for _, next := range a {
		count := next.t - now.t
		distance := abs(next.x-now.x) + abs(next.y-now.y)
		if count-distance < 0 || (count-distance)%2 == 1 {
			ok = false
			break
		}
		now = next
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
