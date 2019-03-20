package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		time.Sleep(time.Second * 1)
		for i := 0; i < 10; i++ {
			ch <- i
			log.Println("生产者1生产了", i)
		}
		close(ch)
		log.Println("ch通道关闭了")
	}()

	for {
		select {
		case item, ok := <-ch:
			if ok {
				log.Println("------消费者消费了ch:", item)
			} else {
				return
			}

		case <-time.After(time.Second * 3):
			log.Println("超时了")
			return
		}

		time.Sleep(time.Second)
	}
}
