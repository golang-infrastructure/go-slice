package slice


// Join 两个切片join，两个都有的才会被join返回
func Join[T any, K comparable](sliceA []T, sliceB []T, keyFunc KeyFunc[T, K]) map[K][]T {
	sliceAMap := make(map[K]T, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFunc(indexA, itemA)
		sliceAMap[keyA] = itemA
	}
	sliceADistinctSet := make(map[K]struct{}, 0)
	joinResultMap := make(map[K][]T, 0)
	for indexB, itemB := range sliceB {
		keyB := keyFunc(indexB, itemB)
		if itemA, exists := sliceAMap[keyB]; exists {
			if _, exists := sliceADistinctSet[keyB]; !exists {
				joinResultMap[keyB] = append(joinResultMap[keyB], itemA)
				sliceADistinctSet[keyB] = struct{}{}
			}
			joinResultMap[keyB] = append(joinResultMap[keyB], itemB)
		}
	}
	return joinResultMap
}

// LeftJoin 左边全来
func LeftJoin[T any, K comparable](sliceA []T, sliceB []T, keyFunc KeyFunc[T, K]) map[K][]T {
	sliceAMap := make(map[K]T, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFunc(indexA, itemA)
		sliceAMap[keyA] = itemA
	}
	sliceADistinctSet := make(map[K]struct{}, 0)
	joinResultMap := make(map[K][]T, 0)
	for indexB, itemB := range sliceB {
		keyB := keyFunc(indexB, itemB)
		if itemA, exists := sliceAMap[keyB]; exists {
			if _, exists := sliceADistinctSet[keyB]; !exists {
				joinResultMap[keyB] = append(joinResultMap[keyB], itemA)
				sliceADistinctSet[keyB] = struct{}{}
			}
			joinResultMap[keyB] = append(joinResultMap[keyB], itemB)
		}
	}

	// 把没join上的sliceA的元素也加入到结果
	for index, itemA := range sliceA {
		keyA := keyFunc(index, itemA)
		if _, exists := sliceADistinctSet[keyA]; exists {
			continue
		}
		joinResultMap[keyA] = append(joinResultMap[keyA], itemA)
		sliceADistinctSet[keyA] = struct{}{}
	}

	return joinResultMap
}

// RightJoin 右边全来
func RightJoin[T any, K comparable](sliceA []T, sliceB []T, keyFunc KeyFunc[T, K]) map[K][]T {
	return LeftJoin(sliceB, sliceA, keyFunc)
}

// FullJoin 两边全来
func FullJoin[T any, K comparable](sliceA []T, sliceB []T, keyFunc KeyFunc[T, K]) map[K][]T {
	joinResultMap := make(map[K][]T, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFunc(indexA, itemA)
		joinResultMap[keyA] = append(joinResultMap[keyA], itemA)
	}
	for indexB, itemB := range sliceB {
		keyB := keyFunc(indexB, itemB)
		joinResultMap[keyB] = append(joinResultMap[keyB], itemB)
	}
	return joinResultMap
}

func CoGroup[T any, K comparable](keyFunc KeyFunc[T, K], slices ...[]T) map[K][]T {
	resultMap := make(map[K][]T, 0)
	for _, slice := range slices {
		for index, item := range slice {
			key := keyFunc(index, item)
			resultMap[key] = append(resultMap[key], item)
		}
	}
	return resultMap
}

// FoldByKey 根据key折叠合并切片中的元素
func FoldByKey[T any, K comparable](slice []T, keyFunc KeyFunc[T, K], mergeFunc func(a T, b T) T) []T {
	keyMap := make(map[K]T, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
		if lastItem, exists := keyMap[key]; exists {
			keyMap[key] = mergeFunc(lastItem, item)
		} else {
			keyMap[key] = item
		}
	}
	newSlice := make([]T, 0)
	for _, item := range keyMap {
		newSlice = append(newSlice, item)
	}
	return newSlice
}
