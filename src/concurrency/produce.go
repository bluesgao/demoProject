package main

import (
	"log"
	"math/rand"
	"time"
)

func produce(buf chan<- int) {
	//定义随机因子
	rand.Seed(time.Now().UnixNano())

	//将产生的随机数存储到管道中
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		buf <- num
		log.Println("生产者生产了：", num)
	}
}

func consume(buf <-chan int, exitChan chan<- int) {
	for i := 0; i < 5; i++ {
		num := <-buf
		log.Println("---消费者消费了：", num)
		time.Sleep(time.Second * 2)
	}

	exitChan <- 666
}

func main() {
	bufChan := make(chan int, 5)
	exitChan := make(chan int)

	go produce(bufChan)
	go consume(bufChan, exitChan)

	//堵塞
	<-exitChan
}
