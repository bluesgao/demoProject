package bees

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
)

type T func()

type Beehive struct {
	coreSize   int32 //基本bee数量
	runnings   int32 //运行中的bee数量
	bees       []*Bee
	mutex      sync.Mutex
	once       sync.Once
	startBeeId int32 //bee起始编号
}

//新建
func NewBeehive(coreSize int) *Beehive {
	p := Beehive{
		coreSize:   int32(coreSize),
		startBeeId: -1,
	}
	return &p
}

//关闭
func (p *Beehive) ShutDown() {
	p.once.Do(func() {
		p.mutex.Lock()
		//将所有闲置的bee置为空
		for i, w := range p.bees {
			w.taskQue <- nil
			p.bees[i] = nil
		}
		p.bees = nil
		p.mutex.Unlock()
	})
}

//beehive容量
func (p *Beehive) GetCapacity() int {
	return int(atomic.LoadInt32(&p.coreSize))
}

//提交任务
func (p *Beehive) SubmitTask(t T) {
	log.Printf("向beehive提交任务 %+v \n", p)
	p.assignTask(t)
}

//分配task(选择一个bee，将task分配给它)
func (p *Beehive) assignTask(t T) {
	log.Printf("分配task %+v \n", p)
	var b *Bee
	p.mutex.Lock()
	n := len(p.bees) //bee数量
	//从 bees 中取出一个bee
	if n <= 0 || n < int(p.coreSize) { //小于等于0，新建
		log.Printf("新建bee %+v \n", p)
		//新建worker
		b = &Bee{
			id:      atomic.AddInt32(&p.startBeeId, 1),
			beehive: p,
			taskQue: make(chan T, 8),
		}
		//将运行中的bee数量加1
		atomic.AddInt32(&p.runnings, 1)
		p.bees = append(p.bees, b)
	} else { //大于0
		log.Printf("选择bee %+v \n", p)
		//worker选择策略，默认随机
		// todo 优化方案：选择任务最少的bee
		b = p.bees[rand.Intn(n)]
	}
	p.mutex.Unlock()
	log.Printf("选定bee%d, tasksize:%d ,%+v, \n", b.id, len(b.taskQue), b)

	b.addTask(t) //将任务分配给选定的bee
	b.do()       //bee开始执行task
}
