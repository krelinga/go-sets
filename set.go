// The Set package provides a type-agnostic Set data structure.
package set

import (
	"iter"

	"github.com/krelinga/go-genmath"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates a new Set with the provided items.
func New[T comparable](items ...T) *Set[T] {
	return NewCapacity(len(items), items...)
}

// NewCapacity creates a new Set with the specified capacity and initial items.
func NewCapacity[T comparable](capacity int, items ...T) *Set[T] {
	s := &Set[T]{}
	capacity = genmath.Max(capacity, len(items))
	if capacity > 0 {
		s.data = make(map[T]struct{}, capacity)
	}
	for _, item := range items {
		s.data[item] = struct{}{}
	}
	return s
}

// Add adds the given item to the Set.  Returns true if the item was added, false if it was already present.
func (s *Set[T]) Add(item T) bool {
	if s.data == nil {
		s.data = make(map[T]struct{})
		s.data[item] = struct{}{}
		return true
	}
	_, exists := s.data[item]
	if !exists {
		s.data[item] = struct{}{}
	}
	return !exists
}

// Has checks if the Set contains the given item.
func (s *Set[T]) Has(item T) bool {
	if s.data == nil {
		return false
	}
	_, exists := s.data[item]
	return exists
}

// Del removes the given item from the Set. Returns true if the item was present and removed, false if it was not present.
func (s *Set[T]) Del(item T) bool {
	if s.data == nil {
		return false
	}
	_, exists := s.data[item]
	if exists {
		delete(s.data, item)
	}
	return exists
}

// Len returns the number of items in the Set.
func (s *Set[T]) Len() int {
	if s.data == nil {
		return 0
	}
	return len(s.data)
}

// Values returns an iterator that yields all items in the Set.
func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		if s.data == nil {
			return
		}
		for key, _ := range s.data {
			if !yield(key) {
				return
			}
		}
	}
}
