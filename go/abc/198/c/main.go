package main

import (
	"fmt"
	"math"
)

func main() {
	var R, X, Y int
	fmt.Scanf("%d %d %d", &R, &X, &Y)
	r, x, y := float64(R), float64(X), float64(Y)
	distance := math.Sqrt(x*x + y*y)
	var ans float64
	if distance < r {
		ans = 2
	} else {
		ans = math.Ceil(distance / r)
	}
	fmt.Println(ans)
}
