package salesman

import (
	"log"
	"testing"
)

func TestCombination(t *testing.T) {

	res := combination(2, 3)

	log.Println(res)
}

func TestCombination2(t *testing.T) {

	res := combination2Start(2, 3)

	log.Println(res)
}
