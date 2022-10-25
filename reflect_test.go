package slice

import (
	"testing"
)

func TestGetSliceHeader(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	header := GetSliceHeader(slice)
	t.Log(header)
}

func TestIsSliceOrArrayValue(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsSliceOrArrayValue(slice))
	t.Log(IsSliceOrArrayValue(nil))
	t.Log(IsSliceOrArrayValue("CC11001100"))
	t.Log(IsSliceOrArrayValue([3]int{}))
}

func TestIsSliceValue(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsSliceValue(slice))
	t.Log(IsSliceValue(nil))
	t.Log(IsSliceValue("CC11001100"))
	t.Log(IsSliceValue([3]int{}))
}

func TestIsArrayValue(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsArrayValue(slice))
	t.Log(IsArrayValue(nil))
	t.Log(IsArrayValue("CC11001100"))
	t.Log(IsArrayValue([3]int{}))
}

func TestIsSliceOrArrayType(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsSliceOrArrayType(slice))
	t.Log(IsSliceOrArrayType(nil))
	t.Log(IsSliceOrArrayType("CC11001100"))
	t.Log(IsSliceOrArrayType([3]int{}))
}

func TestIsArrayType(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsArrayType(slice))
	t.Log(IsArrayType(nil))
	t.Log(IsArrayType("CC11001100"))
	t.Log(IsArrayType([3]int{}))
}

func TestIsSliceType(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 10)
	t.Log(IsSliceType(slice))
	t.Log(IsSliceType(nil))
	t.Log(IsSliceType("CC11001100"))
	t.Log(IsSliceType([3]int{}))
}