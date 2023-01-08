package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnion(t *testing.T) {

	// case 1: empty
	var s1 []string
	s2 := []string{"a", "b"}
	s3 := Union(s1, s2)
	excepted := []string{"a", "b"}
	assert.Equal(t, excepted, s3)

	// case 2: 合并两个非空切片
	type foo struct {
		bar string
	}
	f1 := []*foo{&foo{bar: "001"}}
	f2 := []*foo{&foo{bar: "002"}}
	f3 := Union(f1, f2)
	exceptedFoo := []*foo{f1[0], f2[0]}
	assert.Equal(t, exceptedFoo, f3)

}

func TestIntersection(t *testing.T) {

	// case 1: 没有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a"}
		s3 := Intersection(s1, s2)
		assert.Equal(t, 0, len(s3))
	}

	// case 2: 有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a", "b"}
		s3 := Intersection(s1, s2)
		excepted := []string{"b"}
		assert.Equal(t, excepted, s3)
	}

}

func TestIntersectionByKeyFunc(t *testing.T) {

	// case 1: 没有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a"}
		s3 := IntersectionByKeyFunc(s1, s2, func(index int, item string) string {
			return item
		})
		assert.Equal(t, 0, len(s3))
	}

	// case 2: 有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a", "b"}
		s3 := IntersectionByKeyFunc(s1, s2, func(index int, item string) string {
			return item
		})
		excepted := []string{"b"}
		assert.Equal(t, excepted, s3)
	}
}

func TestNotIntersection(t *testing.T) {

	// case 1: 没有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a"}
		s3 := NotIntersection(s1, s2)
		Sort(s3)
		assert.Equal(t, []string{"a", "b"}, s3)
	}

	// case 2: 有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a", "b"}
		s3 := NotIntersection(s1, s2)
		assert.Equal(t, []string{"a"}, s3)
	}

}

func TestNotIntersectionByKeyFunc(t *testing.T) {
	// case 1: 没有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a"}
		s3 := NotIntersectionByKeyFunc(s1, s2, func(index int, item string) string {
			return item
		})
		Sort(s3)
		assert.Equal(t, []string{"a", "b"}, s3)
	}

	// case 2: 有交集
	{
		s1 := []string{"b"}
		s2 := []string{"a", "b"}
		s3 := NotIntersectionByKeyFunc(s1, s2, func(index int, item string) string {
			return item
		})
		assert.Equal(t, []string{"a"}, s3)
	}
}

func TestSubtract(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{"b"}
	s3 := Subtract(s1, s2)
	assert.Equal(t, []string{"a"}, s3)
}

func TestSubtractByKeyFunc(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{"b"}
	s3 := SubtractByKeyFunc(s1, s2, func(index int, item string) string {
		return item
	})
	assert.Equal(t, []string{"a"}, s3)
}
