package slice

import (
	"testing"
)

func TestChooseEvenIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseEvenIndexes(slice))
}

func TestChooseIndex(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseIndex(slice, 1))
	t.Log(ChooseIndex(slice, 100))
}

func TestChooseIndexOrDefault(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseIndexOrDefault(slice, 1, 10))
	t.Log(ChooseIndexOrDefault(slice, 100, 100))
}

func TestChooseIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseIndexes(slice))
	t.Log(ChooseIndexes(slice, 1, 100))
}

func TestChooseMiddleIndex(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseMiddleIndex(slice))
}

func TestChooseOddIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	t.Log(ChooseOddIndexes(slice))
}

func TestChooseRandomIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(Random(slice))
}

func TestChooseRandomIndexes(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(Randoms(slice, 1))
	t.Log(Randoms(slice, 4))
	t.Log(Randoms(slice, 100))
	t.Log(Randoms(slice, 0))
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(Filter(slice, func(item int) bool {
		return item%2 == 0
	}))
}

func TestFirstItem(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(FirstItem(slice))
}

func TestFirstItemOrDefault(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(FirstItemOrDefault(slice, 1))
	t.Log(FirstItemOrDefault(slice, 0))
	t.Log(FirstItemOrDefault(slice, -1))
}

func TestFirstItems(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(FirstItems(slice, 1))
	t.Log(FirstItems(slice, 0))
	t.Log(FirstItems(slice, -1))
}

func TestLastItem(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(LastItem(slice))
	var s []int
	t.Log(LastItem(s))
	var s2 []any
	t.Log(LastItem(s2))
}

func TestLastItemOrDefault(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(LastItemOrDefault(slice, 1))
	var s []int
	t.Log(LastItemOrDefault(s, 1))
}

func TestLastItems(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(LastItems(slice, 3))
	t.Log(LastItems(slice, 0))
	t.Log(LastItems(slice, -1))
}

func TestSubSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(SubSlice(slice, 0, 3))
	t.Log(SubSlice(slice, -1, 10000))
}
