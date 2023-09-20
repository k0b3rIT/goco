package collections

import "fmt"

type list[T comparable] struct {
	elements []T
}

func NewList[T comparable](e... T) *list[T] {
	return &list[T]{e}
}

func (l *list[T]) String() string {
	return fmt.Sprintf("%v", l.elements)
}

func (l *list[T]) Add(e T) {
	l.elements = append(l.elements, e)
}

func (l *list[T]) Remove(e T) {
	for i, v := range l.elements {
		if v == e {
			l.elements = append(l.elements[:i], l.elements[i+1:]...)
		}
	}
}

func (l *list[T]) Contains(e T) bool {
	for _, v := range l.elements {
		if v == e {
			return true
		}
	}
	return false
}

func (l *list[T]) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *list[T]) Size() int {
	return len(l.elements)
}

func (l *list[T]) Get(i int) (*T) {
	return &l.elements[i]
}

func (l *list[T]) SafeGet(i int) (*T, error) {
	if (l.IsEmpty() || i >= l.Size()) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return &l.elements[i], nil
}

func (l *list[T]) Clear() {
	l.elements = nil
}