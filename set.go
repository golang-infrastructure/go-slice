package slice

// Set 如果有这个下标存在的话则设置值，返回是否设置成功
func Set[T any](slice []T, index int, value T) bool {
	if index >= len(slice) {
		return false
	}
	slice[index] = value
	return true
}
