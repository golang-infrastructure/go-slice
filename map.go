package slice

// 对数组中的
func Map[T, V any](slice []T, mapFunc func(item T) V) []V {
	newSlice := make([]V, 0)
	for _, item := range slice {
		newSlice = append(newSlice, mapFunc(item))
	}
	return newSlice
}

func AllAdd() {

}

func AllSub() {

}


