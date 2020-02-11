package main

import (
	"fmt"
	"sync"
	"time"
)

type saramaMsgFetcher struct {
	currIdx  int
	outputCh chan int
	lock     sync.Mutex
}

// Start will start producing msgs to output channel
func (s *saramaMsgFetcher) Start() {
	go func() {
		for {
			s.lock.Lock()
			select {
			case s.outputCh <- s.currIdx:
				s.currIdx++
			case <-time.After(time.Millisecond * 100):
				fmt.Printf("give up %v\n", s.currIdx)
			}
			s.lock.Unlock()
		}
	}()
}

// rewind index, and return index before rewind
func (s *saramaMsgFetcher) Rewind(set int) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	prev := s.currIdx
	s.currIdx = set
	return prev
}

func main() {
	a, b := make(chan bool), make(chan bool)

	go func() {
		b <- false
	}()

	go func() {
		for {
			select {
			case <-a:

			case <-b:
				fmt.Println("either a, b")
			}
		}
	}()

	time.Sleep(time.Second)

}

// Rewind has no problem
func withoutSDKLoopTest(testTime time.Duration) {
	saramaCh := make(chan int)
	msgFetcher := &saramaMsgFetcher{
		outputCh: saramaCh,
	}
	msgFetcher.Start()

	startConsumer(saramaCh, msgFetcher)

	<-time.After(testTime)
}

// Rewind will fail
func withSDKLoopTest(testTime time.Duration) {
	saramaCh := make(chan int)
	msgFetcher := &saramaMsgFetcher{
		outputCh: saramaCh,
	}
	msgFetcher.Start()

	sdkCh := make(chan int)
	startSDKLoop(msgFetcher.outputCh, sdkCh)

	startConsumer(sdkCh, msgFetcher)

	<-time.After(testTime)
}

// SDK loop is same as adding a buffered channel of size 1 in between
func startSDKLoop(saramaCh, sdkCh chan int) {
	go func() {
		//Both of these fails

		//for i := range saramaCh {
		//	sdkCh <- i
		//}
		for {
			sdkCh <- <-saramaCh
		}
	}()
}

// consumer consumes msg until alert line
func startConsumer(sdkCh chan int, fetcher *saramaMsgFetcher) {
	alertLine := 10
	go func() {
		for {
			for msg := range sdkCh {
				if msg > alertLine {
					panic(fmt.Sprintf("rewind failed, since got index after alert line: %v\n", msg))
				}
				if msg >= alertLine {
					idx := fetcher.Rewind(0)
					fmt.Printf("index before rewind: %v\n", idx)
				}
			}
		}
	}()
}
