package bees

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type T func()

type BeeHive struct {
	coreSize      int32 //基本bee数量
	runnings      int32 //运行中的bee数量
	taskQueueSize int32 //任务队列大小
	workerBees    []*WorkerBee
	mutex         sync.Mutex
	once          sync.Once
	destroy       int32 //关闭标志
	workerNo      int32 //起始工号
}

//新建
func NewBeeHive(coreSize int, taskQueueSize int, lazy bool) *BeeHive {
	beehive := BeeHive{
		coreSize:      int32(coreSize),
		taskQueueSize: int32(taskQueueSize),
		workerNo:      0,
	}

	if lazy == false {
		for i := 0; i < coreSize; i++ {
			//新建bee
			b := NewWorkerBee(atomic.AddInt32(&beehive.workerNo, 1), &beehive)
			//将运行中的bee数量加1
			atomic.AddInt32(&beehive.runnings, 1)
			//将bee添加到bees中
			beehive.workerBees = append(beehive.workerBees, b)
		}
	}
	go beehive.purge() //防止go runtime系统报deadlock异常
	return &beehive
}

//关闭
func (beehive *BeeHive) Destroy() {
	beehive.once.Do(func() {
		atomic.StoreInt32(&beehive.destroy, 1) //将关闭标志置为1
		log.Printf("Destroy %+v \n", beehive)
		/*beehive.mutex.Lock()
		//让所有的bee退出
		for i, b := range beehive.workerBees {
			b.fire()
			beehive.workerBees[i] = nil
		}
		beehive.workerBees = nil
		beehive.mutex.Unlock()*/
	})
}

//beehive容量
func (beehive *BeeHive) Capacity() int {
	return int(atomic.LoadInt32(&beehive.coreSize))
}

//提交任务
func (beehive *BeeHive) SubmitTask(t T) {
	log.Printf("向beehive提交任务 %+v \n", beehive)
	beehive.assignTask(t)
}

//定时清理idleBee
func (beehive *BeeHive) purge() {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for t := range ticker.C {
		log.Printf("purge time:%+v , beehive:%+v \n", t, beehive)

		beehive.mutex.Lock()

		//判断bee中的taskque是否为0
		idleBees := beehive.workerBees
		log.Printf("清理前 idleBees:%+v\n", beehive.workerBees)
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

		var busyBees []*WorkerBee
		//将 workerBees=nil的删除
		for ii, ib := range idleBees {
			if ib != nil {
				busyBees = append(busyBees, idleBees[ii])
			}
		}
		log.Printf("清理后 idleBees:%+v\n", beehive.workerBees)
		beehive.workerBees = busyBees
		if len(busyBees) <= 0 { //如果没有在工作的bee，且destory标志为1，purge线程退出
			if 1 == atomic.LoadInt32(&beehive.destroy) {
				beehive.mutex.Unlock() //先解锁再退出
				break
			}
		}

		beehive.mutex.Unlock()
	}
}

//分配task(选择一个bee，将task分配给它)
func (beehive *BeeHive) assignTask(t T) {
	log.Printf("分配task %+v \n", beehive)
	var bee *WorkerBee
	beehive.mutex.Lock()
	n := len(beehive.workerBees) //bee数量
	//从 workerBees 中取出一个bee
	if n <= 0 || n < int(beehive.coreSize) { //小于等于0，新建
		log.Printf("新建bee %+v \n", beehive)
		//新建bee
		bee = NewWorkerBee(atomic.AddInt32(&beehive.workerNo, 1), beehive)
		//将运行中的bee数量加1
		atomic.AddInt32(&beehive.runnings, 1)
		//将bee添加到bees中
		beehive.workerBees = append(beehive.workerBees, bee)
	} else { //大于0
		log.Printf("选择bee %+v \n", beehive)
		//worker选择策略，默认随机
		// todo 优化方案：选择任务最少的bee
		bee = beehive.workerBees[rand.Intn(n)]
	}
	beehive.mutex.Unlock()
	log.Printf("选定bee%d, tasksize:%d ,%+v, \n", bee.no, len(bee.taskQueue), bee)

	bee.addTask(t) //将任务分配给选定的bee
	bee.do()       //bee开始干活
}
