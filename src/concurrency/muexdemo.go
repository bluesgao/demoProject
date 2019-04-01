package main

import (
	"fmt"
	"sync"
)

var total struct {
	value int
	sync.Mutex
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		total.Lock()
		total.value = total.value + 1
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go worker(&wg)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println(total.value)
}
