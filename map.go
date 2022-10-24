package slice

import (
	"errors"
)

// ---------------------------------------------------------------------------------------------------------------------

// Map 对数组中的每个元素应用给定行为
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

func AllAdd[T Ordered](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item + n
	}
}

func AllSub[T Integer | Float](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item - n
	}
}

func AllMulti[T Integer | Float](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item * n
	}
}

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
