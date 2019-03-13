package main

import (
	"fmt"
	"time"
)

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second)
		fmt.Println("go:", x, y)
	}(a, counter())

	a += 100
	fmt.Println("main:", a, counter())

	time.Sleep(time.Second * 3)
}
