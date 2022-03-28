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

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	abc := "abcdefghijklmnopqrstuvwxyz"
	sabc := strings.Split(abc, "")
	mabc := make(map[int]string, 0)
	mabcrev := make(map[string]int, 0)
	for i := range sabc {
		mabc[i] = sabc[i]
		mabcrev[sabc[i]] = i
	}

	x := reads()
	sx := strings.Split(x, "")
	mx := make(map[string]int, 0)
	mxrev := make(map[int]string, 0)
	for i := range sx {
		mx[sx[i]] = i
		mxrev[i] = sx[i]
	}

	N := readi()
	// 新しい名前のslice. sortするのに使用する
	newnames := make([]string, 0)
	// 新しい名前と古い名前の対応
	m := make(map[string]string, 0)

	for i := 0; i < N; i++ {
		S := reads()
		ss := strings.Split(S, "")
		newname := ""
		for j := range ss {
			newname = newname + mabc[mx[ss[j]]]
		}
		m[newname] = S
		newnames = append(newnames, newname)
	}

	sort.Strings(newnames)

	for i := range newnames {
		fmt.Println(m[newnames[i]])
	}

}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int    { res, _ := strconv.Atoi(reads()); return res }
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
