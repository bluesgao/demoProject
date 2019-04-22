package bees

import (
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

/**
    go		task	speed(s)
    10		1000	1.05
    20		10000   2.34
	50		10000   3.15
	100		10000   4.59
*/

var n = 100

func TestSubmit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		defaultBees.SubmitTask(func() {
			log.Printf("hello task")
			//time.Sleep(time.Millisecond)
			wg.Done()
		})
	}

	wg.Wait()

	t.Logf("defaultBees:%+v", defaultBees)

	t.Logf("running workerBees number:%d", defaultBees.GetCapacity())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d", mem.TotalAlloc/1024)

	time.Sleep(time.Second * 30)
	defaultBees.Destroy()
}

func TestNomal(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			log.Printf("hello task %d \n", i)
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d", mem.TotalAlloc/1024)

}
