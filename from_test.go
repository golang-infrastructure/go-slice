package slice

import (
	"testing"
)

func TestFromItems(t *testing.T) {
	t.Log(FromItems(1, 2, 3))
}

func TestFromMapByFunc(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	slice := FromMapByFunc(a, func(key string, value int) int {
		return value + 1
	})
	t.Log(slice)
}

func TestFromMapFlatByFunc(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	slice := FromMapFlatByFunc(a, func(key string, value int) []int {
		return []int{value, value + 1}
	})
	t.Log(slice)
}

func TestFromMapKey(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	slice := FromMapKey(a)
	t.Log(slice)
}

func TestFromMapValue(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	slice := FromMapValue(a)
	t.Log(slice)
}

func TestFromSet(t *testing.T) {
	a := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
	}
	slice := FromSet(a)
	t.Log(slice)
}

func TestFromSetByFunc(t *testing.T) {
	a := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
	}
	slice := FromSetByFunc(a, func(key string) string {
		return key + "aaaa"
	})
	t.Log(slice)
}

func TestFromSetFlatByFunc(t *testing.T) {
	a := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
	}
	slice := FromSetFlatByFunc(a, func(key string) []string {
		return []string{key, key + "aaaa"}
	})
	t.Log(slice)
}

func TestFromMapValueFlat(t *testing.T) {
	m := map[string][]int{
		"a": []int{1, 2, 3},
		"b": []int{11, 22, 33},
		"c": []int{91, 72, 63},
	}
	flat := FromMapValueFlat(m)
	t.Log(flat)
}
