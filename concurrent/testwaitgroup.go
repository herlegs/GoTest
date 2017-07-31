package main

import (
	"sync"
	"time"
	"fmt"
)

func testWgWrong(){
	//recover cannot recover panic from its sub-go-routine
	defer func(){
		_= recover()
		fmt.Println("recovered %v")
	}()
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		//wg.Add(1)
		go workerWrong(wg)
	}
	//wg.Wait()
	time.Sleep(time.Second)
}

func workerWrong(wg *sync.WaitGroup){
	//defer wg.Done()
	time.Sleep(1)
	panic("shit")
}

func testWgRight(){
	wg := &sync.WaitGroup{}
	done := make(chan interface{})
	errChan := make(chan error)
	for i := 0; i < 3; i++ {
		j := i
		wg.Add(1)
		go workerRight(wg, errChan, time.Second*time.Duration(j+1))
	}
	go waitJobDone(wg, done)
	select {
	case <- done:
	case err :=  <- errChan:
		fmt.Printf("error in job: %v\n", err)
		return
	}
	fmt.Println("jobs are done")
}

func waitJobDone(wg *sync.WaitGroup, done chan interface{}) {
	fmt.Println("start waiting")
	wg.Wait()
	done <- struct {}{}
}

func recoverFromPanic(errChan chan error){
	if r := recover(); r != nil {
		switch r := r.(type) {
		case error:
			errChan <- r
		default:
			errChan <- fmt.Errorf("%v", r)
		}
	}
}

func workerRight(wg *sync.WaitGroup, errChan chan error, i time.Duration){
	defer recoverFromPanic(errChan)
	defer wg.Done()
	time.Sleep(i)
	if time.Now().Unix() % 1 != 0{
		panic(fmt.Sprintf("panic from worker %v", i))
	}
	fmt.Printf("job %v is done\n", i)
}

func main(){
	testWgRight()
	//time.Sleep(time.Second)
}
