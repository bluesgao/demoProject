package main

import (
	"log"
	"math/rand"
	"time"
)

//定义管道模拟缓冲区
var buf = make(chan int, 10)

func produce() {
	//定义随机因子
	rand.Seed(time.Now().UnixNano())

	//将产生的随机数存储到管道中
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		buf <- num
		log.Println("生产者生产了：", num)
	}
}

func consume() {
	for i := 0; i < 10; i++ {
		num := <-buf
		log.Println("---消费者消费了：", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go produce()
	go produce()

	go consume()
	go consume()

	for {

	}
}
