func permuteUnique(nums []int) [][]int {
    slices.Sort(nums)
    var res [][]int
    visit := map[int]bool{}

    var backtrack func(i int, subset []int, path map[int]bool)
    backtrack = func(i int, subset []int, path map[int]bool) {
        if len(subset) == len(nums) {
            
            dst := make([]int, len(nums))
            copy(dst, subset)
            res = append(res, dst)
            return
        } else if i >= len(nums) || len(subset) > len(nums) {
            return
        }

        for j:=0; j < len(nums); j++ {
            if !path[j] {
                subset = append(subset, nums[j])
                path[j]  = true
                backtrack(i+1, subset, path)
                subset = subset[:len(subset)-1]
                path[j] = false
                for j+1 < len(nums) && nums[j] == nums[j+1] {
                    j++
                }
            }
        }
    }   

    backtrack(0, make([]int, 0), visit)
    return res 
}