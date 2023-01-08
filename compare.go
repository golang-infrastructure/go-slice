package slice

import (
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-maths"
)

// ---------------------------------------------------------------------------------------------------------------------

// Equals 比较两个切片是否相等，需要切片中的元素是comparable的
func Equals[T comparable](sliceA, sliceB []T) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, itemA := range sliceA {
		itemB := sliceB[index]
		if itemA != itemB {
			return false
		}
	}
	return true
}

// EqualsByFunc 比较两个切片是否相等，根据自定义的比较器来决定是否相等，切片中的元素可以是任意类型
func EqualsByFunc[T any](sliceA, sliceB []T, equalsFunc compare_anything.EqualsFunc[T]) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for index, itemA := range sliceA {
		itemB := sliceB[index]
		if !equalsFunc(itemA, itemB) {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------

// Compare 比较两个切片的大小，需要切片中的元素时是Ordered的
func Compare[T gtypes.Ordered](sliceA, sliceB []T) int {
	// 先按照都有的长度进行比较
	endIndex := maths.Min(len(sliceA), len(sliceB))
	for index := 0; index < endIndex; index++ {
		itemA := sliceA[index]
		itemB := sliceB[index]
		if itemA == itemB {
			continue
		}
		if itemA < itemB {
			return -1
		} else {
			return 1
		}
	}
	// 如果元素比较完了都没法分出大小，则认为还有剩下的那个是较大的一个
	return len(sliceA) - len(sliceB)
}

// CompareByFunc 使用自定义的比较器比较两个函数是否相等，切片中的元素可以是任意类型
func CompareByFunc[T any](sliceA, sliceB []T, comparator compare_anything.Comparator[T]) int {
	endIndex := maths.Min(len(sliceA), len(sliceB))
	for index := 0; index < endIndex; index++ {
		r := comparator(sliceA[index], sliceB[index])
		if r != 0 {
			return r
		}
	}
	return len(sliceA) - len(sliceB)
}

// ---------------------------------------------------------------------------------------------------------------------
