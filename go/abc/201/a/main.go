//26 minutes
package main

import (
	"fmt"
)

func main() {
	v := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &v[i])
	}

	s1 := abs(v[0] - v[1])
	s2 := abs(v[1] - v[2])
	s3 := abs(v[2] - v[0])

	ans := "No"
	switch {
	case s1 == s2 && s1+s2 == s3:
		ans = "Yes"
	case s1 == s3 && s1+s3 == s2:
		ans = "Yes"
	case s2 == s3 && s2+s3 == s1:
		ans = "Yes"
	}
	fmt.Println(ans)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
