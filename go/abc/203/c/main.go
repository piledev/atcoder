package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	N := readInt()
	K := readInt()
	g := make(map[int]int)
	for i := 0; i < N; i++ {
		g[readInt()] += readInt()
	}
	keys := make([]int, len(g))
	i := 0
	for key := range g {
		keys[i] = key
		i++
	}
	sort.Ints(keys)

	fmt.Println(run(K, g, keys))
}

func run(K int, g map[int]int, keys []int) int {
	res := K
	for _, k := range keys {
		if k <= res {
			res += g[k]
		} else {
			break
		}
	}
	return res
}
