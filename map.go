package slice

import (
	"errors"
)

// ---------------------------------------------------------------------------------------------------------------------

// Map 对数组中的每个元素应用给定行为
func Map[T, V any](slice []T, mapFunc func(item T) V) []V {
	newSlice := make([]V, 0)
	for _, item := range slice {
		newSlice = append(newSlice, mapFunc(item))
	}
	return newSlice
}

func FlatMap[T any](slice []T, flatFunc func(index int, item T) []T) []T {
	newSlice := make([]T, 0)
	for index, item := range slice {
		newSlice = append(newSlice, flatFunc(index, item)...)
	}
	return newSlice
}

// ---------------------------------------------------------------------------------------------------------------------

func AllAdd[T string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item + n
	}
}

func AllSub[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item - n
	}
}

func AllMulti[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, n T) {
	for index, item := range slice {
		slice[index] = item * n
	}
}

func AllDivision[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](slice []T, n T) error {
	if n == 0 {
		return errors.New("can not division zero")
	}
	for index, item := range slice {
		slice[index] = item / n
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
