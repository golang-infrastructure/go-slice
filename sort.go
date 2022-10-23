package slice

import "sort"

// SortByKey 根据key对切片排序
func SortByKey[T any, K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](slice []T, keyFunc KeyFunc[T, K]) {
	sort.Slice(slice, func(i, j int) bool {
		a := slice[i]
		b := slice[j]
		return keyFunc(a) < keyFunc(b)
	})
}
