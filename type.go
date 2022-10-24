package slice

// ------------------------------------------------ ---------------------------------------------------------------------

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface {
	Signed | Unsigned
}
type Float interface {
	~float32 | ~float64
}
type Complex interface {
	~complex64 | ~complex128
}
type Ordered interface {
	Integer | Float | ~string
}

// ------------------------------------------------ ---------------------------------------------------------------------

// KeyFunc 为元素生成一个Key
type KeyFunc[T any, K any] func(index int, item T) K

// ComparableKeyFunc 为元素生成一个key，这个key要能够参与比较是否相等
type ComparableKeyFunc[T any, K comparable] func(index int, item T) K

// OrderedKeyFunc ComparableKeyFunc 为元素生成一个key，这个key要能够参与排序
type OrderedKeyFunc[T any, K Ordered] func(index int, item T) K

// ------------------------------------------------ ---------------------------------------------------------------------

// UniqIDFunc 为元素生成一个唯一ID，这个唯一ID要能够参与比较，不然怎么证明它是唯一的呢
type UniqIDFunc[T any, K comparable] func(item T) K

// ------------------------------------------------ ---------------------------------------------------------------------
