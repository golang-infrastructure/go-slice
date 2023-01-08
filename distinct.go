package slice

// ---------------------------------------------------------------------------------------------------------------------

// Compact Compact是Distinct的一个别名
func Compact[T comparable](slice []T) []T {
	return Distinct(slice)
}

func CompactByUniqIDFunc[T any, K comparable](slice []T, idFunc UniqIDFunc[T, K]) []T {
	return DistinctByUniqIDFunc(slice, idFunc)
}

// ------------------------------------------------- ------------------------------------------------------------------------

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

// DistinctByUniqIDFunc 根据自定义的ID去重
func DistinctByUniqIDFunc[T any, K comparable](slice []T, uniqIDFunc UniqIDFunc[T, K]) []T {
	uniqIDSet := make(map[K]struct{}, 0)
	newSlice := make([]T, 0)
	for index, item := range slice {
		id := uniqIDFunc(index, item)
		if _, exists := uniqIDSet[id]; exists {
			continue
		}
		newSlice = append(newSlice, item)
		uniqIDSet[id] = struct{}{}
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

// IsUniq 判断切片中的元素是否不重复
func IsUniq[T comparable](slice []T) bool {
	distinctSet := make(map[T]struct{}, 0)
	for _, item := range slice {
		if _, exists := distinctSet[item]; exists {
			return false
		}
		distinctSet[item] = struct{}{}
	}
	return true
}

// IsUniqByUniqIDFunc 判断切片中的元素是否不重复，以自定义的函数为每个元素生成一个唯一标识
func IsUniqByUniqIDFunc[T any, K comparable](slice []T, idFunc UniqIDFunc[T, K]) bool {
	distinctSet := make(map[K]struct{})
	for index, item := range slice {
		uniqId := idFunc(index, item)
		if _, exists := distinctSet[uniqId]; exists {
			return false
		}
		distinctSet[uniqId] = struct{}{}
	}
	return true
}

// ------------------------------------------------- ------------------------------------------------------------------------
