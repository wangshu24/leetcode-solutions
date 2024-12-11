func maximumBeauty(nums []int, k int) (ans int) {
	mx := slices.Max(nums)
	diff := make([]int, mx+2)
	for _, x := range nums {
		diff[max(x-k, 0)]++
		diff[min(x+k+1, mx+1)]--
	}
	count := 0
	for _, x := range diff {
		count += x
		ans = max(ans, count)
	}
	return
}