package slice

import (
	"testing"
)

// TODO 待测试

func TestDistinct(t *testing.T) {
	slice := []int{1, 1, 2, 3, 4, 4, 7, 7, 10, 10}
	t.Log(Distinct(slice))
}

func TestDistinctByUniqIDFunc(t *testing.T) {
	slice := []int{1, 1, 2, 3, 4, 4, 7, 7, 10, 10}
	t.Log(DistinctByUniqIDFunc(slice, func(index int, item int) int { return item }))
}
