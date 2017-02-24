package skiplist

import (
	"math/rand"
	"time"
)

// DefaultProbability the default level probability for new lists.
var DefaultProbability = 0.25

type node struct {
	k    string
	v    interface{}
	next []*node
}

// List represents a normal skiplist keyed by a string.
type List struct {
	head   *node
	rnd    *rand.Rand
	update []*node // update buffer
	prob   float64

	len   int
	level int
}

// New is an alias for NewCustom(maxlevel, DefaultProbability, cmpFn, time.Now().Unix()).
func New(maxlevel int) *List {
	return NewCustom(maxlevel, DefaultProbability, time.Now().Unix())
}

// NewCustom returns a new skiplist with the specified max level and random seed.
func NewCustom(maxlevel int, prob float64, seed int64) *List {
	return &List{
		head:   &node{next: make([]*node, maxlevel)},
		rnd:    rand.New(rand.NewSource(seed)),
		update: make([]*node, maxlevel),
		prob:   prob,
	}
}

// MaxLevel returns the list's max level
func (sl *List) MaxLevel() int { return cap(sl.update) }

// Level returns the current level
func (sl *List) Level() int { return sl.level }

// Len returns the length of the list.
func (sl *List) Len() int { return sl.len }

func (sl *List) findAndUpdate(k string) *node {
	n := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil && next.k < k; next = n.next[i] {
			n = next
		}
		sl.update[i] = n
	}
	return n.next[0]
}

func (sl *List) find(k string) *node {
	n := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil && next.k < k; next = n.next[i] {
			n = next
		}
	}
	return n.next[0]
}

// Set assigns a key to a value, returns true if the key didn't already exist.
func (sl *List) Set(k string, v interface{}) (added bool) {
	n := sl.findAndUpdate(k)

	if n != nil && k == n.k {
		n.v = v
		return
	}

	n = &node{
		k:    k,
		v:    v,
		next: make([]*node, sl.newLevel()),
	}

	for i := range n.next {
		if up := sl.update[i]; up != nil {
			tmp := up.next[i]
			n.next[i] = tmp
			up.next[i] = n
			sl.update[i] = nil
		}
	}

	sl.len++

	return true
}

// Get returns the value if found, otherwise nil.
func (sl *List) Get(k string) interface{} {
	if n := sl.find(k); n.k == k {
		return n.v
	}
	return nil
}

// ForEach provides an easy way to loop over the list.
// if fn returns true, it breaks early.
func (sl *List) ForEach(fn func(k string, v interface{}) (breakNow bool)) bool {
	for n := sl.head.next[0]; n != nil; n = n.next[0] {
		if fn(n.k, n.v) {
			return true
		}
	}
	return false
}

func (sl *List) newLevel() (nlevel int) {
	mh := sl.MaxLevel()
	nlevel = 1
	for ; nlevel < mh && sl.rnd.Float64() < sl.prob; nlevel++ {
	}

	if nlevel > sl.level {
		for i := sl.level; i < nlevel; i++ {
			sl.update[i] = sl.head
		}
		sl.level = nlevel
	}
	return nlevel
}

// IteratorAt returns an iterator starting at the specific key.
// Example:
//	for it := sl.IteratorAt(0); it.HasMore(); it.Next() {
//		key, value := it.Key(), it.Value()
//	}
func (sl *List) IteratorAt(k string) *Iterator {
	return &Iterator{sl.find(k)}
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
