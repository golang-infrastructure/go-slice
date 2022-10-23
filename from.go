package slice

//import "reflect"
//
//func New[T any](itemType reflect.Type) []T {
//
//}

// From 从一个或多个元素创建切片
func From[T any](items ...T) []T {
	slice := make([]T, len(items), len(items))
	for index, item := range items {
		slice[index] = item
	}
	return slice
}
