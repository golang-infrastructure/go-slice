package slice

import (
	"testing"
)

func TestSort(t *testing.T) {
	slice := NewSliceFillRandomValue(10, 0, 100)
	Sort(slice)
	t.Log(slice)
}

func TestSortByKey(t *testing.T) {
	type Test struct {
		Key int
	}
	slice := make([]*Test, 0)
	slice = append(slice, &Test{Key: 3})
	slice = append(slice, &Test{Key: 1})
	slice = append(slice, &Test{Key: 2})
	SortByKey(slice, func(index int, item *Test) int {
		return item.Key
	})
	t.Logf("%#v", slice)
}

func TestReverse(t *testing.T) {
	slice := []int{
		5, 6, 7, 8, 9, 1, 2, 3,
	}
	reverse := Reverse(slice)
	t.Log(reverse)
	t.Log(slice)

	reverse = Reverse(slice, true)
	t.Log(reverse)
	t.Log(slice)
}

func TestMergeSortedSlices(t *testing.T) {

}

func TestIsReverseSorted(t *testing.T) {

}

func TestIsReverseSortedByKey(t *testing.T) {

}

func TestIsSorted(t *testing.T) {

}

func TestIsSortedByKey(t *testing.T) {

}

func TestMergeSortedSlicesByKey(t *testing.T) {

}

func TestUnionAndSort(t *testing.T) {
}

func TestUnionAndSortByKey(t *testing.T) {

}

func Test_reverseInSitu(t *testing.T) {

}

func Test_reverseNewSlice(t *testing.T) {

}
