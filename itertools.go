package itertools

import (
	"errors"
	"reflect"
)

func slice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

type iterator struct {
	items any
	index int
	cycle bool
}

// Iter creates a new iterator from any slice
func Iter(items any) (*iterator, error) {
	if !slice(items) {
		return nil, errors.New("Iter() requires a slice")
	}
	return &iterator{items, -1, false}, nil
}

// Cycle creates an infinite iterator from any slice
func Cycle(items any) (*iterator, error) {
	if !slice(items) {
		return nil, errors.New("Cycle() requires a slice")
	}
	return &iterator{items, -1, true}, nil
}

// HasNext returns true if the iterator has a next item (always true if cycle mode is enabled)
func (it *iterator) HasNext() bool {
	if it.cycle {
		return true
	}
	return it.index+1 < reflect.ValueOf(it.items).Len()
}

// Next returns the next item (nil if not set)
func (it *iterator) Next() any {
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

// HasPrev returns true if the iterator has a previous item (always true if cycle mode is enabled)
func (it *iterator) HasPrev() bool {
	if it.cycle {
		return true
	}
	return it.index-1 >= 0
}

// Prev returns the previous item (nil if not set)
func (it *iterator) Prev() any {
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
func (it *iterator) Current() any {
	val := reflect.ValueOf(it.items)
	if it.index >= 0 && it.index < val.Len() {
		return val.Index(it.index).Interface()
	}
	return nil
}

// Len returns the length of the iterator
func (it *iterator) Len() int {
	return reflect.ValueOf(it.items).Len()
}

// Index returns the current index (0-based)
func (it *iterator) Index() int {
	return it.index
}

// IsCycle returns true if the iterator is in cycle mode
func (it *iterator) IsCycle() bool {
	return it.cycle
}

// Reset resets the iterator to the beginning (nil item)
func (it *iterator) Reset() {
	it.index = -1
}

// SetIndex sets the iterator to the given index (0-based)
func (it *iterator) SetIndex(index int) error {
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

// SetCycle sets the iterator to cycle mode (true) or not (false)
func (it *iterator) SetCycle(cycle bool) {
	it.cycle = cycle
}
