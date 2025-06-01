package main

import "container/heap"

type StackWithHeap struct {
	itemsCounter int
	values       map[int]int
	heap         *maxHeap
}

func NewStackWithHeap() *StackWithHeap {
	return &StackWithHeap{
		itemsCounter: 0,
		values:       map[int]int{},
		heap:         &maxHeap{},
	}
}

type heapItem struct {
	val   int
	count int
	order int
}

type maxHeap []heapItem

func (m maxHeap) Len() int          { return len(m) }
func (m maxHeap) Swap(i int, j int) { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Less(i int, j int) bool {
	if m[i].count < m[j].count {
		return true
	}
	if m[i].count == m[j].count {
		if m[i].order < m[j].order {
			return true
		}
	}

	return false
}

func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(heapItem))
}

func (s *StackWithHeap) push(v int) {
	s.values[v]++
	heap.Push(s.heap, heapItem{
		val:   v,
		count: s.values[v],
		order: s.itemsCounter,
	})
	s.itemsCounter++
}

func (s *StackWithHeap) pop() *int {
	if len(s.values) == 0 || s.heap.Len() == 0 {
		return nil
	}

	val := heap.Pop(s.heap).(heapItem).val
	s.values[val]--
	if s.values[val] == 0 {
		delete(s.values, val)
	}

	return &val
}
