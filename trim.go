package slice

import "strings"

// ---------------------------------------------------------------------------------------------------------------------

// Trim 去除切片两端的nil元素
func Trim[T any](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item == nil
	})
}

// TrimLeft 去除切片左边的nil元素
func TrimLeft[T any](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item == nil
	})
}

// TrimRight 去除右边的nil元素
func TrimRight[T any](slice []T) []T {
	return TrimRightByFunc[T](slice, func(item T) bool {
		return item == nil
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
		return slice[right:]
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

func TrimIntByZero[T int | int8 | int16 | int32 | int64](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimIntLeftByZero[T int | int8 | int16 | int32 | int64](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimIntRightByZero[T int | int8 | int16 | int32 | int64](slice []T) []T {
	return TrimRightByFunc(slice, func(item T) bool {
		return item == 0
	})
}

// ---------------------------------------------------------------------------------------------------------------------

func TrimUIntByZero[T uint | uint8 | uint16 | uint32 | uint64](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimUIntLeftByZero[T uint | uint8 | uint16 | uint32 | uint64](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item == 0
	})
}

func TrimUIntRightByZero[T uint | uint8 | uint16 | uint32 | uint64](slice []T) []T {
	return TrimRightByFunc(slice, func(item T) bool {
		return item == 0
	})
}

// ---------------------------------------------------------------------------------------------------------------------

func TrimFloatByZero[T float32 | float64](slice []T) []T {
	return TrimByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

func TrimFloatLeftByZero[T float32 | float64](slice []T) []T {
	return TrimLeftByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

func TrimFloatRightByZero[T float32 | float64](slice []T) []T {
	return TrimRightByFunc(slice, func(item T) bool {
		return item < 0.000001
	})
}

// ---------------------------------------------------------------------------------------------------------------------
