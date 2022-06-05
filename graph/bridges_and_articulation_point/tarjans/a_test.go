package tarjans

import (
	g "gotest2/graph"
	"log"
	"testing"
)

func TestA(t *testing.T) {

	gra := g.CreateGraph(6)
	g.AddEdge(gra, 0, 1)
	g.AddDoubleEdge(gra, 1, 2)

	g.AddEdge(gra, 3, 4)
	g.AddEdge(gra, 4, 5)
	g.AddEdge(gra, 5, 3)

	tj := Tarjans{}
	tj.Process(gra)

	log.Println(tj.result)
	log.Println(tj.low)

}

func TestB(t *testing.T) {

	gra := g.CreateGraph(10)
	g.AddEdge(gra, 0, 1)
	g.AddEdge(gra, 1, 2)
	g.AddEdge(gra, 2, 3)
	g.AddEdge(gra, 3, 4)
	g.AddEdge(gra, 4, 5)
	g.AddEdge(gra, 5, 6)
	g.AddEdge(gra, 6, 7)
	g.AddEdge(gra, 7, 8)
	g.AddEdge(gra, 8, 9)
	g.AddEdge(gra, 9, 0)

	tj := Tarjans{}
	tj.Process(gra)

	log.Println(tj.result)
	log.Println(tj.low)

}

func TestC(t *testing.T) {

	gra := g.CreateGraph(10)
	g.AddEdge(gra, 0, 1)
	g.AddEdge(gra, 1, 2)
	g.AddEdge(gra, 2, 3)
	g.AddEdge(gra, 3, 4)
	g.AddEdge(gra, 4, 0)

	g.AddEdge(gra, 5, 0)
	g.AddEdge(gra, 5, 6)
	g.AddEdge(gra, 6, 7)
	g.AddEdge(gra, 7, 8)
	g.AddEdge(gra, 8, 9)
	g.AddEdge(gra, 9, 5)

	tj := Tarjans{}
	tj.Process(gra)

	log.Println(tj.result)
	log.Println(tj.low)

}
