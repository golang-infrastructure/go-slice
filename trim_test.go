package slice

import (
	"fmt"
	"testing"
)

func TestTrim(t *testing.T) {
	slice := []int{0}
	slice = Trim(slice)
	t.Log(slice)
}

func ExampleTrim() {
	fmt.Println("OK")
	// Output:
	// OK
}

func TestTrimRight(t *testing.T) {
	slice := []int{1}
	slice = TrimRight(slice)
	t.Log(slice)
}
