package slice

// ---------------------------------------------------------------------------------------------------------------------

// Clip 把slice的容量压缩到实际元素的个数
func Clip[T any](slice []T) []T {
	return slice[:len(slice):len(slice)]
}

//// Grow TODO 增加slice的容量
//func Grow() {
//
//}

//func Shift() {
//
//}
//
//func UnShift() {
//
//}

// ------------------------------------------------ ---------------------------------------------------------------------

// All 判断切片中的所有元素符合条件
func All[T any](slice []T, judgeFunc func(index int, item T) bool) bool {
	for index, item := range slice {
		if !judgeFunc(index, item) {
			return false
		}
	}
	return true
}

// Any 判断切片中的任意一个元素符合条件
func Any[T any](slice []T, judgeFunc func(index int, item T) bool) bool {
	for index, item := range slice {
		if judgeFunc(index, item) {
			return true
		}
	}
	return false
}

// ------------------------------------------------ ---------------------------------------------------------------------
