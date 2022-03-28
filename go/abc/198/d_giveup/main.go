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

func reads() string {
	sc.Scan()
	res := sc.Text()
	return res
}

func main() {
	var ans string
	sc.Split(bufio.ScanWords)
	S := make([]string, 3)
	for i := 0; i < 3; i++ {
		S[i] = reads()
	}

	for i := 0; i < 3; i++ {
		if len(S[i]) > 10 {

		}
	}

	fmt.Println(ans)

}
