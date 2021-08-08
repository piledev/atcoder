package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var queue []int
var N, M, ans int
var visited []bool
var m map[int][]int
var fixed bool

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N = readi()
	M = readi()
	m = make(map[int][]int)
	queue = make([]int, 0, N)
	for i := 0; i < M; i++ {
		a := readi()
		b := readi()
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}

	visited = make([]bool, N)
	queue = append(queue, 1)
	for len(queue) > 0 {
		bfs(queue[0])
		queue = queue[1:]
	}

	fmt.Println(ans)
}

func bfs(source int) {
	// 終着都市の場合
	if source == N {
		fixed = true
		ans = getAns(ans + 1)
		return
	}
	// すでに訪問済みの場合
	if visited[source-1] {
		return
	}

	if !fixed {
		// 未訪問の場合、訪問可能な都市をキューに設定する
		for i := range m[source] {
			queue = append(queue, m[source][i])
		}
	}
	return
}

func getAns(i int) int {
	return i % 1_000_000_007
}
