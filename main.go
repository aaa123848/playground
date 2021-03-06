package main

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

func main() {
	log.Println(twoSum(21, []int{7, 9, 8, 3}))

}
