package slice

// Insert 往切片中插入元素，会引用原来的数组
func Insert[T any](slice []T, newValue T, insertToIndex int) []T {
	// 保证下标安全
	if insertToIndex < 0 || insertToIndex > len(slice) {
		return slice
	}
	// 先插入到最后
	slice = append(slice, newValue)
	// 然后再一点一点往前挪动
	// TODO
	return slice
}

// 插入一个切片
func InsertSlice() {

}

// InsertWithCopy 往切片中插入元素，会拷贝原切片到一个新的切片中
func InsertWithCopy() {

}
