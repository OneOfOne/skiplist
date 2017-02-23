package skiplist

import "testing"

const N = 1e6

func TestSkipList(t *testing.T) {
	sl := NewCustom(32, 0.25, IntCompareFn, 42)
	for i := 0; i < N; i++ {
		sl.Set(i, i)
	}

	for i := 0; i < N; i++ {
		if v, _ := sl.Get(i).(int); v != i {
			t.Fatalf("%d is not exact: %v", i, v)
		}
	}

	t.Logf("Len: %d, Level: %d, MaxLevel: %d", sl.Len(), sl.Level(), sl.MaxLevel())
}

func BenchmarkGet(b *testing.B) {
	sl := NewCustom(32, 0.25, IntCompareFn, 42)

	b.ResetTimer()
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
