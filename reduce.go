package slice

import compare_anything "github.com/CC11001100/go-compare-anything"

// ---------------------------------------------------------------------------------------------------------------------

type ReduceFunc[T, R any] func(index int, item T, result R) R

func Reduce[T, R any](slice []T, initResult R, reduceFunc ReduceFunc[T, R]) R {
	result := initResult
	for index, item := range slice {
		result = reduceFunc(index, item, result)
	}
	return result
}

// ---------------------------------------------------------------------------------------------------------------------

func Max[T any](slice []T, comparator compare_anything.Comparator[T]) T {
	return Reduce(slice, nil, func(index int, item T, result T) T {
		if comparator(item, result) > 0 {
			return item
		} else {
			return result
		}
	})
}

func Min() {

}

func Avg() {

}

func Sum[T int | int8]() {

}
