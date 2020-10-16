package server

import (
	"time"
	"sync"
	"github.com/golang/glog"
	"runtime"
)

var wg sync.WaitGroup

func PrintString(){
	for i := 0; i<20; i++{
		glog.Info("Hello")
	}

	wg.Done()
	return
}

func Produce(channel chan int){
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		channel <- i
	}

	close(channel)
}

func Consume(channel chan int){
	for data := range channel{
		glog.Info(data)
	}
}

func PrintNum(){
	// wg.Add(1)
	// go PrintString()

	// for i := 0; i<10; i++{
	// 	glog.Info(i)
	// }
	// wg.Wait()

	channel := make (chan int,10)

	go Consume(channel)

	time.Sleep(5*time.Second)

	channel <- 1

	channel <- 2

	//close(channel)

	

	// for value := range channel{
	// 	data := <- channel
	// 	glog.Info(data)
	// 	glog.Info(value)
	// }

	//go Produce(channel)

	// Consume(channel)

	// for value := range channel{
	// 	glog.Info(value)
	// }
	
	cpu :=runtime.NumCPU()
	glog.Info(cpu)
	return
}