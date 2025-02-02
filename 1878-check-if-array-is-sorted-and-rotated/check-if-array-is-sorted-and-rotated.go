import "slices"
func check(nums []int) bool {
    if slices.IsSorted(nums) {return true}

    new := slices.Concat(nums, nums)
    count, small  := 0, slices.Min(nums)
    for  _, val := range new {
        if val >= small {
            small = val 
            count++
        } else {
            small = min(small, val)
            count = 1
        }
        if count == len(nums) {
            return true
        }
    }

    return false
}

func min (a,b int) int {
    if a > b {return b}
    return a
}