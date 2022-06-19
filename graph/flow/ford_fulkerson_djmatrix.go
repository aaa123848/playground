package flow

type FordFulkersonDfsAdjancyMatrix struct {
	visited      []int
	visitedToken int

	cap [][]int
}

func Initffd(cap [][]int) FordFulkersonDfsAdjancyMatrix {
	vidited := make([]int, len(cap))
	visitedTken := 1
	return FordFulkersonDfsAdjancyMatrix{
		visited:      vidited,
		visitedToken: visitedTken,
		cap:          cap,
	}
}

func (f *FordFulkersonDfsAdjancyMatrix) Process(start, end int) (maxFlow int) {

	// 一次只找一條 augmention path, 所以要找很多次, 直到沒有任何可流動的值
	for {
		flow := f.dfs(start, end, maxInt)
		if flow <= 0 {
			break
		}
		f.visitedToken++
		maxFlow += flow
	}

	return maxFlow
}

func (f *FordFulkersonDfsAdjancyMatrix) dfs(start, end, flow int) int {
	if start == end {
		return flow
	}

	tosCap := f.cap[start]
	f.visited[start] = f.visitedToken

	for to := range tosCap {
		if f.visited[to] != f.visitedToken && tosCap[to] > 0 {
			flow = min(tosCap[to], flow)
			dfsFlow := f.dfs(to, end, flow)

			// 若這個 augmenting path 的 flow 為0 則找下一條路徑
			if dfsFlow > 0 {
				f.cap[start][to] -= dfsFlow
				f.cap[to][start] += dfsFlow
				return dfsFlow // 一次只找一條 augmenting path, 所以有就 return
			}
		}
	}
	return 0

}
