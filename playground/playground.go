package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/herlegs/KafkaPlay/benchmark/rate"
)

var counter uint64
var limiter *rate.Limiter

func watcher() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var last uint64
	for {
		<-ticker.C
		val := atomic.LoadUint64(&counter)
		fmt.Printf("%v\n", val-last)
		last = val
	}
}

func adder(limiter *rate.Limiter) {
	for {
		t := time.Now()
		for i := 0; i < 1; i++ {
			if limiter.AllowN(t, 1) {
				atomic.AddUint64(&counter, 1)
			}
		}

	}
}

type aaa struct {
	a int
}

func main() {
	//limiter = rate.NewLimiter(rate.Limit(65), 1)
	//for i := 0; i < 10000; i++ {
	//	go adder(limiter)
	//}
	//
	//go watcher()
	//
	//time.Sleep(time.Minute)
	//6-6 05:03

	c := make(chan *aaa)

	go func() {
		i := 0
		for {
			i++
			c <- &aaa{
				a: i,
			}
			time.Sleep(time.Second)
			if i > 3 {
				close(c)
				return
			}
		}
	}()

	for {
		select {
		case a := <-c:
			fmt.Printf("%v\n", a)
		}
	}

	fmt.Printf("end\n")
}
