package slice

import (
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-maths"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------

// Trim 去除切片两端的零值元素
func Trim[T comparable](slice []T) []T {
	var zero T
	return TrimByFunc(slice, func(index int, item T) bool {
		return item == zero
	})
}

// TrimLeft 去除切片左边的零值元素
func TrimLeft[T comparable](slice []T) []T {
	var zero T
	return TrimLeftByFunc(slice, func(index int, item T) bool {
		return item == zero
	})
}

// TrimRight 去除右边的零值元素
func TrimRight[T comparable](slice []T) []T {
	var zero T
	return TrimRightByFunc[T](slice, func(index int, item T) bool {
		return item == zero
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// TrimByFunc 从切片两侧trim，元素是否需要trim由自定义的trim函数决定
func TrimByFunc[T any](slice []T, itemFilterFunc func(index int, value T) bool) []T {
	t := TrimLeftByFunc(slice, itemFilterFunc)
	return TrimRightByFunc(t, itemFilterFunc)
}

// TrimLeftByFunc 从切片左边trim，根据自定义的函数来判定是否需要trim
func TrimLeftByFunc[T any](slice []T, itemFilterFunc func(index int, value T) bool) []T {
	if len(slice) == 0 {
		return nil
	}
	left := 0
	for left < len(slice) {
		if itemFilterFunc(left, slice[left]) {
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
func TrimRightByFunc[T any](slice []T, trimFunc func(index int, value T) bool) []T {
	if len(slice) == 0 {
		return nil
	}
	right := len(slice) - 1
	for right >= 0 {
		if trimFunc(right, slice[right]) {
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
	return TrimByFunc(slice, func(index int, item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimLeftByBlankString(slice []string) []string {
	return TrimLeftByFunc(slice, func(index int, item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimRightByBlankString(slice []string) []string {
	return TrimRightByFunc(slice, func(index int, item string) bool {
		return strings.TrimSpace(item) == ""
	})
}

func TrimByEmptyString(slice []string) []string {
	return TrimByFunc(slice, func(index int, item string) bool {
		return item == ""
	})
}

func TrimLeftByEmptyString(slice []string) []string {
	return TrimLeftByFunc(slice, func(index int, item string) bool {
		return item == ""
	})
}

func TrimRightByEmptyString(slice []string) []string {
	return TrimRightByFunc(slice, func(index int, item string) bool {
		return item == ""
	})
}

// ---------------------------------------------------------------------------------------------------------------------

func TrimIntegerByZero[T gtypes.Integer](slice []T) []T {
	return TrimByFunc(slice, func(index int, item T) bool {
		return item == 0
	})
}

func TrimIntegerLeftByZero[T gtypes.Integer](slice []T) []T {
	return TrimLeftByFunc(slice, func(index int, item T) bool {
		return item == 0
	})
}

func TrimIntegerRightByZero[T gtypes.Integer](slice []T) []T {
	return TrimRightByFunc(slice, func(index int, item T) bool {
		return item == 0
	})
}

// ---------------------------------------------------------------------------------------------------------------------

// TrimFloatByZero 把浮点型切片的两侧的零值去除
func TrimFloatByZero[T gtypes.Float](slice []T) []T {
	return TrimByFunc(slice, func(index int, item T) bool {
		return maths.IsZero(item)
	})
}

// TrimFloatLeftByZero 把浮点型切片的左侧的零值去除
func TrimFloatLeftByZero[T gtypes.Float](slice []T) []T {
	return TrimLeftByFunc(slice, func(index int, item T) bool {
		return maths.IsZero(item)
	})
}

// TrimFloatRightByZero 把浮点型切片的右侧的零值去除
func TrimFloatRightByZero[T gtypes.Float](slice []T) []T {
	return TrimRightByFunc(slice, func(index int, item T) bool {
		return maths.IsZero(item)
	})
}

// ---------------------------------------------------------------------------------------------------------------------
