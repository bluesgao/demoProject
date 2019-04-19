package executor

import "sync/atomic"

type Worker struct {
	pool     *Pool
	taskChan chan Task
}

func (w *Worker) run() {
	go func() {
		for t := range w.taskChan {
			if t == nil {
				//将忙碌worker数量减1
				atomic.AddInt32(&w.pool.busy, -1)
				return
			}

			//执行task
			t()
			//将worker归还到pool中
			w.pool.putWorker(w)
		}
	}()
}

func (w *Worker) addTask(t Task) {
	w.taskChan <- t
}
