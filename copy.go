package slice

// Copy 复制切片
func Copy[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	for index, value := range slice {
		newSlice[index] = value
	}
	return newSlice
}

// ReSlice 切片重新new，摆脱原来的数组指向
func ReSlice[T any](slice []T) []T {
	newSlice := make([]T, len(slice), len(slice))
	for index, value := range slice {
		newSlice[index] = value
	}
	return newSlice
}
