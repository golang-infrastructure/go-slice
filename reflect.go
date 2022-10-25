package slice

import (
	"reflect"
	"unsafe"
)

// 反射相关的操作

// 访问slice的元数据
// 通过指针直接访问等

// ---------------------------------------------------------------------------------------------------------------------

// IsSliceType 判断是否是切片类型
func IsSliceType(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	return of.Kind() == reflect.Slice
}

// IsArrayType 判断值是否是数组类型
func IsArrayType(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	return of.Kind() == reflect.Array
}

// IsSliceOrArrayType 判断是是否是切片或者数组类型
func IsSliceOrArrayType(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	return of.Kind() == reflect.Slice || of.Kind() == reflect.Array
}

// ---------------------------------------------------------------------------------------------------------------------

// IsSliceValue 判断存储的值是否是切片
func IsSliceValue(v any) bool {
	of := reflect.ValueOf(v)
	if !of.IsValid() {
		return false
	}
	return of.Kind() == reflect.Slice
}

// IsArrayValue 判断存储的值是否是数组
func IsArrayValue(v any) bool {
	of := reflect.ValueOf(v)
	if !of.IsValid() {
		return false
	}
	return of.Kind() == reflect.Array
}

// IsSliceOrArrayValue 判断存储的值是否是切片或数组
func IsSliceOrArrayValue(v any) bool {
	of := reflect.ValueOf(v)
	if !of.IsValid() {
		return false
	}
	return of.Kind() == reflect.Slice || of.Kind() == reflect.Array
}

// ---------------------------------------------------------------------------------------------------------------------

// GetSliceHeader 通过反射获取切片的元数据信息
func GetSliceHeader[T any](slice []T) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&slice))
}

// ---------------------------------------------------------------------------------------------------------------------
