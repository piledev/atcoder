package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%v", &input)
	ret := 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "1" {
			ret++
		}
	}
	fmt.Println(ret)
}
