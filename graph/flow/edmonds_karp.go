package flow

import (
	g "gotest2/graph"
)

type Ek struct {
	que     g.Que
	visited []bool

	capacity [][]int

	nodeNum int
}

type previousArr []int

func makePreviousArr(num int) previousArr {
	res := make([]int, num)
	for i := range res {
		res[i] = -1
	}
	return res
}

func InitEdmondsKarp(capacityGraph [][]int) Ek {
	return Ek{
		que:      make(g.Que, 0),
		visited:  make([]bool, len(capacityGraph)),
		capacity: capacityGraph,
		nodeNum:  len(capacityGraph),
	}
}

func (e *Ek) refresh() {
	e.visited = make([]bool, len(e.capacity))
	e.que = make(g.Que, 0)
}

func (e *Ek) Process(start, end int) (maxFlow int) {

	for {
		e.refresh()
		flow := e.bfs(start, end)
		maxFlow += flow
		if flow <= 0 {
			break
		}
	}
	return maxFlow

}

func (e *Ek) bfs(start, end int) (flow int) {

	pre := makePreviousArr(e.nodeNum)

	e.que.Insert(start)
	e.visited[start] = true
	// 先用 bfs 找出 augumation path, 並用 pre 紀錄 augumation path
	for {
		at, ok := e.que.Pop()
		if !ok {
			break
		}
		if at == end {
			break
		}
		for to, tCap := range e.capacity[at] {
			if tCap > 0 && !e.visited[to] {
				e.visited[to] = true
				e.que.Insert(to)
				pre[to] = at
			}
		}
	}

	// 看是否無法到達 end
	if pre[end] == -1 {
		return 0
	}

	// 求得 augumatino path 上的 bottleneck, 並把該路線上的 capacity 減掉 bottleneck 值
	bottleneck := maxInt
	for nowOn := end; pre[nowOn] != -1; nowOn = pre[nowOn] {
		bottleneck = min(bottleneck, e.capacity[pre[nowOn]][nowOn])
	}

	for nowOn := end; pre[nowOn] != -1; nowOn = pre[nowOn] {
		e.capacity[pre[nowOn]][nowOn] -= bottleneck
	}

	return bottleneck
}
