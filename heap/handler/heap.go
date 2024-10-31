package handler

import (
	"errors"
	"fmt"
)

// MinHeap represents a Min-Heap structure
type MinHeap struct {
	element []int
}

// NewMinHeap initializes a new Min-Heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		element: []int{},
	}
}

// Insert adds a new element to the heap
func (m *MinHeap) Insert(value int) {
	m.element = append(m.element, value)
	m.HeapifyUp(len(m.element) - 1)
}

// heapifyUp ensures the heap properties are maintained after insertion
func (m *MinHeap) HeapifyUp(idx int) {
	for m.element[m.Parent(idx)] > m.element[idx] {
		m.element[m.Parent(idx)], m.element[idx] = m.element[idx], m.element[m.Parent(idx)]
		idx = m.Parent(idx)
	}
}

// ExtractMin removes and returns the smallest element in the heap
func (m *MinHeap) ExtractMin() (int, error) {
	if len(m.element) == 0 {
		return 0, errors.New("heap is empty")
	}
	min := m.element[0]
	m.element[0] = m.element[len(m.element)-1]
	m.element = m.element[:len(m.element)-1]
	m.HeapifyDown(0)
	return min, nil
}

func (m *MinHeap) HeapifyDown(idx int) {
	lastIndex := len(m.element) - 1
	left, right := m.LeftChild(idx), m.RightChild(idx)
	smallest := idx
	if left <= lastIndex && m.element[left] < m.element[smallest] {
		smallest = left
	}
	if right <= lastIndex && m.element[right] < m.element[smallest] {
		smallest = right
	}
	if smallest != idx {
		m.element[smallest], m.element[idx] = m.element[idx], m.element[smallest]
		m.HeapifyDown(smallest)
	}
}

// parent returns the index of the parent node
func (m *MinHeap) Parent(idx int) int {
	return (idx - 1) / 2
}

// leftChild returns the index of the left child
func (m *MinHeap) LeftChild(idx int) int {
	return 2*idx + 1
}

// rightChild returns the index of the right child
func (m *MinHeap) RightChild(idx int) int {
	return 2*idx + 2
}

func (m *MinHeap) Display() {
	fmt.Println("heap elements ", m.element)
}
