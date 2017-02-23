package skiplist

import "testing"

const N = 1e6

func TestSkipList(t *testing.T) {
	sl := NewCustom(32, 0.25, IntLessFn, 42)
	for i := 0; i < N; i++ {
		sl.Set(i, i)
	}

	for i := 0; i < N; i++ {
		if v, _ := sl.Get(i).(int); v != i {
			t.Fatalf("%d is not exact: %v", i, v)
		}
	}

	t.Logf("Len: %d, Level: %d, MaxLevel: %d", sl.Len(), sl.Level(), sl.MaxLevel())
	t.Log(sl.At(55))

}

func BenchmarkGet(b *testing.B) {
	sl := NewCustom(32, 0.25, IntLessFn, 42)
	for i := 0; i < N; i++ {
		sl.Set(i, i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		exp := i % N
		if v, _ := sl.Get(exp).(int); v != exp {
			b.Fatalf("expected %v, got %v", exp, v)
		}
	}
}
