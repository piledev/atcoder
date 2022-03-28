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
	N := readi()
	sum := 0
	i := 0

	for sum < N {
		sum += i
		i++
	}
	ans := i - 1
	fmt.Println(ans)
}
