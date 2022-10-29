package slice

import "math/rand"

// FillSameValue 为切片中的所有下标赋值同一个值
func FillSameValue[T any](slice []T, value T) {
	for index := range slice {
		slice[index] = value
	}
}

// FillSliceRandomValue 为切片中的所有下标赋值随机值
func FillSliceRandomValue(slice []int, low, high int) {
	for index := range slice {
		slice[index] = rand.Int()%(high-low) + low
	}
}

// NewSliceFillRandomValue 创建一个切片，并使用随机值填充
func NewSliceFillRandomValue(len int, low, high int) []int {
	slice := make([]int, len)
	FillSliceRandomValue(slice, low, high)
	return slice
}
