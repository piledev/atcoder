// 0633-0655
package main

import "fmt"

func main() {
	var n, x, y int
	var ax, ay []int
	var rec map[int]int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// 単独で200の倍数があったら即成立
	// 同じ数字があったら即成立
	// 組み合わせを網羅して、それぞれの200で割ったあまりをmap[int]int に記録する
	// map[x] > 1 になった時点で成立する
	// map[x] > 1 になったとき、x の組み合わせはどうやって思い出す？

	if total(a, ax) == total(a, ay) {
		fmt.Println("Yes")
		fmt.Println(x, ax...)
		fmt.Println(y, ay...)
	}

}

func total(a []int, x []int) int {
	ret := 0
	for i := 0; i < x; i++ {
		ret += a[x[i]]
	}
	return ret
}
