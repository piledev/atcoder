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
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	A := readi()
	X := readi()
	Y := readi()
	ans := 0

  if N > A {
    ans = X * A + Y * (N-A)
  } else {
    ans = X*N
  }

	fmt.Println(ans)
}
