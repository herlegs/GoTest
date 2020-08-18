package main

import (
	"fmt"
	"testing"

	"gitlab.myteksi.net/gophers/go/commons/util/random"
)

func TestWorkerUtilisation(t *testing.T) {
	nums := []int{10, 30, 50, 100, 250, 500}
	for _, n := range nums {
		uni := UniqueWorkerBenchmark(n)
		fmt.Printf("[%v] out of [%v] in use, utilisation rate: %v\n", uni, n, float64(uni)/float64(n))
	}
}

// UniqueWorkerBenchmark ...
func UniqueWorkerBenchmark(workerNum int) int {
	testTimes := 10000
	totalCount := 0

	for i := 0; i < testTimes; i++ {
		totalCount += UniqueWorker(workerNum)
	}

	return totalCount / testTimes
}

// UniqueWorker returns number of unique workers until first duplicate
func UniqueWorker(workerNum int) int {
	count, dedup := 0, map[uint64]bool{}

	for {
		index := HashKey(random.String(10)) % uint64(workerNum)
		if !dedup[index] {
			count++
			dedup[index] = true
		} else {
			break
		}
	}
	return count
}

func TestRandString(t *testing.T) {
	times := 10
	worker := uint64(10)
	countM := map[uint64]int{}
	for i := 0; i < times; i++ {
		countM[HashKey(random.String(10))%worker]++
	}
	fmt.Printf("total:%v, dup:%v\n", times, times-len(countM))
}
