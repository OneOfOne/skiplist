package skiplist

// LessFn defines a function type to compare keys.
type LessFn func(a, b interface{}) bool

// StringLessFn is a helper function to compare string keys.
func StringLessFn(a, b interface{}) bool {
	as, _ := a.(string)
	bs, _ := b.(string)
	return as < bs
}

// IntLessFn is a helper function to compare Int keys.
func IntLessFn(a, b interface{}) bool {
	as, _ := a.(int)
	bs, _ := b.(int)
	return as < bs
}

// UintLessFn is a helper function to compare uint keys.
func UintLessFn(a, b interface{}) bool {
	as, _ := a.(uint)
	bs, _ := b.(uint)
	return as < bs
}

// Int64LessFn is a helper function to compare Int64 keys.
func Int64LessFn(a, b interface{}) bool {
	as, _ := a.(int64)
	bs, _ := b.(int64)
	return as < bs
}

// Uint64LessFn is a helper function to compare uint64 keys.
func Uint64LessFn(a, b interface{}) bool {
	as, _ := a.(uint64)
	bs, _ := b.(uint64)
	return as < bs
}
