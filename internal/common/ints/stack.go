package ints

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value int
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// View the top item on the stack
func (s *Stack) Peek() int {
	if s.length == 0 {
		return -1
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() int {
	if s.length == 0 {
		return -1
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack) Push(value int) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func (s *Stack) Copy() *Stack {
	reverse := NewStack()
	for s.Len() > 0 {
		reverse.Push(s.Pop())
	}

	result := NewStack()
	for reverse.Len() > 0 {
		reverse.Push(reverse.Pop())
	}
	return result
}
