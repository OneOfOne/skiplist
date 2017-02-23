package skiplist

import (
	"math/rand"
	"time"
)

// DefaultProbability the default level probability for new lists.
var DefaultProbability = 0.25

type node struct {
	k    interface{}
	v    interface{}
	next []*node
}

// List represents a normal skiplist keyed by a string
type List struct {
	head   *node
	rnd    *rand.Rand
	lessFn LessFn
	update []*node // update buffer
	prob   float64

	len   int
	level int
}

// New is an alias for NewCustom(maxlevel, DefaultProbability, lessFn, time.Now().Unix()).
func New(maxlevel int, lessFn LessFn) *List {
	return NewCustom(maxlevel, DefaultProbability, lessFn, time.Now().Unix())
}

// NewCustom returns a new skiplist with the specified max level and random seed.
func NewCustom(maxlevel int, prob float64, lessFn LessFn, seed int64) *List {
	return &List{
		head:   &node{next: make([]*node, maxlevel)},
		rnd:    rand.New(rand.NewSource(seed)),
		lessFn: lessFn,
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

func (sl *List) findAndUpdate(k interface{}) (n *node) {
	n = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil && sl.lessFn(next.k, k); next = n.next[i] {
			n = next
		}
		sl.update[i] = n
	}
	return
}

// Set assigns a key to a value, returns true if the key didn't already exist.
func (sl *List) Set(k, v interface{}) (added bool) {
	if n := sl.findAndUpdate(k).next[0]; n != nil && !sl.lessFn(k, n.k) {
		n.v = v
		return
	}

	nlevel := sl.newLevel()
	n := &node{
		k:    k,
		v:    v,
		next: make([]*node, nlevel),
	}

	for i := 0; i < nlevel; i++ {
		if up := sl.update[i]; up != nil {
			tmp := up.next[i]
			n.next[i] = tmp
			up.next[i] = n
		}
		sl.update[i] = nil
	}

	sl.len++

	return true
}

// At returns a key/value by index.
func (sl *List) At(idx int) (interface{}, interface{}, bool) {
	n := sl.head
	for i := 0; i < sl.level; i++ {
		for next := n.next[i]; next != nil; next = n.next[i] {
			n = next
			if idx--; idx == 0 {
				return n.k, n.v, true
			}
		}
	}
	return nil, nil, false
}

// Get returns the value if found, otherwise nil.
func (sl *List) Get(k interface{}) interface{} {
	n := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil; next = n.next[i] {
			if sl.lessFn(next.k, k) {
				n = next
			} else {
				break
			}
		}
	}

	if n = n.next[0]; n != nil && !sl.lessFn(k, n.k) {
		return n.v
	}
	return nil
}

// ForEach provides an easy way to loop over the list.
// if fn returns true, it breaks early.
func (sl *List) ForEach(fn func(k interface{}, v interface{}) (breakNow bool)) bool {
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
	for ; nlevel <= mh && sl.rnd.Float64() < sl.prob; nlevel++ {
	}

	if nlevel > sl.level {
		sl.level = nlevel
	}

	return nlevel
}
