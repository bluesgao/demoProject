package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	exit := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("goroutine for i:", i)
		}
		fmt.Println("goroutine done")
		//close(exit)
	}()

	<-exit

	fmt.Println("main end")

}
