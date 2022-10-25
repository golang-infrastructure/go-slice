package slice

import "github.com/CC11001100/go-tuple"

// ---------------------------------------------------------------------------------------------------------------------

// JoinByKeyFunc 两个切片join，两个都有的才会被join返回，注意如果数据有重复会笛卡尔积，与数据库的join行为保持一致
func JoinByKeyFunc[A, B any, K comparable](sliceA []A, keyFuncA KeyFunc[A, K], sliceB []B, keyFuncB KeyFunc[B, K]) map[K][]*tuple.Tuple2[A, B] {
	// 先把A做成一个关联点，因为可能会存在多个相同key的元素，要都能够保留着
	sliceAMap := make(map[K][]A, 0)
	for indexA, itemA := range sliceA {
		keyA := keyFuncA(indexA, itemA)
		sliceAMap[keyA] = append(sliceAMap[keyA], itemA)
	}
	joinResultMap := make(map[K][]*tuple.Tuple2[A, B], 0)
	for indexB, itemB := range sliceB {
		keyB := keyFuncB(indexB, itemB)
		if itemASlice, exists := sliceAMap[keyB]; exists {
			for _, itemA := range itemASlice {
				joinResultMap[keyB] = append(joinResultMap[keyB], tuple.New2(itemA, itemB))
			}
		}
	}
	return joinResultMap
}

// LeftJoinByKeyFunc 左边全来
func LeftJoinByKeyFunc[A, B any, K comparable](sliceA []A, keyFuncA KeyFunc[A, K], sliceB []B, keyFuncB KeyFunc[B, K]) map[K][]*tuple.Tuple2[A, B] {

	// 除了保留切片A的元素，还要把下表也保留着，等下要把被join成功的下标都收集得到
	sliceAMap := make(map[K][]*tuple.Tuple2[A, int], 0)
	for indexA, itemA := range sliceA {
		keyA := keyFuncA(indexA, itemA)
		sliceAMap[keyA] = append(sliceAMap[keyA], tuple.New2(itemA, indexA))
	}
	sliceAJoinSuccessIndexSet := make(map[int]struct{}, 0)
	joinResultMap := make(map[K][]*tuple.Tuple2[A, B], 0)
	for indexB, itemB := range sliceB {
		keyB := keyFuncB(indexB, itemB)
		if itemASlice, exists := sliceAMap[keyB]; exists {
			for _, itemA := range itemASlice {
				joinResultMap[keyB] = append(joinResultMap[keyB], tuple.New2(itemA.V1, itemB))
				sliceAJoinSuccessIndexSet[itemA.V2] = struct{}{}
			}
		}
	}

	// 把没join上的sliceA的元素也加入到结果
	for index, itemA := range sliceA {
		if _, exists := sliceAJoinSuccessIndexSet[index]; exists {
			continue
		}
		keyA := keyFuncA(index, itemA)
		var zero B
		joinResultMap[keyA] = append(joinResultMap[keyA], tuple.New2(itemA, zero))
	}

	return joinResultMap
}

// RightJoinByKeyFunc 右边全来
func RightJoinByKeyFunc[A, B any, K comparable](sliceA []A, keyFuncA KeyFunc[A, K], sliceB []B, keyFuncB KeyFunc[B, K]) map[K][]*tuple.Tuple2[A, B] {
	// 除了保留切片A的元素，还要把下表也保留着，等下要把被join成功的下标都收集得到
	sliceBMap := make(map[K][]*tuple.Tuple2[B, int], 0)
	for indexB, itemB := range sliceB {
		keyB := keyFuncB(indexB, itemB)
		sliceBMap[keyB] = append(sliceBMap[keyB], &tuple.Tuple2[B, int]{itemB, indexB})
	}
	sliceBJoinSuccessIndexSet := make(map[int]struct{}, 0)
	joinResultMap := make(map[K][]*tuple.Tuple2[A, B], 0)
	for indexA, itemA := range sliceA {
		keyA := keyFuncA(indexA, itemA)
		if itemBSlice, exists := sliceBMap[keyA]; exists {
			for _, itemB := range itemBSlice {
				joinResultMap[keyA] = append(joinResultMap[keyA], &tuple.Tuple2[A, B]{itemA, itemB.V1})
				sliceBJoinSuccessIndexSet[itemB.V2] = struct{}{}
			}
		}
	}

	// 把没join上的sliceB的元素也加入到结果
	for index, itemB := range sliceB {
		if _, exists := sliceBJoinSuccessIndexSet[index]; exists {
			continue
		}
		keyB := keyFuncB(index, itemB)
		var zero A
		joinResultMap[keyB] = append(joinResultMap[keyB], &tuple.Tuple2[A, B]{zero, itemB})
	}

	return joinResultMap
}

