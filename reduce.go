package slice

import compare_anything "github.com/CC11001100/go-compare-anything"

// ---------------------------------------------------------------------------------------------------------------------

type ReduceFunc[T, R any] func(index int, item T, result R) R

// Reduce 把切片中的所有元素合并为一个
func Reduce[T, R any](slice []T, initResult R, reduceFunc ReduceFunc[T, R]) R {
	result := initResult
	for index, item := range slice {
		result = reduceFunc(index, item, result)
	}
	return result
}

// ReduceByKey 根据key把切片中的元素分组并合并返回
func ReduceByKey[T any, K comparable, R any](slice []T, keyFunc KeyFunc[T, K], reduceFunc ReduceFunc[T, R]) map[K]R {
	keyMap := make(map[K]R, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
		result := reduceFunc(index, item, keyMap[key])
		keyMap[key] = result
	}
	return keyMap
}

// ---------------------------------------------------------------------------------------------------------------------

// Max 求出切片中的最大值
func Max[T any](slice []T, comparator compare_anything.Comparator[T]) T {
	if len(slice) == 0 {
		return nil
	}
	return Reduce(slice[1:], slice[0], func(index int, item T, result T) T {
		if comparator(item, result) > 0 {
			return item
		} else {
			return result
		}
	})
}

// Min 求出切片中的最小值
func Min[T any](slice []T, comparator compare_anything.Comparator[T]) T {
	if len(slice) == 0 {
		return nil
	}
	return Reduce(slice[1:], slice[0], func(index int, item T, result T) T {
		if comparator(item, result) < 0 {
			return item
		} else {
			return result
		}
	})
}

// TODO 2022-10-24 03:03:24 有点难搞下面这几个，可能要引入新的概念

func Avg() {
}

func Sum() {

}
