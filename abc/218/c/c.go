package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()

	S := make([][]string, N)
	smini, smaxi, sminj, smaxj, scnt := 200, -1, 200, -1, 0
	for i := 0; i < N; i++ {
		str := reads()
		S[i] = strings.Split(str, "")
		for j := 0; j < N; j++ {
			if S[i][j] == "#" {
				scnt++
				smini = min(smini, i)
				smaxi = max(smaxi, i)
				sminj = min(sminj, j)
				smaxj = max(smaxj, j)
			}
		}
	}
	sleni := smaxi - smini + 1
	slenj := smaxj - sminj + 1
	s := make([][]string, sleni)
	for i := range s {
		s[i] = make([]string, slenj)
		for j := range s[i] {
			s[i][j] = S[smini+i][sminj+j]
		}
	}

	T := make([][]string, N)
	tmini, tmaxi, tminj, tmaxj, tcnt := 200, -1, 200, -1, 0
	for i := range T {
		str := reads()
		T[i] = strings.Split(str, "")
		for j := range T[i] {
			if T[i][j] == "#" {
				tcnt++
				tmini = min(tmini, i)
				tmaxi = max(tmaxi, i)
				tminj = min(tminj, j)
				tmaxj = max(tmaxj, j)
			}
		}
	}
	tleni := tmaxi - tmini + 1
	tlenj := tmaxj - tminj + 1
	t := make([][]string, tleni)
	for i := range t {
		t[i] = make([]string, tlenj)
		for j := range t[i] {
			t[i][j] = T[tmini+i][tminj+j]
		}
	}

	// fmt.Println("scnt, tcnt: ", scnt, tcnt)
	if scnt != tcnt {
		fmt.Println("No")
		return
	} else if scnt == N*N {
		fmt.Println("Yes")
		return
	}
	// 縦横が同じ
	if sleni == tleni && slenj == tlenj {
		// 正方形
		if sleni == slenj {
			for i := 0; i < 4; i++ {
				if comp(s, t) {
					fmt.Println("Yes")
					return
				}
				if i < 3 {
					s = rotate(s)
				}
			}
			// 長方形
		} else {
			if comp(s, t) {
				fmt.Println("Yes")
				return
			}
			s = reverse(s)
			if comp(s, t) {
				fmt.Println("Yes")
				return
			}
			fmt.Println("No")
			return
		}
	} else if sleni == tlenj && slenj == tleni {
		s = rotate(s)
		if comp(s, t) {
			fmt.Println("Yes")
			return
		}
		reverse(s)
		if comp(s, t) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
	return
}

func reverse(s [][]string) [][]string {
	r := make([][]string, len(s))
	for i := range r {
		r[i] = make([]string, len(s[0]))
		for j := range r[i] {
			r[i][j] = s[len(s)-1-i][len(s[0])-1-j]
		}
	}
	return r
}

func rotate(s [][]string) [][]string {
	r := make([][]string, len(s[0]))
	for i := range r {
		r[i] = make([]string, len(s))
		for j := range r[i] {
			r[i][j] = s[j][len(s[0])-1-i]
		}
	}
	return r
}

func comp(s, t [][]string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] != t[i][j] {
				return false
			}
		}
	}
	return true
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
