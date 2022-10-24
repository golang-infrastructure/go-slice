package slice

import "testing"

func TestTrim(t *testing.T) {
	slice := []int{0}
	slice = Trim(slice)
	t.Log(slice)
}

func TestTrimRight(t *testing.T) {
	slice := []int{1}
	slice = TrimRight(slice)
	t.Log(slice)
}
