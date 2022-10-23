package slice

// Reverse 反转slice，可以选择把反转后的放在一个新的slice，或者在原slice原地反转
func Reverse[T any](slice []T, isNewSlice ...bool) []T {
	newSlice := make([]T, len(slice))
	for index, item := range slice {
		newSlice[index] = item
	}
	return newSlice
}
