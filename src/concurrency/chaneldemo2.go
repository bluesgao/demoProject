package main

import (
	"fmt"
	"time"
)

var ch = make(chan bool)

func prt(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Microsecond * 2000)
	}
}

func person1() {
	prt("hello person 1")
	//往管道中写入数据
	//只有prt函数执行完,才会往管道中写入数据
	ch <- true
}

func person2() {
	//从管道中读取数据
	//只有管道中有数据才会读取，否则会堵塞
	<-ch
	prt("world person 2")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
