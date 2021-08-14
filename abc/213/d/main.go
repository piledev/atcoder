package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var m map[int][]int
var visited []bool

// var ans string
var ans []int

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readi()
	m = make(map[int][]int, N)
	visited = make([]bool, N)
	for i := 0; i < N-1; i++ {
		a := readi()
		b := readi()
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}

	for i := range m {
		sort.Ints(m[i])
	}

	dfs(1)

	// ans = strings.TrimSpace(ans)
	// fmt.Println(ans)
	fmt.Println(strings.Trim(fmt.Sprint(ans), "[]"))
}

func dfs(now int) {
	// ans += strconv.Itoa(now) + " "
	ans = append(ans, now)
	visited[now-1] = true

	for i := range m[now] {
		if !visited[m[now][i]-1] {
			dfs(m[now][i])
			// ans += strconv.Itoa(now) + " "
			ans = append(ans, now)
		}
	}
}
