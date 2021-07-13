// 二分探索法らしいが、まだ理解できないので一旦諦める
// bit全探索でやってみたらダメだった。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func min(x , y int) int {
 if x > y {
  return y 
 } 
 return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
  ans := 0
	total := 0
  T := make([]int,N)
	for i := range T {
		T[i] = readInt()
		total += T[i]
	}
  p := total/2
  if total%2 == 1 {
    p++
  }

	for bit := 0; bit < (1 << N); bit++ {
    sum := 0
		for i,v := range T {
			if (1 & (bit>>i )) == 1 {
        sum +=v 
			}
		}
    if sum < p {
      sum = total - sum
    }
    if bit > 0 {
      ans = min(ans,sum)
    } else {
      ans = sum
    }
    if ans ==p {
      break
    }
	}
	fmt.Println(ans)
}
