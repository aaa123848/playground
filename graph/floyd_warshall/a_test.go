package floyd

import (
	"log"
	"testing"
)

func TestB(t *testing.T) {
	m := createGraph(7)

	m[0][1] = 2
	m[0][2] = 5
	m[0][6] = 10
	m[1][2] = 2
	m[1][4] = 11
	m[2][6] = 2
	m[6][5] = 11
	m[4][5] = 1
	m[5][4] = -2
	res := Floyd(m)

	for i := range res {
		for j := range res {
			p := res[i][j]
			if p == inf {
				log.Printf("from %v to %v is %v", i, j, "inif")
				continue
			}

			if p == negIng {
				log.Printf("from %v to %v is %v", i, j, "-inif")
				continue
			}

			log.Printf("from %v to %v is %v", i, j, res[i][j])
		}
	}

}
