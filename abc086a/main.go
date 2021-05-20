package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	c := a * b
	var message string
	if c%2 == 1 {
		message = "Odd"
	} else {
		message = "Even"
	}
	fmt.Println(message)
}
