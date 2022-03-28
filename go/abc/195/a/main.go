// 5min
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	M := readi()
	H := readi()

	ans := "No"
	if H%M == 0 {
		ans = "Yes"
	}
	fmt.Println(ans)
}
