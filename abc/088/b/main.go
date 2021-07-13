// "sort" を使ってもう一度挑戦したい。
package main

import (
	"fmt"
)

type Node struct {
	l *Node
	r *Node
	v int
}

func main() {
	var n, a, b int
	fmt.Scanf("%d", &n)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v[i])
	}

	btree := &Node{v: v[0]}
	for i := 1; i < n; i++ {
		makebtree(btree, v[i])
	}
	sorted := getbtree(btree)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a += sorted[i]
		} else {
			b += sorted[i]
		}
	}

	fmt.Println(a - b)
}

func makebtree(n *Node, value int) {
	if n.v < value {
		if n.l != nil {
			makebtree(n.l, value)
		} else {
			n.l = &Node{v: value}
		}
	} else {
		if n.r != nil {
			makebtree(n.r, value)
		} else {
			n.r = &Node{v: value}
		}
	}
	return
}

func getbtree(n *Node) []int {
	var ret []int
	if n.l != nil {
		ret = append(ret, getbtree(n.l)...)
	}
	ret = append(ret, n.v)
	if n.r != nil {
		ret = append(ret, getbtree(n.r)...)
	}
	return ret
}
