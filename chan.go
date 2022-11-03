package slice

import "context"

// SendToChannel 把切片发送到channel中，返回值是发送成功的元素的个数，因为可能会涉及到发送超时、channel关闭之类的
func SendToChannel[T any](ctx context.Context, slice []T, channel chan T) int {

	// channel未初始化或者slice空的话直接返回
	if channel == nil || len(slice) == 0 {
		return 0
	}

	// 防止正发送着呢，channel被关闭了
	defer func() {
		if r := recover(); r != nil {
			// ignored
		}
	}()

	sendSuccessCount := 0
	for _, item := range slice {
		select {
		case channel <- item:
			sendSuccessCount++
			continue
		case <-ctx.Done():
			return sendSuccessCount
		}
	}
	return sendSuccessCount
}

// ReadFromChannel 从channel中读取给定数量的元素放到切片中返回，读取成功的元素的数量可以通过返回的切片的长度来确定，limit为可选参数，如果指定为正数则只会读取limit个元素，不传或者负数则表示一直读取直到channel关闭或超时为止
func ReadFromChannel[T any](ctx context.Context, channel chan T, limit ...int) []T {

	// 防止channel未初始化
	if channel == nil {
		return nil
	}

	// 设置默认值
	limit = append(limit, -1)

	slice := make([]T, 0)
	for limit[0] != 0 {
		select {
		case <-ctx.Done():
			return slice
		case v, ok := <-channel:
			if !ok {
				return slice
			}
			slice = append(slice, v)
			limit[0]--
		}
	}
	return slice
}

// ToChannel 把切片转为channel，channel的缓冲区大小等于切片的元素个数，注意，返回的channel并没有被关闭
func ToChannel[T any](slice []T) chan T {
	resultChannel := make(chan T, len(slice))
	for _, item := range slice {
		resultChannel <- item
	}
	return resultChannel
}

// ToClosedChannel 把切片转为channel，channel的缓冲区大小等于切片的元素个数，并且返回的channel是被关闭的
func ToClosedChannel[T any](slice []T) chan T {
	channel := ToChannel(slice)
	close(channel)
	return channel
}
