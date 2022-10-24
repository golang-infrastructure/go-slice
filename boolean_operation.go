package slice

// Union 并集，把多个切片合并，会保持多个切片中元素的相对顺序
func Union[T any](slices ...[]T) []T {
	newSlice := make([]T, 0)
	for _, slice := range slices {
		newSlice = append(newSlice, slice...)
	}
	return newSlice
}

// Intersection 求两个切片的交集
func Intersection[T comparable](sliceA []T, sliceB []T) []T {
	sliceBSet := ToSet(sliceB)
	newSlice := make([]T, 0)
	for _, itemA := range sliceA {
		if _, exists := sliceBSet[itemA]; !exists {
			continue
		}
		newSlice = append(newSlice, itemA)
	}
	return newSlice
}

// Subtract 切片A减去切片B，sliceA中的元素会保持之前的相对顺序
func Subtract[T comparable](sliceA []T, sliceB []T) []T {
	sliceBSet := ToSet(sliceB)
	newSlice := make([]T, 0)
	for _, itemA := range sliceA {
		if _, exists := sliceBSet[itemA]; exists {
			continue
		}
		newSlice = append(newSlice, itemA)
	}
	return newSlice
}

// NotIntersection 求非交集部分
func NotIntersection[T comparable](sliceA []T, sliceB []T) []T {
	sliceASet := ToSet(sliceA)
	sliceBSet := ToSet(sliceB)

	subtractSlice := make([]T, 0)

	for itemA := range sliceASet {
		if _, exists := sliceBSet[itemA]; exists {
			continue
		}
		subtractSlice = append(subtractSlice, itemA)
	}

	for itemB := range sliceBSet {
		if _, exists := sliceASet[itemB]; exists {
			continue
		}
		subtractSlice = append(subtractSlice, itemB)
	}

	return subtractSlice
}
