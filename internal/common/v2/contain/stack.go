package contain

type (
	Stack[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		value T
		prev  *node[T]
	}
)

// Create a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

// Return the number of items in the stack
func (s *Stack[T]) Len() int {
	return s.length
}

// View the top item on the stack
func (s *Stack[T]) Peek() T {
	if s.length == 0 {
		return GetZero[T]()
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		return GetZero[T]()
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack[T]) Push(value T) {
	n := &node[T]{value, s.top}
	s.top = n
	s.length++
}

func (s *Stack[T]) Copy() *Stack[T] {
	reverse := NewStack[T]()
	for s.Len() > 0 {
		reverse.Push(s.Pop())
	}

	result := NewStack[T]()
	for reverse.Len() > 0 {
		reverse.Push(reverse.Pop())
	}
	return result
}
