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
	Q := readi()
	obv := make(map[int]int, N)
	rev := make(map[int]int, N)
	for i := 0; i < Q; i++ {
		c := readi()
		switch c {
		case 1:
			x := readi()
			y := readi()
			obv[x] = y
			rev[y] = x
		case 2:
			x := readi()
			y := readi()
			obv[x] = 0
			rev[y] = 0
		case 3:
			x := readi()
			top := 0
			for {
				if rev[x] == 0 {
					top = x
					break
				}
				x = rev[x]
			}

			M := 0
			key := top
			keys := make([]int, N)

			for i := 0; i < N; i++ {
				keys[i] = key
				key = obv[key]
				if key == 0 {
					M = i + 1
					// str := strings.Join(keys[:M], " ")
					// fmt.Printf("%d %s\n", M, str)
					fmt.Print(M)
					for j := 0; j < M; j++ {
						fmt.Printf(" %d", keys[j])
					}
					fmt.Print("\n")
					break
				}
			}
		}
	}
}

// my functions
func reads() string { sc.Scan(); return sc.Text() }
func readi() int    { res, _ := strconv.Atoi(reads()); return res }
