func minimumOperations(nums []int) int {
    n := len(nums)
    freq := make([]bool, 101)

    for i:=n-1; i >= 0; i-- {
        if freq[nums[i]] {
            return (i/3)+1
        }
        freq[nums[i]] = true
    }

    return 0
}