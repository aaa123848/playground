package flow

type capacityMatrix [][]int

func CreateCapacityMetrix(nodNum int) (res capacityMatrix) {
	for i := 0; i < nodNum; i++ {
		res = append(res, make([]int, nodNum))
		for j := 0; j < nodNum; j++ {
			res[i][j] = 0
		}
	}
	return res
}

func (ccm capacityMatrix) addCapacity(from, to, cap int) {
	ccm[from][to] = cap
}

func (ccm capacityMatrix) ParseToArray() (res [][]int) {
	for i := range ccm {
		res = append(res, []int{})
		for j := range ccm[i] {
			res[i] = append(res[i], ccm[i][j])
		}
	}

	return res
}
