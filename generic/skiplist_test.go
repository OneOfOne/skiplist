package skiplist

import "testing"

const N = 1e6

func TestSkipList(t *testing.T) {
	sl := NewCustom(32, 0.25, IntCompareFn, 42)
	for i := 0; i < N; i++ {
		sl.Set(i, i)
	}

	for i := 0; i < N; i++ {
		if v, ok := sl.Get(i).(int); !ok || v != i {
			t.Fatalf("%d is not exact: %v", i, sl.Get(i))
		}
	}

	for i, it := 0, sl.IteratorAt(0); it.HasMore(); it.Next() {
		if it.Key() != i || it.Value() != i {
			t.Errorf("expected %d, got (%v, %v)", i, it.Key(), it.Value())
		}
		i++
	}

	t.Logf("Len: %d, Level: %d, MaxLevel: %d", sl.Len(), sl.Level(), sl.MaxLevel())
}

func BenchmarkSetGet(b *testing.B) {
	sl := NewCustom(32, 0.25, IntCompareFn, 42)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := i % N
			sl.Set(kv, kv)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			exp := i % N
			if v, _ := sl.Get(exp).(int); v != exp {
				b.Fatalf("expected %v, got %v", exp, v)
			}
		}
	})

}
