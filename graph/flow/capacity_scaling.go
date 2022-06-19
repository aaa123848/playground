package flow

type CapacityScaling struct {
	visited []bool

	capacity [][]int

	tryFlow int
}

func (cs *CapacityScaling) refreshVisited() {
	cs.visited = make([]bool, len(cs.visited))
	return
}

func InitCapacityScaling(graph [][]int) CapacityScaling {
	visited := make([]bool, len(graph))
	return CapacityScaling{
		visited:  visited,
		capacity: graph,
	}
}

func (cs *CapacityScaling) findMaxCapacity() int {
	mc := 0
	for i := range cs.capacity {
		for j := range cs.capacity[i] {
			mc = max(mc, cs.capacity[i][j])
		}
	}
	return mc
}

func (cs *CapacityScaling) initTryFlow() {
	maxFlow := cs.findMaxCapacity()
	res := 1
	for {
		newRes := res * 2
		if newRes > maxFlow {
			cs.tryFlow = res
			return
		}
		res = newRes
	}
}

func (cs *CapacityScaling) Process(start, end int) int {
	cs.initTryFlow()
	maxFlow := 0
	for {
		if cs.tryFlow < 1 {
			break
		}
		cs.refreshVisited()
		ok := cs.dfs(start, end)
		if !ok {
			cs.tryFlow /= 2
			continue
		}
		maxFlow += cs.tryFlow
	}
	return maxFlow
}

func (cs *CapacityScaling) dfs(at, end int) bool {
	cs.visited[at] = true
	if at == end {
		return true
	}
	for to, cap := range cs.capacity[at] {
		if !cs.visited[to] && cap >= cs.tryFlow {
			ok := cs.dfs(to, end)
			if ok {
				cs.capacity[at][to] -= cs.tryFlow
				cs.capacity[to][at] += cs.tryFlow
				return ok
			}
		}
	}
	return false
}
