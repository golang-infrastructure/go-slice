package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChooseEvenIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, []int{0, 2, 4}, ChooseEvenIndexes(slice))
}

func TestChooseIndex(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, 1, ChooseIndex(slice, 1))
	assert.Equal(t, 0, ChooseIndex(slice, 100))
}

func TestChooseIndexOrDefault(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, 1, ChooseIndexOrDefault(slice, 1, 10))
	assert.Equal(t, 100, ChooseIndexOrDefault(slice, 100, 100))
}

func TestChooseIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, 0, len(ChooseIndexes(slice)))
	assert.Equal(t, []int{1, 0}, ChooseIndexes(slice, 1, 100))
}

func TestChooseMiddleIndex(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	a, b := ChooseMiddleIndex(slice)
	assert.Equal(t, 2, a)
	assert.Equal(t, 3, b)
}

func TestChooseOddIndexes(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, []int{1, 3, 5}, ChooseOddIndexes(slice))
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
	filterResult := Filter(slice, func(index int, item int) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{2, 4}, filterResult)
}

func TestFirstItem(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 1, FirstItem(slice))
}

func TestFirstItemOrDefault(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 1, FirstItemOrDefault(slice, 1))
	assert.Equal(t, 1, FirstItemOrDefault(slice, 0))
	assert.Equal(t, -1, FirstItemOrDefault(nil, -1))
}

func TestFirstItems(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{1}, FirstItems(slice, 1))
	assert.Equal(t, 0, len(FirstItems(slice, 0)))
	assert.Equal(t, []int{5}, FirstItems(slice, -1))
}

func TestLastItem(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 5, LastItem(slice))
	var s []int
	assert.Equal(t, 0, LastItem(s))
	var s2 []any
	assert.Equal(t, nil, LastItem(s2))
}

func TestLastItemOrDefault(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 5, LastItemOrDefault(slice, 1))
	var s []int
	assert.Equal(t, 1, LastItemOrDefault(s, 1))
}

func TestLastItems(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{3, 4, 5}, LastItems(slice, 3))
	assert.Equal(t, 0, len(LastItems(slice, 0)))
	assert.Equal(t, 1, len(LastItems(slice, -1)))
}

func TestSubSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{1, 2, 3}, SubSlice(slice, 0, 3))
	assert.Equal(t, slice, SubSlice(slice, -1, 10000))
}
