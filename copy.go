package slice

// Copy 复制切片
func Copy[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// ReSlice 切片重新new，摆脱原来的数组指向
// 感觉跟上面干的事一件事?
func ReSlice[T any](slice []T) []T {
	newSlice := make([]T, len(slice), len(slice))
	for index, item := range slice {
		newSlice[index] = item
	}
	return newSlice
}
