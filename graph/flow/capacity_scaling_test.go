package flow

import (
	"log"
	"testing"
)

func TestCSMaxFlow(t *testing.T) {
	nodeNum := 6
	s := nodeNum - 1
	e := nodeNum - 2
	gra := CreateCapacityMetrix(6)
	gra.addCapacity(s, 0, 20)
	gra.addCapacity(s, 1, 41)
	gra.addCapacity(2, e, 23)
	gra.addCapacity(3, e, 16)
	gra.addCapacity(0, 1, 2)
	gra.addCapacity(0, 2, 4)
	gra.addCapacity(0, 3, 8)
	gra.addCapacity(1, 3, 9)
	gra.addCapacity(3, 2, 6)

	cs := InitCapacityScaling(gra)
	res := cs.findMaxCapacity()
	log.Println(res)
	log.Println(cs.visited)
}

func TestCSMaxFlowInitTryFlow(t *testing.T) {
	nodeNum := 6
	s := nodeNum - 1
	e := nodeNum - 2
	gra := CreateCapacityMetrix(6)
	gra.addCapacity(s, 0, 20)
	gra.addCapacity(s, 1, 91)
	gra.addCapacity(2, e, 23)
	gra.addCapacity(3, e, 16)
	gra.addCapacity(0, 1, 2)
	gra.addCapacity(0, 2, 4)
	gra.addCapacity(0, 3, 8)
	gra.addCapacity(1, 3, 9)
	gra.addCapacity(3, 2, 6)

	cs := InitCapacityScaling(gra)
	cs.initTryFlow()
	log.Println(cs.tryFlow)
}

func TestCSA(t *testing.T) {
	nodeNum := 6
	s := nodeNum - 1
	e := nodeNum - 2

	gra := CreateCapacityMetrix(6)
	gra.addCapacity(s, 0, 10)
	gra.addCapacity(s, 1, 10)
	gra.addCapacity(2, e, 10)
	gra.addCapacity(3, e, 10)
	gra.addCapacity(0, 1, 2)
	gra.addCapacity(0, 2, 4)
	gra.addCapacity(0, 3, 8)
	gra.addCapacity(1, 3, 9)
	gra.addCapacity(3, 2, 6)

	cs := InitCapacityScaling(gra)
	res := cs.Process(s, e)
	log.Println(res)
}

func TestCSB(t *testing.T) {
	nodeNum := 6
	s := nodeNum - 1
	e := nodeNum - 2

	gra := CreateCapacityMetrix(6)
	gra.addCapacity(s, 0, 6)
	gra.addCapacity(s, 1, 14)

	gra.addCapacity(2, e, 11)
	gra.addCapacity(3, e, 12)

	gra.addCapacity(0, 1, 1)
	gra.addCapacity(2, 3, 1)
	gra.addCapacity(0, 2, 5)
	gra.addCapacity(1, 2, 7)
	gra.addCapacity(1, 3, 10)

	cs := InitCapacityScaling(gra)
	res := cs.Process(s, e)
	log.Println(res)
}
