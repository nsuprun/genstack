// Package stack provides thread and type safe stack implementation.
package stack

import "sync"

// Stack struct stores stack elements in a slice of T
type Stack[T comparable] struct {
	mu     sync.Mutex
	values []T
}

// New returns an empty stack of type T
func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		values: nil,
	}
}

// Push inserts v of type T to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values = append(s.values, v)
}

// Pop removes and return the ast inserted element from the stack.
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	i := s.Size() - 1
	top := s.values[i]
	s.values = s.values[:i]

	return top, true
}

// Top returns the last inserted element from the stack.
func (s *Stack[T]) Top() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	return s.values[s.Size()-1], true
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.values)
}

// IsEmpty checks if there are elements stored in the stack
func (s *Stack[T]) IsEmpty() bool {
	if s.Size() > 0 {
		return false
	}

	return true
}
