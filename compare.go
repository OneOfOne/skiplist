package skiplist

// CompareFn defines a function type to compare keys.
type CompareFn func(a, b interface{}) int

// StringCompareFn is a helper function to compare string keys.
func StringCompareFn(a, b interface{}) int {
	as, _ := a.(string)
	bs, _ := b.(string)
	switch {
	case as < bs:
		return -1
	case as == bs:
		return 0
	default:
		return 1
	}
}

// IntCompareFn is a helper function to compare Int keys.
func IntCompareFn(a, b interface{}) int {
	as, _ := a.(int)
	bs, _ := b.(int)
	switch {
	case as < bs:
		return -1
	case as == bs:
		return 0
	default:
		return 1
	}
}

// UintCompareFn is a helper function to compare uint keys.
func UintCompareFn(a, b interface{}) int {
	as, _ := a.(uint)
	bs, _ := b.(uint)
	switch {
	case as < bs:
		return -1
	case as == bs:
		return 0
	default:
		return 1
	}
}

// Int64CompareFn is a helper function to compare Int64 keys.
func Int64CompareFn(a, b interface{}) int {
	as, _ := a.(int64)
	bs, _ := b.(int64)
	switch {
	case as < bs:
		return -1
	case as == bs:
		return 0
	default:
		return 1
	}
}

// Uint64CompareFn is a helper function to compare uint64 keys.
func Uint64CompareFn(a, b interface{}) int {
	as, _ := a.(uint64)
	bs, _ := b.(uint64)
	switch {
	case as < bs:
		return -1
	case as == bs:
		return 0
	default:
		return 1
	}
}
