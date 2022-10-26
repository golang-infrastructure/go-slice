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

// ToStringSliceByFunc 使用自定义的函数把切片转为string类型的切片
func ToStringSliceByFunc[T any](slice []T, toStringFunc func(index int, item T) string) []string {
	stringSlice := make([]string, 0)
	for index, item := range slice {
		s := toStringFunc(index, item)
		stringSlice = append(stringSlice, s)
	}
	return stringSlice
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

// ToMap 把切片转为map，将索引作为map的key，将对应下标元素作为map的value
func ToMap[T, K int](slice []T) map[int]T {
	resultMap := make(map[int]T, 0)
	for index, item := range slice {
		resultMap[index] = item
	}
	return resultMap
}

// ToMapByFunc 把切片转为map，使用自定义的函数从元素中生成key和value
func ToMapByFunc[T any, K comparable, V any](slice []T, toMapFunc func(index int, item T) (K, V)) map[K]V {
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

// ToSetByFunc 把切片转为set，根据自定义的函数从切片的元素生成set的key
func ToSetByFunc[T any, K comparable](slice []T, keyFunc KeyFunc[T, K]) map[K]struct{} {
	set := make(map[K]struct{}, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
		set[key] = struct{}{}
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

// ------------------------------------------------ ---------------------------------------------------------------------
