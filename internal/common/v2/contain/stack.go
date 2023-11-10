package contain

import (
	"sync"
)

// Stack is a generic array-based stack implementation that is concurrent-safe
type Stack[T any] struct {
	lock  sync.Mutex
	items []T
}

// NewStack initializes a new Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

// Push adds an item on top of the stack
func (stack *Stack[T]) Push(item T) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.items = append(stack.items, item)
}

// PushMany adds an item on top of the stack
func (stack *Stack[T]) PushMany(items ...T) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.items = append(stack.items, items...)
}

// Pop removes and returns the top item on the stack
func (stack *Stack[T]) Pop() T {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	size := len(stack.items)
	if size == 0 {
		var empty T
		return empty
	}

	item := stack.items[size-1]
	stack.items = stack.items[:size-1]
	return item
}

// PopMany removes and returns the top n items on the stack. If n exceeds the
// size of the stack, all remaining items are removed and returned.
func (stack *Stack[T]) PopMany(n int) []T {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	size := len(stack.items)
	if size == 0 {
		var empty []T
		return empty
	}

	items := stack.items[size-n:]
	stack.items = stack.items[:size-n]
	return items
}

// Peek returns the top item on the stack
func (stack *Stack[T]) Peek() T {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	size := len(stack.items)
	if size == 0 {
		var empty T
		return empty
	}

	return stack.items[size-1]
}
