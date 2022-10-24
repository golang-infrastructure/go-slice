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

// DistinctByUniqIDFunc 根据自定义的ID去重
func DistinctByUniqIDFunc[T any, K comparable](slice []T, uniqIDFunc UniqIDFunc[T, K]) []T {
	uniqIDSet := make(map[K]struct{}, 0)
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
