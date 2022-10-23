package slice

func Join[T any, K comparable](sliceA []T, sliceB []T, keyFunc KeyFunc[T, K]) map[K][]T {
	return nil
}

// countByKey

type KeyFunc[T any, K comparable] func(item T) K

func CoGroup[T any, K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](keyFunc KeyFunc[T, K], slices ...[]T) map[K][]T {
	resultMap := make(map[K][]T, 0)
	for _, slice := range slices {
		for _, item := range slice {
			key := keyFunc(item)
			resultMap[key] = append(resultMap[key], item)
		}
	}
	return resultMap
}

// FoldByKey 根据key折叠合并元素
func FoldByKey[T any, K comparable](slice []T, keyFunc KeyFunc[T, K], mergeFunc func(a T, b T) T) []T {
	return nil
}
