package store

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

func RandomUnits(num int) (res []Unit) {
	for i := 0; i < num; i++ {
		u := Unit{
			point:      "a",
			disFromOri: int(RandomInt(0, 1000)),
		}

		res = append(res, u)
	}
	return res
}

func TestPQAppend(t *testing.T) {
	units := RandomUnits(10)
	log.Println(units)

	pq := PriorityQ{
		que: make([]Unit, 0),
	}

	for _, u := range units {
		pq.append(u)

	}
	log.Println(pq)

	log.Println(pq.pop())
	log.Println(pq)

}
