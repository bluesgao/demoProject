package executor

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

var n = 3

func TestSubmit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		Submit(func() {
			t.Logf("hello task %d \n", i)
			time.Sleep(time.Millisecond)
			wg.Done()
		})
	}

	wg.Wait()

	t.Logf("running workers number:%d", defaultExecutors.GetRunnings())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("memory usage:%d", mem.TotalAlloc/1024)

	defaultExecutors.ShutDown()
}
