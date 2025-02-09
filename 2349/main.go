package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

type NumberContainers struct {
	indexToNumber   map[int]int
	numberToIndices map[int]*MinHeap
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexToNumber:   make(map[int]int),
		numberToIndices: make(map[int]*MinHeap),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	if oldNumber, exists := nc.indexToNumber[index]; exists {
		if oldNumber == number {
			return
		}
		nc.removeIndexFromNumber(oldNumber, index)
	}
	nc.indexToNumber[index] = number
	if _, exists := nc.numberToIndices[number]; !exists {
		nc.numberToIndices[number] = &MinHeap{}
		heap.Init(nc.numberToIndices[number])
	}
	heap.Push(nc.numberToIndices[number], index)
}

func (nc *NumberContainers) Find(number int) int {
	if indices, exists := nc.numberToIndices[number]; exists && indices.Len() > 0 {
		for indices.Len() > 0 {
			smallestIndex := heap.Pop(indices).(int)
			if nc.indexToNumber[smallestIndex] == number {
				heap.Push(indices, smallestIndex)
				return smallestIndex
			}
		}
	}
	return -1
}

func (nc *NumberContainers) removeIndexFromNumber(number int, index int) {
	if indices, exists := nc.numberToIndices[number]; exists {
		for i, idx := range *indices {
			if idx == index {
				heap.Remove(indices, i)
				break
			}
		}
	}
}

// MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
