package slice

import (
	"testing"
)

func TestClone(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4}
	s1 := Clone(slice)
	t.Log(s1)
	header := GetSliceHeader(s1)
	t.Log(header)
}
