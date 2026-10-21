package main

import (
	"container/heap"
	"fmt"
)

type IntMaxHeap []int

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *IntMaxHeap) Pop() any {
	oldHeap := *h
	n := len(oldHeap)
	val := oldHeap[n-1]
	*h = oldHeap[0 : n-1]

	return val
}

func main() {
	var itemCount int

	if _, err := fmt.Scan(&itemCount); err != nil {
		fmt.Println("Error reading item count:", err)

		return
	}

	maxHeap := &IntMaxHeap{}
	heap.Init(maxHeap)

	for range itemCount {
		var rating int

		if _, err := fmt.Scan(&rating); err != nil {
			fmt.Println("Error reading rating:", err)

			return
		}

		heap.Push(maxHeap, rating)
	}

	var targetRank int

	if _, err := fmt.Scan(&targetRank); err != nil {
		fmt.Println("Error reading target rank:", err)

		return
	}

	var chosenRating int

	for range targetRank {
		popVal := heap.Pop(maxHeap)

		val, ok := popVal.(int)
		if !ok {
			return
		}

		chosenRating = val
	}

	fmt.Println(chosenRating)
}
