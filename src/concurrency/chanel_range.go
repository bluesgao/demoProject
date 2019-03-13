package main

import "fmt"

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for i := range c {
			fmt.Println("i:", i)
		}
	}()

	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)

	<-done
}
