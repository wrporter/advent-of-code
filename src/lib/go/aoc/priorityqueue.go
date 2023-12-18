package aoc

import "container/heap"

// PriorityQueueItem is an item that can be stored in a PriorityQueue.
type PriorityQueueItem interface {
	Less(item PriorityQueueItem) bool
}

// priorityQueueImpl is a concrete implementation of a priority queue, using a
// slice.
type priorityQueueImpl []PriorityQueueItem

// Len returns the size of the queue.
func (pqi priorityQueueImpl) Len() int {
	return len(pqi)
}

// Less compares two items in the queue to determine their priority.
func (pqi priorityQueueImpl) Less(i, j int) bool {
	return pqi[i].Less(pqi[j])
}

// Swap switches the order or items in the queue.
func (pqi priorityQueueImpl) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[i]
}

// Push adds an item to the queue.
func (pqi *priorityQueueImpl) Push(x interface{}) {
	item := x.(PriorityQueueItem)
	*pqi = append(*pqi, item)
}

// Pop removes the top item from the queue.
func (pqi *priorityQueueImpl) Pop() interface{} {
	old := *pqi
	n := len(old)
	item := old[n-1]
	*pqi = old[0 : n-1]
	return item
}

// PriorityQueue implements priorityQueueImpl
type PriorityQueue struct {
	priorityQueueImpl priorityQueueImpl
}

// NewPriorityQueue returns an empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

// Push adds an item to the queue.
func (pq *PriorityQueue) Push(item PriorityQueueItem) {
	heap.Push(&pq.priorityQueueImpl, item)
}

// Pop removes the top item from the queue.
func (pq *PriorityQueue) Pop() PriorityQueueItem {
	return heap.Pop(&pq.priorityQueueImpl).(PriorityQueueItem)
}

// Front returns the top item from the queue without removing it.
func (pq *PriorityQueue) Front() PriorityQueueItem {
	// The first bit in the queue should be the minimum
	return pq.priorityQueueImpl[0]
}

// Length returns the size of the queue.
func (pq *PriorityQueue) Length() int {
	return pq.priorityQueueImpl.Len()
}
