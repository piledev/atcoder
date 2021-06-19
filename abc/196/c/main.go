package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	for i := 1; i <= 1000000; i++ {
		istr := strconv.Itoa(i)
		ii, _ := strconv.Atoi(istr + istr)
		if ii > N {
			fmt.Println(i - 1)
			break
		}
	}
}
