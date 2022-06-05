package salesman

import (
	g "gotest2/graph"
	"testing"
)

func TestA(t *testing.T) {
	gra := g.CreateGraphMetrix(4)
	gra[0][1] = 9
	gra[0][2] = 2
	gra[0][3] = 5
	gra[1][0] = 3
	gra[1][2] = 6
	gra[1][3] = 7

	gra[2][0] = 9
	gra[2][1] = 2
	gra[2][3] = 5

	gra[3][0] = 9
	gra[3][1] = 2
	gra[3][2] = 5

	tsp(gra, 0)
}
