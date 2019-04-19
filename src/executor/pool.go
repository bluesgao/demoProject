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

type Task func()

type Pool struct {
	capacity     int32 //worker数量
	busy         int32 //运行中的worker数量
	workersCache []*Worker
	destroy      chan bool
	mutex        sync.Mutex
}

func NewPool(size int) *Pool {
	p := Pool{
		capacity: int32(size),
		destroy:  make(chan bool),
	}
	return &p
}

func (p *Pool) ShutDown() {

}

func (p *Pool) GetBusy() int {
	return int(atomic.LoadInt32(&p.busy))
}

func (p *Pool) GetCapacity() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) Execute(task Task) {
	w := p.getWorker()
	w.addTask(task)
}

func (p *Pool) getWorker() *Worker {
	var w *Worker
	p.mutex.Lock()
	workers := p.workersCache
	n := len(p.workersCache) //worker数量
	//从 workersCache 中取出一个worker，如果 workersCache 为空，则新建一个worker
	if n <= 0 { //小于等于0，新建
		if p.busy >= p.capacity {
			//todo 执行task拒绝策略
		} else {
			//新建worker
			w = &Worker{
				pool:     p,
				taskChan: make(chan Task),
			}
			w.run()
			//将忙碌中的woker数量加1
			atomic.AddInt32(&p.busy, 1)
		}
	} else { //大于0直接取
		w = workers[n]
		workers[n] = nil
		p.workersCache = workers[:n]
	}
	p.mutex.Unlock()
	log.Printf("getWorker %v \n", w)
	return w
}

func (p *Pool) putWorker(w *Worker) {
	log.Printf("putWorker %v \n", w)
	p.mutex.Lock()
	p.workersCache = append(p.workersCache, w)
	p.mutex.Unlock()
}
