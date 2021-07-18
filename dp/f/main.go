package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func rs() string{
	sc.Scan()
  return sc.Text()
}

func ri() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	s := rs()
	t := rs()
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

  lens := len(s)
  lent := len(t)

	for i := 0; i < lens; i++ {
		for j := 0; j < lent; j++ {
      // sのi番目の文字とtのj番目の文字が一致するなら、
      // sのi-1番目の文字とtのj-1番目の文字までの一致数に1を加えれば良い
      // そうでない場合は
      // 一つ前のLCSの最大値を継承する。(値は変わらない。)
      if s[i] == t[j] {
        dp[i+1][j+1] = dp[i][j]+1
      } else {
        dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
      }
		}
	}

  fmt.Println("dp:::::::::::")
  for _,v := range dp {
    fmt.Println(v)
  }
  fmt.Println(":::::::::::::")

  lenans := dp[lens][lent]
  i := len

  
}
