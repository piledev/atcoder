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

func readInt() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func readString() string {
	sc.Scan()
	res := sc.Text()
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readInt()
	S := readString()
	s := strings.Split(S, "")
	Q := readInt()
	T := make([]int, Q)
	A := make([]int, Q)
	B := make([]int, Q)
	for i := 0; i < Q; i++ {
		T[i] = readInt()
		A[i] = readInt()
		B[i] = readInt()
	}

	inversion := false
	for i := 0; i < Q; i++ {
		switch T[i] {
		case 1:
			var small, big int
			if inversion {
				if A[i] > N {
					small = A[i] - 1 - N
				} else {
					small = A[i] - 1 + N
				}
				if B[i] > N {
					big = B[i] - 1 - N
				} else {
					big = B[i] - 1 + N
				}
			} else {
				small = A[i] - 1
				big = B[i] - 1
			}
			s[small], s[big] = s[big], s[small]
		case 2:
			inversion = !inversion
		}
	}
	if inversion {
		S = strings.Join(s[N:], "") + strings.Join(s[:N], "")
	} else {
		S = strings.Join(s, "")
	}
	s = strings.Split(S, "")
	fmt.Println(strings.Join(s, ""))
}
