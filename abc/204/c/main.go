// DFS(深さ優先探索)の問題
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
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	ans := 0
	N := readInt()
	M := readInt()
	can := make([][]int, N)
	for i := 0; i < N; i++ {
		// can[i]: i+1 の都市から直接行ける都市の slice
		can[i] = make([]int, 0)
	}
	for i := 0; i < M; i++ {
		a := readInt()
		b := readInt()
		a--
		b--
		can[a] = append(can[a], b)
	}
	// seen: すでに探索した都市のメモ。初期化して使い回す。
	seen := make([]bool, N)
	var dfs func(int)
	dfs = func(i int) {
		if seen[i] {
			return
		}
		seen[i] = true
		for _, v := range can[i] {
			dfs(v)
		}
		return
	}

	for i := range can {
		// seen の初期化
		for j := range seen {
			seen[j] = false
		}
		dfs(i)
		// seen の true をカウント
		for _, v := range seen {
			if v {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
