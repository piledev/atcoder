package main

import (
	"bufio"
	"fmt"
  "math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func rail(C, i1, i2, j1, j2 int) int {
  return C * (abs(i1 - i2) + abs(j1 - j2))
}

func abs(i int) int{
  if i < 0 {
    return -1 * i
  }
  return i
}

func main() {
	sc.Split(bufio.ScanWords)
	H := readi()
	W := readi()
	C := readi()
	A := make([][]int, H)
  for i := range A {
    A[i] = make([]int, W)
    for j := range A[i] {
      A[i][j] = readi()
    }
  }
	ans := math.MaxInt64;
  for i1 := 0; i1 < H; i1++ {
    for j1 := 0; j1 < W; j1++ {
      i2 := 0
      j2 := 0
      if j1 == W-1 {
        i2 = i1+1 
        j2 = 0
      } else {
        i2 = i1
        j2 = j1+1
      }
      for ;i2<H; i2++ {
        for ;j2<W; j2++ {
          // fmt.Printf("(i1,j1)=(%d,%d), (i2,j2)=(%d,%d)\n",i1,j1,i2,j2)
          ans = min(ans,A[i1][j1]+A[i2][j2]+rail(C,i1,i2,j1,j2))
        }
        j2 =0
        // fmt.Println("---------------")
      }
    }
  }
	fmt.Println(ans)
}
