package skiplist

import (
	"fmt"
	"testing"
)

const N = 1e6

func pad(i int) string {
	return fmt.Sprintf("%08d", i)
}

func TestSkipList(t *testing.T) {
	sl := NewCustom(32, 0.25, 42)
	for i := 0; i < N; i++ {
		i := pad(i)
		sl.Set(i, i)
	}

	for i := 0; i < N; i++ {
		i := pad(i)
		if v, ok := sl.Get(i).(string); !ok || v != i {
			t.Fatalf("%v is not exact: %v", i, sl.Get(i))
		}
	}

	for i, it := 0, sl.IteratorAt("00000000"); it.HasMore(); it.Next() {
		is := pad(i)
		if it.Key() != is || it.Value() != is {
			t.Fatalf("expected %d, got (%v, %v)", i, it.Key(), it.Value())
		}
		i++
	}

	t.Logf("Len: %d, Level: %d, MaxLevel: %d", sl.Len(), sl.Level(), sl.MaxLevel())
}

func BenchmarkSetGet(b *testing.B) {
	sl := NewCustom(32, 0.25, 42)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := pad(i % N)
			sl.Set(kv, kv)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			exp := pad(i % N)
			if v, _ := sl.Get(exp).(string); v != exp {
				b.Fatalf("expected %v, got %v", exp, v)
			}
		}
	})

}
