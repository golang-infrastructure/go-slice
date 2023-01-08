package slice

// ------------------------------------------------ ---------------------------------------------------------------------

// Remove 移除slice中的元素，其实就是把原切片复制了一遍，复杂度是O(n)
func Remove[T comparable](slice []T, removeItem T) []T {
	return Filter[T](slice, func(index int, item T) bool {
		return item != removeItem
	})
}

func Removes[T comparable](slice []T, removeItems ...T) []T {
	return Filter[T](slice, func(index int, item T) bool {
		return Contains(removeItems, item)
	})
}

// ------------------------------------------------ ---------------------------------------------------------------------

// RemoveByIndex 根据下标删除元素，其实就是把元素复制了一遍，复杂度是O(n)
func RemoveByIndex[T any](slice []T, deleteIndex int) []T {
	return Filter[T](slice, func(index int, item T) bool {
		return index != deleteIndex
	})
}

func RemoveByIndexes[T any](slice []T, deleteIndexes ...int) []T {
	if len(deleteIndexes) == 0 {
		return slice
	}
	set := ToSet(deleteIndexes)
	return Filter[T](slice, func(index int, item T) bool {
		_, exists := set[index]
		return !exists
	})
}

// ------------------------------------------------ ---------------------------------------------------------------------

func RemoveFirst[T any](slice []T) []T {
	if len(slice) <= 1 {
		return nil
	}
	return slice[1:]
}

func RemoveFirstN[T any](slice []T, n int) []T {
	if len(slice) <= n {
		return nil
	}
	return slice[n:]
}

func RemoveLast[T any](slice []T) []T {
	if len(slice) <= 1 {
		return nil
	}
	return slice[0 : len(slice)-1]
}

func RemoveLastN[T any](slice []T, n int) []T {
	if len(slice) <= n {
		return nil
	}
	return slice[0 : len(slice)-n]
}

// ------------------------------------------------ ---------------------------------------------------------------------
