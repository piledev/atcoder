package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64
	fmt.Scanf("%g", &N)
	dp := make([]bool, int(N))
	ans := int(N)

	for a := float64(2); a <= math.Sqrt(N); a++ {
		if dp[int(a)-1] {
			continue
		}
		for b := float64(2); b <= N; b++ {
			v := math.Pow(a, b)
			if v > N {
				break
			}
			if dp[int(v)-1] {
				break
			}
			dp[int(v)-1] = true
			ans--
		}
	}
	fmt.Println(ans)
}
