package two_sum

import "log"

func twoSum(target int, subArr []int) bool {

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}
	for _, sub := range subArr {
		b := twoSum(target-sub, subArr)
		if b {
			log.Println(sub)
			return true
		}
	}
	return false
}
