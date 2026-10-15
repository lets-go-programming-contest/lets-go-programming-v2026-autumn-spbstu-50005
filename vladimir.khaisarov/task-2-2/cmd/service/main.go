package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *MaxHeap) Pop() any {
	values := *h
	lastIndex := len(values) - 1
	value := values[lastIndex]
	*h = values[:lastIndex]

	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var dishCount int

	if _, err := fmt.Fscan(reader, &dishCount); err != nil {
		return
	}

	ratings := make(MaxHeap, dishCount)

	for index := range ratings {
		if _, err := fmt.Fscan(reader, &ratings[index]); err != nil {
			return
		}
	}

	var position int

	if _, err := fmt.Fscan(reader, &position); err != nil {
		return
	}

	if position < 1 || position > dishCount {
		return
	}

	heap.Init(&ratings)

	var answer int

	for range position {
		answer = heap.Pop(&ratings).(int)
	}

	fmt.Println(answer)
}
