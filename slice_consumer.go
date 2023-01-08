package slice

import "errors"

var (
	ErrSliceEmpty = errors.New("slice empty")
)

// SliceConsumer 用于方便的从slice中按顺序消费数据，即把切片当做一个简单的队列
type SliceConsumer[T any] struct {
	slice []T
	index int

	zero T
}

func NewSliceConsumer[T any](slice []T) *SliceConsumer[T] {
	return &SliceConsumer[T]{
		slice: slice,
		index: 0,
	}
}

func (x *SliceConsumer[T]) IsDone() bool {
	return x.index >= len(x.slice)
}

func (x *SliceConsumer[T]) Peek() T {
	if x.index >= len(x.slice) {
		return x.zero
	}
	return x.slice[x.index]
}

func (x *SliceConsumer[T]) PeekE() (T, error) {
	if x.index >= len(x.slice) {
		return x.zero, ErrSliceEmpty
	}
	return x.slice[x.index], nil
}

func (x *SliceConsumer[T]) Take() T {
	if x.index >= len(x.slice) {
		return x.zero
	}
	result := x.slice[x.index]
	x.index++
	return result
}

func (x *SliceConsumer[T]) TakeE() (T, error) {
	if x.index >= len(x.slice) {
		return x.zero, ErrSliceEmpty
	}
	result := x.slice[x.index]
	x.index++
	return result, nil
}
