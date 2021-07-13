package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func reads() string {
	sc.Scan()
	res := sc.Text()
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer(math.MaxInt64)
	S := reads()

	fmt.Println(run(S))
}

func run(s string) string {
	return s[1:] + s[:1]
}
