// あとまわし
// dp を利用した解法
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var (
	as      [][]int
	c       int
	ans     int
	mincost int
)

func readi() int {
	sc.Scan()
	res, _ := strconv.Atoi(sc.Text())
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func main() {

	sc.Split(bufio.ScanWords)
	h, w := readi(), readi()
	c = readi()
	as = make([][]int, h)
	mincost = math.MaxInt64
	for i := range as {
		as[i] = make([]int, w)
		for j := range as[i] {
			as[i][j] = readi()
			if as[i][j] < mincost {
				mincost = as[i][j]
			}
		}
	}

	// 初期値
	ans = as[0][0] + as[0][1] + c
	for i, row := range as {
		for j := range row {
			fmt.Println("start from ", i, j, " --------------------------")
			dfs(i, j, i, j)
		}
	}
	fmt.Println(ans)
}

func dfs(x1, y1, x2, y2 int) {
	dis := abs(x2-x1) + abs(y2-y1)
	// 処理 //////////////////////////////////////////////////////////////////////
	// 出発駅と到着駅が一致しないとき(dis>0)だけコストを計算する
	if dis > 0 {
		fmt.Printf("to(%d, %d)", x2, y2)
		// 到着駅の建設コストが最小だったとしても、
		// 総コストが最小となりえない場合にはこれ以上計算する必要はない。
		if ans <= c*dis+as[x1][y1]+mincost {
			fmt.Println(" > x")
			return
		}

		// コストを計算し、最小の場合は ans に代入する
		cost := as[x1][y1] + as[x2][y2] + dis*c
		fmt.Println(" cost:", cost)
		if ans > cost {
			ans = cost
		}
	}

	// 移動 //////////////////////////////////////////////////////////////////////
	// 上以外に移動することはない

	// 到着駅のほうが右にあり、右端ではないとき、右に移動する。
	if y1 <= y2 && y2+1 < len(as[0]) {
		fmt.Printf(" > ")
		dfs(x1, y1, x2, y2+1)
	}
	// 行が一致しなくて、到着駅のほうが左側にあり、左端ではないとき、左に移動する。
	// つまり同じ行のときは左側には移動しない。
	// また、出発駅から左側にあるときに限定しているので、"><"みたいな動き方はしない。
	if x1 != x2 && y1 >= y2 && y2 > 0 {
		fmt.Printf(" < ")
		dfs(x1, y1, x2, y2-1)
	}
	// 列が一致して、到着駅が下の端ではないとき、下に移動する
	// 出発駅の真下にしか移動しないので、">v"みたいな動き方はしない。
	if y1 == y2 && x2+1 < len(as) {
		fmt.Printf(" v ")
		dfs(x1, y1, x2+1, y2)
	}
}
