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
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		m[readi()]++
	}
  1 2 3 1 1 2     11 
  13 22 31
  ans :=  
  for v := range m {

  }
	ans := 0

	fmt.Println(ans)
}
