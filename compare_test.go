package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------- Compare -----------------------------------------------------------

func TestCompare(t *testing.T) {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	compare := Compare(sliceA, sliceB)
	assert.Equal(t, -1, compare)
}

func ExampleCompare() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	compare := Compare(sliceA, sliceB)
	fmt.Println(compare)
	// Output:
	// -1
}

// ------------------------------------------------- CompareByFunc -----------------------------------------------------

func TestCompareByFunc(t *testing.T) {

	type Foo struct {
		bar int
	}

	sliceA := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{5},
	}

	sliceB := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{6},
	}

	r := CompareByFunc(sliceA, sliceB, func(a, b Foo) int {
		return a.bar - b.bar
	})
	assert.Equal(t, -1, r)

}

func ExampleCompareByFunc() {

	type Foo struct {
		bar int
	}

	sliceA := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{5},
	}

	sliceB := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{6},
	}

	r := CompareByFunc(sliceA, sliceB, func(a, b Foo) int {
		return a.bar - b.bar
	})
	fmt.Println(r)
	// Output:
	// -1

}

// ------------------------------------------------- Equals ------------------------------------------------------------

func TestEquals(t *testing.T) {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	compare := Equals(sliceA, sliceB)
	assert.Equal(t, true, compare)
}

func ExampleEquals() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	compare := Equals(sliceA, sliceB)
	fmt.Println(compare)
	// Output:
	// true
}

// ------------------------------------------------- EqualsByFunc ------------------------------------------------------

func TestEqualsByFunc(t *testing.T) {

	type Foo struct {
		bar int
	}

	sliceA := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{5},
	}

	sliceB := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{6},
	}

	r := EqualsByFunc(sliceA, sliceB, func(a, b Foo) bool {
		return a.bar == b.bar
	})
	assert.Equal(t, false, r)

}

func ExampleEqualsByFunc() {

	type Foo struct {
		bar int
	}

	sliceA := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{5},
	}

	sliceB := []Foo{
		Foo{1},
		Foo{2},
		Foo{3},
		Foo{4},
		Foo{6},
	}

	r := EqualsByFunc(sliceA, sliceB, func(a, b Foo) bool {
		return a.bar == b.bar
	})
	fmt.Println(r)
	// Output:
	// false

}

// ---------------------------------------------------------------------------------------------------------------------
