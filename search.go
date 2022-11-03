package slice

import compare_anything "github.com/CC11001100/go-compare-anything"

// 这个下面都是从切片中查找元素相关的

// ------------------------------------------------ 无序切片 ---------------------------------------------------------------------

// Contains 判断切片是否包含给定的元素，此切片可以是无序的
func Contains[T comparable](slice []T, exceptedItem T) bool {
	return FindIndex(slice, exceptedItem) != -1
}

func NotContains[T comparable](slice []T, exceptedItem T) bool {
	return !Contains(slice, exceptedItem)
}

// FindIndex 查找第一次出现的下表
func FindIndex[T comparable](slice []T, exceptedItem T) int {
	for index, item := range slice {
		if exceptedItem == item {
			return index
		}
	}
	return -1
}

func FindIndexByFunc[T any](slice []T, exceptedItem T, equalsFunc compare_anything.EqualsFunc[T]) int {
	for index, item := range slice {
		if equalsFunc(exceptedItem, item) {
			return index
		}
	}
	return -1
}

func FindLastIndex[T comparable](slice []T, item T) int {
	for index := len(slice) - 1; index >= 0; index-- {
		if slice[index] == item {
			return index
		}
	}
	return -1
}

func FindAllIndex[T comparable](slice []T, exceptedItem T) []int {
	indexes := make([]int, 0)
	for index, item := range slice {
		if item == exceptedItem {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

// ContainsAll 第一个切片是否包含第二个切片中的全部元素
func ContainsAll[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; !exists {
			return false
		}
	}
	return true
}

// ContainsAny 第一个切片中是否包含第二个切片中的任意元素
func ContainsAny[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; exists {
			return true
		}
	}
	return false
}

// NotContainsAny 检查第一个切片中是否不包含第二个切片中的任何元素
func NotContainsAny[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; exists {
			return false
		}
	}
	return true
}

// NotContainsAll 检查第一个切片是否未包含第二个切片中的全部元素
func NotContainsAll[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; !exists {
			return false
		}
	}
	return true
}

// ------------------------------------------------ 有序切片 ---------------------------------------------------------------------

// OrderedContains 判断切片是否包含给定的元素，此切片可以是无序的
func OrderedContains[T comparable](slice []T, exceptedItem T) bool {
	return FindIndex(slice, exceptedItem) != -1
}

func OrderedNotContains[T comparable](slice []T, exceptedItem T) bool {
	return !Contains(slice, exceptedItem)
}

// OrderedFindIndex 查找第一次出现的下表
func OrderedFindIndex[T comparable](slice []T, exceptedItem T) int {
	for index, item := range slice {
		if exceptedItem == item {
			return index
		}
	}
	return -1
}

func OrderedFindLastIndex[T comparable](slice []T, item T) int {
	for index := len(slice) - 1; index >= 0; index-- {
		if slice[index] == item {
			return index
		}
	}
	return -1
}

func OrderedFindAllIndex[T comparable](slice []T, exceptedItem T) []int {
	indexes := make([]int, 0)
	for index, item := range slice {
		if item == exceptedItem {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

// OrderedContainsAll 第一个切片是否包含第二个切片中的全部元素
func OrderedContainsAll[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; !exists {
			return false
		}
	}
	return true
}

// OrderedContainsAny 第一个切片中是否包含第二个切片中的任意元素
func OrderedContainsAny[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; exists {
			return true
		}
	}
	return false
}

// OrderedNotContainsAny 检查第一个切片中是否不包含第二个切片中的任何元素
func OrderedNotContainsAny[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; exists {
			return false
		}
	}
	return true
}

// OrderedNotContainsAll 检查第一个切片是否未包含第二个切片中的全部元素
func OrderedNotContainsAll[T comparable](slice []T, exceptedItems []T) bool {
	sliceSet := ToSet(slice)
	for _, exceptedItem := range exceptedItems {
		if _, exists := sliceSet[exceptedItem]; !exists {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------
