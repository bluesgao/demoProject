package executor

import "log"

type Worker struct {
	pool     *WorkerPool
	taskChan chan T
}

func (w *Worker) run() {
	go func() {
		//异常处理
		defer func() {
			if p := recover(); p != nil {
				w.pool.decrRunnings()
				log.Printf("worker exits from a panic: %v", p)
			}
		}()

		if len(w.pool.taskQue) > 0 {
			for t := range w.pool.taskQue {
				t()
			}
		}

		for t := range w.taskChan {
			if t == nil {
				//将运行中worker数量减1
				w.pool.decrRunnings()
				return
			}

			//执行task
			log.Printf("执行task: %v", w)
			t()
			//将worker归还到pool中
			w.pool.returnWorker(w)
		}
	}()
}

func (w *Worker) addTask(t T) {
	w.taskChan <- t
}
