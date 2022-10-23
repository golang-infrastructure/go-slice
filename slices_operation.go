package slice

// 这个文件中是多个切片之间的操作

// Union 并集，把两个切片合并，会保持两个切片中元素的相对顺序
func Union[T any](sliceA []T, sliceB []T) []T {
	newSlice := make([]T, 0, len(sliceA)+len(sliceB))
	for i := 0; i < len(sliceA); i++ {
		newSlice[i] = sliceA[i]
	}
	newSliceIndex := len(newSlice)
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

// RemoveAll 切片A减去切片B，sliceA中的元素会保持之前的相对顺序
func RemoveAll[T comparable](sliceA []T, sliceB []T) []T {
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

// Subtract 求非交集部分
func Subtract[T comparable](sliceA []T, sliceB []T) []T {
	//sliceBMap := make(map[T]struct{}, 0)
	//for _, item := range sliceB {
	//	sliceBMap[item] = struct{}{}
	//}
	//newSlice := make([]T, 0)
	//for _, item := range sliceA {
	//	if _, exists := sliceBMap[item]; exists {
	//		continue
	//	}
	//	newSlice = append(newSlice, item)
	//}
	//return newSlice
	return nil
}

// GroupByKey 根据key进行分组
func GroupByKey[T any, K comparable](slice []T, keyFunc KeyFunc[T, K]) map[K][]T {
	keySliceMap := make(map[K][]T, 0)
	for _, item := range slice {
		key := keyFunc(item)
		keySliceMap[key] = append(keySliceMap[key], item)
	}
	return keySliceMap
}

// GroupByKeyThenCount 先分组，再求每一组的count
func GroupByKeyThenCount[T any, K comparable](slice []T, keyFunc KeyFunc[T, K]) map[K]int {
	keySliceMap := GroupByKey(slice, keyFunc)
	countMap := make(map[K]int, 0)
	for key, keySlice := range keySliceMap {
		countMap[key] = len(keySlice)
	}
	return countMap
}

// ------------------------------------------------ ---------------------------------------------------------------------

type itemMeta[T any, K comparable] struct {
	key       K
	count     int
	itemSlice []T
}

// GroupByKeyThenOrderByCount 先分组，再根据每组的count排序
func GroupByKeyThenOrderByCount[T any, K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](slice []T, keyFunc KeyFunc[T, K]) {
	//keySliceMap := GroupByKey(slice, keyFunc)
	//itemMetaSlice := make([]*itemMeta[T, K], 0)
	//for key, itemSlice := range keySliceMap {
	//	itemMetaSlice = append(itemMetaSlice, &itemMeta[T, K]{
	//		key:       key,
	//		count:     len(itemSlice),
	//		itemSlice: itemSlice,
	//	})
	//}
	//SortByKey(itemMetaSlice, func(item string) int {
	//
	//})
	//return countMap
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ReduceByKey[T any, K comparable](slice []T, reduceFunc ReduceFunc[T, K]) map[K]T {
	return nil
}

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
