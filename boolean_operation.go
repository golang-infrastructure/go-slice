package slice

// Union 并集，把两个切片合并，会保持两个切片中元素的相对顺序
func Union[T any](sliceA []T, sliceB []T) []T {
	newSlice := make([]T, len(sliceA)+len(sliceB), len(sliceA)+len(sliceB))
	for i := 0; i < len(sliceA); i++ {
		newSlice[i] = sliceA[i]
	}
	newSliceIndex := len(sliceA)
	for i := 0; i < len(sliceB); i++ {
		newSlice[newSliceIndex] = sliceB[i]
		newSliceIndex++
	}
	return newSlice
}

// Intersection 求两个切片的交集
func Intersection[T comparable](sliceA []T, sliceB []T) []T {
	sliceBMap := make(map[T]struct{}, 0)
	for _, item := range sliceB {
		sliceBMap[item] = struct{}{}
	}
	newSlice := make([]T, 0)
	for _, item := range sliceA {
		if _, exists := sliceBMap[item]; !exists {
			continue
		}
		newSlice = append(newSlice, item)
	}
	return newSlice
}

// Subtract 切片A减去切片B，sliceA中的元素会保持之前的相对顺序
func Subtract[T comparable](sliceA []T, sliceB []T) []T {
	sliceBMap := make(map[T]struct{}, 0)
	for _, item := range sliceB {
		sliceBMap[item] = struct{}{}
	}
	newSlice := make([]T, 0)
	for _, item := range sliceA {
		if _, exists := sliceBMap[item]; exists {
			continue
		}
		newSlice = append(newSlice, item)
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
