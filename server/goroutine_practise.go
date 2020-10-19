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

func server1(channel chan int){
	time.Sleep(6*time.Second)
	channel <- 1

	return
}

func server2(channel chan int){
	time.Sleep(3*time.Second)
	channel <- 2
	
	return
}

func PrintNum(){
	// wg.Add(1)
	// go PrintString()

	// for i := 0; i<10; i++{
	// 	glog.Info(i)
	// }
	// wg.Wait()

	channel1 := make (chan int,10)
	channel2 := make (chan int,10)
	
	go server1(channel1)
	go server2(channel2)
	
	select{
	case s1 := <- channel1:
		glog.Info(s1)
	case s2 := <- channel2:
		glog.Info(s2)
	}

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