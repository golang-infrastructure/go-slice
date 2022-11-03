package slice

// Clone 把切片复制一份
func Clone[T any](slice []T) []T {
	if slice == nil {
		return nil
	}
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	return newSlice
}
