package slice

import "context"

// SendToChannel 把切片发送到给定的channel中，返回值是发送成功的元素的个数
func SendToChannel[T any](ctx context.Context, slice []T, channel chan T) int {
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

// ReadFromChannel 从channel中读取给定数量的元素放到切片中返回，读取成功的元素的数量可以通过返回的切片的长度来确定
func ReadFromChannel[T any](ctx context.Context, channel chan T, limit int) []T {
	slice := make([]T, 0)
	for limit > 0 {
		select {
		case <-ctx.Done():
			break
		case v, ok := <-channel:
			if !ok {
				break
			}
			slice = append(slice, v)
			limit--
		}
	}
	return slice
}

// ToChannel 把切片转为channel
func ToChannel[T any](slice []T) chan T {
	resultChannel := make(chan T, len(slice))
	for _, item := range slice {
		resultChannel <- item
	}
	return resultChannel
}
