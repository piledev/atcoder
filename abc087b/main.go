package main

import (
	"fmt"
)

func main() {

	var a, b, c, x, counter int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	for ta := 0; ta <= a; ta++ {
		for tb := 0; tb <= b; tb++ {
			for tc := 0; tc <= c; tc++ {
				if 500*ta+100*tb+50*tc == x {
					counter++
				}
			}
		}
	}

	fmt.Println(counter)

}
