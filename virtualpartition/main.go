package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"gitlab.myteksi.net/gophers/go/end/common/hash"

	"gitlab.myteksi.net/gophers/go/commons/util/random"
)

func main() {
	eventCh := make(chan *Event)
	var counter int64
	var flowFunc = func(event *Event) {
		time.Sleep(time.Millisecond * 1000)
		atomic.AddInt64(&counter, 1)
	}
	StartDispatcher(eventCh, 50, flowFunc)
	StartProducer(eventCh)
	StartMonitor(&counter)
}

type Event struct {
	Key string
}

type Worker struct {
	ch      chan *Event
	process func(*Event)
}

func (w *Worker) Start() {
	go func() {
		for e := range w.ch {
			w.process(e)
		}
	}()
}

func StartProducer(eventCh chan *Event) {
	go func() {
		for {
			event := &Event{
				Key: random.String(10),
			}
			eventCh <- event
		}
	}()
}

func StartDispatcher(eventCh chan *Event, workerNum int, fn func(event *Event)) {
	workers := make([]*Worker, workerNum)
	for i := 0; i < workerNum; i++ {
		w := &Worker{
			ch:      make(chan *Event),
			process: fn,
		}
		w.Start()
		workers[i] = w
	}
	// dispatcher
	go func() {
		for e := range eventCh {
			// simulate key conversion process
			index := HashKey(e.Key) % uint64(workerNum)
			workers[index].ch <- e
		}
	}()
}

func StartMonitor(addr *int64) {
	duration := time.Second * 5
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	prevCount := atomic.LoadInt64(addr)
	for range ticker.C {
		count := atomic.LoadInt64(addr)
		processed := count - prevCount
		qps := float64(processed) / float64(duration/time.Second)
		fmt.Printf("processed %v in %v. QPS: %v\n", processed, duration, qps)

		prevCount = count
	}
}

func HashKey(keyField string) uint64 {
	// in entity
	streamPartitionID := int64(hash.String(keyField) % 1000000000)
	// in stream collections
	key := strconv.FormatInt(streamPartitionID, 10)
	return hash.String(key)
}
