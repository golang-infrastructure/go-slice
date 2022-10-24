package slice

import "strings"

// ---------------------------------------------------------------------------------------------------------------------

// Trim 去除切片两端的零值元素
func Trim[T comparable](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		var zero T
		return item == zero
	})
}

// TrimLeft 去除切片左边的零值元素
func TrimLeft[T comparable](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		var zero T
		return item == zero
	})
}

// TrimRight 去除右边的零值元素
func TrimRight[T comparable](slice []T) []T {
	return TrimRightByFunc[T](slice, func(item T) bool {
		var zero T
		return item == zero
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// TrimFunc 返回true的元素会被trim掉
type TrimFunc[T any] func(item T) bool

// TrimByFunc 从切片两侧trim，元素是否需要trim由自定义的trim函数决定
func TrimByFunc[T any](slice []T, trimFunc TrimFunc[T]) []T {
	t := TrimLeftByFunc(slice, trimFunc)
	return TrimRightByFunc(t, trimFunc)
}

// TrimLeftByFunc 从切片左边trim，根据自定义的函数来判定是否需要trim
func TrimLeftByFunc[T any](slice []T, trimFunc TrimFunc[T]) []T {
	if len(slice) == 0 {
		return nil
	}
	left := 0
	for left < len(slice) {
		if !trimFunc(slice[left]) {
			break
		}
		left++
	}
	if left < len(slice) {
		return slice[left:]
	} else {
		return nil
	}
}

// TrimRightByFunc 从切片右边trim，根据自定义的函数来判定是否需要trim
func TrimRightByFunc[T any](slice []T, trimFunc TrimFunc[T]) []T {
	if len(slice) == 0 {
		return nil
	}
	right := len(slice) - 1
	for right >= 0 {
		if !trimFunc(slice[right]) {
			break
		}
		right--
	}
	if right >= 0 {
		return slice[:right+1]
	} else {
		return nil
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TrimByBlankString(slice []string) []string {
	return TrimByFunc(slice, func(item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimLeftByBlankString(slice []string) []string {
	return TrimLeftByFunc(slice, func(item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimRightByBlankString(slice []string) []string {
	return TrimRightByFunc(slice, func(item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimByEmptyString(slice []string) []string {
	return TrimByFunc(slice, func(item string) bool {
		return item == ""
	})
}

func TrimLeftByEmptyString(slice []string) []string {
	return TrimLeftByFunc(slice, func(item string) bool {
		return item == ""
	})
}

func TrimRightByEmptyString(slice []string) []string {
	return TrimRightByFunc(slice, func(item string) bool {
		return item == ""
	})
}

// ---------------------------------------------------------------------------------------------------------------------

func TrimIntegerByZero[T Integer](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimIntegerLeftByZero[T Integer](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimIntegerRightByZero[T Integer](slice []T) []T {
	return TrimRightByFunc(slice, func(item T) bool {
		return item == 0
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// TrimFloatByZero 把浮点型切片的两侧的零值去除
func TrimFloatByZero[T Float](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

// TrimFloatLeftByZero 把浮点型切片的左侧的零值去除
func TrimFloatLeftByZero[T Float](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

// TrimFloatRightByZero 把浮点型切片的右侧的零值去除
func TrimFloatRightByZero[T Float](slice []T) []T {
	return TrimRightByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

// ---------------------------------------------------------------------------------------------------------------------
