package main

import (
	"fmt"
)

func main() {
	var input string
	var o, q int
	fmt.Scanf("%v", &input)

	for i := 0; i < 10; i++ {
		if input[i] == 'o' {
			o++
		} else if input[i] == '?' {
			q++
		}
	}
	// oの個数から明らかな場合
	if o > 4 || o+q == 0 {
		fmt.Println(0)
		return
	} else if o == 4 {
		fmt.Println(4 * 3 * 2 * 1)
		return
	}
	// 10000ケース総当りで条件を満たすかどうか確認する
	count := 0
	for i := 0; i < 10000; i++ {
		e := make([]int, 10)
		now := i
		// 暗証番号を下1桁から順に記録する
		for j := 0; j < 4; j++ {
			e[now%10]++
			now /= 10
		}
		// 記録とinputの条件を1件ずつ照合する
		ok := true
		for j, v := range e {
			if (v == 0 && input[j] == 'o') || (v > 0 && input[j] == 'x') {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}
	fmt.Println(count)
}
