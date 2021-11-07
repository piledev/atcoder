package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	M := readi()
	B := make([][]int, N)
	ans := "Yes"

	for i := range B {
		B[i] = make([]int, M)
		for j := range B[i] {
			B[i][j] = readi()

			if i == 0 {
				if j == 0 {
					continue
				} else {
					if B[i][j] != B[i][j-1]+1 {
						ans := "No"
						fmt.Println(ans)
						return
					}
					if B[i][j]%7 == 1 {
						ans := "No"
						fmt.Println(ans)
						return
					}
				}
			} else {
				if B[i][j] != B[i-1][j]+7 {
					ans := "No"
					fmt.Println(ans)
					return
				}
			}
		}
	}

	fmt.Println(ans)
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
