package main

import (
	"fmt"
)

type MinHeap struct {
	array []int
	size  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: []int{},
		size:  0,
	}
}

func (h *MinHeap) ParentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) LeftChildIndex(i int) int {
	return 2*i + 1
}

func (h *MinHeap) RightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *MinHeap) Pop() (int, error) {
	if h.size == 0 {
		return 0, fmt.Errorf("Heap is empty")
	}
	min := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--
	h.heapifyDown(0)
	return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := h.ParentIndex(index)
		if h.array[parentIndex] <= h.array[index] {
			break
		}
		h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
		index = parentIndex
	}
}

func (h *MinHeap) heapifyDown(index int) {
	for {
		leftChildIndex := h.LeftChildIndex(index)
		rightChildIndex := h.RightChildIndex(index)
		smallest := index

		if leftChildIndex < h.size && h.array[leftChildIndex] < h.array[smallest] {
			smallest = leftChildIndex
		}
		if rightChildIndex < h.size && h.array[rightChildIndex] < h.array[smallest] {
			smallest = rightChildIndex
		}

		if smallest == index {
			break
		}

		h.array[index], h.array[smallest] = h.array[smallest], h.array[index]
		index = smallest
	}
}

func (h *MinHeap) Peek() (int, error) {
	if h.size == 0 {
		return 0, fmt.Errorf("Heap is empty")
	}
	return h.array[0], nil
}

func main() {
	heap := NewMinHeap()

	// Inserir elementos no heap
	elements := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for _, element := range elements {
		heap.Insert(element)
	}

	// Retirar elementos do heap em ordem crescente
	fmt.Println("Elementos em ordem crescente:")
	for heap.size > 0 {
		min, err := heap.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%d ", min)
	}
}
