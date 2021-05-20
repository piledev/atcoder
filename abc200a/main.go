// 4 minites
package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)

	ans := (i-1)/100 + 1
	fmt.Println(ans)
}
