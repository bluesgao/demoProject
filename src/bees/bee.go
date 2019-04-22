package bees

import (
	"log"
	"sync/atomic"
)

type Bee struct {
	beehive   *Beehive
	id        int32     //工号
	taskQueue chan T    //任务队列
	quit      chan bool //bee退出标志
}

func NewBee(beeId int32, beehive *Beehive) *Bee {
	b := Bee{
		id:        beeId,
		beehive:   beehive,
		taskQueue: make(chan T, beehive.taskQueueSize),
		quit:      make(chan bool),
	}
	return &b
}

//干活
func (b *Bee) do() {
	go func() {
		//异常处理
		defer func() {
			if p := recover(); p != nil {
				atomic.AddInt32(&b.beehive.runnings, -1)
				log.Printf("bee%d异常退出:%+v \n", b.id, p)
			}
		}()

		for {
			select {
			case t := <-b.taskQueue:
				log.Printf("bee%d执行task开始:%+v \n", b.id, b)
				t() //执行task
				log.Printf("bee%d执行task结束:%+v \n", b.id, b)
			case <-b.quit:
				log.Printf("bee%d收到退出信号:%+v \n", b.id, b)
				return // we have received a signal to stop
			}
		}
	}()
}

//添加任务
func (b *Bee) addTask(t T) {
	b.taskQueue <- t
}

//开除
func (b *Bee) fire() {
	go func() {
		b.quit <- true
	}()
}
