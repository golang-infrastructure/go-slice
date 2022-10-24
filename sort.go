package slice

import (
	"github.com/CC11001100/go-heap"
	"sort"
)

// ---------------------------------------------------------------------------------------------------------------------

// Sort 如果切片中存储的元素本身就是可比较顺序的话，则可以使用这个方法对其排序
func Sort[T Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortByKey 如果切片中的元素本身不是可排序的，但是可以从元素生成一个key，则可以根据key对切片排序
func SortByKey[T any, K Ordered](slice []T, keyFunc OrderedKeyFunc[T, K]) {
	sort.Slice(slice, func(i, j int) bool {
		a := slice[i]
		b := slice[j]
		return keyFunc(-1, a) < keyFunc(-1, b)
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// Reverse 反转slice，可以选择把反转后的放在一个新的slice，或者在原slice原地反转
func Reverse[T any](slice []T, isNewSlice ...bool) []T {
	isNewSlice = append(isNewSlice, false)
	if isNewSlice[0] {
		return reverseInSitu(slice)
	}
	return reverseNewSlice(slice)
}

func reverseNewSlice[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	for index, item := range slice {
		newSlice[len(slice)-index-1] = item
	}
	return newSlice
}
func reverseInSitu[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return slice
}

// ------------------------------------------------ ---------------------------------------------------------------------

// MergeSortedSlices 合并多个已经排好序的切片，需要切片中的元素本身是可比较顺序的，否则请使用 MergeSortedSlicesByKey
func MergeSortedSlices[T Ordered](slice ...[]T) []T {
	return MergeSortedSlicesByKey(func(index int, item T) T {
		return item
	}, slice...)
}

// MergeSortedSlicesByKey 根据key合并多个有序切片，需要多个切片本身就是已经根据key排好序的
func MergeSortedSlicesByKey[T any, K Ordered](keyFunc OrderedKeyFunc[T, K], slices ...[]T) []T {

	consumerSlice := make([]*SliceConsumer[T], 0)
	for _, slice := range slices {
		consumerSlice = append(consumerSlice, NewSliceConsumer(slice))
	}

	resultSlice := make([]T, 0)
	for {
		chooseSliceNo := -1
		for sliceNo, consumer := range consumerSlice {
			if consumer.IsDone() {
				continue
			}
			if chooseSliceNo == -1 {
				chooseSliceNo = sliceNo
			} else {
				keyNow := keyFunc(-1, consumer.Peek())
				keyLast := keyFunc(-1, consumerSlice[chooseSliceNo].Peek())
				if keyNow < keyLast {
					chooseSliceNo = sliceNo
				}
			}
		}
		if chooseSliceNo == -1 {
			break
		}
		result := consumerSlice[chooseSliceNo].Take()
		resultSlice = append(resultSlice, result)
	}
	return resultSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------

// UnionAndSort 多个无序切片合并并排序
func UnionAndSort[T Ordered](slice ...[]T) []T {
	return UnionAndSortByKey(func(index int, item T) T {
		return item
	}, slice...)
}

// UnionAndSortByKey 多个无序切片合并并排序
func UnionAndSortByKey[T any, K Ordered](keyFunc OrderedKeyFunc[T, K], slices ...[]T) []T {

	// 创建一个堆，用来根据key排序
	heap := heap.New(func(a T, b T) int {
		keyA := keyFunc(-1, a)
		keyB := keyFunc(-1, b)
		if keyA == keyB {
			return 0
		} else if keyA < keyB {
			return -1
		} else {
			return 1
		}
	})

	// 然后把所有的元素都放入到堆中
	for _, slice := range slices {
		for _, item := range slice {
			heap.Push(item)
		}
	}

	// 弹出即可
	newSlice := make([]T, 0)
	for heap.IsNotEmpty() {
		newSlice = append(newSlice, heap.Pop())
	}

	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

// IsSorted 判断切片是否已经是有序的，需要切片中的元素本身就是可比较顺序的，否则请使用 IsSortedByKey
func IsSorted[T Ordered](slice []T) bool {
	return IsSortedByKey(slice, func(index int, item T) T {
		return item
	})
}

// IsReverseSorted 判断切片是否是倒序排序的
func IsReverseSorted[T Ordered](slice []T) bool {
	return IsReverseSortedByKey(slice, func(index int, item T) T {
		return item
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// IsReverseSortedByKey 判断切片是否是已经倒序排好序的
func IsReverseSortedByKey[T any, K Ordered](slice []T, keyFunc OrderedKeyFunc[T, K]) bool {
	for i := 0; i < len(slice)-1; i++ {
		keyA := keyFunc(i, slice[i])
		keyB := keyFunc(i+1, slice[i+1])
		if keyA > keyB {
			return false
		}
	}
	return true
}

// IsSortedByKey 判断切片是否是已经排好序的
func IsSortedByKey[T any, K Ordered](slice []T, keyFunc OrderedKeyFunc[T, K]) bool {
	for i := 0; i < len(slice)-1; i++ {
		keyA := keyFunc(i, slice[i])
		keyB := keyFunc(i+1, slice[i+1])
		if keyA > keyB {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------
