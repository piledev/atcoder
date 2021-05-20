//29 mitutes
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	name := make([]string, n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &name[i], &h[i])
	}

	top1, top2 := -1, -1

	if h[0] > h[1] {
		top1 = 0
		top2 = 1
	} else {
		top1 = 1
		top2 = 0
	}
	for i := 2; i < n; i++ {

		if h[i] > h[top1] {
			top2 = top1
			top1 = i
		} else if h[i] > h[top2] {
			top2 = i
		}
	}
	fmt.Println(name[top2])
}
