package floyd

import (
	g "gotest2/graph"
)

// const inf int = math.MaxInt
const inf int = g.Inf
const negIng int = g.NegInf

func createGraph(nodNum int) (res [][]int) {
	return g.CreateGraphMetrix(nodNum)
}

func min(i, j int) int {
	return g.Min(i, j)
}

func copyMetrix(m [][]int) (res [][]int) {
	nodeNum := len(m)

	for t := 0; t < nodeNum; t++ {
		res = append(res, make([]int, nodeNum))
	}

	for i := range res {
		for j := range res[i] {
			res[i][j] = m[i][j]
		}
	}

	return res
}

func Floyd(m [][]int) (res [][]int) {
	nodeNum := len(m)

	res = copyMetrix(m)

	for m := 0; m < nodeNum; m++ {
		for i := 0; i < nodeNum; i++ {
			for j := 0; j < nodeNum; j++ {
				if res[i][m] == inf || res[m][j] == inf {
					continue
				}
				res[i][j] = min(res[i][j], res[i][m]+res[m][j])
			}
		}
	}

	findNegativeCycleAfterFloyd(res, nodeNum)

	return res

}

func findNegativeCycleAfterFloyd(dp [][]int, n int) {

	arr := make([][]int, 0)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] > dp[i][k]+dp[k][j] {
					arr = append(arr, []int{i, j})
				}
			}
		}
	}

	for _, a := range arr {
		dp[a[0]][a[1]] = negIng
	}
}
