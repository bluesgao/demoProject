package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int, 3)

	b <- 1
	b <- 2
	fmt.Println("a: ", len(a), cap(a))
	fmt.Println("b: ", len(b), cap(b))
}
