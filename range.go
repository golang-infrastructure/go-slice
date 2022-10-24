package slice

// Range 生成给定范围 [from, to) 的数值,步进为step
func Range(from, to, step int) []int {
	slice := make([]int, 0)
	for from < to {
		slice = append(slice, from)
		from += step
	}
	return slice
}
