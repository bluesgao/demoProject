package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	fmt.Println("count x:", x)
}

func test1(n int) {
	start := time.Now().Nanosecond()

	for i := 0; i < n; i++ {
		count()
	}
	fmt.Println("test spend time:", time.Now().Nanosecond()-start)
}

func test2(n int) {
	start := time.Now().Nanosecond()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	fmt.Println("test2 spend time:", time.Now().Nanosecond()-start)
	wg.Wait()
}

func main() {
	n := runtime.GOMAXPROCS(0)
	fmt.Println("n:", n)
	//test spend time: 202743500
	//test1(n)
	test2(n)
}
