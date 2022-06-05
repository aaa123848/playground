package eulerian

import (
	g "gotest2/graph"
	"log"
	"testing"
)

func TestA(*testing.T) {
	graph := g.CreateGraphArr(7)
	g.AddEdgeArr(graph, 1, 2)
	g.AddEdgeArr(graph, 1, 3)
	g.AddEdgeArr(graph, 2, 2)
	g.AddEdgeArr(graph, 2, 4)
	g.AddEdgeArr(graph, 2, 4)
	g.AddEdgeArr(graph, 3, 1)
	g.AddEdgeArr(graph, 3, 2)
	g.AddEdgeArr(graph, 3, 5)
	g.AddEdgeArr(graph, 4, 3)
	g.AddEdgeArr(graph, 4, 6)
	g.AddEdgeArr(graph, 5, 6)
	g.AddEdgeArr(graph, 6, 3)

	e := InitEullerianPath(graph)

	isEuPath, start, _ := e.IsEulerianPath()
	log.Println("start at node: ", start)

	if isEuPath {
		e.FindEulerianPath(start)
	}

	log.Println(e.path)

}

func TestB(*testing.T) {
	graph := g.CreateGraphArr(5)
	g.AddEdgeArr(graph, 0, 1)
	g.AddEdgeArr(graph, 1, 2)
	g.AddEdgeArr(graph, 1, 3)
	g.AddEdgeArr(graph, 2, 1)
	g.AddEdgeArr(graph, 3, 4)

	e := InitEullerianPath(graph)

	isEuPath, start, _ := e.IsEulerianPath()
	log.Println("start at node: ", start)

	if isEuPath {
		e.FindEulerianPath(start)
	}

	log.Println(e.path)

}
