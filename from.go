package slice

//import "reflect"
//
//func New[T any](itemType reflect.Type) []T {
//
//}

// FromItems 从一个或多个元素创建切片
func FromItems[T any](items ...T) []T {
	slice := make([]T, len(items), len(items))
	for index, item := range items {
		slice[index] = item
	}
	return slice
}

// FromMap 从map创建一个切片
func FromMap[K comparable, V, T any](fromMap map[K]V, fromMapFunc func(key K, value V) T) []T {
	slice := make([]T, 0)
	for key, value := range fromMap {
		slice = append(slice, fromMapFunc(key, value))
	}
	return slice
}

// FromMapFlat 从map创建一个切片，map的一个键值对可能会有多个返回值，这多个返回会被打平
func FromMapFlat[K comparable, V, T any](fromMap map[K]V, fromMapFunc func(key K, value V) []T) []T {
	slice := make([]T, 0)
	for key, value := range fromMap {
		slice = append(slice, fromMapFunc(key, value)...)
	}
	return slice
}

// FromSet 从set创建一个切片
func FromSet[K comparable, T any](set map[K]struct{}, fromSetFunc func(key K) T) []T {
	slice := make([]T, 0)
	for key := range set {
		slice = append(slice, fromSetFunc(key))
	}
	return slice
}

// FromSetFlat 从set创建一个切片，set的一个键可能会有多个返回值，这多个返回会被打平
func FromSetFlat[K comparable, T any](set map[K]struct{}, fromSetFunc func(key K) []T) []T {
	slice := make([]T, 0)
	for key := range set {
		slice = append(slice, fromSetFunc(key)...)
	}
	return slice
}
