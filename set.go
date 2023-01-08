package slice

// Set 如果有这个下标存在的话则设置值，返回是否设置成功
func Set[T any](slice []T, index int, value T) bool {
	// 空的切片不允许替换值
	if slice == nil {
		return false
	}
	// 下标越界不允许
	if index >= len(slice) {
		return false
	}
	slice[index] = value
	return true
}

// SetFirst 设置切片的第一个下标的值为给定的值，如果数组长度小于1，则设置不成功
func SetFirst[T any](slice []T, value T) bool {
	if len(slice) == 0 {
		return false
	}
	slice[0] = value
	return true
}

// SetLast 设置切片的最后一个下标的值，如果切片为空则设置不成功
func SetLast[T any](slice []T, value T) bool {
	if len(slice) == 0 {
		return false
	}
	slice[len(slice)-1] = value
	return true
}

// Copy 在切片之间互相拷贝值
func Copy[T any](sourceSlice []T, destinationSlice []T, sourceStartIndex, sourceEndIndex, destinationStartIndex, destinationEndIndex int) int {
	copySuccessCount := 0
	for sourceStartIndex < sourceEndIndex && destinationStartIndex < destinationEndIndex {
		sourceSlice[sourceStartIndex] = destinationSlice[destinationStartIndex]
		sourceStartIndex++
		destinationStartIndex++
		copySuccessCount++
	}
	return copySuccessCount
}
