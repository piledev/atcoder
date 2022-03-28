package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readi() int {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	H := readi()
	W := readi()
	X := readi()
	Y := readi()
	S := make([]string, H)
	for i := 0; i < H; i++ {
		S[i] = reads()
	}

	ans := 1
	for i := X - 2; i >= 0; i-- {
		if S[i][Y-1] == '#' {
			break
		}
		ans++
	}
	for i := Y - 2; i >= 0; i-- {
		if S[X-1][i] == '#' {
			break
		}
		ans++
	}
	for i := X; i < H; i++ {
		if S[i][Y-1] == '#' {
			break
		}
		ans++
	}
	for i := Y; i < W; i++ {
		if S[X-1][i] == '#' {
			break
		}
		ans++
	}

	fmt.Println(ans)

}
