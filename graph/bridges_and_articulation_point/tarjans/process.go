package tarjans

import (
	g "gotest2/graph"
)

type Tarjans struct {
	stack   g.Stack
	visited []bool
	ids     []int
	low     []int
	isStack []bool
	graph   map[int][]int

	result [][]int

	nowID int
}

func initTarjans(graph map[int][]int) Tarjans {
	return Tarjans{
		stack:   g.InitStack(),
		visited: make([]bool, len(graph)),
		ids:     make([]int, len(graph)),
		low:     make([]int, len(graph)),
		isStack: make([]bool, len(graph)),
		graph:   graph,
		result:  make([][]int, 0),
		nowID:   0,
	}
}

func (t *Tarjans) Process(graph map[int][]int) {
	*t = initTarjans(graph)
	for n := range t.graph {
		if t.visited[n] {
			continue
		}
		t.dfs(n)
	}
}

func (t *Tarjans) dfs(at int) {
	if t.visited[at] {
		return
	}
	t.ids[at] = t.nowID
	t.low[at] = t.nowID
	t.visited[at] = true
	t.isStack[at] = true
	t.stack.Insert(at)
	t.nowID++

	for _, to := range t.graph[at] {

		t.dfs(to)
		if t.isStack[to] {
			t.low[at] = g.Min(t.low[at], t.low[to])
		}
	}

	if t.low[at] != t.ids[at] {
		return
	}

	tmp := make([]int, 0)

	for {
		v, ok := t.stack.Pop()
		if !ok {
			break
		}
		t.isStack[v] = false
		tmp = append(tmp, v)
		if v == at {
			t.result = append(t.result, tmp)
			break
		}
	}

}
