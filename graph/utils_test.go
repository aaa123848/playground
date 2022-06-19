package graph

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateGraph(t *testing.T) {
	a := CreateGraph(5)
	log.Println(a)
}

func TestStack(t *testing.T) {
	stack := InitStack()

	stack.Insert(0)
	stack.Insert(1)
	stack.Insert(2)

	log.Println(stack)

	log.Println(stack.Pop())
	log.Println(stack.Pop())

	log.Println(stack.Pop())

	log.Println(stack)

}

func TestCreateGraphArr(t *testing.T) {

	graph := CreateGraphArr(5)

	AddEdgeArr(graph, 0, 1)
	log.Println(graph)

}

func TestQue(t *testing.T) {
	q := InitQue()
	for i := 0; i < 10; i++ {
		q.Insert(i)
	}
	require.Equal(t, 10, len(q))
	log.Println(q)

	for {
		if v, ok := q.Pop(); ok {
			log.Println(v)
			continue
		}
		break
	}
}
