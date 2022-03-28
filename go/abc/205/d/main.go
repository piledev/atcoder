package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
	N := readi()
	Q := readi()
	A := make([]int, N)
	// 含まれてほしくない整数
	for i := range A {
		A[i] = readi()
	}

	// クエリ(Aに含まれないK[i]番目の数)
	K := make([]int, Q)
	for i := range K {
		K[i] = readi()
	}

	// A[i]以下のAに含まれな整数の数
	C := make([]int, N)
	for i := 0; i < N; i++ {
		C[i] = A[i] - 1 - i
	}

	// すべてのKについて判定する
	for _, v := range K {
		ans := 0
		// C[N-1]よりも大きい場合は
		if C[N-1] < v {
			ans = A[N-1] + (v - C[N-1])
		} else {
			// C[i]>=v となるような最小の i (v番目の値はA[i-1]とA[i]の間にあることがわかる)
			i := sort.Search(len(C), func(i int) bool { return C[i] >= v })
			// A[i]-1: A[i]以下最大の除かない値
			// C[i]-v: C[i]とv番目の値の差
			// ans: A[i]以下最大の除かない値からC[i]とv番目の値の差だけ戻った値
			ans = (A[i] - 1) - (C[i] - v)
		}
		fmt.Println(ans)
	}

}
