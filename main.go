package main

import (
	"log"
	"time"
)

func a() {
	for {
		log.Println(1)
		time.Sleep(time.Second)
	}
}

func b() {
	go a()
}

func c() {
	go func() {
		for {
			log.Println(2)
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	a := 1

	ap := &a

	c := *ap

	log.Println(c)

	*ap = 3

	log.Println(a)
	log.Println(*ap)
	log.Println(c)

}
