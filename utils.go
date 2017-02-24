package skiplist

import "math"

// GuessLevel returns the maximum level you can pass to new to get enough space for X elements.
func GuessLevel(x int) int {
	for lvl := 1; ; lvl++ {
		if math.Pow(2, float64(lvl)) >= float64(x) {
			return lvl
		}
	}
}

// Iterator represent a forward-only iterator.
// TODO: support backwards operations.
type Iterator struct {
	n *node
}

// HasMore returns true if there are more items in the list.
func (it *Iterator) HasMore() bool {
	return it.n != nil
}

// Next moves the iterator to the next item and returns true if there are more items in the list.
func (it *Iterator) Next() bool {
	it.n = it.n.next[0]
	return it.n != nil
}

// Key is the current iterator key.
func (it *Iterator) Key() string { return it.n.k }

// Value is the current iterator value.
func (it *Iterator) Value() interface{} { return it.n.v }
