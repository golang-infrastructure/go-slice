package slice

import (
	"testing"
)

func TestRange(t *testing.T) {
	t.Log(Range(0, 10, 1))
	t.Log(Range(0, 10, 2))
	t.Log(Range(0, 10, 3))
	t.Log(Range(0, 0, 3))
	t.Log(Range(0, 10, -3))
	t.Log(Range(10, 1, -1))
	t.Log(Range(10, 1, 1))
}
