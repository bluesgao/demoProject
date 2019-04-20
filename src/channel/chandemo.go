package main

import (
	"log"
	"math/rand"
)

func main() {
	/*	ch := make(chan int, 8)
		log.Printf("ch cap:%d \n", cap(ch))
		log.Printf("ch len:%d \n", len(ch))

		ch <- 1
		ch <- 2
		log.Printf("ch cap:%d \n", cap(ch))
		log.Printf("ch len:%d \n", len(ch))*/

	for i := 0; i < 1000; i++ {
		log.Printf("randint:%d \n", rand.Intn(4))

	}
}
