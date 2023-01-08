package slice

import (
	"testing"
)

func TestUnion(t *testing.T) {

	// empty
	var s1 []string
	s2 := []string{"a", "b"}
	s3 := Union(s1, s2)
	t.Log(s3)

	type foo struct {
		bar string
	}
	f1 := []*foo{&foo{bar: "001"}}
	f2 := []*foo{&foo{bar: "002"}}
	f3 := Union(f1, f2)
	t.Log(f3)

}

func TestIntersection(t *testing.T) {
	s1 := []string{"b"}
	s2 := []string{"a", "b"}
	s3 := Intersection(s1, s2)
	t.Log(s3)
}

func TestSubtract(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{"b"}
	s3 := Subtract(s1, s2)
	t.Log(s3)
}

func TestNotIntersection(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{"b"}
	s3 := NotIntersection(s1, s2)
	t.Log(s3)
}
