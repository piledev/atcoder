package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var required []bool
var tree [][]int

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{},math.MaxInt64)
	N := readi()
	T := make([]int64, N)
	//それぞれの技に必要な技の習得
	tree = make([][]int, N)
	//Nに必要なすべての技の列挙
	required = make([]bool, N)
	for i := range T {
		T[i] = readi64()
		K := readi()
		A := make([]int, K)
		for j := range A {
			A[j] = readi() - 1
		}
		tree[i] = A
	}
	// fmt.Println("tree: ", tree)

	addrec(N - 1)

	// fmt.Println("required: ", required)
	ans := int64(0)
	for i := range required {
		if required[i] {
			ans += T[i]
		}
	}

	fmt.Println(ans)
}

func addrec(index int) {
	// fmt.Println("required[", index, "]: ", required[index])
	if required[index] {
		// fmt.Println("return")
		return
	}
	required[index] = true
	// fmt.Println("required[", index, "]: ", required[index])
	for i := range tree[index] {
		// fmt.Println("index:", index, " i:", i, " tree[index]:", tree[index])
		addrec(tree[index][i])
	}
}

// my functions
func reads() string  { sc.Scan(); return sc.Text() }
func readi() int     { res, _ := strconv.Atoi(reads()); return res }
func readi64() int64 { res, _ := strconv.Atoi(reads()); return int64(res) }
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
