func minimizedMaximum(n int, quantities []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(quantities)))
	emptyShop := n - len(quantities)
	// binary-search,  must take care of left-bound and right-bound
	l, r := 1, quantities[0]
	minMax := r
	loop:
	for l <= r {
		m := (l+r)>>1
		resShop := emptyShop
		for _, v := range quantities {
			if v <= m {
				break
			}
			needShop := v / m
			if v % m == 0 {
				needShop--
			}
			resShop -= needShop
			if resShop < 0 {
				l = m + 1
				continue loop
			}
		}
		if m < minMax {
			minMax = m
		}
		r = m - 1
	}
	return minMax
}