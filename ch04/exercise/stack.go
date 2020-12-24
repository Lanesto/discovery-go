package eval

import "errors"

// Stack holds any items but it is user's responsibility
// to check its types
type Stack []interface{}

// ErrStackIsEmpty occurs when stack is empty but
// try to peek or pop item from it.
var ErrStackIsEmpty error = errors.New("stack is empty")

// NewStack returns new(Stack) with items in it.
// First item of slice is pushed at first.
func NewStack(items ...interface{}) *Stack {
	s := new(Stack)
	for _, item := range items {
		s.Push(item)
	}
	return s
}

// Push put a single item to stack
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// Pop take a item from stack
func (s *Stack) Pop() (interface{}, error) {
	ret, err := s.Peek()
	if err != nil {
		return nil, err
	}
	*s = (*s)[:len(*s)-1]
	return ret, nil
}

// Peek just read stack's top without modifying stack
func (s *Stack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, ErrStackIsEmpty
	}
	return (*s)[len(*s)-1], nil
}

// Empty returns whether stack is empty
func (s *Stack) Empty() bool {
	return len(*s) == 0
}
