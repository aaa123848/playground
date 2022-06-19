package graph

import (
	"log"
	"testing"
)

func TestBfsWithQ(t *testing.T) {
	numNodes := 8
	gra := CreateGraph(numNodes)
	AddEdge(gra, 1, 2)
	AddEdge(gra, 1, 2)
	AddEdge(gra, 1, 3)
	AddEdge(gra, 2, 4)
	AddEdge(gra, 2, 5)
	AddEdge(gra, 3, 6)
	AddEdge(gra, 3, 7)
	AddEdge(gra, 2, 2)
	AddEdge(gra, 2, 3)
	AddEdge(gra, 6, 2)
	AddEdge(gra, 1, 6)
	bq := intBfs(gra)
	bq.startBfs()
	log.Println(bq.order)
	log.Println(bq.cycle)

}
