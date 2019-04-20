package bees

import (
	"log"
	"sync/atomic"
)

type Bee struct {
	beehive *Beehive
	id      int32  //编号
	taskQue chan T //任务队列
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

		for t := range b.taskQue {
			if t == nil {
				log.Printf("bee%d任务处理完成，需要给beehive发送任务完成通知 \n", b.id)
				atomic.AddInt32(&b.beehive.runnings, -1) //将运行中worker数量减1
				//todo bee任务处理完成，需要给beehive发送任务完成通知
				return
			}

			log.Printf("bee%d执行task开始:%+v \n", b.id, b)
			t() //执行task
			log.Printf("bee%d执行task结束:%+v \n", b.id, b)
		}
	}()
}

func (b *Bee) addTask(t T) {
	b.taskQue <- t
}
