package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	S := reads()

	ans := "WA"
	if S == "Hello,World!" {
		ans = "AC"
	}

	fmt.Println(ans)
}
