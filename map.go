package slice

// ---------------------------------------------------------------------------------------------------------------------

// Map 对数组中的每个元素应用给定行为
func Map[T, V any](slice []T, mapFunc func(item T) V) []V {
	newSlice := make([]V, 0)
	for _, item := range slice {
		newSlice = append(newSlice, mapFunc(item))
	}
	return newSlice
}

func FlatMap[T any](slice []T, flatFunc func(index int, item T) []T) []T {
	newSlice := make([]T, 0)
	for index, item := range slice {
		newSlice = append(newSlice, flatFunc(index, item)...)
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

func AllAdd() {

}

func AllSub() {

}
