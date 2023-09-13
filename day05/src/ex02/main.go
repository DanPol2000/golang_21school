package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h PresentHeap) Len() int {
	return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Size < h[j].Size
	}
	return h[i].Value > h[j].Value
}

func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x any) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getNCoolestPresents(presents []Present, n int) ([]interface{}, error) {
	h := PresentHeap(presents)
	var coolestPresents []interface{}

	if n > len(presents) {
		return nil, errors.New("`n` is larger than the size of the slice or is negative")
	}

	heap.Init(&h)

	for h.Len() > n {
		coolestPresents = append(coolestPresents, heap.Pop(&h))
	}

	return coolestPresents, nil
}

func main() {
	var n int = 2
	presents := []Present{
		Present{Value: 5, Size: 1},
		Present{Value: 4, Size: 5},
		Present{Value: 3, Size: 1},
		Present{Value: 5, Size: 2},
	}

	coolestPresents, err := getNCoolestPresents(presents, n)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(coolestPresents)
}