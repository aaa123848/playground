package graph

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 製造一個描述圖的 metrix, 擁有 nodeNum 個 Node, 一個 NodeNum * NodeNum 的 metrix
// 其中： 對腳皆為0, 其餘是 inf
func CreateGraphMetrix(nodNum int) (res [][]int) {
	for i := 0; i < nodNum; i++ {
		res = append(res, make([]int, nodNum))
		for j := 0; j < nodNum; j++ {
			if i == j {
				res[i][j] = 0
			} else {
				res[i][j] = Inf
			}
		}
	}
	return res
}

func CreateGraph(nodNum int) map[int][]int {

	res := make(map[int][]int)
	for k := 0; k < nodNum; k++ {
		res[k] = make([]int, 0)
	}

	return res
}

func AddEdge(g map[int][]int, from, to int) {
	g[from] = append(g[from], to)
}

func AddDoubleEdge(g map[int][]int, a, b int) {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
}

type Stack []int

func InitStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Insert(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (popNum int, ok bool) {
	arr := *s
	if len(arr) > 0 {
		popNum := arr[len(arr)-1]
		arr2 := arr[:len(arr)-1]
		*s = arr2

		return popNum, true
	}
	return
}

func (s *Stack) CheckInStake(n int) bool {
	arr := *s
	for i := range arr {
		if arr[i] == n {
			return true
		}
	}
	return false
}

func CreateGraphArr(nodNum int) [][]int {

	res := make([][]int, nodNum)

	return res
}

func AddEdgeArr(graph [][]int, from, to int) {
	graph[from] = append(graph[from], to)
}
