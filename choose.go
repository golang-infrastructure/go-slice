package slice

// 这个文件中的方法应该都是下标安全的

// ---------------------------------------------------------------------------------------------------------------------

// SubSlice 子切片，不使用指针以防止引用释放不掉
func SubSlice[T any](slice []T, from, to int) []T {
	if from < 0 {
		from = 0
	}
	if to > len(slice) {
		to = len(slice)
	}
	if to <= from || len(slice) == 0 {
		return nil
	}
	capacity := to - from
	subSlice := make([]T, capacity, capacity)
	for i := from; i < to && i < len(slice); i++ {
		subSlice[i-from] = slice[i]
	}
	return subSlice
}

// Filter 从切片中选择符合条件的元素
func Filter[T any](slice []T, filterFunc func(index int, item T) bool) []T {
	newSlice := make([]T, 0)
	for index, item := range slice {
		if filterFunc(index, item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

// ------------------------------------------------ By Index ---------------------------------------------------------------------

// ChooseIndex 返回切片中给定下表的元素，如果不存在的话返回对应类型的零值
func ChooseIndex[T any](slice []T, index int) T {
	var zero T
	return ChooseIndexOrDefault(slice, index, zero)
}

// ChooseIndexOrDefault 返回切片中给定下表的元素，如果不存在的话返回给定的默认值
func ChooseIndexOrDefault[T any](slice []T, index int, defaultValue T) T {
	if index < 0 || index >= len(slice) {
		return defaultValue
	}
	return slice[index]
}

// ChooseIndexes 返回给定的几个下标的元素，如果越界的话对应位置会填充为零值，会保证返回的切片的长度等于给的索引的数量
func ChooseIndexes[T any](slice []T, indexes ...int) []T {
	if len(indexes) == 0 {
		return nil
	}
	newSlice := make([]T, len(indexes), len(indexes))
	for newSliceIndex, index := range indexes {
		newSlice[newSliceIndex] = ChooseIndex(slice, index)
	}
	return newSlice
}

// ChooseOddIndexes 选择奇数下标的元素
func ChooseOddIndexes[T any](slice []T) []T {
	newSlice := make([]T, 0)
	for index, item := range slice {
		if index%2 == 1 {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

// ChooseEvenIndexes 选择偶数下标的元素
func ChooseEvenIndexes[T any](slice []T) []T {
	newSlice := make([]T, 0)
	for index, item := range slice {
		if index%2 == 0 {
			newSlice = append(newSlice, item)
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
		var zero T
		return slice[len(slice)/2], zero
	}
}

// FirstItem 选择数组中的第一个元素
func FirstItem[T any](slice []T) T {
	var zero T
	return FirstItemOrDefault(slice, zero)
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
	if n <= 0 {
		return nil
	}
	// 避免越界，如果选多了的话就选到头就算完
	if n > len(slice) {
		n = len(slice)
	}
	return slice[0:n]
}

// LastItem 返回切片的最后一个元素，如果有的话
func LastItem[T any](slice []T) T {
	var zero T
	return LastItemOrDefault(slice, zero)
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
	if n <= 0 {
		return nil
	} else if n >= len(slice) {
		return slice
	} else {
		return slice[len(slice)-n:]
	}
}

// ---------------------------------------------------------------------------------------------------------------------
