package slice

// Range 生成给定范围 [from, to) 的数值,步进为step
func Range(from, to, step int) []int {

	// 非法的step直接返回空切片
	if step == 0 {
		return nil
	}

	slice := make([]int, 0)
	if from <= to {
		// 正序，正序的时候step必须是一个正数
		if step < 0 {
			return nil
		}
		for from < to {
			slice = append(slice, from)
			from += step
		}
	} else {
		// 倒序，倒序的时候step不论是正是负数都认为是每次减去 |step| 步
		if step > 0 {
			step *= -1
		}
		for from > to {
			slice = append(slice, from)
			from += step
		}
	}
	return slice
}
