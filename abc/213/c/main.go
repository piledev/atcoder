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
	H := readi()
	W := readi()
	N := readi()
	mh = make(map[int]bool, H)
	mw = make(map[int]bool, W)
	for i := 0; i < N-1; i++ {
		a := readi()
		b := readi()
		mh[a] = append(mh[a], true)
		mw[b] = append(mw[b], true)
	}

	for i := 0; i < H; i++ {

	}

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
