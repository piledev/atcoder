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
	// sc.Buffer(math.MaxInt64)
	A := float64(readi())
	B := float64(readi())

	ans := A * B / 100
	fmt.Println(ans)
}
