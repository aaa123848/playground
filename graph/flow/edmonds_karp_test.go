package flow

import (
	"log"
	"testing"
)

func TestEdmondsKarps(t *testing.T) {
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

	log.Println(gra)
	ek := InitEdmondsKarp(gra)

	res := ek.Process(s, e)
	log.Println(res)
	log.Println(ek.capacity)
}

func TestEdmondsKarpsB(t *testing.T) {
	gra := CreateCapacityMetrix(4)
	s := 0
	e := 3
	gra.addCapacity(s, 1, 100)
	gra.addCapacity(s, 2, 100)
	gra.addCapacity(1, 2, 1)
	gra.addCapacity(1, e, 100)
	gra.addCapacity(2, e, 100)

	ek := InitEdmondsKarp(gra)

	res := ek.Process(s, e)
	log.Println(res)
}

func TestEdmondsKarpsC(t *testing.T) {
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

	ek := InitEdmondsKarp(gra)
	res := ek.Process(s, e)
	log.Println(res)
}
