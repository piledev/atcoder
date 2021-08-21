package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func reads() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := reads()
	m := make(map[string]bool, 0)
	n := len(s)

	for i := 0; i < (1 << n); i++ {
		// fmt.Println("")
		// fmt.Println("========== i:", i, "==========")
		str := ""
		pre := false
		for j := 0; j < n; j++ {
			// fmt.Println("j:", j, s[j], pre)

			if pre == false {
				// fmt.Println("i:", i, ", 1<<j:", (1 << j), ", i & 1<<j: ", i&1<<j)

				if (i>>j)&1 == 1 {
					// fmt.Println("str:", str, "s[j]:", string(s[j]))
					str = str + string(s[j])
					pre = true
				}
			} else {
				pre = false
			}
		}
		if len(str) > 0 {
			m[str] = true
		}
	}

	fmt.Println(m)
	fmt.Println(len(m))

}
