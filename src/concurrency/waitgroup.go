package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //累加计数

		go func(id int) {
			defer wg.Done() //递减计数
			time.Sleep(time.Second)
			fmt.Println("goroutine ", id, "done.")
		}(i)
	}

	fmt.Println("main...")
	wg.Wait() //堵塞，直到计数为0
	fmt.Println("main exit.")
}
