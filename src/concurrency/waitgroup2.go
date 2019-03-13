package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //累加计数

	go func() {
		//time.Sleep(time.Second)
		fmt.Println("task1 done.")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("task2 done.")
		wg.Done()
	}()

	fmt.Println("main...")
	wg.Wait() //堵塞，直到计数为0
	fmt.Println("main exit.")
}
