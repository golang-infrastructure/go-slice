package slice

import "math/rand"

// Shuffle 对切片中的元素洗牌
func Shuffle[T any](slice []T) []T {
	//indexSet := make(map[int]struct{}, 0)
	//for index := range slice {
	//	indexSet[index] = struct{}{}
	//}
	//left := 0
	//right :=
	//for {
	//	// TODO 2022-10-27 01:23:29
	//}
	return nil
}

// Random 从切片中随机选择一个元素
func Random[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[rand.Int()%len(slice)]
}

// Randoms 从切片中随机选择n个元素，会避免重复选择
func Randoms[T any](slice []T, n int) []T {
	if n == 0 || len(slice) == 0 {
		return nil
	}
	// 抖个机灵，借助于map访问时的随机性来实现随机选择
	indexSet := make(map[int]struct{}, 0)
	for index := range slice {
		indexSet[index] = struct{}{}
	}
	newSlice := make([]T, 0)
	for index := range indexSet {
		newSlice = append(newSlice, slice[index])
		if len(newSlice) == n {
			break
		}
	}
	return newSlice
}
