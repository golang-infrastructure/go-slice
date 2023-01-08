package slice

// 注意，插入的方法大部分都会导致整个切片复制一份，请慎用，确保不会用在超大切片上

// Insert 往切片中插入元素，会引用原来的数组中的元素的地址
func Insert[T any](slice []T, value T, insertToIndex int) []T {
	// 保证下标安全
	if insertToIndex < 0 || insertToIndex > len(slice) {
		return slice
	}
	if len(slice) == 0 {
		return []T{value}
	}
	// 先插入到最后
	slice = append(slice, value)
	// 然后往前蠕动
	index := len(slice) - 1
	for index > insertToIndex {
		slice[index] = slice[index-1]
		index--
	}
	return slice
}

// ------------------------------------------------ ---------------------------------------------------------------------

// InsertBeforeFirst 把元素插入到给定切片的第一个位置
func InsertBeforeFirst[T any](slice []T, value T) []T {
	newSlice := make([]T, len(slice)+1)
	newSlice[0] = value
	for index, value := range slice {
		newSlice[index+1] = value
	}
	return newSlice
}

// InsertAfterFirst 把给定元素插入到切片当前的的第一个元素之后
func InsertAfterFirst[T any](slice []T, value T) []T {
	if len(slice) == 0 {
		return []T{value}
	}
	newSlice := make([]T, len(slice)+1)
	newSlice[0] = slice[0]
	newSlice[1] = value
	for index := 1; index < len(slice); index++ {
		newSlice[index+1] = value
	}
	return newSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------

// InsertBeforeLast 把给定元素插入到切片当前的最后一个元素之前
func InsertBeforeLast[T any](slice []T, value T) []T {
	if len(slice) == 0 {
		return []T{value}
	}
	// 把原来的last替换掉，然后再插入
	last := slice[len(slice)-1]
	slice[len(slice)-1] = value
	return append(slice, last)
}

// InsertAfterLast 把给定元素插入到切片当前的最后一个元素之后，其实等同于append
func InsertAfterLast[T any](slice []T, value T) []T {
	return append(slice, value)
}

// ------------------------------------------------ ---------------------------------------------------------------------

// InsertSlice 把sourceSlice切片插入到destinationSlice的指定位置，从那个位置开始的都往后移动
func InsertSlice[T any](destinationSlice []T, sourceSlice []T, destinationStartIndex int) {
	// TODO
}

// ------------------------------------------------ ---------------------------------------------------------------------
