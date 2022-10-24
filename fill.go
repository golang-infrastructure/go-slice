package slice

import "math/rand"

func FillSliceValue[T any](slice []T, value T) {
	for index := range slice {
		slice[index] = value
	}
}

func FillSliceRandomValue(slice []int, low, high int) {
	for index := range slice {
		slice[index] = rand.Int()%(high-low) + low
	}
}

func NewSliceFillRandomValue(len int, low, high int) []int {
	slice := make([]int, len)
	FillSliceRandomValue(slice, low, high)
	return slice
}
