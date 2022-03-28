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

func max(x,y int) int {
  if x > y {
    return x
  }
  return y
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	K := readi()
  m := make(map[int]int)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		C[i] = readi()
    if i < K {
      m[C[i]]++
    }
	}
	ans := len(m) 

  for i:= 1; i < N-K+1; i++ {
    m[C[K-1+i]]++
    m[C[i-1]]--
    if m[C[i-1]] == 0 {
      delete(m,C[i-1])
    }
    ans = max(ans, len(m))
    if ans == K {
      break
    }
  }

	fmt.Println(ans)
}
