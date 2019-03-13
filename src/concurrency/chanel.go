package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) //定义结束事件
	c := make(chan string)      //数据传输通道

	go func() {
		s := <-c //接收消息
		fmt.Println("go s:", s)
		time.Sleep(time.Second * 2)
		close(done) //关闭通道，作为结束事件
	}()

	c <- "hi" //发送消息
	<-done    //堵塞，直到有数据或者管道关闭
}
