package slice

import (
	"reflect"
	"unsafe"
)

// 反射相关的操作

// 访问slice的元数据
// 通过指针直接访问等

// ---------------------------------------------------------------------------------------------------------------------

func IsSliceType(v any) bool {
	of := reflect.TypeOf(v)
	return of.Kind() == reflect.Slice
}

func IsSliceOrArrayType(v any) bool {
	of := reflect.TypeOf(v)
	return of.Kind() == reflect.Slice || of.Kind() == reflect.Array
}

// ---------------------------------------------------------------------------------------------------------------------

func IsSliceValue(v any) bool {
	of := reflect.ValueOf(v)
	return of.Kind() == reflect.Slice
}

func IsSliceOrArrayValue(v any) bool {
	of := reflect.ValueOf(v)
	return of.Kind() == reflect.Slice || of.Kind() == reflect.Array
}

// ---------------------------------------------------------------------------------------------------------------------

func GetSliceHeader[T any](slice []T) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&slice))
}

// ---------------------------------------------------------------------------------------------------------------------
