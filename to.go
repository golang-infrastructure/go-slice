package slice

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------

// To 类似Map，不同的是更推荐这个方法只用来做做类型转换
func To[From, To any](slice []From, toFunc func(index int, item From) To) []To {
	return Map[From, To](slice, toFunc)
}

// ---------------------------------------------------------------------------------------------------------------------

// ToStringSlice 把切片转为string类型的切片
func ToStringSlice[T any](slice []T) []string {
	return To(slice, func(index int, item T) string {
		return fmt.Sprintf("%#v", item)
	})
}

// ToString 把整个切片转为一个字符串，通常用于打印、记录之类的
func ToString[T any](slice []T) string {
	sb := strings.Builder{}
	for _, item := range slice {
		sb.WriteString(fmt.Sprintf("%#v", item))
	}
	return sb.String()
}

// ---------------------------------------------------------------------------------------------------------------------

// ToMap 把切片转为map
func ToMap[T, K comparable, V any](slice []T, toMapFunc func(index int, item T) (K, V)) map[K]V {
	resultMap := make(map[K]V, 0)
	for index, item := range slice {
		k, v := toMapFunc(index, item)
		resultMap[k] = v
	}
	return resultMap
}

// ---------------------------------------------------------------------------------------------------------------------

// ToSet 把切片转为set
func ToSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{}, 0)
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}

// ToSetWithFunc 把切片转为set
func ToSetWithFunc[T any, S comparable](slice []T, toSetFunc func(index int, item T) S) map[S]struct{} {
	set := make(map[S]struct{}, 0)
	for index, item := range slice {
		set[toSetFunc(index, item)] = struct{}{}
	}
	return set
}

// ---------------------------------------------------------------------------------------------------------------------

// Flat2 二维切片打平为一维切片
func Flat2[T any](slice [][]T) []T {
	resultSlice := make([]T, 0)
	for _, itemSlice := range slice {
		for _, item := range itemSlice {
			resultSlice = append(resultSlice, item)
		}
	}
	return resultSlice
}

// ---------------------------------------------------------------------------------------------------------------------

// ToMatrix 一维切片转换为矩阵
func ToMatrix[T any](slice []T, columnCount int) [][]T {
	matrix := make([][]T, 0)
	consumer := NewSliceConsumer(slice)
	for !consumer.IsDone() {
		row := make([]T, columnCount, columnCount)
		matrix = append(matrix, row)
		for index := range row {
			e, err := consumer.TakeE()
			if err != nil {
				return matrix
			}
			row[index] = e
		}
	}
	return matrix
}

// FromMatrix 矩阵转为一维切片
func FromMatrix[T any](matrix [][]T) []T {
	return Flat2(matrix)
}

// ------------------------------------------------ ---------------------------------------------------------------------
