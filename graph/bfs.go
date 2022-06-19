package graph

type bfsWithQ struct {
	que     Que
	graph   map[int][]int
	visited []bool

	order []int
	cycle int // 有向圖的cycle 和 開始的點位有關
}

func intBfs(g map[int][]int) bfsWithQ {
	return bfsWithQ{
		que:     make(Que, 0),
		graph:   g,
		visited: make([]bool, len(g)),
		order:   make([]int, 0),
	}
}

func (bq *bfsWithQ) startBfs() {
	for i := range bq.graph {
		bq.bfs(i)
	}
}

func (bq *bfsWithQ) bfs(start int) {
	if bq.visited[start] {
		return
	}
	bq.cycle++
	bq.visited[start] = true
	bq.que.Insert(start)
	for {
		at, ok := bq.que.Pop()
		if !ok {
			break
		}

		bq.doThing(at) // do something about this node
		for _, to := range bq.graph[at] {
			if !bq.visited[to] {
				bq.visited[to] = true
				bq.que.Insert(to)
			}
		}
	}
}

func (bq *bfsWithQ) doThing(at int) {
	bq.order = append(bq.order, at)

}
