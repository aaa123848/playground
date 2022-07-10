package two_sum

import "log"

type sumTable []bool

func getSumTable(n int) sumTable {
	arrLen := n + 1
	return make(sumTable, arrLen)
}

func TwoSumTable(target int, subArr []int) bool {
	st := getSumTable(target)
	st[0] = true
	for i, unit := range st {
		if unit {
			for _, sub := range subArr {
				if i+sub > len(st)-1 {
					continue
				}
				st[i+sub] = true
			}
		}
	}

	log.Println(st)
	return st[target]
}
