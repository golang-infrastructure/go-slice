package slice

// Distinct 对切片去重
func Distinct[T comparable](slice []T) []T {
	valueSet := make(map[T]struct{})
	newSlice := make([]T, 0)
	for _, value := range slice {
		if _, exists := valueSet[value]; exists {
			continue
		}
		newSlice = append(newSlice, value)
		valueSet[value] = struct{}{}
	}
	return newSlice
}

// UniqIDFunc 为元素生成一个唯一ID
type UniqIDFunc[T any] func(value T) string

// DistinctByUniqIDFunc 根据自定义的ID去重
func DistinctByUniqIDFunc[T any](slice []T, uniqIDFunc UniqIDFunc[T]) []T {
	uniqIDSet := make(map[string]struct{}, 0)
	newSlice := make([]T, 0)
	for _, value := range slice {
		id := uniqIDFunc(value)
		if _, exists := uniqIDSet[id]; exists {
			continue
		}
		newSlice = append(newSlice, value)
		uniqIDSet[id] = struct{}{}
	}
	return newSlice
}
