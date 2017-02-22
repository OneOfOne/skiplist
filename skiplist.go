package skiplist

import (
	"math/rand"
	"time"
)

type node struct {
	k    string
	v    interface{}
	next []*node
}

// List represents a normal skiplist keyed by a string
type List struct {
	head *node
	len  uint64
	rnd  *rand.Rand

	update []*node // update buffer
}

// New is an alias for NewWithSeed(maxHeight, time.Now().Unix()).
func New(maxHeight int) *List { return NewWithSeed(maxHeight, time.Now().Unix()) }

// NewWithSeed returns a new skiplist with the specified max height and random seed.
func NewWithSeed(maxHeight int, seed int64) *List {
	return &List{
		head:   &node{next: make([]*node, maxHeight)},
		rnd:    rand.New(rand.NewSource(seed)),
		update: make([]*node, maxHeight),
	}
}

func (sl *List) maxHeight() int { return len(sl.update) }

func (sl *List) findNodeAndPrev(k string, prev []*node) (n *node) {
	n = sl.head
	for i := sl.maxHeight() - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil && next.k < k; next = n.next[i] {
			n = next
		}
		prev[i] = n
	}
	return
}

// Set assigns a key to a value, returns true if the key didn't already exist.
func (sl *List) Set(k string, v interface{}) (added bool) {
	var (
		n       = sl.findNodeAndPrev(k, sl.update)
		nheight int
	)

	if n = n.next[0]; n != nil && n.k == k {
		n.v = v
		return
	}

	nheight = sl.newHeight()
	n = &node{
		k:    k,
		v:    v,
		next: make([]*node, nheight),
	}

	for i := 0; i < nheight; i++ {
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

// Get returns the value if found, otherwise nil.
func (sl *List) Get(k string) interface{} {
	n := sl.head
	for i := sl.maxHeight() - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil; next = n.next[i] {
			if next.k <= k {
				n = next
			} else {
				break
			}
		}
	}

	if n.k == k {
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

func (sl *List) newHeight() (nheight int) {
	mh := sl.maxHeight()
	nheight = 1
	for ; nheight <= mh && sl.rnd.Int63()%100 < 50; nheight++ {
	}

	return nheight
}
