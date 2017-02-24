package skiplist

import (
	"fmt"
	"strconv"
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
	keys := make([]string, N)
	for i := range keys {
		keys[i] = pad(i)
	}
	b.ResetTimer()

	for _, n := range [...]int{100, 1e3, 1e5, 1e6} {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			benchmarkSetGet(b, keys[:n])
		})
	}
}

func benchmarkSetGet(b *testing.B, keys []string) {
	sl := NewCustom(32, 0.50, 42)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := keys[i%len(keys)]
			sl.Set(kv, kv)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := keys[i%len(keys)]
			if v, _ := sl.Get(kv).(string); v != kv {
				b.Fatalf("expected %v, got %v", kv, v)
			}
		}
	})
}
