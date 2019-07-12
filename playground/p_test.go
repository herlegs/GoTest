package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"testing"
)

func TestHeap(t *testing.T) {

}

func TestVector(t *testing.T) {
	a := "grab"
	b := "abgr"
	fmt.Printf("%v\n", a[2:4]+a[0:2])
	fmt.Printf("%T[%v]\n", b, b)
}

func BenchmarkList(b *testing.B) {
	a := math.MaxInt32
	c := []string{}
	c[a] = ""
	for i := 0; i < b.N; i++ {
		l := list.New()
		for i := 0; i < 200; i++ {
			l.PushBack(i)
			l.Remove(l.Back())
		}
	}
	//fmt.Printf("len:%v\n",l.Len())
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := make([]int, 0, 300)
		for i := 0; i < b.N; i++ {
			for i := 0; i < 200; i++ {
				l = append(l, i)
				//l = l[:len(l)-1]
			}
		}
	}
	//fmt.Printf("len:%v\n",len(l))
}

func findTheDifference(s string, t string) byte {
	var result byte
	for _, c := range []byte(s) {
		result ^= c
	}
	for _, c := range []byte(t) {
		result ^= c
	}
	return result
}

func BinarySearchSmaller(array []int, start, end, target int) int {
	if start < 0 || end >= len(array) {
		return -1
	}
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end
}

func BinarySearchLarger(array []int, start, end, target int) int {
	if start < 0 || end >= len(array) {
		return -1
	}
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func bsLeftMost(items []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2
		if items[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end + 1
}

func TestSss(t *testing.T) {
	//fmt.Printf("%v\n",kSimilarity("ba", "ab"))
	s := []int{1, 2, 3, 3, 3, 3, 4, 4, 4, 5}
	fmt.Printf("%v\n", bsLeftMost(s, 0, len(s)-1, 1))
}
func kSimilarity(A string, B string) int {
	if A == B {
		return 0
	}
	queue := list.New()
	queue.PushBack([]byte(A))
	swaps := 0
	for queue.Len() != 0 {
		swaps++
		size := queue.Len()
		for i := 0; i < size; i++ {
			s := queue.Front().Value.([]byte)
			queue.Remove(queue.Front())
			i, currPos := 0, len(B)-len(s)
			for i < len(s) && currPos < len(B) && s[i] == B[currPos] {
				i++
				currPos++
			}
			for j := i + 1; j < len(s); j++ {
				if s[j] != B[currPos] || s[j] == B[j-i+currPos] {
					continue
				}
				swapped := make([]byte, len(s[i+1:]))
				copy(swapped, s[i+1:])
				swapped[j-i-1] = s[i]
				if isSuffix(swapped, B) {
					return swaps
				}
				queue.PushBack(swapped)
			}
		}
	}
	return swaps
}

func isSuffix(arr []byte, s string) bool {
	for i, j := len(arr)-1, len(s)-1; i >= 0 && j >= 0; {
		if arr[i] != s[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push(&Ele{3, 3, 3})
	pq.Push(&Ele{16, 16, 16})
	pq.Push(&Ele{2, 2, 2})
	pq.Push(&Ele{1, 1, 1})

	for pq.Len() != 0 {
		ele := pq.Pop()
		fmt.Printf("%#v\n", ele)
	}
}

type Ele struct {
	Val int
	X   int
	Y   int
}

type PQ []*Ele

func (q *PQ) Len() int {
	return len(*q)
}

func (q *PQ) Less(i, j int) bool {
	return (*q)[i].Val < (*q)[j].Val
}

func (q *PQ) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PQ) Push(v interface{}) {
	*q = append(*q, v.(*Ele))
}

func (q *PQ) Pop() interface{} {
	last := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return last
}

type PriorityQueue interface {
	Len() int
	Push(ele *Ele)
	Pop() *Ele
}

type priorityQueueImpl struct {
	*PQ
}

func NewPriorityQueue() PriorityQueue {
	return &priorityQueueImpl{
		PQ: &PQ{},
	}
}

func (pq *priorityQueueImpl) Len() int {
	return pq.PQ.Len()
}

func (pq *priorityQueueImpl) Push(e *Ele) {
	heap.Push(pq.PQ, e)
}

func (pq *priorityQueueImpl) Pop() *Ele {
	return heap.Pop(pq.PQ).(*Ele)
}

type segmentTreeImpl struct {
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qps()
	}
}

func qps() {
	a := uint(0)
	for i := 0; i < 10*50*50*50*1000; i++ {
		a *= 2
	}
}
