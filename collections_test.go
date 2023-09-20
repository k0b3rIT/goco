package goco

import (
	"testing"
)

type Car struct {
	Brand string
	Model string
}

func TestList(t *testing.T) {

	l := NewList[Car](Car{"bmw", "x5"}, Car{"audi", "a4"}, Car{"mercedes", "c180"})

	l.Add(Car{"volkswagen", "golf"})

	if l.Size() != 4 {
		t.Errorf("got %q, wanted %q", l.Size(), 4)
	}

	l.Remove(Car{"bmw", "x5"})

	if l.Size() != 3 {
		t.Errorf("got %q, wanted %q", l.Size(), 3)
	}

	l.Clear()

	if !l.IsEmpty() {
		t.Errorf("got %q, wanted %q", l.Size(), 0)
	}
}

func TestToMap(t *testing.T) {
	l := NewList[Car](Car{"bmw", "x5"}, Car{"audi", "a4"}, Car{"mercedes", "c180"})
	lmap := ListToMap[Car](l, func(c Car) string { return c.Brand })

	if !SlicesHasTheSameElements[string](lmap.Keys(), []string{"bmw", "audi", "mercedes"}) {
		t.Errorf("got %q, wanted %q", lmap.Keys(), []string{"bmw", "audi", "mercedes"})
	}

	lmap.Remove("bmw")

	if !SlicesHasTheSameElements[string](lmap.Keys(), []string{"audi", "mercedes"}) {
		t.Errorf("got %q, wanted %q", lmap.Keys(), []string{"audi", "mercedes"})
	}

}
