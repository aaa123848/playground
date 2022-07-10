package graph

import (
	"log"
	"testing"
)

func TestDfsWithStack(t *testing.T) {
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

	dws := initDfsWithStack(gra)
	dws.process()

	log.Println(dws.order)

}

func TestDfsWithStack2(t *testing.T) {
	numNodes := 5
	gra := CreateGraph(numNodes)
	AddEdge(gra, 0, 1)
	AddEdge(gra, 0, 2)
	AddEdge(gra, 1, 2)
	AddEdge(gra, 1, 3)
	AddEdge(gra, 2, 3)
	AddEdge(gra, 2, 2)

	dws := initDfsWithStack(gra)
	dws.process()
	log.Println(dws.order)

}
