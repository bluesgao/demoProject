package bees

import (
	"log"
	"sync/atomic"
)

type WorkerBee struct {
	beehive   *BeeHive
	id        int32     //工号
	taskQueue chan T    //任务队列
	quit      chan bool //bee退出标志
}

func NewWorkerBee(beeId int32, beehive *BeeHive) *WorkerBee {
	bee := WorkerBee{
		id:        beeId,
		beehive:   beehive,
		taskQueue: make(chan T, beehive.taskQueueSize),
		quit:      make(chan bool),
	}
	return &bee
}

//干活
func (bee *WorkerBee) do() {
	go func() {
		//异常处理
		defer func() {
			if p := recover(); p != nil {
				atomic.AddInt32(&bee.beehive.runnings, -1)
				log.Printf("WorkerBee%d异常退出:%+v \n", bee.id, p)
			}
		}()

		for {
			select {
			case t := <-bee.taskQueue:
				log.Printf("WorkerBee%d执行task开始:%+v \n", bee.id, bee)
				t() //执行task
				log.Printf("WorkerBee%d执行task结束:%+v \n", bee.id, bee)
			case <-bee.quit:
				log.Printf("WorkerBee%d收到退出信号:%+v \n", bee.id, bee)
				return // we have received a signal to stop
			}
		}
	}()
}

//添加任务
func (bee *WorkerBee) addTask(t T) {
	bee.taskQueue <- t
}

//开除
func (bee *WorkerBee) fire() {
	go func() {
		bee.quit <- true
	}()
}
