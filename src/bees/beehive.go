package bees

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type T func()

type Beehive struct {
	coreSize int32 //基本bee数量
	//maxSize     int32 //最大bee数量
	runnings    int32 //运行中的bee数量
	taskqueSize int32 //任务队列大小
	bees        []*Bee
	mutex       sync.Mutex
	once        sync.Once
	startBeeId  int32 //bee起始编号
}

//新建
func NewBeehive(coreSize int, taskqueSize int) *Beehive {
	beehive := Beehive{
		coreSize:    int32(coreSize),
		taskqueSize: int32(taskqueSize),
		startBeeId:  -1,
	}

	/*	for i := 0; i < coreSize; i++ {
		//新建bee
		b := NewBee(atomic.AddInt32(&beehive.startBeeId, 1), &beehive)
		//将运行中的bee数量加1
		atomic.AddInt32(&beehive.runnings, 1)
		//将bee添加到bees中
		beehive.bees = append(beehive.bees, b)
	}*/
	go beehive.purge() //防止go runtime系统报deadlock异常
	return &beehive
}

//关闭
func (beehive *Beehive) ShutDown() {
	beehive.once.Do(func() {
		beehive.mutex.Lock()
		//将所有闲置的bee置为空
		for i, b := range beehive.bees {
			b.Quit()
			beehive.bees[i] = nil
		}
		beehive.bees = nil
		beehive.mutex.Unlock()
	})
}

//定时清理idleBee
func (beehive *Beehive) purge() {
	ticker := time.NewTicker(time.Second * 3)
	for t := range ticker.C {
		log.Printf("purge time:%+v \n", t)

		beehive.mutex.Lock()
		//todo 判断bees是否为0，runnings是否为0
		if len(beehive.bees) <= 0 && beehive.runnings <= 0 {
			break
		}
		//todo 判断bee中的taskque是否为0
		idleBees := beehive.bees
		n := 0
		for i, b := range idleBees {
			l := len(b.taskQue)
			n = i
			if l <= 0 {
				log.Printf("idleBee:%+v被清理\n", idleBees[n])
				beehive.bees = idleBees[n+1:]
			}
		}
		beehive.mutex.Unlock()
	}
	defer ticker.Stop()
}

//beehive容量
func (beehive *Beehive) GetCapacity() int {
	return int(atomic.LoadInt32(&beehive.coreSize))
}

//提交任务
func (beehive *Beehive) SubmitTask(t T) {
	log.Printf("向beehive提交任务 %+v \n", beehive)
	beehive.assignTask(t)
}

//分配task(选择一个bee，将task分配给它)
func (beehive *Beehive) assignTask(t T) {
	log.Printf("分配task %+v \n", beehive)
	var b *Bee
	beehive.mutex.Lock()
	n := len(beehive.bees) //bee数量
	//从 bees 中取出一个bee
	if n <= 0 || n < int(beehive.coreSize) { //小于等于0，新建
		log.Printf("新建bee %+v \n", beehive)
		//新建bee
		b = NewBee(atomic.AddInt32(&beehive.startBeeId, 1), beehive)
		//将运行中的bee数量加1
		atomic.AddInt32(&beehive.runnings, 1)
		//将bee添加到bees中
		beehive.bees = append(beehive.bees, b)
	} else { //大于0
		log.Printf("选择bee %+v \n", beehive)
		//worker选择策略，默认随机
		// todo 优化方案：选择任务最少的bee
		b = beehive.bees[rand.Intn(n)]
	}
	beehive.mutex.Unlock()
	log.Printf("选定bee%d, tasksize:%d ,%+v, \n", b.id, len(b.taskQue), b)

	b.addTask(t) //将任务分配给选定的bee
	b.do()       //bee开始执行task
}
