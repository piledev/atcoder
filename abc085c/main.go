package main

import (
	"fmt"
)

func main() {
	var n, yen int
	fmt.Scanf("%d %d", &n, &yen)
	yen /= 1000
	x, y, z := judge(0, n, 0, yen)
	fmt.Println(x, y, z)
}

func judge(x, y, z, yen int) (int, int, int) {
	if total(x, y, z) > yen {
		y--
		z++
	} else if total(x, y, z) < yen {
		y--
		x++
	} else {
		return x, y, z
	}
	if y < 0 {
		return -1, -1, -1
	}
	x, y, z = judge(x, y, z, yen)
	return x, y, z
}

func total(x, y, z int) int {
	return x*10 + y*5 + z
}
