package server

// import (
// 	"math/rand"
// )

// type Job struct{
// 	Num int
// 	ID int
// }

// type Result struct{
// 	job  *Job
// 	sum int
// }


// func CalcuNum(job *Job, result chan *Result){
// 	var sum int
// 	number := job.Num

// 	for number != 0 {
// 		tmp := number%10
// 		sum += tmp
// 		number /= 10
// 	}

// 	resultElement := &Result{
// 		job:	job,
// 		sum:	sum,
// 	}
	
// 	result <- resultElement

// 	return
// }

// func Worker(){

// }

// func startWorkerPool(workerNum int, jobChannel chan *Job, resultChannel chan *Result){

// }

// func JobProduce(){
// 	jobChannel := make(chan *Job,1000)
// 	resultChannel := make(chan *Result,1000)

// 	var id int = 0
// 	for {
// 		id++
// 		number := rand.Int()

// 		job := &Job{
// 			Num :	number,
// 			ID:	id,
// 		}

// 		jobChannel <- job
// 	}

// }