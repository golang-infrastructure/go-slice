package slice

import "sort"

// ---------------------------------------------------------------------------------------------------------------------

// GroupByKey 根据key进行分组
func GroupByKey[T any, K comparable](slice []T, keyFunc KeyFunc[T, K]) map[K][]T {
	keySliceMap := make(map[K][]T, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
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

// ---------------------------------------------------------------------------------------------------------------------

type GroupByCountContext[T any, K comparable] struct {
	Key       K
	Count     int
	ItemSlice []T
}

// GroupByKeyThenOrderByCount 先分组，再根据每组的count排序
func GroupByKeyThenOrderByCount[T any, K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, keyFunc KeyFunc[T, K]) []*GroupByCountContext[T, K] {
	keySliceMap := GroupByKey(slice, keyFunc)
	itemMetaSlice := make([]*GroupByCountContext[T, K], 0)
	for key, itemSlice := range keySliceMap {
		itemMetaSlice = append(itemMetaSlice, &GroupByCountContext[T, K]{
			Key:       key,
			Count:     len(itemSlice),
			ItemSlice: itemSlice,
		})
	}
	// 排序
	sort.Slice(itemMetaSlice, func(i, j int) bool {
		return itemMetaSlice[i].Key < itemMetaSlice[j].Key
	})
	return itemMetaSlice
}

// ---------------------------------------------------------------------------------------------------------------------
