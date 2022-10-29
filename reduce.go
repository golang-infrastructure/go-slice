package slice

import (
	"fmt"
	compare_anything "github.com/CC11001100/go-compare-anything"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------

type ReduceFunc[T, R any] func(index int, item T, result R) R

// Reduce 把切片中的所有元素合并为一个
func Reduce[T, R any](slice []T, initResult R, reduceFunc ReduceFunc[T, R]) R {
	result := initResult
	for index, item := range slice {
		result = reduceFunc(index, item, result)
	}
	return result
}

// ReduceByKey 根据key把切片中的元素分组并合并返回
func ReduceByKey[T any, K comparable, R any](slice []T, keyFunc KeyFunc[T, K], reduceFunc ReduceFunc[T, R]) map[K]R {
	keyMap := make(map[K]R, 0)
	for index, item := range slice {
		key := keyFunc(index, item)
		result := reduceFunc(index, item, keyMap[key])
		keyMap[key] = result
	}
	return keyMap
}

// ---------------------------------------------------------------------------------------------------------------------

// Max 求切片中的最大值
func Max[T Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

// MaxByComparator 求出切片中的最大值，使用给定的比较器来比较元素之间的大小
func MaxByComparator[T any](slice []T, comparator compare_anything.Comparator[T]) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return Reduce(slice[1:], slice[0], func(index int, item T, result T) T {
		if comparator(item, result) > 0 {
			return item
		} else {
			return result
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// Min 求切片中的最小值
func Min[T Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

// MinByComparator 求出切片中的最小值，使用给定的比较器来比较元素之间的大小
func MinByComparator[T any](slice []T, comparator compare_anything.Comparator[T]) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return Reduce(slice[1:], slice[0], func(index int, item T, result T) T {
		if comparator(item, result) < 0 {
			return item
		} else {
			return result
		}
	})
}

// ------------------------------------------------ ---------------------------------------------------------------------

// TODO 2022-10-24 03:03:24 有点难搞下面这几个，可能要引入新的概念

//// Avg 注意，有溢出风险
//func Avg[T Float](slice []T) T {
//	if len(slice) == 0 {
//		return 0
//	}
//	var sum T
//	for _, item := range slice {
//		sum += item
//	}
//	return sum * 1.0 / len(slice)
//}
//
//// AvgByValueFunc 注意，有溢出风险
//func AvgByValueFunc[T any, V Integer | Float](slice []T, valueFunc func(index int, item T) V) V {
//	var sum V
//	for index, item := range slice {
//		sum += valueFunc(index, item)
//	}
//	return sum
//}

// ------------------------------------------------ ---------------------------------------------------------------------

// Sum 注意，有溢出风险
func Sum[T Integer | Float](slice []T) T {
	var sum T
	for _, item := range slice {
		sum += item
	}
	return sum
}

// SumByValueFunc 注意，有溢出风险
func SumByValueFunc[T any, V Integer | Float](slice []T, valueFunc func(index int, item T) V) V {
	var sum V
	for index, item := range slice {
		sum += valueFunc(index, item)
	}
	return sum
}

// ------------------------------------------------ ---------------------------------------------------------------------

func Join[T any](slice []T, delimiter string) string {
	sb := strings.Builder{}
	for index, item := range slice {
		sb.WriteString(fmt.Sprintf("%v", item))
		if index != len(slice)-1 {
			sb.WriteString(delimiter)
		}
	}
	return sb.String()
}

// ------------------------------------------------ ---------------------------------------------------------------------
