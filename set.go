package goco

import "fmt"

type Set[T comparable] interface {
	Add(e T)
	Remove(e T)
	Contains(e T) bool
	IsEmpty() bool
	Size() int
	Clear()
	ToSlice() []T
}

type set[T comparable] struct {
	elements map[T]T
}

func NewSet[T comparable](e ...T) *set[T] {
	return &set[T]{SliceTomap[T, T](e, func(e T) T { return e })}
}

func (s *set[T]) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func (s *set[T]) Add(e T) {
	s.elements[e] = e
}

func (s *set[T]) Remove(e T) {
	delete(s.elements, e)
}

func (s *set[T]) Contains(e T) bool {
	_, ok := s.elements[e]
	return ok
}

func (s *set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *set[T]) Size() int {
	return len(s.elements)
}

func (s *set[T]) Clear() {
	s.elements = make(map[T]T)
}

func (s *set[T]) ToSlice() []T {
	keys := make([]T, len(s.elements))
	i := 0
	for k := range s.elements {
		keys[i] = k
		i++
	}
	return keys
}
