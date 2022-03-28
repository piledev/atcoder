package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	l := readi()
	r := readi()
	s := reads()

	S := s[l-1 : r]

	mid := make([]byte, 0)
	for i := len(S); i > 0; i-- {
		mid = append(mid, S[i-1])
	}

	left := make([]byte, 0)
	if l > 1 {
		for i := range s[:l-1] {
			left = append(left, s[i])
		}
	}

	right := make([]byte, 0)
	if r < len(s) {
		for i := r; i < len(s); i++ {
			right = append(right, s[i])
		}
	}
	ans := strings.Trim(string(left)+string(mid)+string(right), " ")
	fmt.Println(ans)
}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readl() int64   { res, _ := strconv.ParseInt(reads(), 10, 64); return res }
func readf() float64 { res, _ := strconv.ParseFloat(reads(), 64); return res }
