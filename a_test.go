package main

import (
	"log"
	"math"
)

const (
	inif int = math.MaxInt
)

func setInit(graph [][]int) [][]int {
	res := make([][]int, 0)

	res = append(res, graph...)

	return res

}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func aProcess(res [][]int, src, dst, mid int) {
	a := res[src][dst]
	b := res[src][mid] + res[mid][dst]
	res[src][dst] = getMin(a, b)
}

func printRes(res [][]int) {
	for _, r := range res {
		log.Println(r)
	}
}

// func main() {
// 	graph := [][]int{
// 		{0, 2, 5, 4},
// 		{2, 0, 1, inif},
// 		{5, 1, 0, 5},
// 		{4, inif, 5, 0},
// 	}

// 	nn := len(graph)

// 	res := setInit(graph)
// 	printRes(res)
// 	for k := 0; k < nn; k++ {
// 		for i := 0; i < nn; i++ {
// 			for j := 0; j < nn; j++ {
// 				aProcess(res, i, j, k)
// 			}
// 		}
// 		log.Println("-------- k = ", k, "------------")
// 		printRes(res)
// 	}

// }
