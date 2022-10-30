package slice

// Remove 移除slice中的元素，其实就是把原切片复制了一遍，复杂度是O(n)
func Remove[T comparable](slice []T, removeItem T) []T {
	return Filter[T](slice, func(index int, item T) bool {
		return item != removeItem
	})
}

// RemoveByIndex 根据下标删除元素，其实就是把元素复制了一遍，复杂度是O(n)
func RemoveByIndex[T any](slice []T, deleteIndex int) []T {
	return Filter[T](slice, func(index int, item T) bool {
		return index != deleteIndex
	})
}

// RemoveFirst
// RemoveFirstN
// RemoveLastN

//func Removes[T any](slice []T, removeItems ...T) []T {
//
//}
