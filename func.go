package slice

// 因为Goland的提示不支持这种自定义的类型，所以这里只是做类型定义，并不实际使用
//// ItemFilterFunc 表示对切片中的元素进行过滤操作，返回true表示ok，返回false表示要被过滤掉
//type ItemFilterFunc[T any] func(index int, value T) bool
//
//// ItemMapFunc 表示对切片中的元素进行转换操作
//type ItemMapFunc[From, To any] func(index int, value From) To
//
//// ItemForeachFunc 表示对切片中的每个元素都处理一下，没有返回值
//type ItemForeachFunc[T any] func(index int, value T)
//
//// ItemVisitFunc 表示对切片中的每个元素都处理一下，返回值表示是否继续访问下一个元素
//type ItemVisitFunc[T any] func(index int, value T) bool
//
//// KeyFunc 为元素生成一个Key
//type KeyFunc[T any, K comparable] func(index int, item T) K
//
//// OrderedKeyFunc ComparableKeyFunc 为元素生成一个key，这个key要能够参与排序
//type OrderedKeyFunc[T any, K gtypes.Ordered] func(index int, item T) K
//
//// UniqIDFunc 为元素生成一个唯一ID，这个唯一ID要能够参与比较，不然怎么证明它是唯一的呢
//type UniqIDFunc[T any, K comparable] func(index int, item T) K
