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

func main() {
	const INF int = math.MaxInt32
	sc.Split(bufio.ScanWords)
	N := readi()
	Q := readi()

	// m: 親と子（＝直接つながっている街）の記録。木にあたるもの。
	// 領域を確保しているのは都度のallocateを防ぐため。
	// 入力1行につき街は2つなので、N*2を確保する。
	m := make(map[int][]int, N*2)

	for i := 0; i < N-1; i++ {
		a := readi()
		b := readi()
		// aを親とするケースとbを親とするケースの療法を登録する
		m[a-1] = append(m[a-1], b-1)
		m[b-1] = append(m[b-1], a-1)
	}

	// q: キュー。
	// 子を追加していく
	q := make([]int, 0, N)

	// color: 色のメモ。
	// color[i]: 街(i+1)の色をメモする.
	// 色は起点となる街からの距離が偶数の街と奇数の街で塗り分ける.
	// 0: 起点となる街からの距離が偶数のとき
	// 1: 起点となる街からの距離が奇数のとき
	// -1: 初期値(未着色の状態)
	color := make([]int, N, N)
	for i := 0; i < N; i++ {
		color[i] = -1
	}

	// queueにはとりあえず0だけ入れておく。
	// 後のイテレートで0(=1の街)から処理が開始される。
	// 結果として街1をrootとする木について考えることになる。
	q = append(q, 0)
	// 街1の色を0とする。
	// 街1からの距離が奇数の場合は1,偶数の場合は0になる。
	color[0] = 0

	for len(q) != 0 {
		// queueの先頭の値を取り出し、空いた先頭を詰める。
		t := q[0]
		q = q[1:]
		// 街t+1の宛先ごとに色を確認する。
		// 未着色の場合のみ、tの色と逆の色をつける。
		// 一度ついた色が変ることは考える必要がない。
		// BFSで処理を行っているので初回訪問時の着色が最短距離の色ということになるからだ。
		for _, v := range m[t] {
			if color[v] == -1 {
				// color[v] = 1 - color[t]
				color[v] = 1 ^ color[t]
				q = append(q, v)
			}
		}
	}

	for i := 0; i < Q; i++ {
		c := readi()
		d := readi()
		if color[c-1] == color[d-1] {
			fmt.Println("Town")
		} else {
			fmt.Println("Road")
		}
	}
}
