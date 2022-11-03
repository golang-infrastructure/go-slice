package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4}
	copySlice := Clone(slice)
	assert.Equal(t, slice, copySlice)
}

func ExampleClone() {
	slice := []int{0, 1, 2, 3, 4}
	copySlice := Clone(slice)
	fmt.Println(copySlice)
	// Output:
	// [0 1 2 3 4]
}
