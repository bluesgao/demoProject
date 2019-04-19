package main

import "log"

func main() {
	ch := make(chan int, 8)
	log.Printf("ch cap:%d \n", cap(ch))
	log.Printf("ch len:%d \n", len(ch))

	ch <- 1
	ch <- 2
	log.Printf("ch cap:%d \n", cap(ch))
	log.Printf("ch len:%d \n", len(ch))
}
