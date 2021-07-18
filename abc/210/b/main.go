package main

import (
	"bufio"
	"fmt"
	"os"
  "math"
	"strconv"
	"strings"
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
	sc.Buffer([]byte{},math.MaxInt64)
	_ = readi()
	S := reads()
	ans := "Takahashi" 

  tmp := strings.Split(S,"1")
  m := len(tmp[0])%2 
  if m == 1 {
    ans = "Aoki"
  }
  fmt.Println(ans)
}
