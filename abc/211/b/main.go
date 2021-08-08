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
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := make([]string, 4)
	m := make(map[string]int)
	for i := 0; i < 4; i++ {
		s[i] = reads()
		m[s[i]]++
	}

	ans := "No"
	if len(m) == 4 {
		ans = "Yes"
	}

	fmt.Println(ans)
}
