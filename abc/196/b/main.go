package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Scanf("%s", &x)
	ans := strings.Split(x, ".")[0]
	fmt.Println(ans)
}
