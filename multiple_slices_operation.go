package slice

// 这个文件中是多个切片之间的操作


// aggregateByKey

// Cartesian 笛卡尔积，把两个切片相乘
func Cartesian[A, B, C any](sliceA []A, sliceB []B, multiplicationFunc func(a A, b B) C) []C {
	slice := make([]C, 0)
	for _, itemA := range sliceA {
		for _, itemB := range sliceB {
			slice = append(slice, multiplicationFunc(itemA, itemB))
		}
	}
	return slice
}

// pipe

// Foreach
func Foreach[T any](slice []T, eachFunc func(index int, item T) bool) {
	for index, item := range slice {
		if !eachFunc(index, item) {
			break
		}
	}
}

// aggregate

// Zip 把两个切片压成k/v的形式
func Zip[K comparable, V any](keySlice []K, valueSlice []V) map[K]V {
	if len(keySlice) != len(valueSlice) {
		return nil
	}
	resultMap := make(map[K]V, 0)
	for index, keyItem := range keySlice {
		resultMap[keyItem] = valueSlice[index]
	}
	return resultMap
}
