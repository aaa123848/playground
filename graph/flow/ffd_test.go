package flow

import (
	g "gotest2/graph"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	gra := g.CreateGraphMetrix(6)

	for i := range gra {
		for j := range gra[i] {
			gra[i][j] = 0
		}
	}

	gra[0][1] = 10
	gra[1][0] = 0

	gra[0][2] = 10
	gra[2][0] = 0

	gra[1][3] = 15
	gra[3][1] = 0

	gra[3][2] = 6
	gra[2][3] = 0

	gra[2][4] = 25
	gra[4][2] = 0

	gra[4][5] = 10
	gra[5][4] = 0

	gra[3][5] = 10
	gra[5][3] = 0

	log.Println(gra)

	f := Initffd(gra)
	log.Println(f.Process(0, 5))
	log.Println(f.cap)

}
