package heaps

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_IntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Printf("heap.Pop(h): %v\n", heap.Pop(h))
}
