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
	cmpFn  CompareFn
	update []*node // update buffer
	prob   float64

	len   int
	level int
}

// New is an alias for NewCustom(maxlevel, DefaultProbability, cmpFn, time.Now().Unix()).
func New(maxlevel int, cmpFn CompareFn) *List {
	return NewCustom(maxlevel, DefaultProbability, cmpFn, time.Now().Unix())
}

// NewCustom returns a new skiplist with the specified max level and random seed.
func NewCustom(maxlevel int, prob float64, cmpFn CompareFn, seed int64) *List {
	return &List{
		head:   &node{next: make([]*node, maxlevel)},
		rnd:    rand.New(rand.NewSource(seed)),
		cmpFn:  cmpFn,
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
	var checked *node
	n = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := n.next[i]; next != nil && next != checked && sl.cmpFn(next.k, k) < 0; next = n.next[i] {
			n = next
		}
		checked = n.next[i]
		sl.update[i] = n
	}
	return
}

// Set assigns a key to a value, returns true if the key didn't already exist.
func (sl *List) Set(k, v interface{}) (added bool) {
	if n := sl.findAndUpdate(k).next[0]; n != nil && sl.cmpFn(k, n.k) == 0 {
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
			sl.update[i] = nil
		}
	}

	sl.len++

	return true
}

// Get returns the value if found, otherwise nil.
func (sl *List) Get(k interface{}) interface{} {
	for n, i := sl.head, sl.level-1; i >= 0; i-- {
	L:
		for next := n.next[i]; next != nil; next = n.next[i] {
			switch sl.cmpFn(next.k, k) {
			case -1:
				n = next
			case 0:
				return next.v
			default:
				break L
			}
		}
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
