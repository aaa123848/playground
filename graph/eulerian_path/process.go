package eulerian

type EullerianPath struct {
	g    [][]int
	path []int
	in   []int
	out  []int
}

func InitEullerianPath(graph [][]int) EullerianPath {
	out := make([]int, len(graph))
	in := make([]int, len(graph))
	for from := range graph {
		for _, to := range graph[from] {
			out[from]++
			in[to]++
		}
	}

	return EullerianPath{
		g:   graph,
		in:  in,
		out: out,
	}
}

func (e *EullerianPath) IsEulerianPath() (yes bool, start, end int) {
	var isCycle bool = true
	start, end = -1, -1
	evenCount := 0
	startCount := 0
	endCount := 0
	nodeNum := len(e.g)
	for i := 0; i < nodeNum; i++ {
		isCycle = isCycle && e.in[i] == e.out[i]

		if e.in[i] == e.out[i] {
			evenCount++
		}

		if e.out[i]-e.in[i] == 1 {
			startCount++
			start = i
		}
		if e.in[i]-e.out[i] == 1 {
			endCount++
			end = i
		}
	}

	if isCycle {
		return true, -1, -1
	}

	return startCount == 1 && endCount == 1 && evenCount == nodeNum-2, start, end

}

func (e *EullerianPath) FindEulerianPath(start int) {

	e.dfs(start)

	res := make([]int, 0)
	for i := len(e.path) - 1; i >= 0; i-- {
		res = append(res, e.path[i])
	}

	e.path = res
}

func (e *EullerianPath) dfs(at int) {
	for {
		outNum := e.out[at]
		if outNum < 1 {
			break
		}
		to := e.g[at][outNum-1]
		e.out[at]--
		e.dfs(to)
	}
	e.path = append(e.path, at)
}
