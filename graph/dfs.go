package graph

type dfsWithStack struct {
	stack   Stack
	visited []bool
	graph   map[int][]int

	order []int
}

func initDfsWithStack(graph map[int][]int) dfsWithStack {
	return dfsWithStack{
		stack:   make(Stack, 0),
		visited: make([]bool, len(graph)),
		graph:   graph,

		order: make([]int, 0),
	}
}

func (dws *dfsWithStack) process() {
	for start := range dws.graph {
		dws.dfs(start)
	}
}

// need test
func (dws *dfsWithStack) dfs(start int) {
	if dws.visited[start] {
		return
	}
	dws.stack.Insert(start)
	for {
		at, ok := dws.stack.Pop()
		if !ok {
			break
		}
		dws.visited[at] = true
		dws.aProcess(at)

		for _, to := range dws.graph[at] {
			if !dws.visited[to] {
				dws.visited[to] = true
				dws.stack.Insert(to)
			}
		}
	}
}

func (dws *dfsWithStack) aProcess(at int) {
	dws.order = append(dws.order, at)
}
