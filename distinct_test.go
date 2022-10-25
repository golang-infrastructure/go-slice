package slice

import (
	"testing"
)

func TestDistinct(t *testing.T) {
	slice := []int{1, 1, 2, 3, 4, 4, 7, 7, 10, 10}
	t.Log(Distinct(slice))
}

func TestDistinctByUniqIDFunc(t *testing.T) {
	slice := []int{1, 1, 2, 3, 4, 4, 7, 7, 10, 10}
	t.Log(DistinctByUniqIDFunc(slice, func(item int) int { return item }))
}
