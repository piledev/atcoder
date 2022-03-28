package main

import (
	"fmt"
)

func main() {
	var N string
	fmt.Scanf("%s", &N)
	// s := strings.Split(N, "")
	s := N
	cnt := 0

	for i := len(s); i > 0; i-- {
		if s[i-1] != '0' {
			break
		}
		cnt++
	}

	s = s[:len(s)-cnt]

	ans := "Yes"
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}
