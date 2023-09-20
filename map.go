package goco

import "fmt"

type simpleMap[K comparable, V any] struct {
	elements map[K]V
}

func NewMap[K comparable, V any]() *simpleMap[K, V] {
	return NewMapWithSize[K, V](50)
}

func NewMapWithSize[K comparable, V any](size int) *simpleMap[K, V] {
	return &simpleMap[K, V]{make(map[K]V, size)}
}

func (m *simpleMap[K, V]) String() string {
	return fmt.Sprintf("%v", m.elements)
}

func (m *simpleMap[K, V]) Put(k K, v V) {
	m.elements[k] = v
}

func (m *simpleMap[K, V]) Remove(k K) {
	delete(m.elements, k)
}

func (m *simpleMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.elements[k]
	return ok
}

func (m *simpleMap[K, V]) Keys() []K {
	keys := make([]K, len(m.elements))
	i := 0
	for k := range m.elements {
		keys[i] = k
		i++
	}
	return keys
}

func (m *simpleMap[K, V]) Clear() {
	m.elements = nil
}