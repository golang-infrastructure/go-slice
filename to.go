package slice

// To 类似Map，不同的是更推荐这个方法只用来做做类型转换
func To[From, To any](slice []From, toFunc func(item From) To) []To {
	return Map[From, To](slice, toFunc)
}

// ToMap 把切片转为map
func ToMap[T, K comparable, V any](slice []T, toMapFunc func(index int, item T) (K, V)) map[K]V {
	resultMap := make(map[K]V, 0)
	for index, item := range slice {
		k, v := toMapFunc(index, item)
		resultMap[k] = v
	}
	return resultMap
}

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

// 多维数组转为一维数组

// 一维数组转为多维数组

// 一维数组切割为矩阵

// 矩阵转为一维数组
