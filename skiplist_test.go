package skiplist

import (
	"strconv"
	"testing"
)

func TestSkipList(t *testing.T) {
	const N = 1e6
	sl := NewWithSeed(32, 42)
	for i := 0; i < N; i++ {
		sl.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < N; i++ {
		if v, _ := sl.Get(strconv.Itoa(i)).(int); v != i {
			t.Fatalf("%d is not exact: %v", i, v)
		}
	}

}
