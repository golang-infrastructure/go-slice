package slice

// To 类似Map，不同的是更推荐这个方法只用来做做类型转换
func To[From, To any](slice []From, toFunc func(item From) To) []To {
	return Map[From, To](slice, toFunc)
}

func ToMap[From, To any](slice []From) {}

func ToSet() {}

// 多维数组转为一维数组

// 一维数组转为多维数组

// 一维数组切割为矩阵

// 矩阵转为一维数组
