package slice

import (
	"errors"
)

// ---------------------------------------------------------------------------------------------------------------------

// Map 对数组中的每个元素应用给定行为，并返回
func Map[T, V any](slice []T, mapFunc func(index int, item T) V) []V {
	newSlice := make([]V, 0)
	for index, item := range slice {
		newSlice = append(newSlice, mapFunc(index, item))
	}
	return newSlice
}

func FlatMap[T, V any](slice []T, flatFunc func(index int, item T) []V) []V {
	newSlice := make([]V, 0)
	for index, item := range slice {
		newSlice = append(newSlice, flatFunc(index, item)...)
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

// AllAdd 切片中所有元素加上一个值
func AllAdd[T Ordered](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item + n
	}
}

// AllSub 切片中所有元素减去一个值
func AllSub[T Integer | Float](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item - n
	}
}

// AllMulti 切片中所有元素乘以一个值
func AllMulti[T Integer | Float](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item * n
	}
}

// AllDivision 切片中所有元素除以一个值
func AllDivision[T Integer | Float](slice []T, n T) error {
	if n == 0 {
		return errors.New("can not division zero")
	}
	for index, item := range slice {
		slice[index] = item / n
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
