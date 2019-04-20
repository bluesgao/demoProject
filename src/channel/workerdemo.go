package main

import (
	"log"
	"strconv"
	"time"
)

type Job struct {
	JobId  string //任务id
	Input  string //任务输入
	Output string //任务输出
}

type Worker struct {
	WorkerId string    //工人id
	JobPool  chan Job  //任务队列
	exit chan bool //退出chan
}

func NewWorker(wid string, jq chan Job) Worker {
	w := Worker{
		WorkerId: wid,
		JobPool:  jq,
		exit: make(chan bool),
	}

	return w
}

func (w Worker) Do() {
	go func() {
		for {
			select {
			case job := <-w.JobPool:
				job.Output = "job complete"
				log.Printf("workerId[%s] do work complete. job[%+v] \n", w.WorkerId, job)
			case <-w.exit:
				log.Printf("workerId[%s] quit. \n", w.WorkerId)
				return
			}
		}
	}()
}

func (w Worker) Quit() {
	go func() {
		w.exit <- true
	}()
}

//任务池
var JobPool chan Job

func init() {
	log.Println("init")
	JobPool = make(chan Job, 1000)
	go createJods()
}

func main() {

	exit := make(chan bool)

	maxWorker := 100
	for i := 0; i < maxWorker; i++ {
		w := NewWorker("w"+strconv.Itoa(i), JobPool)
		w.Do()
	}

	<-exit
}

func createJods() {
	log.Printf("create job start \n")

	id := 0
	timer := time.NewTicker(time.Microsecond * 100)
endLoop:
	for {
		select {
		case <-timer.C:
			if id >= 1000 {
				break endLoop
			}
			job := Job{
				JobId: "j" + strconv.Itoa(id),
				Input: "job" + strconv.Itoa(id),
			}
			log.Printf("create job:%v \n", job)
			JobPool <- job
			id++
		}
	}

	log.Printf("create job end \n")

}
