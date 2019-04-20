package bees

import (
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

var n = 1000

func TestSubmit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		Submit(func() {
			log.Printf("hello task %d \n", i)
			time.Sleep(time.Millisecond)
			wg.Done()
		})
	}

	wg.Wait()
	time.Sleep(time.Second * 3)

	t.Logf("defaultBees:%+v", defaultBees)

	t.Logf("running bees number:%d", defaultBees.GetCapacity())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d", mem.TotalAlloc/1024)

	defaultBees.ShutDown()
}
