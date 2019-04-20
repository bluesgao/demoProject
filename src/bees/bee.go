package bees

import (
	"log"
	"sync/atomic"
)

type Bee struct {
	beehive *Beehive
	id      int32     //编号
	taskQue chan T    //任务队列
	quit    chan bool //bee退出标志
}

func NewBee(beeId int32, beehive *Beehive) *Bee {
	b := Bee{
		id:      beeId,
		beehive: beehive,
		taskQue: make(chan T, beehive.taskqueSize),
	}
	return &b
}

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
			case t := <-b.taskQue:
				log.Printf("bee%d执行task开始:%+v \n", b.id, b)
				t() //执行task
				log.Printf("bee%d执行task结束:%+v \n", b.id, b)
			case <-b.quit:
				return // we have received a signal to stop
			}
		}
	}()
}

func (b *Bee) addTask(t T) {
	b.taskQue <- t
}

func (b *Bee) Quit() {
	go func() {
		b.quit <- true
	}()
}
