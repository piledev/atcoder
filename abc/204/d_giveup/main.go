// hand_01.txt だけがどうしてもWAになる...
// テストデータが公開されたらチェックしてみよう。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "math"
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

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
  ans := math.MaxInt64
	total := 0
  T := make([]int,N)
	for i := range T {
		T[i] = readInt()
		total += T[i]
	}

  // dp の作成
  // index は合計時間
  // d[index] に格納するのはindex分ピッタリの組み合わせの有無
  dp := make([]bool, total+1)
  dp[0] = true
  // 料理を一つずつチェックしていく
  for i := 0;i < N; i++ {
    // possible_minutes : 組み合わせ可能な分 
    possible_minutes := make([]int,0)
    // dp(メモ)を一つずつチェックする
    for j, b := range dp {
      // j 分ピッタリの組み合わせがあるならば
      // j 分+(今チェックしているi 番目の料理にかかる時間)もあり得るので、
      // その合計時間を組み合わせ可能な時間リストに追加しておく
      if b {
        possible_minutes = append(possible_minutes,j+T[i])
      }
    }
    // 組み合わせ可能な時間リストに記載のある分数すべてについてtrueをメモしていく
    for _,v := range possible_minutes {
      dp[v] = true
    }
  }

  // 目標値の設定(オーブン2つを使った場合の最短時間はtotalの半分になるはず)
  p := total/2 + total%2
  // dp より目標値以上の適切な値を探す
  for i := p; i<total; i++ {
    if dp[i] {
      ans = i
      break
    }
  }

	fmt.Println(ans)
}
