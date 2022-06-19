package salesman

import "log"

func combination(subNum, totalNum int) []int {
	arr := make([]int, totalNum)
	for i := range arr {
		arr[i] = i
	}

	count := 0
	allSet := dfs(arr, count, subNum)
	log.Println(allSet)
	res := make([]int, 0)

	for _, set := range allSet {
		setNum := 0
		for _, s := range set {
			tmp := uint(1)
			tmp = tmp << s
			setNum += int(tmp)
		}
		res = append(res, setNum)
	}
	return res
}

func dfs(arr []int, count, countLimit int) [][]int {
	if count >= countLimit {
		return [][]int{{}}
	}

	res := make([][]int, 0)

	for i, v := range arr {
		newArr := make([]int, 0, len(arr)-1)
		for j, vv := range arr {
			if i < j {
				newArr = append(newArr, vv)
			}
		}
		tmpArr := dfs(newArr, count+1, countLimit)
		for _, sub := range tmpArr {
			sub := append(sub, v)
			res = append(res, sub)
		}

	}
	return res
}

func combination2Start(r, n int) []int {
	subset := make([]int, 0)
	combination2(0, 0, r, n, &subset)
	return subset
}

func combination2(set uint, at, r, n int, subset *[]int) {
	if r == 0 {
		subsetV := *subset
		subsetV = append(subsetV, int(set))
		*subset = subsetV
		return
	}

	for i := at; i < n; i++ {
		set2 := set | uint(1)<<i

		combination2(set2, i+1, r-1, n, subset)

		//set = set & ^(uint(1) << i) // 下次迴圈不能再考慮 i, 要還原

	}

}
