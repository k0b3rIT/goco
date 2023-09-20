package goco

import "golang.org/x/exp/slices"

func SliceTomap[T any, V comparable](src []T, key func(T) V) map[V]T {
	var result = make(map[V]T)
	for _, v := range src {
		result[key(v)] = v
	}
	return result
}

func SliceToMap[T any, V comparable](src []T, key func(T) V) *simpleMap[V, T] {
	var result = make(map[V]T)
	for _, v := range src {
		result[key(v)] = v
	}
	return NewMapFrom[V, T](result)
}

func ListToMap[T comparable, V comparable](src List[T], key func(T) V) *simpleMap[V, T] {
	var result = make(map[V]T)
	for _, v := range src.ToSlice() {
		result[key(v)] = v
	}
	return NewMapFrom[V, T](result)
}

func ListHasTheSameElements[T comparable](a List[T], b List[T]) bool {
	if a.Size() != b.Size() {
		return false
	}
	for _, v := range a.ToSlice() {
		if !b.Contains(v) {
			return false
		}
	}
	return true
}

func SlicesHasTheSameElements[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		if !slices.Contains(b, v) {
			return false
		}
	}
	return true
}
