package main

import "runtime"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var hello string
var updated bool

type aaa map[string]bool

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		panic("already call")
	}()

	for {
	}
}
