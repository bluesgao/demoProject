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
	coreSize      int32 //基本bee数量
	runnings      int32 //运行中的bee数量
	taskQueueSize int32 //任务队列大小
	bees          []*Bee
	mutex         sync.Mutex
	once          sync.Once
	shutdown      int32 //关闭标志
	startBeeId    int32 //bee起始工号
}

//新建
func NewBeehive(coreSize int, taskQueueSize int, lazy bool) *Beehive {
	beehive := Beehive{
		coreSize:      int32(coreSize),
		taskQueueSize: int32(taskQueueSize),
		startBeeId:    -1,
	}

	if lazy == false {
		for i := 0; i < coreSize; i++ {
			//新建bee
			b := NewBee(atomic.AddInt32(&beehive.startBeeId, 1), &beehive)
			//将运行中的bee数量加1
			atomic.AddInt32(&beehive.runnings, 1)
			//将bee添加到bees中
			beehive.bees = append(beehive.bees, b)
		}
	}
	go beehive.purge() //防止go runtime系统报deadlock异常
	return &beehive
}

//关闭
func (beehive *Beehive) ShutDown() {
	beehive.once.Do(func() {
		atomic.StoreInt32(&beehive.shutdown, 1) //将关闭标志置为1
		log.Printf("ShutDown %+v \n", beehive)
		beehive.mutex.Lock()
		//将所有闲置的bee置为空
		for i, b := range beehive.bees {
			b.fire()
			beehive.bees[i] = nil
		}
		beehive.bees = nil
		beehive.mutex.Unlock()
	})
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

//定时清理idleBee
func (beehive *Beehive) purge() {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for t := range ticker.C {
		log.Printf("purge time:%+v , beehive:%+v \n", t, beehive)
		if 1 == atomic.LoadInt32(&beehive.shutdown) {
			break
		}

		beehive.mutex.Lock()

		//判断bee中的taskque是否为0
		idleBees := beehive.bees
		log.Printf("清理前 idleBees:%+v\n", beehive.bees)
		for i, b := range idleBees {
			l := len(b.taskQueue)
			if l <= 0 {
				log.Printf("被清理 idleBee:%+v\n", idleBees[i])
				//给bee发送quit信号，让go程退出
				idleBees[i].fire()
				//运行中的bee数量减一
				atomic.AddInt32(&beehive.runnings, -1)
				//将需要被清理的bee从bees中删除
				idleBees[i] = nil
			}
		}

		var newBees []*Bee
		//将 bees=nil的删除
		for ii, ib := range idleBees {
			if ib != nil {
				newBees = append(newBees, idleBees[ii])
			}
		}
		log.Printf("清理后 idleBees:%+v\n", beehive.bees)
		beehive.bees = newBees

		beehive.mutex.Unlock()
	}
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
	log.Printf("选定bee%d, tasksize:%d ,%+v, \n", b.id, len(b.taskQueue), b)

	b.addTask(t) //将任务分配给选定的bee
	b.do()       //bee开始执行task
}
