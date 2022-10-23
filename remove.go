package slice

// Remove 移除slice中的元素
func Remove[T comparable](slice []T, removeItem T) []T {
	return Filter[T](slice, func(item T) bool {
		return item != removeItem
	})
}

func Removes() {

}