// FullJoinByKeyFunc 两边全来
func FullJoinByKeyFunc[A, B any, K comparable](sliceA []A, keyFuncA KeyFunc[A, K], sliceB []B, keyFuncB KeyFunc[B, K]) map[K][]*tuple.Tuple2[A, B] {

	// 除了保留切片A的元素，还要把下表也保留着，等下要把被join成功的下标都收集得到
	sliceAMap := make(map[K][]*tuple.Tuple2[A, int], 0)
	for indexA, itemA := range sliceA {
		keyA := keyFuncA(indexA, itemA)
		sliceAMap[keyA] = append(sliceAMap[keyA], &tuple.Tuple2[A, int]{itemA, indexA})
	}
	sliceAJoinSuccessIndexSet := make(map[int]struct{}, 0)
	sliceBJoinSuccessIndexSet := make(map[int]struct{}, 0)
	joinResultMap := make(map[K][]*tuple.Tuple2[A, B], 0)
	for indexB, itemB := range sliceB {
		keyB := keyFuncB(indexB, itemB)
		if itemASlice, exists := sliceAMap[keyB]; exists {
			for _, itemA := range itemASlice {
				joinResultMap[keyB] = append(joinResultMap[keyB], &tuple.Tuple2[A, B]{itemA.V1, itemB})
				sliceAJoinSuccessIndexSet[itemA.V2] = struct{}{}
				sliceBJoinSuccessIndexSet[indexB] = struct{}{}
			}
		}
	}

	// 把没join上的sliceA的元素也加入到结果
	for index, itemA := range sliceA {
		if _, exists := sliceAJoinSuccessIndexSet[index]; exists {
			continue
		}
		keyA := keyFuncA(index, itemA)
		var zero B
		joinResultMap[keyA] = append(joinResultMap[keyA], &tuple.Tuple2[A, B]{itemA, zero})
	}

	// 把没join上的sliceB的元素也加入到结果
	for index, itemB := range sliceB {
		if _, exists := sliceBJoinSuccessIndexSet[index]; exists {
			continue
		}
		keyB := keyFuncB(index, itemB)
		var zero A
		joinResultMap[keyB] = append(joinResultMap[keyB], &tuple.Tuple2[A, B]{zero, itemB})
	}

	return joinResultMap
}

// ---------------------------------------------------------------------------------------------------------------------

// CoGroup 将多个切片中的元素以key做关联，相同的key会被放到同一个key下的切片中
func CoGroup[A, B any, K comparable](sliceA []A, keyFuncA KeyFunc[A, K], sliceB []B, keyFuncB KeyFunc[B, K]) map[K]*tuple.Tuple2[[]A, []B] {
	resultMap := make(map[K]*tuple.Tuple2[[]A, []B], 0)

	// 切片A
	for indexA, itemA := range sliceA {
		keyA := keyFuncA(indexA, itemA)
		if _, exists := resultMap[keyA]; !exists {
			resultA := make([]A, 0)
			resultB := make([]B, 0)
			resultMap[keyA] = tuple.New2[[]A, []B](resultA, resultB)
		}
		resultMap[keyA].V1 = append(resultMap[keyA].V1, itemA)
	}

	// 切片B
	for indexB, itemB := range sliceB {
		keyA := keyFuncB(indexB, itemB)
		if _, exists := resultMap[keyA]; !exists {
			resultA := make([]A, 0)
			resultB := make([]B, 0)
			resultMap[keyA] = tuple.New2[[]A, []B](resultA, resultB)
		}
		resultMap[keyA].V2 = append(resultMap[keyA].V2, itemB)
	}

	return resultMap
}

// ---------------------------------------------------------------------------------------------------------------------

// FoldByKeyFunc 根据key折叠合并切片中的元素，注意折叠完元素在数组中的相对位置可能会被改变，不是稳定的折叠
func FoldByKeyFunc[T any, K comparable](slice []T, keyFunc KeyFunc[T, K], mergeFunc func(source T, destination T) T) []T {
	keyMap := make(map[K]T, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
		if lastItem, exists := keyMap[key]; exists {
			keyMap[key] = mergeFunc(item, lastItem)
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

// ---------------------------------------------------------------------------------------------------------------------
