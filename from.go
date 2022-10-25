package slice

// TODO 需要思考下是否要加入对这个方法的支持，会有什么实际场景用得到吗？
//func New[T any](itemType reflect.Type) []T {
//
//}

// ---------------------------------------------------------------------------------------------------------------------

// Of 从一个或多个元素创建切片，是FromItems的别名
func Of[T any](items ...T) []T {
	return FromItems(items...)
}

// FromItems 从一个或多个元素创建切片
func FromItems[T any](items ...T) []T {
	slice := make([]T, len(items))
	for index, item := range items {
		slice[index] = item
	}
	return slice
}

// ---------------------------------------------------------------------------------------------------------------------

// FromMapByFunc 从map创建一个切片，使用自定义的函数从map中抽取值
func FromMapByFunc[K comparable, V, T any](fromMap map[K]V, fromMapFunc func(key K, value V) T) []T {
	slice := make([]T, 0)
	for key, value := range fromMap {
		slice = append(slice, fromMapFunc(key, value))
	}
	return slice
}

// FromMapFlatByFunc 从map创建一个切片，map的一个键值对可能会有多个返回值，这多个返回会被打平
func FromMapFlatByFunc[K comparable, V, T any](fromMap map[K]V, fromMapFunc func(key K, value V) []T) []T {
	slice := make([]T, 0)
	for key, value := range fromMap {
		slice = append(slice, fromMapFunc(key, value)...)
	}
	return slice
}

// FromMapKey 把map的key转为切片
func FromMapKey[K comparable, V any](fromMap map[K]V) []K {
	slice := make([]K, 0)
	for key := range fromMap {
		slice = append(slice, key)
	}
	return slice
}

// FromMapValue 把map的value转为切片
func FromMapValue[K comparable, V any](fromMap map[K]V) []V {
	slice := make([]V, 0)
	for _, value := range fromMap {
		slice = append(slice, value)
	}
	return slice
}

// FromMapValueFlat FromMapValue 把map的value转为切片，当value是一个切片的时候使用
func FromMapValueFlat[K comparable, V any](fromMap map[K][]V) []V {
	slice := make([]V, 0)
	for _, value := range fromMap {
		slice = append(slice, value...)
	}
	return slice
}

// ---------------------------------------------------------------------------------------------------------------------

// FromSet 把se转换为切片
func FromSet[K comparable](set map[K]struct{}) []K {
	slice := make([]K, 0)
	for key := range set {
		slice = append(slice, key)
	}
	return slice
}

// FromSetByFunc 从set创建一个切片
func FromSetByFunc[K comparable, T any](set map[K]struct{}, fromSetFunc func(key K) T) []T {
	slice := make([]T, 0)
	for key := range set {
		slice = append(slice, fromSetFunc(key))
	}
	return slice
}

// FromSetFlatByFunc 从set创建一个切片，set的一个键可能会有多个返回值，这多个返回会被打平
func FromSetFlatByFunc[K comparable, T any](set map[K]struct{}, fromSetFunc func(key K) []T) []T {
	slice := make([]T, 0)
	for key := range set {
		slice = append(slice, fromSetFunc(key)...)
	}
	return slice
}

// ---------------------------------------------------------------------------------------------------------------------
