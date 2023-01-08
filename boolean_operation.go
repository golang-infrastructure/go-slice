package slice

// ------------------------------------------------- ------------------------------------------------------------------------

// Union 并集，把多个切片合并，会保持多个切片中元素的相对顺序
func Union[T any](slices ...[]T) []T {
	newSlice := make([]T, 0)
	for _, slice := range slices {
		newSlice = append(newSlice, slice...)
	}
	return newSlice
}

// ------------------------------------------------- ------------------------------------------------------------------------

// Intersection 求两个切片的交集
func Intersection[T comparable](sliceA, sliceB []T) []T {
	// 任何切片与空切片没有交集
	if len(sliceA) == 0 || len(sliceB) == 0 {
		return nil
	}
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

// IntersectionByKeyFunc 求两个切片的交集，计算交集的时候每个元素返回一个key用来做关联
func IntersectionByKeyFunc[T any, K comparable](sliceA, sliceB []T, keyFunc func(index int, item T) K) []T {
	// 任何切片与空切片没有交集
	if len(sliceA) == 0 || len(sliceB) == 0 {
		return nil
	}
	sliceBSet := ToSetByFunc(sliceB, keyFunc)
	newSlice := make([]T, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFunc(indexA, itemA)
		if _, exists := sliceBSet[keyA]; !exists {
			continue
		}
		newSlice = append(newSlice, itemA)
	}
	return newSlice
}

// ------------------------------------------------- ------------------------------------------------------------------------

// NotIntersection 求非交集部分
func NotIntersection[T comparable](sliceA, sliceB []T) []T {
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

// NotIntersectionByKeyFunc 求两个切片不相交的部分
func NotIntersectionByKeyFunc[T any, K comparable](sliceA, sliceB []T, keyFunc func(index int, item T) K) []T {
	toMapFunc := func(index int, item T) (K, T) {
		return keyFunc(index, item), item
	}
	sliceAMap := ToMapByFunc(sliceA, toMapFunc)
	sliceBMap := ToMapByFunc(sliceB, toMapFunc)

	subtractSlice := make([]T, 0)

	for itemKeyA, itemValueA := range sliceAMap {
		if _, exists := sliceBMap[itemKeyA]; exists {
			continue
		}
		subtractSlice = append(subtractSlice, itemValueA)
	}

	for itemKeyB, itemValueB := range sliceBMap {
		if _, exists := sliceAMap[itemKeyB]; exists {
			continue
		}
		subtractSlice = append(subtractSlice, itemValueB)
	}

	return subtractSlice
}

// ------------------------------------------------- ------------------------------------------------------------------------

// Subtract 切片A减去切片B，sliceA中的元素会保持之前的相对顺序
func Subtract[T comparable](sliceA, sliceB []T) []T {
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

// SubtractByKeyFunc 切片A减去切片B，sliceA中的元素会保持之前的顺序，是否匹配是靠自定义的函数的返回的key决定的
func SubtractByKeyFunc[T any, K comparable](sliceA, sliceB []T, keyFunc func(index int, item T) K) []T {
	sliceBSet := ToSetByFunc(sliceB, keyFunc)
	newSlice := make([]T, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFunc(indexA, itemA)
		if _, exists := sliceBSet[keyA]; exists {
			continue
		}
		newSlice = append(newSlice, itemA)
	}
	return newSlice
}

// ------------------------------------------------- ------------------------------------------------------------------------
