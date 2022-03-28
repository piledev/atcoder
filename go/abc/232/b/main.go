package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)

	S := reads()
	T := reads()

	var K int
	K = int(T[0]) - int(S[0])
	if K < 0 {
		K += 26
	}
	if K == 0 {
		if S == T {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}
	for i := range S {
		t := int(T[i])
		s := int(S[i])
		if K != t-s && K != t-s+26 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int    { res, _ := strconv.Atoi(reads()); return res }
