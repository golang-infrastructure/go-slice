package slice

// ---------------------------------------------------------------------------------------------------------------------

// Distinct 对切片去重
func Distinct[T comparable](slice []T) []T {
	itemSet := make(map[T]struct{})
	newSlice := make([]T, 0)
	for _, item := range slice {
		if _, exists := itemSet[item]; exists {
			continue
		}
		newSlice = append(newSlice, item)
		itemSet[item] = struct{}{}
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

// UniqIDFunc 为元素生成一个唯一ID
type UniqIDFunc[T any] func(item T) string

// DistinctByUniqIDFunc 根据自定义的ID去重
func DistinctByUniqIDFunc[T any](slice []T, uniqIDFunc UniqIDFunc[T]) []T {
	uniqIDSet := make(map[string]struct{}, 0)
	newSlice := make([]T, 0)
	for _, item := range slice {
		id := uniqIDFunc(item)
		if _, exists := uniqIDSet[id]; exists {
			continue
		}
		newSlice = append(newSlice, item)
		uniqIDSet[id] = struct{}{}
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------
