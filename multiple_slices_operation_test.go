package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeOrderedSlices(t *testing.T) {

	sliceA := []int{1, 3, 66, 888, 9999}
	sliceB := []int{-1111, -6, -1, 23123}
	sliceC := []int{-100, 12, 1231, 123123}

	slices := MergeOrderedSlices(sliceA, sliceB, sliceC)

	excepted := Clone(slices)
	Sort(excepted)

	assert.Equal(t, excepted, slices)

}
