package executor

import (
	"log"
	"sync"
	"sync/atomic"
)

//拒绝策略
const (
	ABORT_POLICY = iota
	DISCARD_POLICY
	DISCARD_OLDEST_POLICY
)

type T func()

type WorkerPool struct {
	capacity    int32 //worker数量
	runnings    int32 //运行中的worker数量
	idleWorkers []*Worker
	destroy     chan bool
	mutex       sync.Mutex
	once        sync.Once
	taskQue     chan T //任务队列
}

//新建
func NewWorkerPool(size int) *WorkerPool {
	p := WorkerPool{
		capacity: int32(size),
		destroy:  make(chan bool),
		taskQue:  make(chan T, 8),
	}
	return &p
}

//关闭
func (p *WorkerPool) ShutDown() {
	p.once.Do(func() {
		p.mutex.Lock()
		//将所有闲置的worker置为空
		for i, w := range p.idleWorkers {
			w.taskChan <- nil
			p.idleWorkers[i] = nil
		}
		p.idleWorkers = nil
		p.mutex.Unlock()
	})
}

//运行中worker数量
func (p *WorkerPool) GetRunnings() int {
	return int(atomic.LoadInt32(&p.runnings))
}

//runnings worker 加1
func (p *WorkerPool) incrRunnings() {
	atomic.AddInt32(&p.runnings, 1)
}

//runnings worker 减1
func (p *WorkerPool) decrRunnings() {
	atomic.AddInt32(&p.runnings, -1)
}

//workerpool容量
func (p *WorkerPool) GetCapacity() int {
	return int(atomic.LoadInt32(&p.capacity))
}

//提交任务
func (p *WorkerPool) SubmitTask(t T) {
	log.Printf("向workerPool提交任务 %+v \n", p)

	w := p.borrowWorker()
	if w != nil {
		w.addTask(t)
	} else {
		if p.runnings >= p.capacity {
			log.Printf("将task 缓存到 taskque中 %+v \n", p)
			// 将task 缓存到 taskque中
			p.taskQue <- t
		}

	}
}

//借
func (p *WorkerPool) borrowWorker() *Worker {
	var w *Worker
	p.mutex.Lock()
	workers := p.idleWorkers
	n := len(p.idleWorkers) //worker数量
	//从 idleWorkers 中取出一个worker，如果 idleWorkers 为空，则新建一个worker
	if n <= 0 { //小于等于0，新建
		log.Printf("小于等于0，新建worker %+v \n", p)
		//新建worker
		w = &Worker{
			pool:     p,
			taskChan: make(chan T),
		}
		w.run()
		//将运行中的woker数量加1
		p.incrRunnings()
	} else { //大于0，取尾部
		log.Printf("大于0，取尾部 %+v \n", p)
		w = workers[n-1]
		workers[n] = nil
		p.idleWorkers = workers[:n-1]
	}
	p.mutex.Unlock()
	log.Printf("borrowWorker %+v \n", w)
	return w
}

//还
func (p *WorkerPool) returnWorker(w *Worker) {
	log.Printf("returnWorker %+v \n", w)
	p.mutex.Lock()
	p.idleWorkers = append(p.idleWorkers, w)
	p.mutex.Unlock()
}
