package slice

import (
	"github.com/golang-infrastructure/go-gtypes"
)

// 这个文件中是多个切片之间的操作

// aggregate
// aggregateByKey

// Cartesian 笛卡尔积，把两个切片相乘
func Cartesian[A, B, C any](sliceA []A, sliceB []B, multiplicationFunc func(a A, b B) C) []C {
	slice := make([]C, 0)
	for _, itemA := range sliceA {
		for _, itemB := range sliceB {
			slice = append(slice, multiplicationFunc(itemA, itemB))
		}
	}
	return slice
}

// pipe

// Foreach 对数组中的每个元素处理
func Foreach[T any](slice []T, eachFunc func(index int, item T) bool) {
	for index, item := range slice {
		if !eachFunc(index, item) {
			break
		}
	}
}

// Zip 把两个切片压成k/v的形式
func Zip[K comparable, V any](keySlice []K, valueSlice []V) map[K]V {
	if len(keySlice) != len(valueSlice) {
		return nil
	}
	resultMap := make(map[K]V, 0)
	for index, keyItem := range keySlice {
		resultMap[keyItem] = valueSlice[index]
	}
	return resultMap
}

// ------------------------------------------------ ---------------------------------------------------------------------

// MergeOrderedSlices 合并任意多个有序数组
func MergeOrderedSlices[T gtypes.Ordered](slices ...[]T) []T {
	return MergeOrderedSlicesByComparator(func(a, b T) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		} else {
			return -1
		}
	}, slices...)
}

// MergeOrderedSlicesByComparator 合并任意多个有序数组
func MergeOrderedSlicesByComparator[T any](comparator func(a, b T) int, slices ...[]T) []T {

	// 先把所有的切片搞成可以消费的
	sliceConsumerSet := make(map[*SliceConsumer[T]]struct{})
	for _, slice := range slices {
		sliceConsumerSet[NewSliceConsumer(slice)] = struct{}{}
	}

	// 然后开始消费
	newSlice := make([]T, 0)
	for len(sliceConsumerSet) != 0 {

		// 找出最小的那个
		var minSliceConsumer *SliceConsumer[T]
		for sliceConsumer := range sliceConsumerSet {
			if sliceConsumer.IsDone() {
				delete(sliceConsumerSet, sliceConsumer)
				continue
			}
			if minSliceConsumer == nil || comparator(minSliceConsumer.Peek(), sliceConsumer.Peek()) > 0 {
				minSliceConsumer = sliceConsumer
			}
		}

		take := minSliceConsumer.Take()
		newSlice = append(newSlice, take)
		if minSliceConsumer.IsDone() {
			delete(sliceConsumerSet, minSliceConsumer)
		}
	}

	return newSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------
