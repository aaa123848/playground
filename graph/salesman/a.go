package salesman

import (
	g "gotest2/graph"
	"log"
	"math"
)

const (
	inf int = g.Inf
)

type memoType [][]int

func makeMemo(num int) memoType {
	memo := make(memoType, num)

	arrLen := int(math.Pow(float64(num), 2))

	for i := range memo {
		for n := 0; n < arrLen; n++ {
			memo[i] = append(memo[i], inf)
		}
	}
	return memo
}

func tsp(m [][]int, start int) {
	n := len(m)
	memo := makeMemo(n)
	setUp(m, memo, start, n)

	log.Println(memo)

}

func setUp(m [][]int, memo memoType, s, n int) {
	for i := 0; i < n; i++ {
		if i == s {
			continue
		}
		memo[i][uint(1)<<s|uint(1)<<i] = m[s][i]
	}
}
