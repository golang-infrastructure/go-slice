package slice

import "math/rand"

// ------------------------------------------------ ---------------------------------------------------------------------

// SubSlice 子切片，不使用指针以防止引用释放不掉
func SubSlice[T any](slice []T, from, to int) []T {
	if to <= from || to < 0 || from < 0 {
		return nil
	}
	capacity := to - from
	subSlice := make([]T, 0, capacity)
	for i := from; i < to && i < len(slice); i++ {
		subSlice[i-from] = slice[from]
	}
	return subSlice
}

// Filter 从切片中选择符合条件的元素
func Filter[T any](slice []T, filterFunc func(value T) bool) []T {
	newSlice := make([]T, 0)
	for _, value := range slice {
		if filterFunc(value) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

// ------------------------------------------------ By Index ---------------------------------------------------------------------

func ChooseIndex[T any](slice []T, index int) T {
	return ChooseIndexOrDefault(slice, index, nil)
}

func ChooseIndexOrDefault[T any](slice []T, index int, defaultValue T) T {
	if index < 0 || index >= len(slice) {
		return defaultValue
	}
	return slice[index]
}

// ChooseIndexes 选择给定的下标的元素
func ChooseIndexes[T any](slice []T, indexes ...int) []T {

	// indexes --> set
	indexSet := make(map[int]struct{}, 0)
	for _, index := range indexes {
		indexSet[index] = struct{}{}
	}

	// filter
	newSlice := make([]T, 0)
	for index, value := range slice {
		if _, exists := indexSet[index]; exists {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

// ChooseOddIndexes 选择奇数下标的元素
func ChooseOddIndexes[T any](slice []T) []T {
	newSlice := make([]T, 0)
	for index, value := range slice {
		if index%2 == 1 {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

// ChooseEvenIndexes 选择偶数下标的元素
func ChooseEvenIndexes[T any](slice []T) []T {
	newSlice := make([]T, 0)
	for index, value := range slice {
		if index%2 == 0 {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

// ------------------------------------------------ Random ---------------------------------------------------------------------

// ChooseRandomIndex 从切片中随机选择一个元素
func ChooseRandomIndex[T any](slice []T) T {
	if len(slice) == 0 {
		return nil
	}
	return slice[rand.Int()%len(slice)]
}

// ChooseRandomIndexes 从切片中随机选择n个元素，会避免重复选择
func ChooseRandomIndexes[T any](slice []T, n int) []T {
	if n == 0 {
		return nil
	}
	// 抖个机灵，借助于map访问时的随机性来实现随机选择
	indexSet := make(map[int]struct{}, 0)
	for index := range slice {
		indexSet[index] = struct{}{}
	}
	newSlice := make([]T, 0)
	for index := range indexSet {
		newSlice = append(newSlice, slice[index])
		if len(newSlice) == n {
			break
		}
	}
	return newSlice
}

// ChooseMiddleIndex 选择最中间的一个下标指向的元素，当数组长度为奇数时有一个返回值，数组长度为偶数时有两个返回值
func ChooseMiddleIndex[T any](slice []T) (T, T) {
	if len(slice)%2 == 0 {
		// 偶数
		return slice[len(slice)/2], slice[len(slice)/2-1]
	} else {
		// 奇数
		return slice[len(slice)/2], nil
	}
}

// FirstItem 选择数组中的第一个元素
func FirstItem[T any](slice []T) T {
	return FirstItemOrDefault(slice, nil)
}

// FirstItemOrDefault 选择数组中的第一个元素，如果数组为空的话则返回默认值
func FirstItemOrDefault[T any](slice []T, defaultValue T) T {
	if len(slice) > 0 {
		return slice[0]
	} else {
		return defaultValue
	}
}

// FirstItems 选择切片最前面几个元素
func FirstItems[T any](slice []T, n int) []T {
	newSlice := make([]T, 0)
	for i := 0; i < n && i < len(slice); i++ {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

// LastItem 返回切片的最后一个元素，如果有的话
func LastItem[T any](slice []T) T {
	return LastItemOrDefault(slice, nil)
}

// LastItemOrDefault 返回切片的最后一个元素，如果切片为空返回给定的默认值
func LastItemOrDefault[T any](slice []T, defaultValue T) T {
	if len(slice) > 0 {
		return slice[len(slice)-1]
	} else {
		return defaultValue
	}
}

// LastItems 选择切片的最后几个元素，会保持其顺序
func LastItems[T any](slice []T, n int) []T {
	newSlice := make([]T, 0)
	beginIndex := len(slice) - n - 1
	if beginIndex < 0 {
		beginIndex = 0
	}
	for ; beginIndex < len(slice); beginIndex++ {
		newSlice = append(newSlice, slice[beginIndex])
	}
	return newSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------
