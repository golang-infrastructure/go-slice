package slice

import (
	"testing"
)

func TestFillSliceRandomValue(t *testing.T) {
	slice := make([]int, 10)
	FillSliceRandomValue(slice, 100, 0)
	t.Log(slice)
}

func TestFillSliceValue(t *testing.T) {
	slice := make([]int, 10)
	FillSameValue(slice, 1)
	t.Log(slice)
}

func TestNewSliceFillRandomValue(t *testing.T) {
	t.Log(NewSliceFillRandomValue(100, 0, 10))
}
