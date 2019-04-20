package bees

import (
	"log"
	"sync"
	"testing"
)

const (
	RunTimes = 1000
)

func task() {
	log.Printf("work... \n")
}

func BenchmarkBees(b *testing.B) {
	var wg sync.WaitGroup
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			defaultBees.SubmitTask(task)
		}
		wg.Wait()
	}
	b.StopTimer()
	defer defaultBees.ShutDown()
}

//BenchmarkGoroutine-4   	      50	  22300364 ns/op
func BenchmarkGoroutine(b *testing.B) {
	var wg sync.WaitGroup
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			go func() {
				task()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	b.StopTimer()
}
