package slice

import (
	compare_anything "github.com/CC11001100/go-compare-anything"
)

// ---------------------------------------------------------------------------------------------------------------------

func Shift() {

}

func UnShift() {

}

// ------------------------------------------------ ---------------------------------------------------------------------

// All 判断切片中的所有元素符合条件
func All[T any](slice []T, judgeFunc func(index int, item T) bool) bool {
	for index, item := range slice {
		if !judgeFunc(index, item) {
			return false
		}
	}
	return true
}

// Any 判断切片中的任意一个元素符合条件
func Any[T any](slice []T, judgeFunc func(index int, item T) bool) bool {
	for index, item := range slice {
		if judgeFunc(index, item) {
			return true
		}
	}
	return false
}

// ------------------------------------------------ ---------------------------------------------------------------------

// IsUniq 判断切片中的元素是否不重复
func IsUniq[T comparable](slice []T) bool {
	distinctSet := make(map[T]struct{}, 0)
	for _, item := range slice {
		if _, exists := distinctSet[item]; exists {
			return false
		}
		distinctSet[item] = struct{}{}
	}
	return true
}

// IsUniqByFunc 判断切片中的元素是否不重复，以自定义的函数为每个元素生成一个唯一标识
func IsUniqByFunc[T any, K comparable](slice []T, idFunc UniqIDFunc[T, K]) bool {
	distinctSet := make(map[K]struct{})
	for _, item := range slice {
		uniqId := idFunc(item)
		if _, exists := distinctSet[uniqId]; exists {
			return false
		}
		distinctSet[uniqId] = struct{}{}
	}
	return true
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Equals 比较两个切片是否相等
func Equals[T comparable](sliceA []T, sliceB []T) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, itemA := range sliceA {
		itemB := sliceB[index]
		if itemA != itemB {
			return false
		}
	}
	return true
}

// EqualsByFunc 比较两个切片是否相等
func EqualsByFunc[T any](sliceA []T, sliceB []T, comparator compare_anything.Comparator[T]) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, itemA := range sliceA {
		itemB := sliceB[index]
		if comparator(itemA, itemB) != 0 {
			return false
		}
	}
	return true
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Insert 往切片中插入元素，会引用原来的数组
func Insert[T any](slice []T, newValue T, insertToIndex int) []T {
	// 保证类型安全
	if insertToIndex < 0 || insertToIndex > len(slice) {
		return slice
	}
	// 先插入到最后
	slice = append(slice, newValue)
	// 然后再一点一点往前挪动

	return slice
	for {

	}
}

// InsertWithCopy 往切片中插入元素，会拷贝原切片到一个新的切片中
func InsertWithCopy() {

}
