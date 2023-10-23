package itertools

import (
	"errors"
	"reflect"
)

func slice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

type Iterator struct {
	items any
	index int
	cycle bool
}

// Iter creates a new Iterator from any slice
func Iter(items any) (*Iterator, error) {
	if !slice(items) {
		return nil, errors.New("Iter() requires a slice")
	}
	return &Iterator{items, -1, false}, nil
}

// Cycle creates an infinite Iterator from any slice
func Cycle(items any) (*Iterator, error) {
	if !slice(items) {
		return nil, errors.New("Cycle() requires a slice")
	}
	return &Iterator{items, -1, true}, nil
}

// HasNext returns true if the Iterator has a next item (always true if cycle mode is enabled)
func (it *Iterator) HasNext() bool {
	if it.cycle {
		return true
	}
	return it.index+1 < reflect.ValueOf(it.items).Len()
}

// Next returns the next item (nil if not set)
func (it *Iterator) Next() any {
	val := reflect.ValueOf(it.items)
	if it.cycle {
		it.index = (it.index + 1) % val.Len()
		return val.Index(it.index).Interface()
	}
	if it.HasNext() {
		it.index++
		return val.Index(it.index).Interface()
	}
	return nil
}

// HasPrev returns true if the Iterator has a previous item (always true if cycle mode is enabled)
func (it *Iterator) HasPrev() bool {
	if it.cycle {
		return true
	}
	return it.index-1 >= 0
}

// Prev returns the previous item (nil if not set)
func (it *Iterator) Prev() any {
	val := reflect.ValueOf(it.items)
	if it.index-1 >= 0 {
		it.index--
		return val.Index(it.index).Interface()
	}
	if it.cycle && val.Len() > 0 {
		it.index = val.Len() - 1
		return val.Index(it.index).Interface()
	}
	return nil
}

// Current returns the current item (nil if not set)
func (it *Iterator) Current() any {
	val := reflect.ValueOf(it.items)
	if it.index >= 0 && it.index < val.Len() {
		return val.Index(it.index).Interface()
	}
	return nil
}

// Len returns the length of the Iterator
func (it *Iterator) Len() int {
	return reflect.ValueOf(it.items).Len()
}

// Index returns the current index (0-based)
func (it *Iterator) Index() int {
	return it.index
}

// IsCycle returns true if the Iterator is in cycle mode
func (it *Iterator) IsCycle() bool {
	return it.cycle
}

// Reset resets the Iterator to the beginning (nil item)
func (it *Iterator) Reset() {
	it.index = -1
}

// SetIndex sets the Iterator to the given index (0-based)
func (it *Iterator) SetIndex(index int) error {
	if it.cycle {
		it.index = index % reflect.ValueOf(it.items).Len()
		return nil
	}
	if index >= 0 && index < reflect.ValueOf(it.items).Len() {
		it.index = index
		return nil
	}

	return errors.New("Index out of range")
}

// SetCycle sets the Iterator to cycle mode (true) or not (false)
func (it *Iterator) SetCycle(cycle bool) {
	it.cycle = cycle
}
