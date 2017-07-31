package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

var counter = 0

func get() int{
	//fmt.Println("counter is",counter)
	return counter
}

func set(i int){
	counter = i
}

type Runs struct {
	m    sync.Mutex
	done uint32
	sync.Once
}

func (o *Runs) Do(f func()) {
	if old := atomic.LoadUint32(&o.done); old >= 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done < 1 {
		o.done++
		f()
	}
}

func testRunLimitedTimes(){
	once := &Runs{}
	for i := 0; i < 10; i++ {
		j := i
		go once.Do(func(){
			fmt.Println("called",j)
		})
	}
	time.Sleep(time.Second*1)
}

func main(){
	testRunLimitedTimes()
}


