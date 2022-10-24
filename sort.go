package slice

import "sort"

// SortByKey 根据key对切片排序
func SortByKey[T any, K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, keyFunc KeyFunc[T, K]) {
	sort.Slice(slice, func(i, j int) bool {
		a := slice[i]
		b := slice[j]
		return keyFunc(-1, a) < keyFunc(-1, b)
	})
}

// Reverse 反转slice，可以选择把反转后的放在一个新的slice，或者在原slice原地反转
func Reverse[T any](slice []T, isNewSlice ...bool) []T {
	isNewSlice = append(isNewSlice, false)
	if isNewSlice[0] {
		return reverseInSitu(slice)
	}
	return reverseNewSlice(slice)
}

func reverseNewSlice[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	for index, item := range slice {
		newSlice[index] = item
	}
	return newSlice
}
func reverseInSitu[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return slice
}
