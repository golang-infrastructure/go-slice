package slice

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------- ReadFromChannel ---------------------------------------------------

func TestReadFromChannel(t *testing.T) {
	channel := make(chan int, 10)
	for i := 0; i < 3; i++ {
		channel <- i
	}
	close(channel)
	slice := ReadFromChannel(context.Background(), channel)
	exceptedSlice := []int{0, 1, 2}
	assert.Equal(t, exceptedSlice, slice)
}

func ExampleReadFromChannel() {
	channel := make(chan int, 10)
	for i := 0; i < 3; i++ {
		channel <- i
	}
	close(channel)
	slice := ReadFromChannel(context.Background(), channel)
	fmt.Println(slice)
	// Output:
	// [0 1 2]
}

// ------------------------------------------------- SendToChannel -----------------------------------------------------

func TestSendToChannel(t *testing.T) {

	// 发送到正常的channel
	channel := make(chan int, 10)
	slice := []int{0, 1, 2}
	sendSuccessCount := SendToChannel(context.Background(), slice, channel)
	assert.Equal(t, 3, sendSuccessCount)

	// 发送到关闭的channel
	close(channel)
	sendSuccessCount = SendToChannel(context.Background(), slice, channel)
	assert.Equal(t, 0, sendSuccessCount)

}

func ExampleSendToChannel() {

	// 发送到正常的channel
	channel := make(chan int, 10)
	slice := []int{0, 1, 2}
	sendSuccessCount := SendToChannel(context.Background(), slice, channel)
	fmt.Println(sendSuccessCount)
	// Output:
	// 3

}

// ------------------------------------------------- ToChannel ---------------------------------------------------------

func TestToChannel(t *testing.T) {
	slice := []int{0, 1, 2}
	channel := ToChannel(slice)
	close(channel)
	fromChannelSlice := ReadFromChannel(context.Background(), channel)
	assert.Equal(t, slice, fromChannelSlice)
}

func ExampleToChannel() {
	slice := []int{0, 1, 2}
	channel := ToChannel(slice)
	close(channel)
	fromChannelSlice := ReadFromChannel(context.Background(), channel)
	fmt.Println(fromChannelSlice)
	// Output:
	// [0 1 2]
}

// ------------------------------------------------- ToClosedChannel ---------------------------------------------------

func TestToClosedChannel(t *testing.T) {
	slice := []int{0, 1, 2}
	channel := ToClosedChannel(slice)
	fromChannelSlice := ReadFromChannel(context.Background(), channel)
	assert.Equal(t, slice, fromChannelSlice)
}

func ExampleToClosedChannel() {
	slice := []int{0, 1, 2}
	channel := ToClosedChannel(slice)
	fromChannelSlice := ReadFromChannel(context.Background(), channel)
	fmt.Println(fromChannelSlice)
	// Output:
	// [0 1 2]
}

// ---------------------------------------------------------------------------------------------------------------------
