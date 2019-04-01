package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value int64

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()
	for i:=0;i<100000 ;i++  {
		atomic.AddInt64(&value, 1)
	}
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(3)

	go worker(&wg)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println(value)
}