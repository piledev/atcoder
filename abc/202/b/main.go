package main

import (
	"fmt"
	"strings"
)

// var sc = bufio.NewScanner(os.Stdin)

// func readStr() string {
// 	sc.Scan()
// 	ret := sc.Text()
// 	return ret
// }

func main() {
	// sc.Split(bufio.ScanWords)
	// s := readStr()
	var s string
	fmt.Scanf("%v", &s)
	a := strings.Split(s, "")
	l := len(a)
	ans := make([]string, l)
	for i := 0; i < l; i++ {
		switch a[i] {
		case "0", "1", "8":
			ans[l-i-1] = a[i]
		case "6":
			ans[l-i-1] = "9"
		case "9":
			ans[l-i-1] = "6"
		}
	}
	fmt.Println(strings.Join(ans, ""))
}
