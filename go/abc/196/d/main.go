package main

import (
	"fmt"
)

var ans, H, W, A int

func main() {
	var B int
	fmt.Scanf("%d %d %d %d", &H, &W, &A, &B)

	dp := make([][]int, H+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	rec(dp, 1, 1, 0)
	fmt.Println(ans)
}

// ac: 2*1パネルの累積枚数
func rec(dp [][]int, h, w, ac int) {
	if ac == A {
		fmt.Println("ans++:")
		for _, v := range dp {
			fmt.Println("      ", v)
		}
		ans++
		return
	}

	// 縦置きできるか
	if h < H && dp[h-1][w] != 1 && dp[h][w-1] != 2 {
		dp[h][w] = 1
		ac++
		fmt.Println("tate:", h, w)
		for _, v := range dp {
			fmt.Println("      ", v)
		}
		nexth, nextw, incremented := nextIndex(h, w)
		if incremented {
			rec(dp, nexth, nextw, ac)
		}
	}

	// 横置きできるか
	if w < W && dp[h-1][w] != 1 && dp[h][w-1] != 2 {
		dp[h][w] = 2
		ac++
		fmt.Println("yoko:", h, w)
		for _, v := range dp {
			fmt.Println("      ", v)
		}
		nexth, nextw, incremented := nextIndex(h, w)
		if incremented {
			rec(dp, nexth, nextw, ac)
		}
	}

	// 置かない
	dp[h][w] = 0
	nexth, nextw, incremented := nextIndex(h, w)
	if incremented {
		rec(dp, nexth, nextw, ac)
	}

}

func nextIndex(h, w int) (int, int, bool) {
	res := true
	if w < W {
		w++
	} else if h < H {
		w = 1
		h++
	} else {
		res = false
	}
	return h, w, res
}
