package ba

import (
	"log"
	"testing"

	g "gotest2/graph"
)

func printAriti(arr []bool) {

	for i := range arr {
		if arr[i] {
			log.Printf("node: %v is articulation point", i)
		}
	}
}

func TestA(t *testing.T) {
	gra := g.CreateGraph(9)

	g.AddDoubleEdge(gra, 0, 1)
	g.AddDoubleEdge(gra, 0, 2)
	g.AddDoubleEdge(gra, 1, 2)
	g.AddDoubleEdge(gra, 2, 3)
	g.AddDoubleEdge(gra, 3, 4)
	g.AddDoubleEdge(gra, 2, 5)
	g.AddDoubleEdge(gra, 5, 6)
	g.AddDoubleEdge(gra, 6, 7)
	g.AddDoubleEdge(gra, 7, 8)
	g.AddDoubleEdge(gra, 8, 5)

	a := initAritculation(gra)
	a.Process()
	printAriti(a.arPoint)

}
