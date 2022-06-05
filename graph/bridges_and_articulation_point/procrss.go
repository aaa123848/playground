package ba

import (
	g "gotest2/graph"
)

type articulation struct {
	visited []bool
	lowest  []int
	ids     []int
	arPoint []bool

	graph   map[int][]int
	nowId   int
	outEdge int
}

func initAritculation(graph map[int][]int) articulation {
	nNum := len(graph)

	return articulation{
		visited: make([]bool, nNum),
		lowest:  make([]int, nNum),
		ids:     make([]int, nNum),
		arPoint: make([]bool, nNum),
		graph:   graph,
	}
}

func (a *articulation) Process() {
	for n := 0; n < len(a.graph); n++ {
		if a.visited[n] {
			continue
		}
		a.outEdge = 0
		a.dfs(n, n, -1)

		// 無向圖中， 若從某點出發的超過兩個, 則該點也是
		if a.outEdge > 1 {
			a.arPoint[n] = true
		}

	}
}

func (a *articulation) dfs(root, at, parent int) {

	// 應該用不到
	if a.visited[at] {
		return
	}

	a.ids[at] = a.nowId
	a.lowest[at] = a.nowId
	a.nowId++
	a.visited[at] = true
	if parent == root {
		a.outEdge++
	}

	for _, to := range a.graph[at] {
		if to == parent {
			continue
		}
		if !a.visited[to] {
			a.dfs(root, to, at)
			a.lowest[at] = g.Min(a.lowest[at], a.lowest[to])

			// !!!! 重要
			if a.lowest[to] >= a.ids[at] && at != root {
				a.arPoint[at] = true
			}
			continue
		}

		if a.visited[to] {
			a.lowest[at] = g.Min(a.lowest[at], a.ids[to])
		}
	}
}
