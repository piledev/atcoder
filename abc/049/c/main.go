package main

import (
	"fmt"
	"strings"
)

func main() {
	words := [4]string{"dream", "dreamer", "erase", "eraser"}
	var s string
	fmt.Scanf("%v", &s)
	for i := 0; i < len(words); i++ {
		s = strings.Replace(s, words[4-i-1], " ", -1)
	}
	if strings.TrimSpace(s) == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
